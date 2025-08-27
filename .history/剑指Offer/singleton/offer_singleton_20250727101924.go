package main

type Singleton struct {
}

var instance *Singleton

func init() {
	instance = &Singleton{}
}

func GetSingleton() *Singleton {
	return instance
}
