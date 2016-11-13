package main

import (
	"github.com/polariseye/statisticcenter/src/dal"
	"github.com/polariseye/statisticcenter/src/model"
)

// 配置初始化
func InitConfig() {

	//todo: 初始化其他配置

	// 初始化数据库连接
	dal.InitDb(func(connectionName string) (driverName string, connectionString string) {
		return "", ""
	})
}

// 数据初始化
func InitData() []string {

	errorList := make([]string, 100, 100)

	// 模块初始化
	if tmpErrorList := model.GetModelInitManager().StartInit(model.GetModelInstance()); len(tmpErrorList) > 0 {
		errorList = append(errorList, tmpErrorList...)
	}

	// 动态对象初始化
	if tmpErrorList := model.GetDanymicInitManager().StartInit(model.GetDanymic()); len(tmpErrorList) > 0 {
		errorList = append(errorList, tmpErrorList...)
	}

	// 玩家数据初始化
	if tmpErrorList := model.GetPlayerInitManager().StartInit(model.GetPlayerInstance()); len(tmpErrorList) > 0 {
		errorList = append(errorList, tmpErrorList...)
	}

	// 全局对象初始化
	if tmpErrorList := model.GetGlobalInitManager().StartInit(model.GetGlobalInstance()); len(tmpErrorList) > 0 {
		errorList = append(errorList, tmpErrorList...)
	}

	return errorList
}
