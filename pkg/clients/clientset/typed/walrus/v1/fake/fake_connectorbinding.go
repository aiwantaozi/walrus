// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walrusv1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/walrus/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConnectorBindings implements ConnectorBindingInterface
type FakeConnectorBindings struct {
	Fake *FakeWalrusV1
	ns   string
}

var connectorbindingsResource = v1.SchemeGroupVersion.WithResource("connectorbindings")

var connectorbindingsKind = v1.SchemeGroupVersion.WithKind("ConnectorBinding")

// Get takes name of the connectorBinding, and returns the corresponding connectorBinding object, and an error if there is any.
func (c *FakeConnectorBindings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ConnectorBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(connectorbindingsResource, c.ns, name), &v1.ConnectorBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConnectorBinding), err
}

// List takes label and field selectors, and returns the list of ConnectorBindings that match those selectors.
func (c *FakeConnectorBindings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConnectorBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(connectorbindingsResource, connectorbindingsKind, c.ns, opts), &v1.ConnectorBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ConnectorBindingList{ListMeta: obj.(*v1.ConnectorBindingList).ListMeta}
	for _, item := range obj.(*v1.ConnectorBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested connectorBindings.
func (c *FakeConnectorBindings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(connectorbindingsResource, c.ns, opts))

}

// Create takes the representation of a connectorBinding and creates it.  Returns the server's representation of the connectorBinding, and an error, if there is any.
func (c *FakeConnectorBindings) Create(ctx context.Context, connectorBinding *v1.ConnectorBinding, opts metav1.CreateOptions) (result *v1.ConnectorBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(connectorbindingsResource, c.ns, connectorBinding), &v1.ConnectorBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConnectorBinding), err
}

// Update takes the representation of a connectorBinding and updates it. Returns the server's representation of the connectorBinding, and an error, if there is any.
func (c *FakeConnectorBindings) Update(ctx context.Context, connectorBinding *v1.ConnectorBinding, opts metav1.UpdateOptions) (result *v1.ConnectorBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(connectorbindingsResource, c.ns, connectorBinding), &v1.ConnectorBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConnectorBinding), err
}

// Delete takes name of the connectorBinding and deletes it. Returns an error if one occurs.
func (c *FakeConnectorBindings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(connectorbindingsResource, c.ns, name, opts), &v1.ConnectorBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConnectorBindings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(connectorbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ConnectorBindingList{})
	return err
}

// Patch applies the patch and returns the patched connectorBinding.
func (c *FakeConnectorBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConnectorBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(connectorbindingsResource, c.ns, name, pt, data, subresources...), &v1.ConnectorBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConnectorBinding), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied connectorBinding.
func (c *FakeConnectorBindings) Apply(ctx context.Context, connectorBinding *walrusv1.ConnectorBindingApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ConnectorBinding, err error) {
	if connectorBinding == nil {
		return nil, fmt.Errorf("connectorBinding provided to Apply must not be nil")
	}
	data, err := json.Marshal(connectorBinding)
	if err != nil {
		return nil, err
	}
	name := connectorBinding.Name
	if name == nil {
		return nil, fmt.Errorf("connectorBinding.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(connectorbindingsResource, c.ns, *name, types.ApplyPatchType, data), &v1.ConnectorBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConnectorBinding), err
}
