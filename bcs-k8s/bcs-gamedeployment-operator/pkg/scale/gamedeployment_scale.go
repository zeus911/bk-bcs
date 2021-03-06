/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package scale

import (
	"fmt"
	"sync"

	tkexv1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-gamedeployment-operator/pkg/apis/tkex/v1alpha1"
	tkexclientset "github.com/Tencent/bk-bcs/bcs-k8s/bcs-gamedeployment-operator/pkg/client/clientset/versioned"
	gdcore "github.com/Tencent/bk-bcs/bcs-k8s/bcs-gamedeployment-operator/pkg/core"
	"github.com/Tencent/bk-bcs/bcs-k8s/bcs-gamedeployment-operator/pkg/util"
	"github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/common/expectations"
	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
)

const (
	// LengthOfInstanceID is the length of instance-id
	LengthOfInstanceID = 5

	// When batching pod creates, initialBatchSize is the size of the initial batch.
	initialBatchSize = 1
)

// Interface for managing replicas including create and delete pod/pvc.
type Interface interface {
	Manage(
		currentCS, updateCS *tkexv1alpha1.GameDeployment,
		currentRevision, updateRevision string,
		pods []*v1.Pod,
	) (bool, error)
}

// New returns a scale control.
func New(kubeClient clientset.Interface, tkexClient tkexclientset.Interface, recorder record.EventRecorder, exp expectations.ScaleExpectations) Interface {
	return &realControl{kubeClient: kubeClient, tkexClient: tkexClient, recorder: recorder, exp: exp}
}

type realControl struct {
	kubeClient clientset.Interface
	tkexClient tkexclientset.Interface
	recorder   record.EventRecorder
	exp        expectations.ScaleExpectations
}

func (r *realControl) Manage(
	currentGD, updateGD *tkexv1alpha1.GameDeployment,
	currentRevision, updateRevision string,
	pods []*v1.Pod,
) (bool, error) {
	if updateGD.Spec.Replicas == nil {
		return false, fmt.Errorf("spec.Replicas is nil")
	}

	controllerKey := util.GetControllerKey(updateGD)
	coreControl := gdcore.New(updateGD)
	if !coreControl.IsReadyToScale() {
		klog.Warningf("GameDeployment %s skip scaling for not ready to scale", controllerKey)
		return false, nil
	}

	if podsToDelete := getPodsToDelete(updateGD, pods); len(podsToDelete) > 0 {
		klog.V(3).Infof("GameDeployment %s begin to delete pods in podsToDelete: %v", controllerKey, podsToDelete)
		return r.deletePods(updateGD, podsToDelete)
	}

	updatedPods, notUpdatedPods := util.SplitPodsByRevision(pods, updateRevision)

	diff, currentRevDiff := calculateDiffs(updateGD, updateRevision == currentRevision, len(pods), len(notUpdatedPods))

	if diff < 0 {
		// total number of this creation
		expectedCreations := diff * -1
		// lack number of current version
		expectedCurrentCreations := 0
		if currentRevDiff < 0 {
			expectedCurrentCreations = currentRevDiff * -1
		}

		klog.V(3).Infof("GameDeployment %s begin to scale out %d pods including %d (current rev)",
			controllerKey, expectedCreations, expectedCurrentCreations)

		// generate available ids
		availableIDs := genAvailableIDs(expectedCreations, pods)

		return r.createPods(expectedCreations, expectedCurrentCreations,
			currentGD, updateGD, currentRevision, updateRevision, availableIDs.List())

	} else if diff > 0 {
		klog.V(3).Infof("GameDeployment %s begin to scale in %d pods including %d (current rev)",
			controllerKey, diff, currentRevDiff)

		podsToDelete := choosePodsToDelete(diff, currentRevDiff, notUpdatedPods, updatedPods)

		return r.deletePods(updateGD, podsToDelete)
	}

	return false, nil
}

func (r *realControl) createPods(
	expectedCreations, expectedCurrentCreations int,
	currentGD, updateGD *tkexv1alpha1.GameDeployment,
	currentRevision, updateRevision string,
	availableIDs []string,
) (bool, error) {
	// new all pods need to create
	coreControl := gdcore.New(updateGD)
	newPods, err := coreControl.NewVersionedPods(currentGD, updateGD, currentRevision, updateRevision,
		expectedCreations, expectedCurrentCreations, availableIDs)
	if err != nil {
		return false, err
	}

	podsCreationChan := make(chan *v1.Pod, len(newPods))
	for _, p := range newPods {
		r.exp.ExpectScale(util.GetControllerKey(updateGD), expectations.Create, p.Name)
		podsCreationChan <- p
	}

	var created bool
	successPodNames := sync.Map{}
	_, err = util.DoItSlowly(len(newPods), initialBatchSize, func() error {
		pod := <-podsCreationChan

		cs := updateGD
		if pod.Labels[apps.ControllerRevisionHashLabelKey] == currentRevision {
			cs = currentGD
		}

		var createErr error
		if createErr = r.createOnePod(cs, pod); createErr != nil {
			return createErr
		}
		created = true
		successPodNames.Store(pod.Name, struct{}{})
		return nil
	})

	// rollback to ignore failure pods because the informer won't observe these pods
	for _, pod := range newPods {
		if _, ok := successPodNames.Load(pod.Name); !ok {
			r.exp.ObserveScale(util.GetControllerKey(updateGD), expectations.Create, pod.Name)
		}
	}

	return created, err
}

func (r *realControl) createOnePod(deploy *tkexv1alpha1.GameDeployment, pod *v1.Pod) error {
	if _, err := r.kubeClient.CoreV1().Pods(deploy.Namespace).Create(pod); err != nil {
		r.recorder.Eventf(deploy, v1.EventTypeWarning, "FailedCreate", "failed to create pod: %v, pod: %v", err, util.DumpJSON(pod))
		return err
	}

	r.recorder.Eventf(deploy, v1.EventTypeNormal, "SuccessfulCreate", "succeed to create pod %s", pod.Name)
	return nil
}

func (r *realControl) deletePods(deploy *tkexv1alpha1.GameDeployment, podsToDelete []*v1.Pod) (bool, error) {
	var deleted bool
	for _, pod := range podsToDelete {
		r.exp.ExpectScale(util.GetControllerKey(deploy), expectations.Delete, pod.Name)
		if err := r.kubeClient.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{}); err != nil {
			r.exp.ObserveScale(util.GetControllerKey(deploy), expectations.Delete, pod.Name)
			r.recorder.Eventf(deploy, v1.EventTypeWarning, "FailedDelete", "failed to delete pod %s: %v", pod.Name, err)
			return deleted, err
		}
		deleted = true
		r.recorder.Event(deploy, v1.EventTypeNormal, "SuccessfulDelete", fmt.Sprintf("succeed to delete pod %s", pod.Name))
	}

	return deleted, nil
}
