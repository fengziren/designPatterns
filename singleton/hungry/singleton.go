package hungry

// singleton 未导出，外部无法初始化。
type singleton struct {
	Value int
}

var instance *singleton

func init() {
	instance = new(singleton)
}

// ------------or-----------
// var instance *singleton = new(singleton)

// ------------or-----------

func GetInstance() *singleton {
	return instance
}
