package yiuAll

import (
	yiuVar "github.com/fidelyiu/yiu-go-han-zi/yiu_var"
	yiuRuneList "github.com/fidelyiu/yiu-go-tool/rune_list"
	yiuStr "github.com/fidelyiu/yiu-go-tool/string"
	"strings"
)

// YiuPinYinIsGt 判断拼音字符串是否大于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func YiuPinYinIsGt(s1, s2 string) bool {
	r1 := yiuStr.ToRuneList(s1)
	r2 := yiuStr.ToRuneList(s2)
	for i := range r1 {
		tr := yiuRuneList.GetElByIndex(r2, i)
		if r1[i] == tr {
			continue
		} else {
			// rune不一样了
			return pinYinRuneIsGt(r1[i], tr)
		}
	}
	return true
}

// YiuPinYinIsGe 判断拼音字符串是否大于等于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func YiuPinYinIsGe(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return YiuPinYinIsGt(s1, s2)
}

// YiuPinYinIsLt 判断拼音字符串是否小于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func YiuPinYinIsLt(s1, s2 string) bool {
	r1 := yiuStr.ToRuneList(s1)
	r2 := yiuStr.ToRuneList(s2)
	for i := range r1 {
		tr := yiuRuneList.GetElByIndex(r2, i)
		if r1[i] == tr {
			continue
		} else {
			// rune不一样了
			return !pinYinRuneIsGt(r1[i], tr)
		}
	}
	return false
}

// YiuPinYinIsLe 判断拼音字符串是否小于等于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func YiuPinYinIsLe(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return YiuPinYinIsLt(s1, s2)
}

func pinYinRuneIsGt(r1, r2 rune) bool {
	// 空数据要排在后面
	if r1 == 0 {
		return true
	}
	s1, s2 := string(r1), string(r2)
	p1 := yiuVar.PinYinToneMap[s1]
	p2 := yiuVar.PinYinToneMap[s2]
	if yiuStr.IsEmpty(p1) {
		p1 = s1
	}
	if yiuStr.IsEmpty(p2) {
		p2 = s2
	}
	if p1 == p2 {
		// 去掉音调后是一样的，那么就判断音调
		return yiuVar.PinYinToneIndexMap[s1] > yiuVar.PinYinToneIndexMap[s2]
	} else {
		// 去掉音调后不一样，那么就比较unicode码
		// 大小写一样
		l1 := strings.ToLower(p1)
		l2 := strings.ToLower(p2)
		if l1 == l2 {
			// 转成小写一样，如果1是大写（排在后面）就大于
			return yiuStr.IsUpperLetter(l1)
		} else {
			// 转成小写不一样，那么判断unicode
			return yiuRuneList.IsGtByUnicode(yiuStr.ToRuneList(l1), yiuStr.ToRuneList(l2))
		}
	}
}
