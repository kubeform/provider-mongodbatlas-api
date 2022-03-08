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
	advancedv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/advanced/v1alpha1"
	alertv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/alert/v1alpha1"
	auditingv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/auditing/v1alpha1"
	cloudv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cloud/v1alpha1"
	clusterv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cluster/v1alpha1"
	customv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/custom/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/data/v1alpha1"
	databasev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/database/v1alpha1"
	encryptionv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/encryption/v1alpha1"
	eventv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/event/v1alpha1"
	globalv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/global/v1alpha1"
	ldapv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/ldap/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/maintenance/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/network/v1alpha1"
	onlinev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/online/v1alpha1"
	orgv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/org/v1alpha1"
	privatev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/private/v1alpha1"
	privatelinkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/privatelink/v1alpha1"
	projectv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/project/v1alpha1"
	searchv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/search/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/team/v1alpha1"
	teamsv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/teams/v1alpha1"
	thirdv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/third/v1alpha1"
	x509v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/x509/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	advancedv1alpha1.AddToScheme,
	alertv1alpha1.AddToScheme,
	auditingv1alpha1.AddToScheme,
	cloudv1alpha1.AddToScheme,
	clusterv1alpha1.AddToScheme,
	customv1alpha1.AddToScheme,
	datav1alpha1.AddToScheme,
	databasev1alpha1.AddToScheme,
	encryptionv1alpha1.AddToScheme,
	eventv1alpha1.AddToScheme,
	globalv1alpha1.AddToScheme,
	ldapv1alpha1.AddToScheme,
	maintenancev1alpha1.AddToScheme,
	networkv1alpha1.AddToScheme,
	onlinev1alpha1.AddToScheme,
	orgv1alpha1.AddToScheme,
	privatev1alpha1.AddToScheme,
	privatelinkv1alpha1.AddToScheme,
	projectv1alpha1.AddToScheme,
	searchv1alpha1.AddToScheme,
	teamv1alpha1.AddToScheme,
	teamsv1alpha1.AddToScheme,
	thirdv1alpha1.AddToScheme,
	x509v1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
