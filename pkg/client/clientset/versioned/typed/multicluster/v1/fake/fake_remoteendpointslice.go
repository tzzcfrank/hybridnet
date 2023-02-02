/*
Copyright 2021 The Hybridnet Authors.

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
	"context"

	multiclusterv1 "github.com/alibaba/hybridnet/pkg/apis/multicluster/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRemoteEndpointSlices implements RemoteEndpointSliceInterface
type FakeRemoteEndpointSlices struct {
	Fake *FakeMulticlusterV1
}

var remoteendpointslicesResource = schema.GroupVersionResource{Group: "multicluster", Version: "v1", Resource: "remoteendpointslices"}

var remoteendpointslicesKind = schema.GroupVersionKind{Group: "multicluster", Version: "v1", Kind: "RemoteEndpointSlice"}

// Get takes name of the remoteEndpointSlice, and returns the corresponding remoteEndpointSlice object, and an error if there is any.
func (c *FakeRemoteEndpointSlices) Get(ctx context.Context, name string, options v1.GetOptions) (result *multiclusterv1.RemoteEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(remoteendpointslicesResource, name), &multiclusterv1.RemoteEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterv1.RemoteEndpointSlice), err
}

// List takes label and field selectors, and returns the list of RemoteEndpointSlices that match those selectors.
func (c *FakeRemoteEndpointSlices) List(ctx context.Context, opts v1.ListOptions) (result *multiclusterv1.RemoteEndpointSliceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(remoteendpointslicesResource, remoteendpointslicesKind, opts), &multiclusterv1.RemoteEndpointSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &multiclusterv1.RemoteEndpointSliceList{ListMeta: obj.(*multiclusterv1.RemoteEndpointSliceList).ListMeta}
	for _, item := range obj.(*multiclusterv1.RemoteEndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested remoteEndpointSlices.
func (c *FakeRemoteEndpointSlices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(remoteendpointslicesResource, opts))
}

// Create takes the representation of a remoteEndpointSlice and creates it.  Returns the server's representation of the remoteEndpointSlice, and an error, if there is any.
func (c *FakeRemoteEndpointSlices) Create(ctx context.Context, remoteEndpointSlice *multiclusterv1.RemoteEndpointSlice, opts v1.CreateOptions) (result *multiclusterv1.RemoteEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(remoteendpointslicesResource, remoteEndpointSlice), &multiclusterv1.RemoteEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterv1.RemoteEndpointSlice), err
}

// Update takes the representation of a remoteEndpointSlice and updates it. Returns the server's representation of the remoteEndpointSlice, and an error, if there is any.
func (c *FakeRemoteEndpointSlices) Update(ctx context.Context, remoteEndpointSlice *multiclusterv1.RemoteEndpointSlice, opts v1.UpdateOptions) (result *multiclusterv1.RemoteEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(remoteendpointslicesResource, remoteEndpointSlice), &multiclusterv1.RemoteEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterv1.RemoteEndpointSlice), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRemoteEndpointSlices) UpdateStatus(ctx context.Context, remoteEndpointSlice *multiclusterv1.RemoteEndpointSlice, opts v1.UpdateOptions) (*multiclusterv1.RemoteEndpointSlice, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(remoteendpointslicesResource, "status", remoteEndpointSlice), &multiclusterv1.RemoteEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterv1.RemoteEndpointSlice), err
}

// Delete takes name of the remoteEndpointSlice and deletes it. Returns an error if one occurs.
func (c *FakeRemoteEndpointSlices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(remoteendpointslicesResource, name), &multiclusterv1.RemoteEndpointSlice{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRemoteEndpointSlices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(remoteendpointslicesResource, listOpts)

	_, err := c.Fake.Invokes(action, &multiclusterv1.RemoteEndpointSliceList{})
	return err
}

// Patch applies the patch and returns the patched remoteEndpointSlice.
func (c *FakeRemoteEndpointSlices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *multiclusterv1.RemoteEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(remoteendpointslicesResource, name, pt, data, subresources...), &multiclusterv1.RemoteEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterv1.RemoteEndpointSlice), err
}
