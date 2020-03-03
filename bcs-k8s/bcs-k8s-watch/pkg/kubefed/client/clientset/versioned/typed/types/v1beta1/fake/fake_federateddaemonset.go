/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedDaemonSets implements FederatedDaemonSetInterface
type FakeFederatedDaemonSets struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federateddaemonsetsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federateddaemonsets"}

var federateddaemonsetsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedDaemonSet"}

// Get takes name of the federatedDaemonSet, and returns the corresponding federatedDaemonSet object, and an error if there is any.
func (c *FakeFederatedDaemonSets) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federateddaemonsetsResource, c.ns, name), &v1beta1.FederatedDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDaemonSet), err
}

// List takes label and field selectors, and returns the list of FederatedDaemonSets that match those selectors.
func (c *FakeFederatedDaemonSets) List(opts v1.ListOptions) (result *v1beta1.FederatedDaemonSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federateddaemonsetsResource, federateddaemonsetsKind, c.ns, opts), &v1beta1.FederatedDaemonSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedDaemonSetList{}
	for _, item := range obj.(*v1beta1.FederatedDaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedDaemonSets.
func (c *FakeFederatedDaemonSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federateddaemonsetsResource, c.ns, opts))

}

// Create takes the representation of a federatedDaemonSet and creates it.  Returns the server's representation of the federatedDaemonSet, and an error, if there is any.
func (c *FakeFederatedDaemonSets) Create(federatedDaemonSet *v1beta1.FederatedDaemonSet) (result *v1beta1.FederatedDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federateddaemonsetsResource, c.ns, federatedDaemonSet), &v1beta1.FederatedDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDaemonSet), err
}

// Update takes the representation of a federatedDaemonSet and updates it. Returns the server's representation of the federatedDaemonSet, and an error, if there is any.
func (c *FakeFederatedDaemonSets) Update(federatedDaemonSet *v1beta1.FederatedDaemonSet) (result *v1beta1.FederatedDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federateddaemonsetsResource, c.ns, federatedDaemonSet), &v1beta1.FederatedDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDaemonSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedDaemonSets) UpdateStatus(federatedDaemonSet *v1beta1.FederatedDaemonSet) (*v1beta1.FederatedDaemonSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federateddaemonsetsResource, "status", c.ns, federatedDaemonSet), &v1beta1.FederatedDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDaemonSet), err
}

// Delete takes name of the federatedDaemonSet and deletes it. Returns an error if one occurs.
func (c *FakeFederatedDaemonSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federateddaemonsetsResource, c.ns, name), &v1beta1.FederatedDaemonSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedDaemonSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federateddaemonsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedDaemonSetList{})
	return err
}

// Patch applies the patch and returns the patched federatedDaemonSet.
func (c *FakeFederatedDaemonSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federateddaemonsetsResource, c.ns, name, data, subresources...), &v1beta1.FederatedDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDaemonSet), err
}
