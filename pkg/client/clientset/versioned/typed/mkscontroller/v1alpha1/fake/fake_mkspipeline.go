/*
Copyright 2020 The Knative Authors

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

	v1alpha1 "github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMksPipelines implements MksPipelineInterface
type FakeMksPipelines struct {
	Fake *FakeMkscontrollerV1alpha1
	ns   string
}

var mkspipelinesResource = schema.GroupVersionResource{Group: "mkscontroller.example.mks", Version: "v1alpha1", Resource: "mkspipelines"}

var mkspipelinesKind = schema.GroupVersionKind{Group: "mkscontroller.example.mks", Version: "v1alpha1", Kind: "MksPipeline"}

// Get takes name of the mksPipeline, and returns the corresponding mksPipeline object, and an error if there is any.
func (c *FakeMksPipelines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MksPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mkspipelinesResource, c.ns, name), &v1alpha1.MksPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MksPipeline), err
}

// List takes label and field selectors, and returns the list of MksPipelines that match those selectors.
func (c *FakeMksPipelines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MksPipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mkspipelinesResource, mkspipelinesKind, c.ns, opts), &v1alpha1.MksPipelineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MksPipelineList{ListMeta: obj.(*v1alpha1.MksPipelineList).ListMeta}
	for _, item := range obj.(*v1alpha1.MksPipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mksPipelines.
func (c *FakeMksPipelines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mkspipelinesResource, c.ns, opts))

}

// Create takes the representation of a mksPipeline and creates it.  Returns the server's representation of the mksPipeline, and an error, if there is any.
func (c *FakeMksPipelines) Create(ctx context.Context, mksPipeline *v1alpha1.MksPipeline, opts v1.CreateOptions) (result *v1alpha1.MksPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mkspipelinesResource, c.ns, mksPipeline), &v1alpha1.MksPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MksPipeline), err
}

// Update takes the representation of a mksPipeline and updates it. Returns the server's representation of the mksPipeline, and an error, if there is any.
func (c *FakeMksPipelines) Update(ctx context.Context, mksPipeline *v1alpha1.MksPipeline, opts v1.UpdateOptions) (result *v1alpha1.MksPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mkspipelinesResource, c.ns, mksPipeline), &v1alpha1.MksPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MksPipeline), err
}

// Delete takes name of the mksPipeline and deletes it. Returns an error if one occurs.
func (c *FakeMksPipelines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(mkspipelinesResource, c.ns, name, opts), &v1alpha1.MksPipeline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMksPipelines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mkspipelinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MksPipelineList{})
	return err
}

// Patch applies the patch and returns the patched mksPipeline.
func (c *FakeMksPipelines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MksPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mkspipelinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MksPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MksPipeline), err
}
