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
	"context"

	v1alpha1 "github.com/kubernetes-sigs/container-object-storage-interface-api/apis/objectstorage.k8s.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBucketClasses implements BucketClassInterface
type FakeBucketClasses struct {
	Fake *FakeObjectstorageV1alpha1
}

var bucketclassesResource = schema.GroupVersionResource{Group: "objectstorage.k8s.io", Version: "v1alpha1", Resource: "bucketclasses"}

var bucketclassesKind = schema.GroupVersionKind{Group: "objectstorage.k8s.io", Version: "v1alpha1", Kind: "BucketClass"}

// Get takes name of the bucketClass, and returns the corresponding bucketClass object, and an error if there is any.
func (c *FakeBucketClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BucketClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(bucketclassesResource, name), &v1alpha1.BucketClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketClass), err
}

// List takes label and field selectors, and returns the list of BucketClasses that match those selectors.
func (c *FakeBucketClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BucketClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(bucketclassesResource, bucketclassesKind, opts), &v1alpha1.BucketClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BucketClassList{ListMeta: obj.(*v1alpha1.BucketClassList).ListMeta}
	for _, item := range obj.(*v1alpha1.BucketClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bucketClasses.
func (c *FakeBucketClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(bucketclassesResource, opts))
}

// Create takes the representation of a bucketClass and creates it.  Returns the server's representation of the bucketClass, and an error, if there is any.
func (c *FakeBucketClasses) Create(ctx context.Context, bucketClass *v1alpha1.BucketClass, opts v1.CreateOptions) (result *v1alpha1.BucketClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(bucketclassesResource, bucketClass), &v1alpha1.BucketClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketClass), err
}

// Update takes the representation of a bucketClass and updates it. Returns the server's representation of the bucketClass, and an error, if there is any.
func (c *FakeBucketClasses) Update(ctx context.Context, bucketClass *v1alpha1.BucketClass, opts v1.UpdateOptions) (result *v1alpha1.BucketClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(bucketclassesResource, bucketClass), &v1alpha1.BucketClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketClass), err
}

// Delete takes name of the bucketClass and deletes it. Returns an error if one occurs.
func (c *FakeBucketClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(bucketclassesResource, name), &v1alpha1.BucketClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBucketClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(bucketclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BucketClassList{})
	return err
}

// Patch applies the patch and returns the patched bucketClass.
func (c *FakeBucketClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BucketClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(bucketclassesResource, name, pt, data, subresources...), &v1alpha1.BucketClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketClass), err
}
