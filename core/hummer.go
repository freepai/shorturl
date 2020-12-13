package core

import (
	"github.com/freepai/hummer/core/config"
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
	"github.com/freepai/hummer/core/shorturl"
	"log"
)

type Hummer struct {
	config *config.HummerConfig
	beans  map[string]interface{}
}

func NewHummer(path string) *Hummer {
	cfg, _ := config.LoadFromYamlFile(path)
	beans := make(map[string]interface{}, 0)

	return &Hummer{
		config: cfg,
		beans:  beans,
	}
}

func (h *Hummer) SetBean(name string, bean interface{}) error {
	h.beans[name] = bean
	return nil
}

func (h *Hummer) GetBean(name string) interface{} {
	return h.beans[name]
}

func (h *Hummer) ApplyPlugins() {
	// shorturl and server plugin
	h.ApplyPlugin(server.PluginName, h.config.Server)
	h.ApplyPlugin(shorturl.PluginName, h.config.ShortUrl)

	// others plugins
	plugins := h.config.Plugins
	for _, cfg := range plugins {
		h.ApplyPluginWithConfig(cfg)
	}
}

func (h *Hummer) ApplyPlugin(name string, config interface{}) error {
	plug := plugin.Get(name)

	if plug != nil {
		ctx := plugin.NewContext(h, config)
		plug.Setup(ctx)
	} else {
		log.Fatal("not found plugin with name: " + name)
	}

	return nil
}

func (h *Hummer) ApplyPluginWithConfig(cfg *plugin.Config) error {
	return h.ApplyPlugin(cfg.Name, cfg.Params)
}

func (h *Hummer) Start() {
	h.ApplyPlugins()

	mgr := server.GetManagerFromContainer(h)
	mgr.ListenAndServe()
}
