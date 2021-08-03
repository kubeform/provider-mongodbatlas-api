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

type EndpointService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointServiceSpec   `json:"spec,omitempty"`
	Status            EndpointServiceStatus `json:"status,omitempty"`
}

type EndpointServiceSpec struct {
	State *EndpointServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource EndpointServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type EndpointServiceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AwsConnectionStatus *string `json:"awsConnectionStatus,omitempty" tf:"aws_connection_status"`
	// +optional
	AzureStatus *string `json:"azureStatus,omitempty" tf:"azure_status"`
	// +optional
	DeleteRequested   *bool   `json:"deleteRequested,omitempty" tf:"delete_requested"`
	EndpointServiceID *string `json:"endpointServiceID" tf:"endpoint_service_id"`
	// +optional
	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message"`
	// +optional
	InterfaceEndpointID *string `json:"interfaceEndpointID,omitempty" tf:"interface_endpoint_id"`
	// +optional
	PrivateEndpointConnectionName *string `json:"privateEndpointConnectionName,omitempty" tf:"private_endpoint_connection_name"`
	// +optional
	PrivateEndpointIPAddress *string `json:"privateEndpointIPAddress,omitempty" tf:"private_endpoint_ip_address"`
	// +optional
	PrivateEndpointResourceID *string `json:"privateEndpointResourceID,omitempty" tf:"private_endpoint_resource_id"`
	PrivateLinkID             *string `json:"privateLinkID" tf:"private_link_id"`
	ProjectID                 *string `json:"projectID" tf:"project_id"`
	ProviderName              *string `json:"providerName" tf:"provider_name"`
}

type EndpointServiceStatus struct {
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

// EndpointServiceList is a list of EndpointServices
type EndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EndpointService CRD objects
	Items []EndpointService `json:"items,omitempty"`
}