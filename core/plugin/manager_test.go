package plugin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupFunc(ctx *Context) error {
	return nil
}

func tearDownFunc(ctx *Context) error {
	return nil
}

func TestRegister(t *testing.T) {
	err := Register("test", setupFunc)
	assert.Nil(t, err, "Register a plugin func should be nil")
}

func TestRegisterWithTearDown(t *testing.T) {
	err := RegisterWithTearDown("test", setupFunc, tearDownFunc)
	assert.Nil(t, err, "Register a plugin func should be nil")
}

func TestGet(t *testing.T) {
	Register("test", setupFunc)
	plugin := Get("test")
	assert.NotNil(t, plugin, "Get a plugin should not be nil")
}
