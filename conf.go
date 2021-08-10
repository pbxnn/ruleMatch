package rulematch

import (
    "encoding/json"
    "fmt"
    "sync"
    "time"
)

func InitConfCache() {
    ConfCache.Mu = &sync.RWMutex{}
    ConfCache.ConfMap = map[string]RuleConf{}
    ConfCache.VersionMap = map[string]int64{}
}

func LoadConf(ruleName string) RuleConf {
    conf := GetFromCache(ruleName)
    if conf != nil {
        return conf
    }

    return LoadConfFromDB(ruleName)
}

func GetFromCache(ruleName string) RuleConf {
    if len(ConfCache.ConfMap) == 0 {
        return nil
    }

    if len(ConfCache.VersionMap) == 0 {
        return nil
    }

    ConfCache.Mu.Lock()
    defer ConfCache.Mu.Unlock()
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

// SetCache 更新cache + version
func SetCache(ruleName string, conf RuleConf) {
    ConfCache.Mu.Lock()
    ConfCache.ConfMap[ruleName] = conf
    ConfCache.VersionMap[ruleName] = time.Now().Unix()
    ConfCache.Mu.Unlock()
}

func LoadConfFromDB(ruleName string) RuleConf {

    var conf RuleConf
    //读取配置(mock)
    confStr := getConf()
    if err := json.Unmarshal([]byte(confStr), &conf); err != nil {
        fmt.Println(err)
        return nil
    }
    SetCache(ruleName, conf)
    return ConfCache.ConfMap[ruleName]
}

func getConf() string {
    var testConfJson = `
[
    {
        "condList":[
            {
                "key":"studentUid",
                "op":"in",
                "value":[
                    "100",
                    "200",
                    "300"
                ]
            },
            {
                "key":"sbschoolId",
                "op":"notIn",
                "value":[
                    "1",
                    "2",
                    "3"
                ]
            },
            {
                "key":"studentUid",
                "op":"randBy",
                "value":[
                    "100"
                ]
            },
            {
                "key":"studentUid",
                "op":">=",
                "value":[
                    "100"
                ]
            },
            {
                "key":"studentUid",
                "op":">",
                "value":[
                    "99"
                ]
            },
            {
                "key":"studentUid",
                "op":"<=",
                "value":[
                    "100"
                ]
            },
            {
                "key":"studentUid",
                "op":"<",
                "value":[
                    "101"
                ]
            },
            {
                "key":"vc",
                "op":"version>",
                "value":[
                    "1.0.9"
                ]
            },
            {
                "key":"vc",
                "op":"version>=",
                "value":[
                    "1.1.0"
                ]
            },
            {
                "key":"vc",
                "op":"version<",
                "value":[
                    "1.1.1"
                ]
            },
            {
                "key":"vc",
                "op":"version<=",
                "value":[
                    "1.1.0"
                ]
            }
        ],
        "seq":1
    },
    {
        "condList":[
            {
                "key":"sbschoolId",
                "op":"notIn",
                "value":[
                    "2",
                    "3"
                ]
            },
            {
                "key":"studentUid",
                "op":"rand",
                "value":[
                    "100"
                ]
            }
        ],
        "seq":2
    }
]
`

    return testConfJson
}
