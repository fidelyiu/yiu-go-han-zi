package yiuHanZi

import yiuAll "github.com/fidelyiu/yiu-go-han-zi/yiu_all"

// IsGe 字符串是否大于等于
//
// - 汉字在前，非汉字在后
// - 按拼音、音调、部首、笔画排序
func IsGe(s1, s2 string) bool {
	return yiuAll.YiuHanZiIsGe(s1, s2)
}

// IsGt 字符串是否大于
//
// - 汉字在前，非汉字在后
// - 按拼音、音调、部首、笔画排序
func IsGt(s1, s2 string) bool {
	return yiuAll.YiuHanZiIsGt(s1, s2)
}

// IsHanZi 是否是汉字
func IsHanZi(i rune) bool {
	return yiuAll.YiuHanZiIsHanZi(i)
}

// IsLe 字符串是否小于等于
//
// - 汉字在前，非汉字在后
// - 按拼音、音调、部首、笔画排序
func IsLe(s1, s2 string) bool {
	return yiuAll.YiuHanZiIsLe(s1, s2)
}

// IsLt 字符串是否小于
//
// - 汉字在前，非汉字在后
// - 按拼音、音调、部首、笔画排序
func IsLt(s1, s2 string) bool {
	return yiuAll.YiuHanZiIsLt(s1, s2)
}

// IsNotHanZi 是否不是汉字
func IsNotHanZi(i rune) bool {
	return yiuAll.YiuHanZiIsNotHanZi(i)
}
