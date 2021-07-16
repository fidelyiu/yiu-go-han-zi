package yiuAll

import yiuVar "github.com/fidelyiu/yiu-go-han-zi/yiu_var"

// YiuBuShouGetByRune 根据rune获取部首
func YiuBuShouGetByRune(r rune) string {
	if YiuHanZiIsNotHanZi(r) {
		return ""
	}
	return yiuVar.FuncHanZiGetBuShou(r)
}
