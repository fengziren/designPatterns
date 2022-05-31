package lazy

// singleton 未导出，外部无法初始化。
type singleton struct {
	Value int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
