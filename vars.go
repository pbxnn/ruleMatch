package rulematch

import "sync"

const (
	OP_RAND        = "rand"
	OP_RAND_BY     = "randBy"
	OP_IN          = "in"
	OP_NOT_IN      = "notIn"
	OP_GT          = ">"
	OP_GE          = ">="
	OP_LT          = "<"
	OP_LE          = "<="
	OP_VERSION_GT  = "version>"
	OP_VERSION_GE  = "version>="
	OP_VERSION_LT  = "version<"
	OP_VERSION_LE  = "version<="

	CONF_CACHE_EXPIRE = 10 //本地缓存过期时间（秒）
)

var OpFuncMap = map[string]func(IChecker, *Cond, string) bool{
	OP_RAND:        IChecker.Rand,
	OP_RAND_BY:     IChecker.RandBy,
	OP_IN:          IChecker.In,
	OP_NOT_IN:      IChecker.NotIn,
	OP_GT:          IChecker.GT,
	OP_GE:          IChecker.GE,
	OP_LT:          IChecker.LT,
	OP_LE:          IChecker.LE,
	OP_VERSION_GT:  IChecker.VersionGT,
	OP_VERSION_GE:  IChecker.VersionGE,
	OP_VERSION_LT:  IChecker.VersionLT,
	OP_VERSION_LE:  IChecker.VersionLE,
}

var ConfCache struct {
	Mu         *sync.RWMutex
	ConfMap    map[string]RuleConf
	VersionMap map[string]int64
}
