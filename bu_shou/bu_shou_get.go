package yiuBuShou

import yiuAll "github.com/fidelyiu/yiu-go-han-zi/yiu_all"

// GetByRune 根据rune获取部首
func GetByRune(r rune) string {
	return yiuAll.YiuBuShouGetByRune(r)
}
