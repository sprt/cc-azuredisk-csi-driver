/*
Copyright The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "sigs.k8s.io/azuredisk-csi-driver/pkg/apis/azuredisk/v1alpha2"
)

// FakeAzVolumes implements AzVolumeInterface
type FakeAzVolumes struct {
	Fake *FakeDiskV1alpha2
	ns   string
}

var azvolumesResource = schema.GroupVersionResource{Group: "disk.csi.azure.com", Version: "v1alpha2", Resource: "azvolumes"}

var azvolumesKind = schema.GroupVersionKind{Group: "disk.csi.azure.com", Version: "v1alpha2", Kind: "AzVolume"}

// Get takes name of the azVolume, and returns the corresponding azVolume object, and an error if there is any.
func (c *FakeAzVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.AzVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azvolumesResource, c.ns, name), &v1alpha2.AzVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AzVolume), err
}

// List takes label and field selectors, and returns the list of AzVolumes that match those selectors.
func (c *FakeAzVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.AzVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azvolumesResource, azvolumesKind, c.ns, opts), &v1alpha2.AzVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.AzVolumeList{ListMeta: obj.(*v1alpha2.AzVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha2.AzVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azVolumes.
func (c *FakeAzVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azvolumesResource, c.ns, opts))

}

// Create takes the representation of a azVolume and creates it.  Returns the server's representation of the azVolume, and an error, if there is any.
func (c *FakeAzVolumes) Create(ctx context.Context, azVolume *v1alpha2.AzVolume, opts v1.CreateOptions) (result *v1alpha2.AzVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azvolumesResource, c.ns, azVolume), &v1alpha2.AzVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AzVolume), err
}

// Update takes the representation of a azVolume and updates it. Returns the server's representation of the azVolume, and an error, if there is any.
func (c *FakeAzVolumes) Update(ctx context.Context, azVolume *v1alpha2.AzVolume, opts v1.UpdateOptions) (result *v1alpha2.AzVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azvolumesResource, c.ns, azVolume), &v1alpha2.AzVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AzVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzVolumes) UpdateStatus(ctx context.Context, azVolume *v1alpha2.AzVolume, opts v1.UpdateOptions) (*v1alpha2.AzVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azvolumesResource, "status", c.ns, azVolume), &v1alpha2.AzVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AzVolume), err
}

// Delete takes name of the azVolume and deletes it. Returns an error if one occurs.
func (c *FakeAzVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(azvolumesResource, c.ns, name, opts), &v1alpha2.AzVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azvolumesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.AzVolumeList{})
	return err
}

// Patch applies the patch and returns the patched azVolume.
func (c *FakeAzVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.AzVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azvolumesResource, c.ns, name, pt, data, subresources...), &v1alpha2.AzVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AzVolume), err
}
