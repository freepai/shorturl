package service

import (
	"fmt"
	"github.com/freepai/hummer/core/shorturl/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestIdGen struct {
	index uint64
}

func (t *TestIdGen) NextUniqueId(ns string) (uint64, error) {
	t.index++
	return t.index, nil
}

type TestIdEncode struct {}

func (t *TestIdEncode) EncodeId(ns string, id uint64) (string, error) {
	return fmt.Sprintf("%d", id), nil
}

type TestIDStore struct {
	shorturls map[string]*domain.ShortUrl
}

func NewTestIDStore() *TestIDStore {
	tids := &TestIDStore {
		shorturls: make(map[string]*domain.ShortUrl, 0),
	}

	return tids
}

func (t *TestIDStore) Save(ns string, code string, longUrl string) (*domain.ShortUrl, error) {
	su := domain.NewShortUrl(ns, code, longUrl)
	t.shorturls[code] = su
	return su, nil
}

func (t *TestIDStore) Get(ns string, code string) (*domain.ShortUrl, error) {
	su := t.shorturls[code]
	return su, nil
}

func TestNewManager(t *testing.T) {
	mgr := NewManager()
	assert.NotNil(t, mgr, "new manager should not be nil")
}

func TestRegisterIdGen(t *testing.T) {
	mgr := NewManager()
	unreg := mgr.RegisterIdGen(&TestIdGen{})
	assert.NotNil(t, unreg, "register IdEncode ok, unreg should not be nil")
}

func TestRegisterIdEncode(t *testing.T) {
	mgr := NewManager()
	unreg := mgr.RegisterIdEncode(&TestIdEncode{})
	assert.NotNil(t, unreg, "register IdEncode ok, unreg should not be nil")
}

func TestRegisterIdStore(t *testing.T) {
	mgr := NewManager()
	unreg := mgr.RegisterIdStore(&TestIDStore{})
	assert.NotNil(t, unreg, "register IdEncode ok, unreg should not be nil")
}

func TestPost(t *testing.T) {
	mgr := NewManager()
	mgr.RegisterIdGen(&TestIdGen{})
	mgr.RegisterIdEncode(&TestIdEncode{})
	mgr.RegisterIdStore(NewTestIDStore())

	su, err := mgr.Post("default", "http://www.baidu.com")

	assert.Nil(t, err, "Post error should be nil")
	assert.NotNil(t, su, "shorturl not be nil")
}

func TestGet(t *testing.T) {
	mgr := NewManager()
	mgr.RegisterIdGen(&TestIdGen{})
	mgr.RegisterIdEncode(&TestIdEncode{})
	mgr.RegisterIdStore(NewTestIDStore())

	su, _ := mgr.Post("default", "http://www.baidu.com")
	su2, err2 := mgr.Get("default", su.Code)

	assert.Nil(t, err2, "Post error should be nil")
	assert.NotNil(t, su2, "shorturl not be nil")
}