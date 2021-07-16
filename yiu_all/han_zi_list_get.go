package yiuAll

import (
	yiuStrList "github.com/fidelyiu/yiu-go-tool/string_list"
	"sort"
)

// YiuHanZiListGetDesc 按照拼音将字符串降序排序
func YiuHanZiListGetDesc(list []string) []string {
	if yiuStrList.IsEmpty(list) {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		return YiuHanZiIsGe(list[i], list[j])
	})
	return list
}

// YiuHanZiListGetAsc 按照拼音将字符串升序排序
func YiuHanZiListGetAsc(list []string) []string {
	if yiuStrList.IsEmpty(list) {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		return YiuHanZiIsLe(list[i], list[j])
	})
	return list
}
