package rulematch

type IChecker interface {
	In(cond *Cond, input string) bool
	NotIn(cond *Cond, input string) bool
	GT(cond *Cond, input string) bool
	GE(cond *Cond, input string) bool
	LT(cond *Cond, input string) bool
	LE(cond *Cond, input string) bool
	Rand(cond *Cond, input string) bool
	RandBy(cond *Cond, input string) bool
	VersionGT(cond *Cond, input string) bool
	VersionGE(cond *Cond, input string) bool
	VersionLT(cond *Cond, input string) bool
	VersionLE(cond *Cond, input string) bool
	TimeBefore(cond *Cond, input string) bool
	TimeAfter(cond *Cond, input string) bool
}

type Checker struct {
	RuleName   string
	Conf       RuleConf
	Params     map[string]string
	IsMatch    bool
	ReasonList []FailedReason
}

type FailedReason struct {
	*Cond
	Value string
}

type RuleConf []struct {
	CondList []Cond `json:"condList"`
	Seq      int    `json:"seq"`
}

type Cond struct {
	Key   string   `json:"key"`
	Op    string   `json:"op"`
	Value []string `json:"value"`
}

func (r RuleConf) Len() int {
	return len(r)
}

func (r RuleConf) Less(i, j int) bool {
	return r[i].Seq <= r[j].Seq
}

func (r RuleConf) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
