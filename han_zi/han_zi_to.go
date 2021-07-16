package yiuHanZi

import yiuAll "github.com/fidelyiu/yiu-go-han-zi/yiu_all"

// ToBiHuaList 汉字转笔画List，没有笔画的就为空
func ToBiHuaList(s string) []uint {
	return yiuAll.YiuHanZiToBiHuaList(s)
}

func ToBuShouList(s string) []string {
	return yiuAll.YiuHanZiToBuShouList(s)
}

// ToPinYinFirstLetterList 获取拼音首字母大写
func ToPinYinFirstLetterList(s string) []string {
	return yiuAll.YiuHanZiToPinYinFirstLetterList(s)
}

// ToPinYinList 汉字转拼音List，没有拼音的就为空
func ToPinYinList(s string) []string {
	return yiuAll.YiuHanZiToPinYinList(s)
}
