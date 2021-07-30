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

type PartyIntegration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PartyIntegrationSpec   `json:"spec,omitempty"`
	Status            PartyIntegrationStatus `json:"status,omitempty"`
}

type PartyIntegrationSpec struct {
	State *PartyIntegrationSpecResource `json:"state,omitempty" tf:"-"`

	Resource PartyIntegrationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type PartyIntegrationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// +optional
	ApiKey *string `json:"-" sensitive:"true" tf:"api_key"`
	// +optional
	ApiToken *string `json:"-" sensitive:"true" tf:"api_token"`
	// +optional
	ChannelName *string `json:"channelName,omitempty" tf:"channel_name"`
	// +optional
	FlowName *string `json:"flowName,omitempty" tf:"flow_name"`
	// +optional
	LicenseKey *string `json:"-" sensitive:"true" tf:"license_key"`
	// +optional
	OrgName   *string `json:"orgName,omitempty" tf:"org_name"`
	ProjectID *string `json:"projectID" tf:"project_id"`
	// +optional
	ReadToken *string `json:"-" sensitive:"true" tf:"read_token"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// +optional
	RoutingKey *string `json:"-" sensitive:"true" tf:"routing_key"`
	// +optional
	Secret *string `json:"-" sensitive:"true" tf:"secret"`
	// +optional
	ServiceKey *string `json:"-" sensitive:"true" tf:"service_key"`
	// +optional
	TeamName *string `json:"teamName,omitempty" tf:"team_name"`
	Type     *string `json:"type" tf:"type"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
	// +optional
	WriteToken *string `json:"-" sensitive:"true" tf:"write_token"`
}

type PartyIntegrationStatus struct {
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

// PartyIntegrationList is a list of PartyIntegrations
type PartyIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PartyIntegration CRD objects
	Items []PartyIntegration `json:"items,omitempty"`
}
