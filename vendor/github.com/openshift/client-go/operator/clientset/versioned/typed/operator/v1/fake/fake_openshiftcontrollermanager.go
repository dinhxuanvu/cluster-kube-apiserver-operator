// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpenShiftControllerManagers implements OpenShiftControllerManagerInterface
type FakeOpenShiftControllerManagers struct {
	Fake *FakeOperatorV1
}

var openshiftcontrollermanagersResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "openshiftcontrollermanagers"}

var openshiftcontrollermanagersKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "OpenShiftControllerManager"}

// Get takes name of the openShiftControllerManager, and returns the corresponding openShiftControllerManager object, and an error if there is any.
func (c *FakeOpenShiftControllerManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.OpenShiftControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(openshiftcontrollermanagersResource, name), &operatorv1.OpenShiftControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftControllerManager), err
}

// List takes label and field selectors, and returns the list of OpenShiftControllerManagers that match those selectors.
func (c *FakeOpenShiftControllerManagers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.OpenShiftControllerManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(openshiftcontrollermanagersResource, openshiftcontrollermanagersKind, opts), &operatorv1.OpenShiftControllerManagerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.OpenShiftControllerManagerList{ListMeta: obj.(*operatorv1.OpenShiftControllerManagerList).ListMeta}
	for _, item := range obj.(*operatorv1.OpenShiftControllerManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openShiftControllerManagers.
func (c *FakeOpenShiftControllerManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(openshiftcontrollermanagersResource, opts))
}

// Create takes the representation of a openShiftControllerManager and creates it.  Returns the server's representation of the openShiftControllerManager, and an error, if there is any.
func (c *FakeOpenShiftControllerManagers) Create(ctx context.Context, openShiftControllerManager *operatorv1.OpenShiftControllerManager, opts v1.CreateOptions) (result *operatorv1.OpenShiftControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(openshiftcontrollermanagersResource, openShiftControllerManager), &operatorv1.OpenShiftControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftControllerManager), err
}

// Update takes the representation of a openShiftControllerManager and updates it. Returns the server's representation of the openShiftControllerManager, and an error, if there is any.
func (c *FakeOpenShiftControllerManagers) Update(ctx context.Context, openShiftControllerManager *operatorv1.OpenShiftControllerManager, opts v1.UpdateOptions) (result *operatorv1.OpenShiftControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(openshiftcontrollermanagersResource, openShiftControllerManager), &operatorv1.OpenShiftControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftControllerManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpenShiftControllerManagers) UpdateStatus(ctx context.Context, openShiftControllerManager *operatorv1.OpenShiftControllerManager, opts v1.UpdateOptions) (*operatorv1.OpenShiftControllerManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(openshiftcontrollermanagersResource, "status", openShiftControllerManager), &operatorv1.OpenShiftControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftControllerManager), err
}

// Delete takes name of the openShiftControllerManager and deletes it. Returns an error if one occurs.
func (c *FakeOpenShiftControllerManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(openshiftcontrollermanagersResource, name, opts), &operatorv1.OpenShiftControllerManager{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenShiftControllerManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(openshiftcontrollermanagersResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.OpenShiftControllerManagerList{})
	return err
}

// Patch applies the patch and returns the patched openShiftControllerManager.
func (c *FakeOpenShiftControllerManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.OpenShiftControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(openshiftcontrollermanagersResource, name, pt, data, subresources...), &operatorv1.OpenShiftControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftControllerManager), err
}
