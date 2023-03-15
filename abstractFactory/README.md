# 抽象工厂设计模式

本文件夹的代码抄袭自 https://refactoringguru.cn/design-patterns/abstract-factory/go/example ，但进行了分文件夹处理

不分文件夹，会有以下两个问题：
1. 不分文件夹，按原来方式保存的代码，在 VScode 下会出现不在 main.go 文件中的方法找不到

2. 代码放同文件夹，不能体现封装设计（或代码层次不明）

分文件后，代码也出现了报错。表面上是找不到某方法，其实是 golang 的包可见性规则引起的问题：标识小写开头，则对包外是不可见的。  
知道问题出在哪，那就把要改成大写的都改吧。
```
type Shoe struct {
	Logo string // 原本是 logo , 以下同样原本是小写
	Size int
}

type Shirt struct {
	Logo string
	Size int
}
```
很明显，这样改就对 Shoe 与 Shirt 白封装了。不行，要再改为小写，并把错误解决。(修改方法看代码)