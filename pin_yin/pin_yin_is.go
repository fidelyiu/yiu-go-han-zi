package yiuPinYin

import yiuAll "github.com/fidelyiu/yiu-go-han-zi/yiu_all"

// IsGe 判断拼音字符串是否大于等于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func IsGe(s1, s2 string) bool {
	return yiuAll.YiuPinYinIsGe(s1, s2)
}

// IsGt 判断拼音字符串是否大于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func IsGt(s1, s2 string) bool {
	return yiuAll.YiuPinYinIsGt(s1, s2)
}

// IsLe 判断拼音字符串是否小于等于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func IsLe(s1, s2 string) bool {
	return yiuAll.YiuPinYinIsLe(s1, s2)
}

// IsLt 判断拼音字符串是否小于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func IsLt(s1, s2 string) bool {
	return yiuAll.YiuPinYinIsLt(s1, s2)
}
