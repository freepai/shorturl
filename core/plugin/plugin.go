package plugin

import (
	"errors"
	"fmt"
	"log"
)

// Config is plugin's config object
type Config struct {
	Name   string            `yaml:"name,omitempty"` // Name is plugin's name
	Params map[string]string `yaml:"params,omitempty"` // Params is plugin's params
}

// Container is the space to extension
type Container interface {
	SetBean(name string, bean interface{}) error
	GetBean(name string) interface{}
	ApplyPlugin(name string, config interface{}) error
}

// Context is plugin context
// which can add bean or get bean
// which can get config for current plugin
type Context struct {
	container Container
	config    interface{}
}

// Plugin is a function which register bean
// in container, or extension bean's interface
type Plugin struct {
	Name string
	Setup func(ctx *Context) error
	TearDown func(ctx *Context) error
}

// NewContext is construct func for Context
// param container is Container
// param config is the config
func NewContext(container Container, config interface{}) *Context {
	if container==nil {
		log.Print("container should not be nil")
		return nil
	}

	return &Context{
		container: container,
		config:    config,
	}
}

func (c *Context) Register(name string, bean interface{}) error {
	b := c.container.GetBean(name)
	if b != nil {
		return errors.New(fmt.Sprintf("bean with name %s already registered", name))
	}

	return c.SetBean(name, bean)
}

func (c *Context) SetBean(name string, bean interface{}) error {
	c.container.SetBean(name, bean)
	return nil
}

func (c *Context) GetBean(name string) interface{} {
	return c.container.GetBean(name)
}

func (c *Context) GetConfig() interface{} {
	return c.config
}

func (c *Context) GetParam(key string) (string, bool) {
	if c.config == nil {
		return "", false
	}

	params, ok := c.config.(map[string]string)
	if !ok {
		return "", false
	}

	return params[key], true
}

func (c *Context) ApplyPlugin(name string, params map[string]string) error {
	return c.container.ApplyPlugin(name, params)
}
