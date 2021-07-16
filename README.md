## 1.YiuGo工具类

`yiu-go-tool`的汉字库

## 2.安装

```bash
go get -u github.com/fidelyiu/yiu-go-han_zi
```

如 goproxy.cn 更新失败，尝试指定版本

```bash
go get github.com/fidelyiu/yiu-go-han_zi@v1.0.1
```

## 3.使用

使用方式和`yiu-go-tool`一样，引入后第一次编译会很慢。

## 4.目前支持

| 对象类型    | 工具包名       | 说明                                      |
| ----------- | -------------- | ----------------------------------------- |
| `汉字`      | `yiuHanZi`     | 字符串和汉字相关的方法                    |
| `汉字_list` | `yiuHanZiList` | 字符串数组和汉字相关方法                  |
| `拼音`      | `yiuPinYin`    | 汉字字符串转拼音相关方法                  |
| `笔画`      | `yiuBiHua`     | 汉字字符串转笔画相关方法                  |
| `部首`      | `yiuBuShou`    | 汉字字符串转部首相关方法                  |
