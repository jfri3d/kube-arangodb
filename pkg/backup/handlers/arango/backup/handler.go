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
// Author Adam Janikowski
//

package backup

import (
	"fmt"
	"reflect"
	"time"

	"github.com/arangodb/kube-arangodb/pkg/backup/operator/event"
	"github.com/arangodb/kube-arangodb/pkg/backup/operator/operation"

	"k8s.io/client-go/kubernetes"

	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/rs/zerolog/log"

	backupApi "github.com/arangodb/kube-arangodb/pkg/apis/backup/v1alpha"
	database "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1alpha"
	arangoClientSet "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultArangoClientTimeout = 30 * time.Second

	StateChange     = "StateChange"
	FinalizerChange = "FinalizerChange"
)

type handler struct {
	client     arangoClientSet.Interface
	kubeClient kubernetes.Interface

	eventRecorder event.EventRecorderInstance

	arangoClientFactory ArangoClientFactory
	arangoClientTimeout time.Duration
}

func (h *handler) Name() string {
	return backupApi.ArangoBackupResourceKind
}

func (h *handler) Handle(item operation.Item) error {
	// Get Backup object. It also cover NotFound case
	b, err := h.client.BackupV1alpha().ArangoBackups(item.Namespace).Get(item.Name, meta.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}

		return err
	}

	// Check if we should start finalizer
	if b.DeletionTimestamp != nil {
		log.Debug().Msgf("Finalizing %s %s/%s",
			item.Kind,
			item.Namespace,
			item.Name)

		return h.finalize(b)
	}

	// Do not act on delete event, finalizer should be used
	if item.Operation == operation.OperationDelete {
		return nil
	}

	// Add finalizers
	if !hasFinalizers(b) {
		b.Finalizers = appendFinalizers(b)
		log.Info().Msgf("Updating finalizers %s %s/%s",
			item.Kind,
			item.Namespace,
			item.Name)

		if _, err = h.client.BackupV1alpha().ArangoBackups(item.Namespace).Update(b); err != nil {
			return err
		}

		return nil
	}

	status, err := h.processArangoBackup(b.DeepCopy())
	if err != nil {
		return err
	}

	// Nothing to update, objects are equal
	if reflect.DeepEqual(b.Status, status) {
		return nil
	}

	// Ensure that transit is possible
	if err = backupApi.ArangoBackupStateMap.Transit(b.Status.State, status.State); err != nil {
		return err
	}

	// Log message about state change
	if b.Status.State != status.State {
		if status.State == backupApi.ArangoBackupStateFailed {
			h.eventRecorder.Warning(b, StateChange, "Transiting from %s to %s with error: %s",
				b.Status.State,
				status.State,
				status.Message)
		} else {
			h.eventRecorder.Normal(b, StateChange, "Transiting from %s to %s",
				b.Status.State,
				status.State)
		}
	} else {
		// Keep old time in case when object did not change
		status.Time = b.Status.Time
	}

	b.Status = status

	log.Debug().Msgf("Updating %s %s/%s",
		item.Kind,
		item.Namespace,
		item.Name)

	// Update status on object
	if _, err = h.client.BackupV1alpha().ArangoBackups(item.Namespace).UpdateStatus(b); err != nil {
		return err
	}

	return nil
}

func (h *handler) processArangoBackup(backup *backupApi.ArangoBackup) (backupApi.ArangoBackupStatus, error) {
	if err := backup.Validate(); err != nil {
		return createFailedState(err, backup.Status), nil
	}

	if f, ok := stateHolders[backup.Status.State]; !ok {
		return backupApi.ArangoBackupStatus{}, fmt.Errorf("state %s is not supported", backup.Status.State)
	} else {
		return f(h, backup)
	}
}

func (h *handler) CanBeHandled(item operation.Item) bool {
	return item.Group == database.SchemeGroupVersion.Group &&
		item.Version == database.SchemeGroupVersion.Version &&
		item.Kind == backupApi.ArangoBackupResourceKind
}

func (h *handler) getArangoDeploymentObject(backup *backupApi.ArangoBackup) (*database.ArangoDeployment, error) {
	if backup.Spec.Deployment.Name == "" {
		return nil, fmt.Errorf("deployment ref is not specified for backup %s/%s", backup.Namespace, backup.Name)
	}

	return h.client.DatabaseV1alpha().ArangoDeployments(backup.Namespace).Get(backup.Spec.Deployment.Name, meta.GetOptions{})
}
