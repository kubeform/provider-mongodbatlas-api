/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cloud/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupSnapshotExportBuckets implements BackupSnapshotExportBucketInterface
type FakeBackupSnapshotExportBuckets struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var backupsnapshotexportbucketsResource = schema.GroupVersionResource{Group: "cloud.mongodbatlas.kubeform.com", Version: "v1alpha1", Resource: "backupsnapshotexportbuckets"}

var backupsnapshotexportbucketsKind = schema.GroupVersionKind{Group: "cloud.mongodbatlas.kubeform.com", Version: "v1alpha1", Kind: "BackupSnapshotExportBucket"}

// Get takes name of the backupSnapshotExportBucket, and returns the corresponding backupSnapshotExportBucket object, and an error if there is any.
func (c *FakeBackupSnapshotExportBuckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupSnapshotExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupsnapshotexportbucketsResource, c.ns, name), &v1alpha1.BackupSnapshotExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportBucket), err
}

// List takes label and field selectors, and returns the list of BackupSnapshotExportBuckets that match those selectors.
func (c *FakeBackupSnapshotExportBuckets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupSnapshotExportBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupsnapshotexportbucketsResource, backupsnapshotexportbucketsKind, c.ns, opts), &v1alpha1.BackupSnapshotExportBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupSnapshotExportBucketList{ListMeta: obj.(*v1alpha1.BackupSnapshotExportBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupSnapshotExportBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupSnapshotExportBuckets.
func (c *FakeBackupSnapshotExportBuckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupsnapshotexportbucketsResource, c.ns, opts))

}

// Create takes the representation of a backupSnapshotExportBucket and creates it.  Returns the server's representation of the backupSnapshotExportBucket, and an error, if there is any.
func (c *FakeBackupSnapshotExportBuckets) Create(ctx context.Context, backupSnapshotExportBucket *v1alpha1.BackupSnapshotExportBucket, opts v1.CreateOptions) (result *v1alpha1.BackupSnapshotExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupsnapshotexportbucketsResource, c.ns, backupSnapshotExportBucket), &v1alpha1.BackupSnapshotExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportBucket), err
}

// Update takes the representation of a backupSnapshotExportBucket and updates it. Returns the server's representation of the backupSnapshotExportBucket, and an error, if there is any.
func (c *FakeBackupSnapshotExportBuckets) Update(ctx context.Context, backupSnapshotExportBucket *v1alpha1.BackupSnapshotExportBucket, opts v1.UpdateOptions) (result *v1alpha1.BackupSnapshotExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupsnapshotexportbucketsResource, c.ns, backupSnapshotExportBucket), &v1alpha1.BackupSnapshotExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupSnapshotExportBuckets) UpdateStatus(ctx context.Context, backupSnapshotExportBucket *v1alpha1.BackupSnapshotExportBucket, opts v1.UpdateOptions) (*v1alpha1.BackupSnapshotExportBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupsnapshotexportbucketsResource, "status", c.ns, backupSnapshotExportBucket), &v1alpha1.BackupSnapshotExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportBucket), err
}

// Delete takes name of the backupSnapshotExportBucket and deletes it. Returns an error if one occurs.
func (c *FakeBackupSnapshotExportBuckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupsnapshotexportbucketsResource, c.ns, name), &v1alpha1.BackupSnapshotExportBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupSnapshotExportBuckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupsnapshotexportbucketsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupSnapshotExportBucketList{})
	return err
}

// Patch applies the patch and returns the patched backupSnapshotExportBucket.
func (c *FakeBackupSnapshotExportBuckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupSnapshotExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupsnapshotexportbucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackupSnapshotExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportBucket), err
}
