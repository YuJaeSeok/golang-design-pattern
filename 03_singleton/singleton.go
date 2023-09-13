package singleton

import "sync"

// Singleton 是单例模式接口
// 通过该接口可以避免GetInstance 返回一个包私有类型的指针
type Singleton interface {
	foo()
}

// singleton 是单例模式类，包私有的
type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
