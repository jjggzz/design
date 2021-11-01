package singleton

type Sing struct {
	Name string
	Num  int
}

var instance *Sing

// 利用golang中init方法只会在导入包的时候执行一次的特性初始化实例
// 注意保存实例的地址而不要保存实例,否则在获取时,会将实例拷贝,不是单例模式了
func init() {
	instance = &Sing{
		Name: "singleton instance",
		Num:  0,
	}
}

func GetInstance() *Sing {
	return instance
}
