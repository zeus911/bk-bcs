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

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"
	scheme "github.com/Tencent/bk-bcs/bcs-mesos/pkg/client/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BcsServicesGetter has a method to return a BcsServiceInterface.
// A group's client should implement this interface.
type BcsServicesGetter interface {
	BcsServices(namespace string) BcsServiceInterface
}

// BcsServiceInterface has methods to work with BcsService resources.
type BcsServiceInterface interface {
	Create(*v2.BcsService) (*v2.BcsService, error)
	Update(*v2.BcsService) (*v2.BcsService, error)
	UpdateStatus(*v2.BcsService) (*v2.BcsService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.BcsService, error)
	List(opts v1.ListOptions) (*v2.BcsServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.BcsService, err error)
	BcsServiceExpansion
}

// bcsServices implements BcsServiceInterface
type bcsServices struct {
	client rest.Interface
	ns     string
}

// newBcsServices returns a BcsServices
func newBcsServices(c *BkbcsV2Client, namespace string) *bcsServices {
	return &bcsServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bcsService, and returns the corresponding bcsService object, and an error if there is any.
func (c *bcsServices) Get(name string, options v1.GetOptions) (result *v2.BcsService, err error) {
	result = &v2.BcsService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BcsServices that match those selectors.
func (c *bcsServices) List(opts v1.ListOptions) (result *v2.BcsServiceList, err error) {
	result = &v2.BcsServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bcsServices.
func (c *bcsServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bcsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a bcsService and creates it.  Returns the server's representation of the bcsService, and an error, if there is any.
func (c *bcsServices) Create(bcsService *v2.BcsService) (result *v2.BcsService, err error) {
	result = &v2.BcsService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bcsservices").
		Body(bcsService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bcsService and updates it. Returns the server's representation of the bcsService, and an error, if there is any.
func (c *bcsServices) Update(bcsService *v2.BcsService) (result *v2.BcsService, err error) {
	result = &v2.BcsService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsservices").
		Name(bcsService.Name).
		Body(bcsService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bcsServices) UpdateStatus(bcsService *v2.BcsService) (result *v2.BcsService, err error) {
	result = &v2.BcsService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsservices").
		Name(bcsService.Name).
		SubResource("status").
		Body(bcsService).
		Do().
		Into(result)
	return
}

// Delete takes name of the bcsService and deletes it. Returns an error if one occurs.
func (c *bcsServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bcsServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bcsService.
func (c *bcsServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.BcsService, err error) {
	result = &v2.BcsService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bcsservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
