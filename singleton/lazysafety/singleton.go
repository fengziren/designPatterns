package lazysafety

import "sync"

// singleton 未导出，外部无法初始化。
type singleton struct {
	Value int
}

var mux sync.Mutex
var instance *singleton

func GetInstance() *singleton {
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
