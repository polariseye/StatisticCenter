package model

import (
	"github.com/polariseye/goutil/initutil"
)

// 全局信息结构体
type GlobalStruct struct {
}

// 初始化
func init() {
	globalInstance = new(GlobalStruct)
	globalInitManager = new(initutil.InitManager)
}

var (

	// 全局实例对象
	globalInstance *GlobalStruct

	// 全局初始化对象
	globalInitManager *initutil.InitManager
)

// 获取全局初始化对象
// *initutil.InitManager:初始化管理对象
func GetGlobalInitManager() *initutil.InitManager {
	return globalInitManager
}

// 获取全局实例对象
// 返回值：
// *GlobalStruct:全局对象
func GetGlobalInstance() *GlobalStruct {
	return globalInstance
}
