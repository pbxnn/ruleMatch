package rulematch

import "time"

func LoadConf(ruleName string) *RuleConf {
	conf := LoadConfFromCache(ruleName)
	if conf != nil {
		return conf
	}

}

func LoadConfFromCache(ruleName string) *RuleConf {
	if len(ConfCache.ConfMap) == 0 {
		return nil
	}

	if len(ConfCache.VersionMap) == 0 {
		return nil
	}

	conf, ok := ConfCache.ConfMap[ruleName]
	if !ok || conf == nil {
		return nil
	}

	version, ok := ConfCache.VersionMap[ruleName]
	curTime := time.Now().Unix()
	if !ok || version <= curTime-CONF_CACHE_EXPIRE {
		return nil
	}

	return conf
}

func LoadConfFromDB(ruleName string) *RuleConf {
	return nil
}
