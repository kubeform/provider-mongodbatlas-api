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

type ProviderSnapshotBackupPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProviderSnapshotBackupPolicySpec   `json:"spec,omitempty"`
	Status            ProviderSnapshotBackupPolicyStatus `json:"status,omitempty"`
}

type ProviderSnapshotBackupPolicySpecPoliciesPolicyItem struct {
	FrequencyInterval *int64  `json:"frequencyInterval" tf:"frequency_interval"`
	FrequencyType     *string `json:"frequencyType" tf:"frequency_type"`
	ID                *string `json:"ID" tf:"id"`
	RetentionUnit     *string `json:"retentionUnit" tf:"retention_unit"`
	RetentionValue    *int64  `json:"retentionValue" tf:"retention_value"`
}

type ProviderSnapshotBackupPolicySpecPolicies struct {
	ID         *string                                              `json:"ID" tf:"id"`
	PolicyItem []ProviderSnapshotBackupPolicySpecPoliciesPolicyItem `json:"policyItem" tf:"policy_item"`
}

type ProviderSnapshotBackupPolicySpec struct {
	State *ProviderSnapshotBackupPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource ProviderSnapshotBackupPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ProviderSnapshotBackupPolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ClusterID   *string `json:"clusterID,omitempty" tf:"cluster_id"`
	ClusterName *string `json:"clusterName" tf:"cluster_name"`
	// +optional
	NextSnapshot *string                                    `json:"nextSnapshot,omitempty" tf:"next_snapshot"`
	Policies     []ProviderSnapshotBackupPolicySpecPolicies `json:"policies" tf:"policies"`
	ProjectID    *string                                    `json:"projectID" tf:"project_id"`
	// +optional
	ReferenceHourOfDay *int64 `json:"referenceHourOfDay,omitempty" tf:"reference_hour_of_day"`
	// +optional
	ReferenceMinuteOfHour *int64 `json:"referenceMinuteOfHour,omitempty" tf:"reference_minute_of_hour"`
	// +optional
	RestoreWindowDays *int64 `json:"restoreWindowDays,omitempty" tf:"restore_window_days"`
	// +optional
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty" tf:"update_snapshots"`
}

type ProviderSnapshotBackupPolicyStatus struct {
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

// ProviderSnapshotBackupPolicyList is a list of ProviderSnapshotBackupPolicys
type ProviderSnapshotBackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProviderSnapshotBackupPolicy CRD objects
	Items []ProviderSnapshotBackupPolicy `json:"items,omitempty"`
}
