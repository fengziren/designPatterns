package dcl

import "sync"

type singleton struct {
	Value int
}

var instance *singleton
var mux sync.Mutex

// GetInstance 第一次判断不加锁，第二次加锁保证线程安全，一旦对象建立后，获取对象就不用加锁了。
func GetInstance() *singleton {
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = new(singleton)
		}
	}
	return instance
}

// ------------------------------------------------------------
// sync.Once 来确保创建单例对象的方法只执行一次。
var once sync.Once

func GetInstance2() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
