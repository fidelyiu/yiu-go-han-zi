package yiuHanZiGen

import yiuGen "github.com/fidelyiu/yiu-go-tool/yiu_gen"

func GetModuleList() []yiuGen.YiuModule {
	return []yiuGen.YiuModule{
		{
			PackageName: "yiuHanZi",
			DirName:     "han_zi",
			MethodModule: []yiuGen.YiuMethodModule{
				{
					Type:     yiuGen.Is,
					FileName: "han_zi_is.go",
				},
				{
					Type:     yiuGen.Get,
					FileName: "han_zi_get.go",
				},
				{
					Type:     yiuGen.To,
					FileName: "han_zi_to.go",
				},
			},
		},
		{
			PackageName: "yiuHanZiList",
			DirName:     "han_zi_list",
			MethodModule: []yiuGen.YiuMethodModule{
				{
					Type:     yiuGen.Get,
					FileName: "han_zi_list_get.go",
				},
				{
					Type:     yiuGen.Op,
					FileName: "han_zi_list_op.go",
				},
			},
		},
		{
			PackageName: "yiuPinYin",
			DirName:     "pin_yin",
			MethodModule: []yiuGen.YiuMethodModule{
				{
					Type:     yiuGen.Get,
					FileName: "pin_yin_get.go",
				},
				{
					Type:     yiuGen.Is,
					FileName: "pin_yin_is.go",
				},
				{
					Type:     yiuGen.To,
					FileName: "pin_yin_to.go",
				},
			},
		},
		{
			PackageName: "yiuBiHua",
			DirName:     "bi_hua",
			MethodModule: []yiuGen.YiuMethodModule{
				{
					Type:     yiuGen.Get,
					FileName: "bi_hua_get.go",
				},
			},
		},
		{
			PackageName: "yiuBuShou",
			DirName:     "bu_shou",
			MethodModule: []yiuGen.YiuMethodModule{
				{
					Type:     yiuGen.Get,
					FileName: "bu_shou_get.go",
				},
			},
		},
	}
}
