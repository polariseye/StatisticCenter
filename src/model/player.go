package model

import (
	"github.com/polariseye/goutil/initutil"
)

// 玩家对象
type playerStruct struct {
}

// 初始化
func init() {
	playerInstance = new(playerStruct)
	playerInitManager = new(initutil.InitManager)
}

var (

	// 模型实例对象
	playerInstance *playerStruct

	// 玩家初始化对象
	playerInitManager *initutil.InitManager
)

// 获取玩家初始化对象
// *initutil.InitManager:初始化管理对象
func GetPlayerInitManager() *initutil.InitManager {
	return playerInitManager
}

// 获取玩家实例对象
// 返回值：
// *playerStruct:模型对象
func GetPlayerInstance() *playerStruct {
	return playerInstance
}
