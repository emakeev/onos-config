// Copyright 2019-present Open Networking Foundation.
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

package manager

import (
	"fmt"
	"github.com/onosproject/onos-config/pkg/events"
	"github.com/onosproject/onos-config/pkg/listener"
	"github.com/onosproject/onos-config/pkg/southbound/synchronizer"
	"github.com/onosproject/onos-config/pkg/southbound/topocache"
	"github.com/onosproject/onos-config/pkg/store"
	"github.com/onosproject/onos-config/pkg/store/change"
	"strings"
	"log"
)

var mgr Manager

// Manager single point of entry for the config system
type Manager struct {
	ConfigStore    *store.ConfigurationStore
	ChangeStore    *store.ChangeStore
	deviceStore    *topocache.DeviceStore
	NetworkStore   *store.NetworkStore
	topoChannel    chan events.Event
	changesChannel chan events.Event
}

// NewManager initializes the network config manager subsystem.
func NewManager(configs *store.ConfigurationStore, changes *store.ChangeStore, device *topocache.DeviceStore,
	network *store.NetworkStore, topoCh chan events.Event) *Manager {
	log.Println("Creating Manager")
	mgr = Manager{
		ConfigStore:    configs,
		ChangeStore:    changes,
		deviceStore:    device,
		NetworkStore:   network,
		topoChannel:    topoCh,
		changesChannel: make(chan events.Event, 10),
	}
	return &mgr

}

// Run starts a synchronizer based on the devices and the northbound services
func (m *Manager) Run() {
	log.Println("Starting Manager")
	// Start the main listener system
	go listener.Listen(m.changesChannel)
	go synchronizer.Factory(m.ChangeStore, m.deviceStore, m.topoChannel)
}

//Close kills the channels and manager related objects
func (m *Manager) Close() {
	log.Println("Closing Manager")
	close(m.changesChannel)
}

// GetNetworkConfig returns a set of change values given a target, a configuration name, a path and a layer.
// The layer is the numbers of config changes we want to go back in time for. 0 is the latest
func (m *Manager) GetNetworkConfig(target string, configname string, path string, layer int) ([]change.ConfigValue, error){
	if _, ok := m.deviceStore.Store[target]; !ok {
		return nil, fmt.Errorf("Device not present %s", target)
	}
	fmt.Println("Getting config for", target, path)
	//TODO the key of the config store should be a tuple of (devicename, configname) use the param
	var config store.Configuration
	for configID, cfg := range m.ConfigStore.Store {
		if cfg.Device == target {
			configname = configID
			config = cfg
			break
		}
	}
	configValues := config.ExtractFullConfig(m.ChangeStore.Store, layer)
	if len(configValues) == 0 || path == "/*" {
		return configValues, nil
	}
	filteredValues := make([]change.ConfigValue, 0)
	for _, cv := range configValues {
		if strings.Contains(cv.Path, path) {
			filteredValues = append(filteredValues, cv)
		}
	}

	return filteredValues, nil
}

// GetManager retuns the initialized and running instance of manager.
// Should be called only after NewManager and Run are done.
func GetManager() *Manager {
	return &mgr
}
