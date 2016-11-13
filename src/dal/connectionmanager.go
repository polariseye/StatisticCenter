package dal

import (
	"errors"
	"sync"
)

// 连接信息获取函数
// connectionName:连接名称
// 返回值：
// driverName:驱动名
// connectionString:数据库字符串
type connectionFun func(connectionName string) (driverName string, connectionString string)

var (
	// 连接对象集合
	connectionDic map[string]*DbConnection

	// 数据库集合锁对象
	connectionLock sync.RWMutex
)

const (

	// 游戏model库
	Con_Db_GameModel string = "GameModel"

	// 游戏数据库
	Con_Db_Game string = "GameDb"
)

func init() {

	// 初始化集合对象
	connectionDic = make(map[string]*DbConnection, 2)
}

// 初始化数据库连接
func InitDb(getConnection connectionFun) {

	// 初始化game库
	if errMsg := addConnection(Con_Db_Game, getConnection); errMsg != nil {
		panic(errors.New("数据库初始化失败" + Con_Db_Game + ":" + errMsg.Error()))
	}

	// 初始化model库
	if errMsg := addConnection(Con_Db_GameModel, getConnection); errMsg != nil {
		panic(errors.New("数据库初始化失败" + Con_Db_GameModel + ":" + errMsg.Error()))
	}
}

// 添加数据库连接
// dbName:数据库名
// getConnection:用于获取数据库连接的函数指针
// 返回值：
// error:错误数据
func addConnection(dbName string, getConnection connectionFun) error {

	// 获取连接信息
	driverName, connectionString := getConnection(dbName)

	// 创建数据库连接
	connection := DbConnection{
		Name: dbName,
	}
	errMsg := connection.Init(driverName, connectionString)
	if errMsg != nil {
		return errMsg
	}

	connectionLock.Lock()
	defer connectionLock.Unlock()

	// 添加到集合
	connectionDic[dbName] = &connection

	return nil
}

// 删除数据库连接
// dbName:数据库名
func removeConnection(dbName string) {
	connectionLock.Lock()
	defer connectionLock.Unlock()

	// 获取连接对象
	item, isFind := connectionDic[dbName]
	if isFind == false {
		return
	}

	// 删除连接
	delete(connectionDic, dbName)

	// 释放连接
	item.Dispose()
}

// 获取连接对象
// connectionName:连接名
// *DbConnection:数据库连接工具类
func GetConnection(connectionName string) *DbConnection {
	connectionLock.RLock()
	defer connectionLock.RUnlock()

	// 获取连接对象
	item, isFind := connectionDic[connectionName]
	if isFind == false {
		return nil
	}

	return item
}
