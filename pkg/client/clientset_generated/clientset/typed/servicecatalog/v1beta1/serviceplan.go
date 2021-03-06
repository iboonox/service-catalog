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

package v1beta1

import (
	"time"

	v1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	scheme "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServicePlansGetter has a method to return a ServicePlanInterface.
// A group's client should implement this interface.
type ServicePlansGetter interface {
	ServicePlans(namespace string) ServicePlanInterface
}

// ServicePlanInterface has methods to work with ServicePlan resources.
type ServicePlanInterface interface {
	Create(*v1beta1.ServicePlan) (*v1beta1.ServicePlan, error)
	Update(*v1beta1.ServicePlan) (*v1beta1.ServicePlan, error)
	UpdateStatus(*v1beta1.ServicePlan) (*v1beta1.ServicePlan, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ServicePlan, error)
	List(opts v1.ListOptions) (*v1beta1.ServicePlanList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServicePlan, err error)
	ServicePlanExpansion
}

// servicePlans implements ServicePlanInterface
type servicePlans struct {
	client rest.Interface
	ns     string
}

// newServicePlans returns a ServicePlans
func newServicePlans(c *ServicecatalogV1beta1Client, namespace string) *servicePlans {
	return &servicePlans{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the servicePlan, and returns the corresponding servicePlan object, and an error if there is any.
func (c *servicePlans) Get(name string, options v1.GetOptions) (result *v1beta1.ServicePlan, err error) {
	result = &v1beta1.ServicePlan{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServicePlans that match those selectors.
func (c *servicePlans) List(opts v1.ListOptions) (result *v1beta1.ServicePlanList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ServicePlanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested servicePlans.
func (c *servicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a servicePlan and creates it.  Returns the server's representation of the servicePlan, and an error, if there is any.
func (c *servicePlans) Create(servicePlan *v1beta1.ServicePlan) (result *v1beta1.ServicePlan, err error) {
	result = &v1beta1.ServicePlan{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceplans").
		Body(servicePlan).
		Do().
		Into(result)
	return
}

// Update takes the representation of a servicePlan and updates it. Returns the server's representation of the servicePlan, and an error, if there is any.
func (c *servicePlans) Update(servicePlan *v1beta1.ServicePlan) (result *v1beta1.ServicePlan, err error) {
	result = &v1beta1.ServicePlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceplans").
		Name(servicePlan.Name).
		Body(servicePlan).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *servicePlans) UpdateStatus(servicePlan *v1beta1.ServicePlan) (result *v1beta1.ServicePlan, err error) {
	result = &v1beta1.ServicePlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceplans").
		Name(servicePlan.Name).
		SubResource("status").
		Body(servicePlan).
		Do().
		Into(result)
	return
}

// Delete takes name of the servicePlan and deletes it. Returns an error if one occurs.
func (c *servicePlans) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceplans").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *servicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceplans").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched servicePlan.
func (c *servicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServicePlan, err error) {
	result = &v1beta1.ServicePlan{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceplans").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
