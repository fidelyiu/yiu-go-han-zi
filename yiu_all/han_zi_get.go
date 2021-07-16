package yiuAll

import yiuStrList "github.com/fidelyiu/yiu-go-tool/string_list"

// YiuHanZiGetFirstLetterStr 获取首字母大写字符串
func YiuHanZiGetFirstLetterStr(s string) string {
	return yiuStrList.ToStr(yiuStrList.GetDeleteEmptyEl(YiuHanZiToPinYinFirstLetterList(s)))
}
