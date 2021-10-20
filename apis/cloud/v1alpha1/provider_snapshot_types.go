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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ProviderSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProviderSnapshotSpec   `json:"spec,omitempty"`
	Status            ProviderSnapshotStatus `json:"status,omitempty"`
}

type ProviderSnapshotSpec struct {
	State *ProviderSnapshotSpecResource `json:"state,omitempty" tf:"-"`

	Resource ProviderSnapshotSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ProviderSnapshotSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterName *string `json:"clusterName" tf:"cluster_name"`
	// +optional
	CreatedAt   *string `json:"createdAt,omitempty" tf:"created_at"`
	Description *string `json:"description" tf:"description"`
	// +optional
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at"`
	// +optional
	MasterKeyUUID *string `json:"masterKeyUUID,omitempty" tf:"master_key_uuid"`
	// +optional
	MongodVersion   *string `json:"mongodVersion,omitempty" tf:"mongod_version"`
	ProjectID       *string `json:"projectID" tf:"project_id"`
	RetentionInDays *int64  `json:"retentionInDays" tf:"retention_in_days"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
	// +optional
	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StorageSizeBytes *int64 `json:"storageSizeBytes,omitempty" tf:"storage_size_bytes"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ProviderSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProviderSnapshotList is a list of ProviderSnapshots
type ProviderSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProviderSnapshot CRD objects
	Items []ProviderSnapshot `json:"items,omitempty"`
}
