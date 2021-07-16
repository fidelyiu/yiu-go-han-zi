package yiuHanZiList

import yiuAll "github.com/fidelyiu/yiu-go-han-zi/yiu_all"

// GetAsc 按照拼音将字符串升序排序
func GetAsc(list []string) []string {
	return yiuAll.YiuHanZiListGetAsc(list)
}

// GetDesc 按照拼音将字符串降序排序
func GetDesc(list []string) []string {
	return yiuAll.YiuHanZiListGetDesc(list)
}
