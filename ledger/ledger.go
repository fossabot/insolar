/*
 *    Copyright 2019 Insolar Technologies
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package ledger

import (
	"context"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/ledger/artifactmanager"
	"github.com/insolar/insolar/ledger/exporter"
	"github.com/insolar/insolar/ledger/heavyserver"
	"github.com/insolar/insolar/ledger/jetcoordinator"
	"github.com/insolar/insolar/ledger/localstorage"
	"github.com/insolar/insolar/ledger/pulsemanager"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/log"
)

// Ledger is the global ledger handler. Other system parts communicate with ledger through it.
type Ledger struct {
	db              storage.DBContext
	ArtifactManager core.ArtifactManager `inject:""`
	PulseManager    core.PulseManager    `inject:""`
	JetCoordinator  core.JetCoordinator  `inject:""`
	LocalStorage    core.LocalStorage    `inject:""`
}

// Deprecated: remove after deleting TmpLedger
// GetPulseManager returns PulseManager.
func (l *Ledger) GetPulseManager() core.PulseManager {
	log.Warn("GetPulseManager is deprecated. Use component injection.")
	return l.PulseManager
}

// Deprecated: remove after deleting TmpLedger
// GetJetCoordinator returns JetCoordinator.
func (l *Ledger) GetJetCoordinator() core.JetCoordinator {
	log.Warn("GetJetCoordinator is deprecated. Use component injection.")
	return l.JetCoordinator
}

// Deprecated: remove after deleting TmpLedger
// GetArtifactManager returns artifact manager to work with.
func (l *Ledger) GetArtifactManager() core.ArtifactManager {
	log.Warn("GetArtifactManager is deprecated. Use component injection.")
	return l.ArtifactManager
}

// Deprecated: remove after deleting TmpLedger
// GetLocalStorage returns local storage to work with.
func (l *Ledger) GetLocalStorage() core.LocalStorage {
	log.Warn("GetLocalStorage is deprecated. Use component injection.")
	return l.LocalStorage
}

// NewTestLedger is the util function for creation of Ledger with provided
// private members (suitable for tests).
func NewTestLedger(
	db storage.DBContext,
	am *artifactmanager.LedgerArtifactManager,
	pm *pulsemanager.PulseManager,
	jc core.JetCoordinator,
	ls *localstorage.LocalStorage,
) *Ledger {
	return &Ledger{
		db:              db,
		ArtifactManager: am,
		PulseManager:    pm,
		JetCoordinator:  jc,
		LocalStorage:    ls,
	}
}

// GetLedgerComponents returns ledger components.
func GetLedgerComponents(conf configuration.Ledger, certificate core.Certificate) []interface{} {
	db, err := storage.NewDB(conf, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize DB"))
	}

	return []interface{}{
		db,
		storage.NewCleaner(),
		storage.NewPulseTracker(),
		storage.NewPulseStorage(),
		storage.NewJetStorage(),
		storage.NewDropStorage(conf.JetSizesHistoryDepth),
		storage.NewNodeStorage(),
		storage.NewObjectStorage(),
		storage.NewReplicaStorage(),
		storage.NewGenesisInitializer(),
		storage.NewRecentStorageProvider(conf.RecentStorage.DefaultTTL),
		artifactmanager.NewHotDataWaiterConcrete(),
		artifactmanager.NewArtifactManger(),
		jetcoordinator.NewJetCoordinator(conf.LightChainLimit),
		pulsemanager.NewPulseManager(conf),
		artifactmanager.NewMessageHandler(&conf, certificate),
		localstorage.NewLocalStorage(db),
		heavyserver.NewSync(db),
		exporter.NewExporter(conf.Exporter),
	}
}

// Start stub.
func (l *Ledger) Start(ctx context.Context) error {
	return nil
}

// Stop stops Ledger gracefully.
func (l *Ledger) Stop(ctx context.Context) error {
	return nil
}
