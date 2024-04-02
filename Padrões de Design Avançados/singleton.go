package singleton

import "sync"

type singleton struct {
	data string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{"some data"}
	})
	return instance
}
