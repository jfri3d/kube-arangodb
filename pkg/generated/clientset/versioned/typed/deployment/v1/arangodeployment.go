//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoDeploymentsGetter has a method to return a ArangoDeploymentInterface.
// A group's client should implement this interface.
type ArangoDeploymentsGetter interface {
	ArangoDeployments(namespace string) ArangoDeploymentInterface
}

// ArangoDeploymentInterface has methods to work with ArangoDeployment resources.
type ArangoDeploymentInterface interface {
	Create(*v1.ArangoDeployment) (*v1.ArangoDeployment, error)
	Update(*v1.ArangoDeployment) (*v1.ArangoDeployment, error)
	UpdateStatus(*v1.ArangoDeployment) (*v1.ArangoDeployment, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ArangoDeployment, error)
	List(opts metav1.ListOptions) (*v1.ArangoDeploymentList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ArangoDeployment, err error)
	ArangoDeploymentExpansion
}

// arangoDeployments implements ArangoDeploymentInterface
type arangoDeployments struct {
	client rest.Interface
	ns     string
}

// newArangoDeployments returns a ArangoDeployments
func newArangoDeployments(c *DatabaseV1Client, namespace string) *arangoDeployments {
	return &arangoDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the arangoDeployment, and returns the corresponding arangoDeployment object, and an error if there is any.
func (c *arangoDeployments) Get(name string, options metav1.GetOptions) (result *v1.ArangoDeployment, err error) {
	result = &v1.ArangoDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangodeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoDeployments that match those selectors.
func (c *arangoDeployments) List(opts metav1.ListOptions) (result *v1.ArangoDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ArangoDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangodeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoDeployments.
func (c *arangoDeployments) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("arangodeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a arangoDeployment and creates it.  Returns the server's representation of the arangoDeployment, and an error, if there is any.
func (c *arangoDeployments) Create(arangoDeployment *v1.ArangoDeployment) (result *v1.ArangoDeployment, err error) {
	result = &v1.ArangoDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("arangodeployments").
		Body(arangoDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a arangoDeployment and updates it. Returns the server's representation of the arangoDeployment, and an error, if there is any.
func (c *arangoDeployments) Update(arangoDeployment *v1.ArangoDeployment) (result *v1.ArangoDeployment, err error) {
	result = &v1.ArangoDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangodeployments").
		Name(arangoDeployment.Name).
		Body(arangoDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *arangoDeployments) UpdateStatus(arangoDeployment *v1.ArangoDeployment) (result *v1.ArangoDeployment, err error) {
	result = &v1.ArangoDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangodeployments").
		Name(arangoDeployment.Name).
		SubResource("status").
		Body(arangoDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the arangoDeployment and deletes it. Returns an error if one occurs.
func (c *arangoDeployments) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangodeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoDeployments) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangodeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched arangoDeployment.
func (c *arangoDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ArangoDeployment, err error) {
	result = &v1.ArangoDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("arangodeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
