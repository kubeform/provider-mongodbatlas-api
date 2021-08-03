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

type Index struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IndexSpec   `json:"spec,omitempty"`
	Status            IndexStatus `json:"status,omitempty"`
}

type IndexSpecAnalyzersCharFilters struct {
	// +optional
	IgnoreTags []string `json:"ignoreTags,omitempty" tf:"ignore_tags"`
	// +optional
	Mappings *string `json:"mappings,omitempty" tf:"mappings"`
	Type     *string `json:"type" tf:"type"`
}

type IndexSpecAnalyzersTokenFilters struct {
	// +optional
	IgnoreCase *bool `json:"ignoreCase,omitempty" tf:"ignore_case"`
	// +optional
	Matches *string `json:"matches,omitempty" tf:"matches"`
	// +optional
	Max *int64 `json:"max,omitempty" tf:"max"`
	// +optional
	MaxGram *int64 `json:"maxGram,omitempty" tf:"max_gram"`
	// +optional
	MaxShingleSize *int64 `json:"maxShingleSize,omitempty" tf:"max_shingle_size"`
	// +optional
	Min *int64 `json:"min,omitempty" tf:"min"`
	// +optional
	MinGram *int64 `json:"minGram,omitempty" tf:"min_gram"`
	// +optional
	MinShingleSize *int64 `json:"minShingleSize,omitempty" tf:"min_shingle_size"`
	// +optional
	NormalizationForm *string `json:"normalizationForm,omitempty" tf:"normalization_form"`
	// +optional
	OriginalTokens *string `json:"originalTokens,omitempty" tf:"original_tokens"`
	// +optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern"`
	// +optional
	Replacement *string `json:"replacement,omitempty" tf:"replacement"`
	// +optional
	StemmerName *string `json:"stemmerName,omitempty" tf:"stemmer_name"`
	// +optional
	TermsNotInBounds *string `json:"termsNotInBounds,omitempty" tf:"terms_not_in_bounds"`
	// +optional
	Tokens []string `json:"tokens,omitempty" tf:"tokens"`
	Type   *string  `json:"type" tf:"type"`
}

type IndexSpecAnalyzersTokenizer struct {
	// +optional
	Group *int64 `json:"group,omitempty" tf:"group"`
	// +optional
	MaxGram *int64 `json:"maxGram,omitempty" tf:"max_gram"`
	// +optional
	MaxTokenLength *int64 `json:"maxTokenLength,omitempty" tf:"max_token_length"`
	// +optional
	MinGram *int64 `json:"minGram,omitempty" tf:"min_gram"`
	// +optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern"`
	Type    *string `json:"type" tf:"type"`
}

type IndexSpecAnalyzers struct {
	// +optional
	CharFilters []IndexSpecAnalyzersCharFilters `json:"charFilters,omitempty" tf:"char_filters"`
	Name        *string                         `json:"name" tf:"name"`
	// +optional
	TokenFilters []IndexSpecAnalyzersTokenFilters `json:"tokenFilters,omitempty" tf:"token_filters"`
	Tokenizer    *IndexSpecAnalyzersTokenizer     `json:"tokenizer" tf:"tokenizer"`
}

type IndexSpec struct {
	State *IndexSpecResource `json:"state,omitempty" tf:"-"`

	Resource IndexSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IndexSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Analyzer *string `json:"analyzer" tf:"analyzer"`
	// +optional
	Analyzers      []IndexSpecAnalyzers `json:"analyzers,omitempty" tf:"analyzers"`
	ClusterName    *string              `json:"clusterName" tf:"cluster_name"`
	CollectionName *string              `json:"collectionName" tf:"collection_name"`
	Database       *string              `json:"database" tf:"database"`
	// +optional
	IndexID *string `json:"indexID,omitempty" tf:"index_id"`
	// +optional
	MappingsDynamic *bool `json:"mappingsDynamic,omitempty" tf:"mappings_dynamic"`
	// +optional
	MappingsFields *string `json:"mappingsFields,omitempty" tf:"mappings_fields"`
	Name           *string `json:"name" tf:"name"`
	ProjectID      *string `json:"projectID" tf:"project_id"`
	// +optional
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" tf:"search_analyzer"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type IndexStatus struct {
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

// IndexList is a list of Indexs
type IndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Index CRD objects
	Items []Index `json:"items,omitempty"`
}