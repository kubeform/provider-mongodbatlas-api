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

// FakeBackupSnapshotExportJobs implements BackupSnapshotExportJobInterface
type FakeBackupSnapshotExportJobs struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var backupsnapshotexportjobsResource = schema.GroupVersionResource{Group: "cloud.mongodbatlas.kubeform.com", Version: "v1alpha1", Resource: "backupsnapshotexportjobs"}

var backupsnapshotexportjobsKind = schema.GroupVersionKind{Group: "cloud.mongodbatlas.kubeform.com", Version: "v1alpha1", Kind: "BackupSnapshotExportJob"}

// Get takes name of the backupSnapshotExportJob, and returns the corresponding backupSnapshotExportJob object, and an error if there is any.
func (c *FakeBackupSnapshotExportJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupSnapshotExportJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupsnapshotexportjobsResource, c.ns, name), &v1alpha1.BackupSnapshotExportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportJob), err
}

// List takes label and field selectors, and returns the list of BackupSnapshotExportJobs that match those selectors.
func (c *FakeBackupSnapshotExportJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupSnapshotExportJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupsnapshotexportjobsResource, backupsnapshotexportjobsKind, c.ns, opts), &v1alpha1.BackupSnapshotExportJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupSnapshotExportJobList{ListMeta: obj.(*v1alpha1.BackupSnapshotExportJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupSnapshotExportJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupSnapshotExportJobs.
func (c *FakeBackupSnapshotExportJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupsnapshotexportjobsResource, c.ns, opts))

}

// Create takes the representation of a backupSnapshotExportJob and creates it.  Returns the server's representation of the backupSnapshotExportJob, and an error, if there is any.
func (c *FakeBackupSnapshotExportJobs) Create(ctx context.Context, backupSnapshotExportJob *v1alpha1.BackupSnapshotExportJob, opts v1.CreateOptions) (result *v1alpha1.BackupSnapshotExportJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupsnapshotexportjobsResource, c.ns, backupSnapshotExportJob), &v1alpha1.BackupSnapshotExportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportJob), err
}

// Update takes the representation of a backupSnapshotExportJob and updates it. Returns the server's representation of the backupSnapshotExportJob, and an error, if there is any.
func (c *FakeBackupSnapshotExportJobs) Update(ctx context.Context, backupSnapshotExportJob *v1alpha1.BackupSnapshotExportJob, opts v1.UpdateOptions) (result *v1alpha1.BackupSnapshotExportJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupsnapshotexportjobsResource, c.ns, backupSnapshotExportJob), &v1alpha1.BackupSnapshotExportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupSnapshotExportJobs) UpdateStatus(ctx context.Context, backupSnapshotExportJob *v1alpha1.BackupSnapshotExportJob, opts v1.UpdateOptions) (*v1alpha1.BackupSnapshotExportJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupsnapshotexportjobsResource, "status", c.ns, backupSnapshotExportJob), &v1alpha1.BackupSnapshotExportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportJob), err
}

// Delete takes name of the backupSnapshotExportJob and deletes it. Returns an error if one occurs.
func (c *FakeBackupSnapshotExportJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupsnapshotexportjobsResource, c.ns, name), &v1alpha1.BackupSnapshotExportJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupSnapshotExportJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupsnapshotexportjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupSnapshotExportJobList{})
	return err
}

// Patch applies the patch and returns the patched backupSnapshotExportJob.
func (c *FakeBackupSnapshotExportJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupSnapshotExportJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupsnapshotexportjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackupSnapshotExportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSnapshotExportJob), err
}