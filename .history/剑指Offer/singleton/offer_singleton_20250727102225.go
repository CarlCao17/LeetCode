package main

import (
	"sync"

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
var once sync.Once

func GetSingletonOnce() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

// 单线程：不推荐
func GetSingleton3() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

func GetSingleton4() *Singleton{
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &Singleton{}
		}
		lock.Unlock()
	}
	return instance
}
