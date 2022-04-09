package common

import (
	"github.com/panjf2000/ants"
)

//创建线程池，threadFunc为要执行的线程函数，maxPoolSize为最大线程籹
func ThreadPoolCreate(maxPoolSize int, threadFunc func(i interface{})) (*ants.PoolWithFunc, error) {
	threadPool, err := ants.NewPoolWithFunc(maxPoolSize, func(i interface{}) {
		threadFunc(i)
	})
	return threadPool, err
}
