/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Cluster
func (mg *Cluster) GetTerraformResourceType() string {
	return "rancher2_cluster"
}

// GetConnectionDetailsMapping for this Cluster
func (tr *Cluster) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"aks_config[*].aad_server_app_secret": "spec.forProvider.aksConfig[*].aadServerAppSecretSecretRef", "aks_config[*].aad_tenant_id": "spec.forProvider.aksConfig[*].aadTenantIdSecretRef", "aks_config[*].add_client_app_id": "spec.forProvider.aksConfig[*].addClientAppIdSecretRef", "aks_config[*].add_server_app_id": "spec.forProvider.aksConfig[*].addServerAppIdSecretRef", "aks_config[*].client_id": "spec.forProvider.aksConfig[*].clientIdSecretRef", "aks_config[*].client_secret": "spec.forProvider.aksConfig[*].clientSecretSecretRef", "ca_cert": "status.atProvider.caCert", "eks_config[*].access_key": "spec.forProvider.eksConfig[*].accessKeySecretRef", "eks_config[*].secret_key": "spec.forProvider.eksConfig[*].secretKeySecretRef", "eks_config[*].session_token": "spec.forProvider.eksConfig[*].sessionTokenSecretRef", "gke_config[*].credential": "spec.forProvider.gkeConfig[*].credentialSecretRef", "gke_config_v2[*].google_credential_secret": "spec.forProvider.gkeConfigV2[*].googleCredentialSecretSecretRef", "kube_config": "status.atProvider.kubeConfig", "oke_config[*].kms_key_id": "spec.forProvider.okeConfig[*].kmsKeyIdSecretRef", "oke_config[*].private_key_contents": "spec.forProvider.okeConfig[*].privateKeyContentsSecretRef", "oke_config[*].private_key_passphrase": "spec.forProvider.okeConfig[*].privateKeyPassphraseSecretRef", "rke_config[*].bastion_host[*].ssh_key": "spec.forProvider.rkeConfig[*].bastionHost[*].sshKeySecretRef", "rke_config[*].cloud_provider[*].azure_cloud_provider[*].aad_client_cert_password": "spec.forProvider.rkeConfig[*].cloudProvider[*].azureCloudProvider[*].aadClientCertPasswordSecretRef", "rke_config[*].cloud_provider[*].azure_cloud_provider[*].aad_client_id": "spec.forProvider.rkeConfig[*].cloudProvider[*].azureCloudProvider[*].aadClientIdSecretRef", "rke_config[*].cloud_provider[*].azure_cloud_provider[*].aad_client_secret": "spec.forProvider.rkeConfig[*].cloudProvider[*].azureCloudProvider[*].aadClientSecretSecretRef", "rke_config[*].cloud_provider[*].azure_cloud_provider[*].subscription_id": "spec.forProvider.rkeConfig[*].cloudProvider[*].azureCloudProvider[*].subscriptionIdSecretRef", "rke_config[*].cloud_provider[*].azure_cloud_provider[*].tenant_id": "spec.forProvider.rkeConfig[*].cloudProvider[*].azureCloudProvider[*].tenantIdSecretRef", "rke_config[*].cloud_provider[*].openstack_cloud_provider[*].global[*].domain_id": "spec.forProvider.rkeConfig[*].cloudProvider[*].openstackCloudProvider[*].global[*].domainIdSecretRef", "rke_config[*].cloud_provider[*].openstack_cloud_provider[*].global[*].password": "spec.forProvider.rkeConfig[*].cloudProvider[*].openstackCloudProvider[*].global[*].passwordSecretRef", "rke_config[*].cloud_provider[*].openstack_cloud_provider[*].global[*].tenant_id": "spec.forProvider.rkeConfig[*].cloudProvider[*].openstackCloudProvider[*].global[*].tenantIdSecretRef", "rke_config[*].cloud_provider[*].openstack_cloud_provider[*].global[*].trust_id": "spec.forProvider.rkeConfig[*].cloudProvider[*].openstackCloudProvider[*].global[*].trustIdSecretRef", "rke_config[*].cloud_provider[*].openstack_cloud_provider[*].global[*].username": "spec.forProvider.rkeConfig[*].cloudProvider[*].openstackCloudProvider[*].global[*].usernameSecretRef", "rke_config[*].cloud_provider[*].vsphere_cloud_provider[*].global[*].password": "spec.forProvider.rkeConfig[*].cloudProvider[*].vsphereCloudProvider[*].global[*].passwordSecretRef", "rke_config[*].cloud_provider[*].vsphere_cloud_provider[*].global[*].user": "spec.forProvider.rkeConfig[*].cloudProvider[*].vsphereCloudProvider[*].global[*].userSecretRef", "rke_config[*].cloud_provider[*].vsphere_cloud_provider[*].virtual_center[*].password": "spec.forProvider.rkeConfig[*].cloudProvider[*].vsphereCloudProvider[*].virtualCenter[*].passwordSecretRef", "rke_config[*].cloud_provider[*].vsphere_cloud_provider[*].virtual_center[*].user": "spec.forProvider.rkeConfig[*].cloudProvider[*].vsphereCloudProvider[*].virtualCenter[*].userSecretRef", "rke_config[*].network[*].aci_network_provider[*].apic_user_crt": "spec.forProvider.rkeConfig[*].network[*].aciNetworkProvider[*].apicUserCrtSecretRef", "rke_config[*].network[*].aci_network_provider[*].apic_user_key": "spec.forProvider.rkeConfig[*].network[*].aciNetworkProvider[*].apicUserKeySecretRef", "rke_config[*].network[*].aci_network_provider[*].token": "spec.forProvider.rkeConfig[*].network[*].aciNetworkProvider[*].tokenSecretRef", "rke_config[*].nodes[*].ssh_key": "spec.forProvider.rkeConfig[*].nodes[*].sshKeySecretRef", "rke_config[*].nodes[*].user": "spec.forProvider.rkeConfig[*].nodes[*].userSecretRef", "rke_config[*].private_registries[*].ecr_credential_plugin[*].aws_secret_access_key": "spec.forProvider.rkeConfig[*].privateRegistries[*].ecrCredentialPlugin[*].awsSecretAccessKeySecretRef", "rke_config[*].private_registries[*].ecr_credential_plugin[*].aws_session_token": "spec.forProvider.rkeConfig[*].privateRegistries[*].ecrCredentialPlugin[*].awsSessionTokenSecretRef", "rke_config[*].private_registries[*].password": "spec.forProvider.rkeConfig[*].privateRegistries[*].passwordSecretRef", "rke_config[*].private_registries[*].user": "spec.forProvider.rkeConfig[*].privateRegistries[*].userSecretRef", "rke_config[*].services[*].etcd[*].backup_config[*].s3_backup_config[*].access_key": "spec.forProvider.rkeConfig[*].services[*].etcd[*].backupConfig[*].s3BackupConfig[*].accessKeySecretRef", "rke_config[*].services[*].etcd[*].backup_config[*].s3_backup_config[*].secret_key": "spec.forProvider.rkeConfig[*].services[*].etcd[*].backupConfig[*].s3BackupConfig[*].secretKeySecretRef", "rke_config[*].services[*].etcd[*].cert": "spec.forProvider.rkeConfig[*].services[*].etcd[*].certSecretRef", "rke_config[*].services[*].etcd[*].key": "spec.forProvider.rkeConfig[*].services[*].etcd[*].keySecretRef"}
}

// GetObservation of this Cluster
func (tr *Cluster) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Cluster
func (tr *Cluster) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Cluster
func (tr *Cluster) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Cluster
func (tr *Cluster) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Cluster
func (tr *Cluster) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Cluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Cluster) LateInitialize(attrs []byte) (bool, error) {
	params := &ClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Cluster) GetTerraformSchemaVersion() int {
	return 1
}
