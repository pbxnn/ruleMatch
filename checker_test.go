package rulematch

import (
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func getTestParams() map[string]string {
    return map[string]string{
        "studentUid": "100",
        "sbschoolId": "100",
        "vc":         "1.1.0",
    }
}

func getTestChecker() IChecker {
    params := getTestParams()
    conf := getTestConf()
    return NewChecker("testConf", conf, params)
}

func TestChecker_DoCheck(t *testing.T) {
    checker := getTestChecker()
    r := checker.DoCheck()
    t.Log(checker.GetFailReason())
    assert.True(t, r)
}

func TestChecker_In(t *testing.T) {
    cond := &Cond{Key: "studentUid", Op: "in", Value: []string{"1", "2", "3"}}
    checker := getTestChecker()
    r := checker.In(cond, "1")
    assert.True(t, r)
}

func TestChecker_NotIn(t *testing.T) {
    cond := &Cond{Key: "studentUid", Op: "notIn", Value: []string{"1", "2", "3"}}
    checker := getTestChecker()
    r := checker.NotIn(cond, "100")
    assert.True(t, r)
}

func TestChecker_GT(t *testing.T) {
    cond := &Cond{Key: "studentUid", Op: ">", Value: []string{"1"}}
    checker := getTestChecker()
    r := checker.GT(cond, "100")
    assert.True(t, r)
}

func TestChecker_GE(t *testing.T) {
    cond := &Cond{Key: "studentUid", Op: ">=", Value: []string{"1"}}
    checker := getTestChecker()
    r := checker.GE(cond, "1")
    assert.True(t, r)
}

func TestChecker_LT(t *testing.T) {
    cond := &Cond{Key: "studentUid", Op: "<", Value: []string{"2"}}
    checker := getTestChecker()
    r := checker.LT(cond, "1")
    assert.True(t, r)
}

func TestChecker_LE(t *testing.T) {
    cond := &Cond{Key: "studentUid", Op: "<=", Value: []string{"1"}}
    checker := getTestChecker()
    r := checker.LE(cond, "1")
    assert.True(t, r)
}

func TestChecker_Rand(t *testing.T) {
    cond := &Cond{Key: "studentUid", Op: "rand", Value: []string{"100"}}
    checker := getTestChecker()
    r := checker.Rand(cond, "")
    assert.True(t, r)
}

func TestChecker_RandBy(t *testing.T) {
    cond := &Cond{Key: "studentUid", Op: "rand", Value: []string{"100"}}
    checker := getTestChecker()
    r := checker.RandBy(cond, "99")
    assert.True(t, r)
}

func TestChecker_VersionGT(t *testing.T) {
    cond := &Cond{Key: "vc", Op: "version>", Value: []string{"1.0.9"}}
    checker := getTestChecker()
    r := checker.VersionGT(cond, "1.1.0")
    assert.True(t, r)
}

func TestChecker_VersionGE(t *testing.T) {
    cond := &Cond{Key: "vc", Op: "version>=", Value: []string{"1.1.0"}}
    checker := getTestChecker()
    r := checker.VersionGE(cond, "1.1.0")
    assert.True(t, r)
}

func TestChecker_VersionLT(t *testing.T) {
    cond := &Cond{Key: "vc", Op: "version<", Value: []string{"1.1.1"}}
    checker := getTestChecker()
    r := checker.VersionLT(cond, "1.1.0")
    assert.True(t, r)
}

func TestChecker_VersionLE(t *testing.T) {
    cond := &Cond{Key: "vc", Op: "version>", Value: []string{"1.1.0"}}
    checker := getTestChecker()
    r := checker.VersionLE(cond, "1.1.0")
    assert.True(t, r)
}

func getTestConf() RuleConf {
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

    var conf RuleConf
    if err := json.Unmarshal([]byte(testConfJson), &conf); err != nil {
        fmt.Println(err)
        return nil
    }
    return conf
}
