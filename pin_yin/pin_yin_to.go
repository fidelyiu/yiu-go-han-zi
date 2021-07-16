package yiuPinYin

import yiuAll "github.com/fidelyiu/yiu-go-han-zi/yiu_all"

// ToLower 将带音调的字符替换成没有音调的,并全换成小写
func ToLower(str string) string {
	return yiuAll.YiuPinYinToLower(str)
}

// ToNoTone 将带音调的字符替换成没有音调的
func ToNoTone(str string) string {
	return yiuAll.YiuPinYinToNoTone(str)
}

// ToUpper 将带音调的字符替换成没有音调的,并全换成大写
func ToUpper(str string) string {
	return yiuAll.YiuPinYinToUpper(str)
}
