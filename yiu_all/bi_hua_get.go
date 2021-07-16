package yiuAll

import yiuVar "github.com/fidelyiu/yiu-go-han-zi/yiu_var"

func YiuBiHuaGetByRune(r rune) uint {
	if YiuHanZiIsNotHanZi(r) {
		return 0
	}
	return yiuVar.FuncHanZiGetBiHua(r)
}
