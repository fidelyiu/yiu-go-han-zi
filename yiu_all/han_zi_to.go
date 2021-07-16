package yiuAll

import yiuStr "github.com/fidelyiu/yiu-go-tool/string"

// YiuHanZiToPinYinList 汉字转拼音List，没有拼音的就为空
func YiuHanZiToPinYinList(s string) []string {
	var result []string
	rList := yiuStr.ToRuneList(s)
	for i := range rList {
		result = append(result, YiuPinYinGetDefByRune(rList[i]))
	}
	return result
}

// YiuHanZiToPinYinFirstLetterList 获取拼音首字母大写
func YiuHanZiToPinYinFirstLetterList(s string) []string {
	var result []string
	rList := yiuStr.ToRuneList(s)
	for i := range rList {
		result = append(result, YiuPinYinGetFirstLetterByRune(rList[i]))
	}
	return result
}

// YiuHanZiToBiHuaList 汉字转笔画List，没有笔画的就为空
func YiuHanZiToBiHuaList(s string) []uint {
	var result []uint
	rList := yiuStr.ToRuneList(s)
	for i := range rList {
		result = append(result, YiuBiHuaGetByRune(rList[i]))
	}
	return result
}

func YiuHanZiToBuShouList(s string) []string {
	var result []string
	rList := yiuStr.ToRuneList(s)
	for i := range rList {
		result = append(result, YiuBuShouGetByRune(rList[i]))
	}
	return result
}
