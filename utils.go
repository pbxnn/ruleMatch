package rulematch

import "strings"

func InStringSlice(collection []string, target string) bool {
	for _, v := range collection {
		if v == target {
			return true
		}
	}

	return false
}

// VersionCompare 该函数比较两个版本号是否相等，是否大于或小于的关系
// 返回值：0表示v1与v2相等；1表示v1大于v2；2表示v1小于v2
func VersionCompare(v1, v2 string) int {
	// 替换一些常见的版本符号
	replaceMap := map[string]string{"V": "", "v": "", "-": "."}
	for k, v := range replaceMap {
		if strings.Contains(v1, k) {
			strings.Replace(v1, k, v, -1)
		}
		if strings.Contains(v2, k) {
			strings.Replace(v2, k, v, -1)
		}
	}

	ver1 := strings.Split(v1, ".")
	ver2 := strings.Split(v2, ".")
	// 找出v1和v2哪一个最短
	var shorter int
	if len(ver1) > len(ver2) {
		shorter = len(ver2)
	} else {
		shorter = len(ver1)
	}

	for i := 0; i < shorter; i++ {
		if ver1[i] == ver2[i] {
			if shorter-1 == i {
				if len(ver1) == len(ver2) {
					return 0
				} else {
					if len(ver1) > len(ver2) {
						return 1
					} else {
						return -1
					}
				}
			}
		} else if ver1[i] > ver2[i] {
			return 1
		} else {
			return -1
		}
	}

	if len(ver1) > len(ver2) {
		return 1
	}

	if len(ver2) > len(ver1) {
		return -1
	}

	return 0
}
