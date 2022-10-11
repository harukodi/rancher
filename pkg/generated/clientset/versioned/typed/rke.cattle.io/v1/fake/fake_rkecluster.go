/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	rkecattleiov1 "github.com/rancher/rancher/pkg/apis/rke.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRKEClusters implements RKEClusterInterface
type FakeRKEClusters struct {
	Fake *FakeRkeV1
	ns   string
}

var rkeclustersResource = schema.GroupVersionResource{Group: "rke.cattle.io", Version: "v1", Resource: "rkeclusters"}

var rkeclustersKind = schema.GroupVersionKind{Group: "rke.cattle.io", Version: "v1", Kind: "RKECluster"}

// Get takes name of the rKECluster, and returns the corresponding rKECluster object, and an error if there is any.
func (c *FakeRKEClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *rkecattleiov1.RKECluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rkeclustersResource, c.ns, name), &rkecattleiov1.RKECluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rkecattleiov1.RKECluster), err
}

// List takes label and field selectors, and returns the list of RKEClusters that match those selectors.
func (c *FakeRKEClusters) List(ctx context.Context, opts v1.ListOptions) (result *rkecattleiov1.RKEClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rkeclustersResource, rkeclustersKind, c.ns, opts), &rkecattleiov1.RKEClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rkecattleiov1.RKEClusterList{ListMeta: obj.(*rkecattleiov1.RKEClusterList).ListMeta}
	for _, item := range obj.(*rkecattleiov1.RKEClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rKEClusters.
func (c *FakeRKEClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rkeclustersResource, c.ns, opts))

}

// Create takes the representation of a rKECluster and creates it.  Returns the server's representation of the rKECluster, and an error, if there is any.
func (c *FakeRKEClusters) Create(ctx context.Context, rKECluster *rkecattleiov1.RKECluster, opts v1.CreateOptions) (result *rkecattleiov1.RKECluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rkeclustersResource, c.ns, rKECluster), &rkecattleiov1.RKECluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rkecattleiov1.RKECluster), err
}

// Update takes the representation of a rKECluster and updates it. Returns the server's representation of the rKECluster, and an error, if there is any.
func (c *FakeRKEClusters) Update(ctx context.Context, rKECluster *rkecattleiov1.RKECluster, opts v1.UpdateOptions) (result *rkecattleiov1.RKECluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rkeclustersResource, c.ns, rKECluster), &rkecattleiov1.RKECluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rkecattleiov1.RKECluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRKEClusters) UpdateStatus(ctx context.Context, rKECluster *rkecattleiov1.RKECluster, opts v1.UpdateOptions) (*rkecattleiov1.RKECluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rkeclustersResource, "status", c.ns, rKECluster), &rkecattleiov1.RKECluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rkecattleiov1.RKECluster), err
}

// Delete takes name of the rKECluster and deletes it. Returns an error if one occurs.
func (c *FakeRKEClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rkeclustersResource, c.ns, name, opts), &rkecattleiov1.RKECluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRKEClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rkeclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &rkecattleiov1.RKEClusterList{})
	return err
}

// Patch applies the patch and returns the patched rKECluster.
func (c *FakeRKEClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *rkecattleiov1.RKECluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rkeclustersResource, c.ns, name, pt, data, subresources...), &rkecattleiov1.RKECluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rkecattleiov1.RKECluster), err
}
