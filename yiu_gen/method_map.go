package yiuHanZiGen

import (
	yiuStrList "github.com/fidelyiu/yiu-go-tool/string_list"
)

func GetMethodMap() map[string]map[string][]string {
	allMethodMap := make(map[string]map[string][]string)
	allMethodMap["yiuHanZi"] = map[string][]string{
		"Get": {
			"GetFirstLetterStr",
		},
		"Is": {
			"IsHanZi",
			"IsNotHanZi",
			"IsGt",
			"IsGe",
			"IsLt",
			"IsLe",
		},
		"To": {
			"ToPinYinList",
			"ToPinYinFirstLetterList",
			"ToBiHuaList",
			"ToBuShouList",
		},
	}
	allMethodMap["yiuHanZiList"] = map[string][]string{
		"Get": {
			"GetDesc",
			"GetAsc",
		},
		"Op": {
			"OpDesc",
			"OpAsc",
		},
	}
	allMethodMap["yiuBiHua"] = map[string][]string{
		"Get": {
			"GetByRune",
		},
	}
	allMethodMap["yiuBuShou"] = map[string][]string{
		"Get": {
			"GetByRune",
		},
	}
	allMethodMap["yiuPinYin"] = map[string][]string{
		"Get": {
			"GetByRune",
			"GetDefByRune",
			"GetFirstLetterByRune",
		},
		"Is": {
			"IsGt",
			"IsGe",
			"IsLt",
			"IsLe",
		},
		"To": {
			"ToNoTone",
			"ToLower",
			"ToUpper",
		},
	}
	// 名字排序
	for s1 := range allMethodMap {
		for s2 := range allMethodMap[s1] {
			allMethodMap[s1][s2] = yiuStrList.GetAscUnicode(allMethodMap[s1][s2])
		}
	}
	return allMethodMap
}
