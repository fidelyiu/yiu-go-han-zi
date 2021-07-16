package yiuAll

// YiuHanZiListOpDesc 将字符串数组按拼音降序排序
func YiuHanZiListOpDesc(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuHanZiListGetDesc(*list)
}

// YiuHanZiListOpAsc 将字符串数组按拼音升序排序
func YiuHanZiListOpAsc(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuHanZiListGetAsc(*list)
}
