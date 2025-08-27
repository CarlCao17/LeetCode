package main

type Singleton struct {
}

var instance *Singleton

// 方法一：通过 init 尽早初始化单例
func init() {
	instance = &Singleton{}
}

func GetSingleton() *Singleton {
	return instance
}

// 方法二：使用 sync.Once 确保单例只被创建一次
