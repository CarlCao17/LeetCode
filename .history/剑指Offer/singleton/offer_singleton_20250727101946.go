package main

type Singleton struct {
}

var instance *Singleton

// 方法一：通过 init 尽早初始化
func init() {
	instance = &Singleton{}
}

func GetSingleton() *Singleton {
	return instance
}
