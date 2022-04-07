package config

import (
	"sync/atomic"
	"sync"
	"patterns/observer/config_example"
)

type Config struct {
	CAFile atomic.Value
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.CAFile.Store("")
	return cfg
}

type ConfigManager struct {
	cfg *Config
	observers []config_example.Observer
	observersMu sync.Mutex
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		cfg: NewConfig(),
		observers: make([]config_example.Observer, 0),
	}
}

func (m *ConfigManager) Add(o config_example.Observer) {
	m.observersMu.Lock()
	defer m.observersMu.Unlock()
	m.observers = append(m.observers, o)
}

func (m *ConfigManager) Remove(o config_example.Observer) {
	m.observersMu.Lock()
	defer m.observersMu.Unlock()
	for i, obs := range m.observers {
		if obs.ID() != o.ID() {
			continue
		}
		m.observers = append(m.observers[:i], m.observers[i+1:]...)
		break
	}
}

func (m *ConfigManager) Notify() {
	m.observersMu.Lock()
	defer m.observersMu.Unlock()
	for _, obs := range m.observers {
		obs.Update(m.cfg.CAFile.Load().(string))
	}
}

func (m *ConfigManager) UpdateCAFile(newFile string) {
	m.cfg.CAFile.Store(newFile)
}
