package yiuAll

import (
	yiuVar "github.com/fidelyiu/yiu-go-han-zi/yiu_var"
	yiuStr "github.com/fidelyiu/yiu-go-tool/string"
	"strings"
)

// YiuPinYinGetByRune 根据rune获取拼音
func YiuPinYinGetByRune(r rune) []string {
	if YiuHanZiIsNotHanZi(r) {
		return nil
	}
	return yiuVar.FuncHanZiGetPinYin(r)
}

// YiuPinYinGetDefByRune 根据rune获取默认拼音
func YiuPinYinGetDefByRune(r rune) string {
	if YiuHanZiIsNotHanZi(r) {
		return ""
	}
	pList := yiuVar.FuncHanZiGetPinYin(r)
	if len(pList) == 0 {
		return ""
	} else {
		return pList[0]
	}
}

// YiuPinYinGetFirstLetterByRune 根据rune获取拼音的首字母大写
func YiuPinYinGetFirstLetterByRune(r rune) string {
	ts := yiuStr.GetFirstRuneStr(YiuPinYinToNoTone(YiuPinYinGetDefByRune(r)))
	if yiuStr.IsLowerLetter(ts) {
		return strings.ToUpper(ts)
	}
	if yiuStr.IsUpperLetter(ts) {
		return ts
	}
	return ""
}
