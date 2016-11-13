package model

import (
	"github.com/polariseye/goutil/initutil"
)

// 玩家对象
type modelStruct struct {
}

// 初始化
func init() {
	modelInstance = new(modelStruct)
	modelInitManager = new(initutil.InitManager)
}

var (

	// 模型实例对象
	modelInstance *modelStruct

	// 模型初始化对象
	modelInitManager *initutil.InitManager
)

// 获取模型初始化对象
// *initutil.InitManager:初始化管理对象
func GetModelInitManager() *initutil.InitManager {
	return modelInitManager
}

// 获取模型实例对象
// 返回值：
// *GlobalStruct:模型对象
func GetModelInstance() *modelStruct {
	return modelInstance
}
