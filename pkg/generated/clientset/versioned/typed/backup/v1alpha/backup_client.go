//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by client-gen. DO NOT EDIT.

package v1alpha

import (
	v1alpha "github.com/arangodb/kube-arangodb/pkg/apis/backup/v1alpha"
	"github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type BackupV1alphaInterface interface {
	RESTClient() rest.Interface
	ArangoBackupsGetter
	ArangoBackupPoliciesGetter
}

// BackupV1alphaClient is used to interact with features provided by the backup.arangodb.com group.
type BackupV1alphaClient struct {
	restClient rest.Interface
}

func (c *BackupV1alphaClient) ArangoBackups(namespace string) ArangoBackupInterface {
	return newArangoBackups(c, namespace)
}

func (c *BackupV1alphaClient) ArangoBackupPolicies(namespace string) ArangoBackupPolicyInterface {
	return newArangoBackupPolicies(c, namespace)
}

// NewForConfig creates a new BackupV1alphaClient for the given config.
func NewForConfig(c *rest.Config) (*BackupV1alphaClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BackupV1alphaClient{client}, nil
}

// NewForConfigOrDie creates a new BackupV1alphaClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BackupV1alphaClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BackupV1alphaClient for the given RESTClient.
func New(c rest.Interface) *BackupV1alphaClient {
	return &BackupV1alphaClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BackupV1alphaClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
