package yiuAll

import (
	yiuVar "github.com/fidelyiu/yiu-go-han-zi/yiu_var"
	yiuStr "github.com/fidelyiu/yiu-go-tool/string"
	"unicode"
)

// YiuPinYinToNoTone 将带音调的字符替换成没有音调的
func YiuPinYinToNoTone(str string) string {
	tpList := yiuStr.ToStrList(str)
	tp := ""
	for i := range tpList {
		if yiuVar.PinYinToneMap[tpList[i]] != "" {
			tp += yiuVar.PinYinToneMap[tpList[i]]
		} else {
			tp += tpList[i]
		}
	}
	return tp
}

// YiuPinYinToLower 将带音调的字符替换成没有音调的,并全换成小写
func YiuPinYinToLower(str string) string {
	str = YiuPinYinToNoTone(str)
	runeList := yiuStr.ToRuneList(str)
	for i := range runeList {
		runeList[i] = unicode.ToLower(runeList[i])
	}
	return string(runeList)
}

// YiuPinYinToUpper 将带音调的字符替换成没有音调的,并全换成大写
func YiuPinYinToUpper(str string) string {
	str = YiuPinYinToNoTone(str)
	runeList := yiuStr.ToRuneList(str)
	for i := range runeList {
		runeList[i] = unicode.ToUpper(runeList[i])
	}
	return string(runeList)
}
