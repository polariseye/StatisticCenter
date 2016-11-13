package model

import (
	"github.com/polariseye/goutil/initutil"
)

// 动态信息结构体
type DanymicStruct struct {
}

var (

	// 动态对象
	danymicInstance *DanymicStruct

	// 动态初始化对象
	danymicInitManager *initutil.InitManager
)

// 初始化动态对象数据
func (this *DanymicStruct) initDanymic() {

}

// 初始化
func init() {

	// 初始化初始对象
	danymicInstance = new(DanymicStruct)
	danymicInitManager = new(initutil.InitManager)

	// 初始化动态对象
	danymicInstance.initDanymic()
}

// 获取动态初始化对象
// *initutil.InitManager:初始化管理对象
func GetDanymicInitManager() *initutil.InitManager {
	return danymicInitManager
}

// 获取一个新对象
func CreateDanymic() *DanymicStruct {
	return new(DanymicStruct)
}

// 获取动态对象
// 返回值：
// *DanymicStruct:动态对象
func GetDanymic() *DanymicStruct {
	return danymicInstance
}

// 设置动态对象值
// 返回值：
// newDanymic:新的动态对象
func SetDanymic(newDanymic *DanymicStruct) {
	danymicInstance = newDanymic
}
