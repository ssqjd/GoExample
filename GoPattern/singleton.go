package gopattern

import (
	"sync"
)

//单例模式

type singleton struct {
}

var ins *singleton
var once sync.Once

func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
