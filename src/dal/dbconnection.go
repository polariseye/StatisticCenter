package dal

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
)

// 数据库连接结构体
type DbConnection struct {

	// 数据库对象
	db *sql.DB

	// 数据库连接名
	Name string
}

// 初始化数据库连接信息
// driver：驱动名
// connectionString:链接字符串
// 返回值：
// error:初始化的错误信息
func (this *DbConnection) Init(driver string, connectionString string) error {

	// 打开连接对象
	db, errMsg := sql.Open(driver, connectionString)
	if errMsg != nil {
		return errMsg
	}

	// 检查是否存在链接错误
	if errMsg = db.Ping(); errMsg != nil {
		return errMsg
	}

	this.db = db

	return nil
}

// 获取数据库连接对象
// 返回值：
// *sql.DB:数据库对象
func (this *DbConnection) GetDb() *sql.DB {
	return this.db
}

// 释放所有资源
func (this *DbConnection) Dispose() {
	if this.db != nil {
		this.db.Close()
	}
}

// 查询
// sql:待查询的sql语句
// args:参数
// 返回值：
// *sql.Rows:结果数据
// error:错误数据
func (this *DbConnection) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return this.db.Query(sql, args...)
}

// 查询并获得一条数据
// sql:待查询的sql语句
// args:参数
// 返回值：
// *sql.Row:结果数据
func (this *DbConnection) QueryRow(sql string, args ...interface{}) *sql.Row {
	return this.db.QueryRow(sql, args...)
}

// 执行一条sql语句
// sql:待执行的sql语句
// args:参数
// 返回值：
// int64:影响记录数
// error:错误信息
func (this *DbConnection) Execute(sql string, args ...interface{}) (int64, error) {

	// 执行sql
	result, errMsg := this.db.Exec(sql, args...)
	if errMsg != nil {
		return 0, errMsg
	}

	// 返回影响记录数
	return result.RowsAffected()
}

// 批量执行sql语句（内部会采用事务提交）
// sqlContent:待执行的sql语句
// paramList:sql语句对应的所有参数列表
// 返回值：
// int64：总共影响的记录集合
// error：错误信息
func (this *DbConnection) ExecuteList(sqlContent string, paramList [][]interface{}) (int64, error) {

	// 开启事务
	tran, errMsg := this.db.Begin()
	if errMsg != nil {
		return 0, errMsg
	}
	defer tran.Rollback()

	// 初始化sql语句执行
	var db *sql.Stmt
	db, errMsg = tran.Prepare(sqlContent)
	if errMsg != nil {
		return 0, errMsg
	}

	var (
		result   sql.Result
		count    int64
		tmpCount int64
	)

	// 批量提交数据
	for _, paramsItem := range paramList {

		// 执行
		result, errMsg = db.Exec(paramsItem...)
		if errMsg != nil {
			return 0, errMsg
		}

		// 获取影响数据行数
		tmpCount, errMsg = result.RowsAffected()
		if errMsg != nil {
			return 0, errMsg
		}

		// 记录影响记录数
		count += tmpCount
	}

	// 提交事务
	tran.Commit()

	return count, nil
}
