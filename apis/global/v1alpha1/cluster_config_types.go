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

type ClusterConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterConfigSpec   `json:"spec,omitempty"`
	Status            ClusterConfigStatus `json:"status,omitempty"`
}

type ClusterConfigSpecCustomZoneMappings struct {
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type ClusterConfigSpecManagedNamespaces struct {
	Collection     *string `json:"collection" tf:"collection"`
	CustomShardKey *string `json:"customShardKey" tf:"custom_shard_key"`
	Db             *string `json:"db" tf:"db"`
	// +optional
	IsCustomShardKeyHashed *bool `json:"isCustomShardKeyHashed,omitempty" tf:"is_custom_shard_key_hashed"`
	// +optional
	IsShardKeyUnique *bool `json:"isShardKeyUnique,omitempty" tf:"is_shard_key_unique"`
}

type ClusterConfigSpec struct {
	State *ClusterConfigSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterConfigSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterConfigSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterName *string `json:"clusterName" tf:"cluster_name"`
	// +optional
	CustomZoneMapping map[string]string `json:"customZoneMapping,omitempty" tf:"custom_zone_mapping"`
	// +optional
	CustomZoneMappings []ClusterConfigSpecCustomZoneMappings `json:"customZoneMappings,omitempty" tf:"custom_zone_mappings"`
	// +optional
	ManagedNamespaces []ClusterConfigSpecManagedNamespaces `json:"managedNamespaces,omitempty" tf:"managed_namespaces"`
	ProjectID         *string                              `json:"projectID" tf:"project_id"`
}

type ClusterConfigStatus struct {
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

// ClusterConfigList is a list of ClusterConfigs
type ClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterConfig CRD objects
	Items []ClusterConfig `json:"items,omitempty"`
}
