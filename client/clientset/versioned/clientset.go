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

package versioned

import (
	"fmt"

	advancedv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/advanced/v1alpha1"
	alertv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/alert/v1alpha1"
	auditingv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/auditing/v1alpha1"
	cloudv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/cloud/v1alpha1"
	clusterv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/cluster/v1alpha1"
	customv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/custom/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/data/v1alpha1"
	databasev1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/database/v1alpha1"
	encryptionv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/encryption/v1alpha1"
	eventv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/event/v1alpha1"
	globalv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/global/v1alpha1"
	ldapv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/ldap/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/maintenance/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/network/v1alpha1"
	onlinev1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/online/v1alpha1"
	orgv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/org/v1alpha1"
	privatev1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/private/v1alpha1"
	privatelinkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/privatelink/v1alpha1"
	projectv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/project/v1alpha1"
	searchv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/search/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/team/v1alpha1"
	teamsv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/teams/v1alpha1"
	thirdv1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/third/v1alpha1"
	x509v1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/typed/x509/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AdvancedV1alpha1() advancedv1alpha1.AdvancedV1alpha1Interface
	AlertV1alpha1() alertv1alpha1.AlertV1alpha1Interface
	AuditingV1alpha1() auditingv1alpha1.AuditingV1alpha1Interface
	CloudV1alpha1() cloudv1alpha1.CloudV1alpha1Interface
	ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface
	CustomV1alpha1() customv1alpha1.CustomV1alpha1Interface
	DataV1alpha1() datav1alpha1.DataV1alpha1Interface
	DatabaseV1alpha1() databasev1alpha1.DatabaseV1alpha1Interface
	EncryptionV1alpha1() encryptionv1alpha1.EncryptionV1alpha1Interface
	EventV1alpha1() eventv1alpha1.EventV1alpha1Interface
	GlobalV1alpha1() globalv1alpha1.GlobalV1alpha1Interface
	LdapV1alpha1() ldapv1alpha1.LdapV1alpha1Interface
	MaintenanceV1alpha1() maintenancev1alpha1.MaintenanceV1alpha1Interface
	NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface
	OnlineV1alpha1() onlinev1alpha1.OnlineV1alpha1Interface
	OrgV1alpha1() orgv1alpha1.OrgV1alpha1Interface
	PrivateV1alpha1() privatev1alpha1.PrivateV1alpha1Interface
	PrivatelinkV1alpha1() privatelinkv1alpha1.PrivatelinkV1alpha1Interface
	ProjectV1alpha1() projectv1alpha1.ProjectV1alpha1Interface
	SearchV1alpha1() searchv1alpha1.SearchV1alpha1Interface
	TeamV1alpha1() teamv1alpha1.TeamV1alpha1Interface
	TeamsV1alpha1() teamsv1alpha1.TeamsV1alpha1Interface
	ThirdV1alpha1() thirdv1alpha1.ThirdV1alpha1Interface
	X509V1alpha1() x509v1alpha1.X509V1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	advancedV1alpha1    *advancedv1alpha1.AdvancedV1alpha1Client
	alertV1alpha1       *alertv1alpha1.AlertV1alpha1Client
	auditingV1alpha1    *auditingv1alpha1.AuditingV1alpha1Client
	cloudV1alpha1       *cloudv1alpha1.CloudV1alpha1Client
	clusterV1alpha1     *clusterv1alpha1.ClusterV1alpha1Client
	customV1alpha1      *customv1alpha1.CustomV1alpha1Client
	dataV1alpha1        *datav1alpha1.DataV1alpha1Client
	databaseV1alpha1    *databasev1alpha1.DatabaseV1alpha1Client
	encryptionV1alpha1  *encryptionv1alpha1.EncryptionV1alpha1Client
	eventV1alpha1       *eventv1alpha1.EventV1alpha1Client
	globalV1alpha1      *globalv1alpha1.GlobalV1alpha1Client
	ldapV1alpha1        *ldapv1alpha1.LdapV1alpha1Client
	maintenanceV1alpha1 *maintenancev1alpha1.MaintenanceV1alpha1Client
	networkV1alpha1     *networkv1alpha1.NetworkV1alpha1Client
	onlineV1alpha1      *onlinev1alpha1.OnlineV1alpha1Client
	orgV1alpha1         *orgv1alpha1.OrgV1alpha1Client
	privateV1alpha1     *privatev1alpha1.PrivateV1alpha1Client
	privatelinkV1alpha1 *privatelinkv1alpha1.PrivatelinkV1alpha1Client
	projectV1alpha1     *projectv1alpha1.ProjectV1alpha1Client
	searchV1alpha1      *searchv1alpha1.SearchV1alpha1Client
	teamV1alpha1        *teamv1alpha1.TeamV1alpha1Client
	teamsV1alpha1       *teamsv1alpha1.TeamsV1alpha1Client
	thirdV1alpha1       *thirdv1alpha1.ThirdV1alpha1Client
	x509V1alpha1        *x509v1alpha1.X509V1alpha1Client
}

// AdvancedV1alpha1 retrieves the AdvancedV1alpha1Client
func (c *Clientset) AdvancedV1alpha1() advancedv1alpha1.AdvancedV1alpha1Interface {
	return c.advancedV1alpha1
}

// AlertV1alpha1 retrieves the AlertV1alpha1Client
func (c *Clientset) AlertV1alpha1() alertv1alpha1.AlertV1alpha1Interface {
	return c.alertV1alpha1
}

// AuditingV1alpha1 retrieves the AuditingV1alpha1Client
func (c *Clientset) AuditingV1alpha1() auditingv1alpha1.AuditingV1alpha1Interface {
	return c.auditingV1alpha1
}

// CloudV1alpha1 retrieves the CloudV1alpha1Client
func (c *Clientset) CloudV1alpha1() cloudv1alpha1.CloudV1alpha1Interface {
	return c.cloudV1alpha1
}

// ClusterV1alpha1 retrieves the ClusterV1alpha1Client
func (c *Clientset) ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface {
	return c.clusterV1alpha1
}

// CustomV1alpha1 retrieves the CustomV1alpha1Client
func (c *Clientset) CustomV1alpha1() customv1alpha1.CustomV1alpha1Interface {
	return c.customV1alpha1
}

// DataV1alpha1 retrieves the DataV1alpha1Client
func (c *Clientset) DataV1alpha1() datav1alpha1.DataV1alpha1Interface {
	return c.dataV1alpha1
}

// DatabaseV1alpha1 retrieves the DatabaseV1alpha1Client
func (c *Clientset) DatabaseV1alpha1() databasev1alpha1.DatabaseV1alpha1Interface {
	return c.databaseV1alpha1
}

// EncryptionV1alpha1 retrieves the EncryptionV1alpha1Client
func (c *Clientset) EncryptionV1alpha1() encryptionv1alpha1.EncryptionV1alpha1Interface {
	return c.encryptionV1alpha1
}

// EventV1alpha1 retrieves the EventV1alpha1Client
func (c *Clientset) EventV1alpha1() eventv1alpha1.EventV1alpha1Interface {
	return c.eventV1alpha1
}

// GlobalV1alpha1 retrieves the GlobalV1alpha1Client
func (c *Clientset) GlobalV1alpha1() globalv1alpha1.GlobalV1alpha1Interface {
	return c.globalV1alpha1
}

// LdapV1alpha1 retrieves the LdapV1alpha1Client
func (c *Clientset) LdapV1alpha1() ldapv1alpha1.LdapV1alpha1Interface {
	return c.ldapV1alpha1
}

// MaintenanceV1alpha1 retrieves the MaintenanceV1alpha1Client
func (c *Clientset) MaintenanceV1alpha1() maintenancev1alpha1.MaintenanceV1alpha1Interface {
	return c.maintenanceV1alpha1
}

// NetworkV1alpha1 retrieves the NetworkV1alpha1Client
func (c *Clientset) NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface {
	return c.networkV1alpha1
}

// OnlineV1alpha1 retrieves the OnlineV1alpha1Client
func (c *Clientset) OnlineV1alpha1() onlinev1alpha1.OnlineV1alpha1Interface {
	return c.onlineV1alpha1
}

// OrgV1alpha1 retrieves the OrgV1alpha1Client
func (c *Clientset) OrgV1alpha1() orgv1alpha1.OrgV1alpha1Interface {
	return c.orgV1alpha1
}

// PrivateV1alpha1 retrieves the PrivateV1alpha1Client
func (c *Clientset) PrivateV1alpha1() privatev1alpha1.PrivateV1alpha1Interface {
	return c.privateV1alpha1
}

// PrivatelinkV1alpha1 retrieves the PrivatelinkV1alpha1Client
func (c *Clientset) PrivatelinkV1alpha1() privatelinkv1alpha1.PrivatelinkV1alpha1Interface {
	return c.privatelinkV1alpha1
}

// ProjectV1alpha1 retrieves the ProjectV1alpha1Client
func (c *Clientset) ProjectV1alpha1() projectv1alpha1.ProjectV1alpha1Interface {
	return c.projectV1alpha1
}

// SearchV1alpha1 retrieves the SearchV1alpha1Client
func (c *Clientset) SearchV1alpha1() searchv1alpha1.SearchV1alpha1Interface {
	return c.searchV1alpha1
}

// TeamV1alpha1 retrieves the TeamV1alpha1Client
func (c *Clientset) TeamV1alpha1() teamv1alpha1.TeamV1alpha1Interface {
	return c.teamV1alpha1
}

// TeamsV1alpha1 retrieves the TeamsV1alpha1Client
func (c *Clientset) TeamsV1alpha1() teamsv1alpha1.TeamsV1alpha1Interface {
	return c.teamsV1alpha1
}

// ThirdV1alpha1 retrieves the ThirdV1alpha1Client
func (c *Clientset) ThirdV1alpha1() thirdv1alpha1.ThirdV1alpha1Interface {
	return c.thirdV1alpha1
}

// X509V1alpha1 retrieves the X509V1alpha1Client
func (c *Clientset) X509V1alpha1() x509v1alpha1.X509V1alpha1Interface {
	return c.x509V1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.advancedV1alpha1, err = advancedv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.alertV1alpha1, err = alertv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.auditingV1alpha1, err = auditingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cloudV1alpha1, err = cloudv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.clusterV1alpha1, err = clusterv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.customV1alpha1, err = customv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.dataV1alpha1, err = datav1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.databaseV1alpha1, err = databasev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.encryptionV1alpha1, err = encryptionv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.eventV1alpha1, err = eventv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.globalV1alpha1, err = globalv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.ldapV1alpha1, err = ldapv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.maintenanceV1alpha1, err = maintenancev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkV1alpha1, err = networkv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.onlineV1alpha1, err = onlinev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.orgV1alpha1, err = orgv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.privateV1alpha1, err = privatev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.privatelinkV1alpha1, err = privatelinkv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.projectV1alpha1, err = projectv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.searchV1alpha1, err = searchv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.teamV1alpha1, err = teamv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.teamsV1alpha1, err = teamsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.thirdV1alpha1, err = thirdv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.x509V1alpha1, err = x509v1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.advancedV1alpha1 = advancedv1alpha1.NewForConfigOrDie(c)
	cs.alertV1alpha1 = alertv1alpha1.NewForConfigOrDie(c)
	cs.auditingV1alpha1 = auditingv1alpha1.NewForConfigOrDie(c)
	cs.cloudV1alpha1 = cloudv1alpha1.NewForConfigOrDie(c)
	cs.clusterV1alpha1 = clusterv1alpha1.NewForConfigOrDie(c)
	cs.customV1alpha1 = customv1alpha1.NewForConfigOrDie(c)
	cs.dataV1alpha1 = datav1alpha1.NewForConfigOrDie(c)
	cs.databaseV1alpha1 = databasev1alpha1.NewForConfigOrDie(c)
	cs.encryptionV1alpha1 = encryptionv1alpha1.NewForConfigOrDie(c)
	cs.eventV1alpha1 = eventv1alpha1.NewForConfigOrDie(c)
	cs.globalV1alpha1 = globalv1alpha1.NewForConfigOrDie(c)
	cs.ldapV1alpha1 = ldapv1alpha1.NewForConfigOrDie(c)
	cs.maintenanceV1alpha1 = maintenancev1alpha1.NewForConfigOrDie(c)
	cs.networkV1alpha1 = networkv1alpha1.NewForConfigOrDie(c)
	cs.onlineV1alpha1 = onlinev1alpha1.NewForConfigOrDie(c)
	cs.orgV1alpha1 = orgv1alpha1.NewForConfigOrDie(c)
	cs.privateV1alpha1 = privatev1alpha1.NewForConfigOrDie(c)
	cs.privatelinkV1alpha1 = privatelinkv1alpha1.NewForConfigOrDie(c)
	cs.projectV1alpha1 = projectv1alpha1.NewForConfigOrDie(c)
	cs.searchV1alpha1 = searchv1alpha1.NewForConfigOrDie(c)
	cs.teamV1alpha1 = teamv1alpha1.NewForConfigOrDie(c)
	cs.teamsV1alpha1 = teamsv1alpha1.NewForConfigOrDie(c)
	cs.thirdV1alpha1 = thirdv1alpha1.NewForConfigOrDie(c)
	cs.x509V1alpha1 = x509v1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.advancedV1alpha1 = advancedv1alpha1.New(c)
	cs.alertV1alpha1 = alertv1alpha1.New(c)
	cs.auditingV1alpha1 = auditingv1alpha1.New(c)
	cs.cloudV1alpha1 = cloudv1alpha1.New(c)
	cs.clusterV1alpha1 = clusterv1alpha1.New(c)
	cs.customV1alpha1 = customv1alpha1.New(c)
	cs.dataV1alpha1 = datav1alpha1.New(c)
	cs.databaseV1alpha1 = databasev1alpha1.New(c)
	cs.encryptionV1alpha1 = encryptionv1alpha1.New(c)
	cs.eventV1alpha1 = eventv1alpha1.New(c)
	cs.globalV1alpha1 = globalv1alpha1.New(c)
	cs.ldapV1alpha1 = ldapv1alpha1.New(c)
	cs.maintenanceV1alpha1 = maintenancev1alpha1.New(c)
	cs.networkV1alpha1 = networkv1alpha1.New(c)
	cs.onlineV1alpha1 = onlinev1alpha1.New(c)
	cs.orgV1alpha1 = orgv1alpha1.New(c)
	cs.privateV1alpha1 = privatev1alpha1.New(c)
	cs.privatelinkV1alpha1 = privatelinkv1alpha1.New(c)
	cs.projectV1alpha1 = projectv1alpha1.New(c)
	cs.searchV1alpha1 = searchv1alpha1.New(c)
	cs.teamV1alpha1 = teamv1alpha1.New(c)
	cs.teamsV1alpha1 = teamsv1alpha1.New(c)
	cs.thirdV1alpha1 = thirdv1alpha1.New(c)
	cs.x509V1alpha1 = x509v1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
