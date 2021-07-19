package rulematch

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestLoadConfFromDB(t *testing.T) {
    ruleName := "testConf"
    conf := LoadConfFromDB(ruleName)
    t.Log(conf)
    t.Log(ConfCache.ConfMap, ConfCache.VersionMap)
    assert.NotNil(t, conf)
    assert.NotNil(t, ConfCache)
    assert.NotNil(t, ConfCache.VersionMap)
    assert.Equal(t, 2, len(conf))
    assert.Equal(t, 2, len(ConfCache.ConfMap[ruleName]))
}

func TestLoadConfFromCache(t *testing.T) {
    ruleName := "testConf"
    conf := LoadConfFromCache(ruleName)
    t.Log(conf)
    t.Log(ConfCache.ConfMap, ConfCache.VersionMap)
}

func TestLoadConf(t *testing.T) {
    ruleName := "testConf"
    conf := LoadConf(ruleName)
    t.Log(conf)
    t.Log(ConfCache.ConfMap, ConfCache.VersionMap)
}
