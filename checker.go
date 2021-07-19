package rulematch

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

func NewChecker(ruleName string, ruleConf RuleConf, params map[string]string) IChecker {
	return &Checker{
		RuleName: ruleName,
		Conf:     ruleConf,
		Params:   params,
		IsMatch:  false,
	}
}

func (c *Checker) DoCheck() {

	if len(c.Params) == 0 {
		c.fail(nil, "empty params")
		return
	}

	if len(c.Conf) == 0 {
		c.fail(nil, "empty conf")
		return
	}

	sort.Sort(c.Conf)
	for _, item := range c.Conf {
		if len(item.CondList) == 0 {
			c.fail(nil, fmt.Sprintf("empty condList, seq=%d", item.Seq))
			continue
		}

		if c.check(item.CondList) {
			c.IsMatch = true
			return
		}
	}
}

func (c *Checker) check(condList []Cond) bool {
	if len(condList) == 0 {
		c.fail(nil, "empty condList")
		return false
	}

	for _, cond := range condList {
		opFunc, ok := OpFuncMap[cond.Op]
		if !ok {
			c.fail(&cond, fmt.Sprintf("invaild op:%s", cond.Op))
			return false
		}

		input := ""
		hasInputKey := false
		if len(cond.Key) > 0 {
			input, hasInputKey = c.Params[cond.Key]
			if !hasInputKey {
				c.fail(&cond, fmt.Sprintf("empty input key:%s", cond.Key))
			}
		}

		if !opFunc(c, &cond, input) {
			c.fail(&cond, fmt.Sprintf("input=%s", input))
			return false
		}
	}
	return true
}

func (c *Checker) fail(cond *Cond, value string) {
	c.ReasonList = append(c.ReasonList, FailedReason{cond, value})
}

func (c *Checker) In(cond *Cond, input string) bool {
	return InStringSlice(cond.Value, input)
}

func (c *Checker) NotIn(cond *Cond, input string) bool {
	return !InStringSlice(cond.Value, input)
}

func (c *Checker) GT(cond *Cond, input string) bool {
	condNum, inputNum, err := convert(cond.Value, input)
	if err != nil {
		return false
	}
	return inputNum > condNum
}

func (c *Checker) GE(cond *Cond, input string) bool {
	condNum, inputNum, err := convert(cond.Value, input)
	if err != nil {
		return false
	}
	return inputNum >= condNum
}

func (c *Checker) LT(cond *Cond, input string) bool {
	condNum, inputNum, err := convert(cond.Value, input)
	if err != nil {
		return false
	}
	return inputNum < condNum
}

func (c *Checker) LE(cond *Cond, input string) bool {
	condNum, inputNum, err := convert(cond.Value, input)
	if err != nil {
		return false
	}
	return inputNum <= condNum
}

func (c *Checker) Rand(cond *Cond, input string) bool {
	condNum, _, _ := convert(cond.Value, input)
	inputNum := rand.Intn(100)
	return inputNum < condNum
}

func (c *Checker) RandBy(cond *Cond, input string) bool {
	condNum, inputNum, err := convert(cond.Value, input)
	if err != nil {
		return false
	}
	return inputNum%100 < condNum
}

func (c *Checker) VersionGT(cond *Cond, input string) bool {
	if len(cond.Value) != 1 {
		return false
	}
	return VersionCompare(input, cond.Value[0]) > 0
}

func (c *Checker) VersionGE(cond *Cond, input string) bool {
	if len(cond.Value) != 1 {
		return false
	}
	return VersionCompare(input, cond.Value[0]) >= 0
}

func (c *Checker) VersionLT(cond *Cond, input string) bool {
	if len(cond.Value) != 1 {
		return false
	}
	return VersionCompare(input, cond.Value[0]) < 0
}

func (c *Checker) VersionLE(cond *Cond, input string) bool {
	if len(cond.Value) != 1 {
		return false
	}
	return VersionCompare(input, cond.Value[0]) <= 0
}

func (c *Checker) TimeBefore(cond *Cond, input string) bool {
	condTime, inputTime, err := convert(cond.Value, input)
	if err != nil {
		return false
	}
	return inputTime < condTime
}

func (c *Checker) TimeAfter(cond *Cond, input string) bool {
	condTime, inputTime, err := convert(cond.Value, input)
	if err != nil {
		return false
	}
	return inputTime > condTime
}

func convert(condValue []string, input string) (int, int, error) {
	if len(condValue) != 1 {
		return 0, 0, fmt.Errorf("invaild condValue:%v", condValue)
	}

	condNum, err := strconv.Atoi(condValue[0])
	if err != nil {
		return 0, 0, err
	}

	inputNum, err := strconv.Atoi(input)
	if err != nil {
		return 0, 0, err
	}

	return condNum, inputNum, nil
}
