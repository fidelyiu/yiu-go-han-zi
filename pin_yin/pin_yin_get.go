package yiuPinYin

import yiuAll "github.com/fidelyiu/yiu-go-han-zi/yiu_all"

// GetByRune 根据rune获取拼音
func GetByRune(r rune) []string {
	return yiuAll.YiuPinYinGetByRune(r)
}

// GetDefByRune 根据rune获取默认拼音
func GetDefByRune(r rune) string {
	return yiuAll.YiuPinYinGetDefByRune(r)
}

// GetFirstLetterByRune 根据rune获取拼音的首字母大写
func GetFirstLetterByRune(r rune) string {
	return yiuAll.YiuPinYinGetFirstLetterByRune(r)
}
