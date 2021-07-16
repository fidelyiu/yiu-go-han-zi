package yiuHanZiList

import yiuAll "github.com/fidelyiu/yiu-go-han-zi/yiu_all"

// OpAsc 将字符串数组按拼音升序排序
func OpAsc(list *[]string) {
	yiuAll.YiuHanZiListOpAsc(list)
}

// OpDesc 将字符串数组按拼音降序排序
func OpDesc(list *[]string) {
	yiuAll.YiuHanZiListOpDesc(list)
}
