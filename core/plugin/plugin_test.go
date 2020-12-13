package plugin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestBean struct {
	Name string
}

type TestContainer struct {
	beans map[string]interface{}
}

func (t *TestContainer) SetBean(name string, bean interface{}) error {
	t.beans[name] = bean
	return nil
}

func (t *TestContainer) GetBean(name string) interface{} {
	return t.beans[name]
}

func (t *TestContainer) ApplyPlugin(name string, config interface{}) error {
	return nil
}

func TestNewConfig(t *testing.T) {
	cfg := &Config{
		Name:   "test",
		Params: nil,
	}

	assert.NotNil(t, cfg, "New Config should not be nil")
}

func TestNewContext(t *testing.T) {
	c := &TestContainer{}
	ctx := NewContext(c, nil)
	assert.NotNil(t, ctx, "ctx should not be nil")
}

func TestNewContextWithContainerIsNil(t *testing.T) {
	ctx := NewContext(nil, nil)
	assert.Nil(t, ctx, "ctx should be nil, when container is nil")
}

func TestContext_Register(t *testing.T) {
	c := &TestContainer{
		beans: make(map[string]interface{}, 0),
	}
	ctx := NewContext(c, nil)

	bean := &TestBean{
		Name: "test",
	}

	err := ctx.Register("testbean", bean)
	assert.Nil(t, err, "register normal bean should has not error")
}

func TestContext_RegisterRepeatReg(t *testing.T) {
	c := &TestContainer{
		beans: make(map[string]interface{}, 0),
	}
	ctx := NewContext(c, nil)

	bean := &TestBean{
		Name: "test",
	}

	err := ctx.Register("testbean", bean)
	assert.Nil(t, err, "register normal bean should has not error")

	err2 := ctx.Register("testbean", bean)
	assert.NotNil(t, err2, "repeat register normal bean should has error")
}

func TestContext_GetParam(t *testing.T) {
	c := &TestContainer{
		beans: make(map[string]interface{}, 0),
	}
	ctx := NewContext(c, nil)

	_, ok := ctx.GetParam("abc")
	assert.False(t, ok, "get params from nil should be not ok")
}

func TestContext_GetParamFromConfig(t *testing.T) {
	c := &TestContainer{
		beans: make(map[string]interface{}, 0),
	}

	bean := &TestBean{
		Name: "test",
	}

	ctx := NewContext(c, bean)

	_, ok := ctx.GetParam("Name")
	assert.False(t, ok, "get params from bean should be not ok")
}

func TestContext_GetParamFromMap(t *testing.T) {
	c := &TestContainer{
		beans: make(map[string]interface{}, 0),
	}

	params := make(map[string]string, 0)
	params["Name"] = "test"
	ctx := NewContext(c, params)

	_, ok := ctx.GetParam("Name")
	assert.True(t, ok, "get params from map should be ok")
}