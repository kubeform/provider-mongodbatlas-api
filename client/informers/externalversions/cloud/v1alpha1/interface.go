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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BackupSchedules returns a BackupScheduleInformer.
	BackupSchedules() BackupScheduleInformer
	// BackupSnapshots returns a BackupSnapshotInformer.
	BackupSnapshots() BackupSnapshotInformer
	// BackupSnapshotExportBuckets returns a BackupSnapshotExportBucketInformer.
	BackupSnapshotExportBuckets() BackupSnapshotExportBucketInformer
	// BackupSnapshotExportJobs returns a BackupSnapshotExportJobInformer.
	BackupSnapshotExportJobs() BackupSnapshotExportJobInformer
	// BackupSnapshotRestoreJobs returns a BackupSnapshotRestoreJobInformer.
	BackupSnapshotRestoreJobs() BackupSnapshotRestoreJobInformer
	// ProviderAccesses returns a ProviderAccessInformer.
	ProviderAccesses() ProviderAccessInformer
	// ProviderAccessAuthorizations returns a ProviderAccessAuthorizationInformer.
	ProviderAccessAuthorizations() ProviderAccessAuthorizationInformer
	// ProviderAccessSetups returns a ProviderAccessSetupInformer.
	ProviderAccessSetups() ProviderAccessSetupInformer
	// ProviderSnapshots returns a ProviderSnapshotInformer.
	ProviderSnapshots() ProviderSnapshotInformer
	// ProviderSnapshotBackupPolicies returns a ProviderSnapshotBackupPolicyInformer.
	ProviderSnapshotBackupPolicies() ProviderSnapshotBackupPolicyInformer
	// ProviderSnapshotRestoreJobs returns a ProviderSnapshotRestoreJobInformer.
	ProviderSnapshotRestoreJobs() ProviderSnapshotRestoreJobInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BackupSchedules returns a BackupScheduleInformer.
func (v *version) BackupSchedules() BackupScheduleInformer {
	return &backupScheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupSnapshots returns a BackupSnapshotInformer.
func (v *version) BackupSnapshots() BackupSnapshotInformer {
	return &backupSnapshotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupSnapshotExportBuckets returns a BackupSnapshotExportBucketInformer.
func (v *version) BackupSnapshotExportBuckets() BackupSnapshotExportBucketInformer {
	return &backupSnapshotExportBucketInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupSnapshotExportJobs returns a BackupSnapshotExportJobInformer.
func (v *version) BackupSnapshotExportJobs() BackupSnapshotExportJobInformer {
	return &backupSnapshotExportJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupSnapshotRestoreJobs returns a BackupSnapshotRestoreJobInformer.
func (v *version) BackupSnapshotRestoreJobs() BackupSnapshotRestoreJobInformer {
	return &backupSnapshotRestoreJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProviderAccesses returns a ProviderAccessInformer.
func (v *version) ProviderAccesses() ProviderAccessInformer {
	return &providerAccessInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProviderAccessAuthorizations returns a ProviderAccessAuthorizationInformer.
func (v *version) ProviderAccessAuthorizations() ProviderAccessAuthorizationInformer {
	return &providerAccessAuthorizationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProviderAccessSetups returns a ProviderAccessSetupInformer.
func (v *version) ProviderAccessSetups() ProviderAccessSetupInformer {
	return &providerAccessSetupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProviderSnapshots returns a ProviderSnapshotInformer.
func (v *version) ProviderSnapshots() ProviderSnapshotInformer {
	return &providerSnapshotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProviderSnapshotBackupPolicies returns a ProviderSnapshotBackupPolicyInformer.
func (v *version) ProviderSnapshotBackupPolicies() ProviderSnapshotBackupPolicyInformer {
	return &providerSnapshotBackupPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProviderSnapshotRestoreJobs returns a ProviderSnapshotRestoreJobInformer.
func (v *version) ProviderSnapshotRestoreJobs() ProviderSnapshotRestoreJobInformer {
	return &providerSnapshotRestoreJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
