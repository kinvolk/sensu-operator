// Copyright 2017 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SensuRestoreList is a list of SensuRestore.
type SensuRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SensuRestore `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SensuRestore represents a Kubernetes SensuRestore Custom Resource.
// The SensuRestore CR name will be used as the name of the new restored cluster.
type SensuRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              RestoreSpec   `json:"spec"`
	Status            RestoreStatus `json:"status,omitempty"`
}

// RestoreSpec defines how to restore a sensu cluster from existing backup.
type RestoreSpec struct {
	// BackupStorageType is the type of the backup storage which is used as RestoreSource.
	BackupStorageType BackupStorageType `json:"backupStorageType"`
	// RestoreSource tells the where to get the backup and restore from.
	RestoreSource `json:",inline"`
	// SensuCluster references an SensuCluster resource whose metadata and spec
	// will be used to create the new restored SensuCluster CR.
	// This reference SensuCluster CR and all its resources will be deleted before the
	// restored SensuCluster CR is created.
	SensuCluster SensuClusterRef `json:"sensuCluster"`
}

// SensuCluster references an SensuCluster resource whose metadata and spec
// will be used to create the new restored SensuCluster CR.
// This reference SensuCluster CR and all its resources will be deleted before the
// restored SensuCluster CR is created.
type SensuClusterRef struct {
	// Name is the SensuCluster resource name.
	// This reference SensuCluster must be present in the same namespace as the restore-operator
	Name string `json:"name"`
}

type RestoreSource struct {
	// S3 tells where on S3 the backup is saved and how to fetch the backup.
	S3 *S3RestoreSource `json:"s3,omitempty"`

	// ABS tells where on ABS the backup is saved and how to fetch the backup.
	ABS *ABSRestoreSource `json:"abs,omitempty"`
}

type S3RestoreSource struct {
	// Path is the full s3 path where the backup is saved.
	// The format of the path must be: "<s3-bucket-name>/<path-to-backup-file>"
	// e.g: "mybucket/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the AWS credential and config files.
	// The file name of the credential MUST be 'credentials'.
	// The file name of the config MUST be 'config'.
	// The profile to use in both files will be 'default'.
	//
	// AWSSecret overwrites the default sensu operator wide AWS credential and config.
	AWSSecret string `json:"awsSecret"`

	// Endpoint if blank points to aws. If specified, can point to s3 compatible object
	// stores.
	Endpoint string `json:"endpoint"`

	ForcePathStyle bool
	DisableSSL     bool
}

type ABSRestoreSource struct {
	// Path is the full abs path where the backup is saved.
	// The format of the path must be: "<abs-container-name>/<path-to-backup-file>"
	// e.g: "myabscontainer/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Azure Blob Storage credential.
	ABSSecret string `json:"absSecret"`
}

// RestoreStatus reports the status of this restore operation.
type RestoreStatus struct {
	// Succeeded indicates if the backup has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any backup related failures.
	Reason string `json:"reason,omitempty"`
}
