package dal

import (
	"database/sql"
)

type DbTransaction struct {
	transaction *sql.Tx
}

type ExecuteFun func(tran *DbTransaction) (isCommit bool, errMsg error)

// 以事务方式执行
// connectionName:连接名
// executeFun:执行函数
// 返回值：
// error：执行的错误信息
func ExecuteByTran(connectionName string, executeFun ExecuteFun) error {

	// 获取连接,并创建事务
	connection := GetConnection(connectionName)
	db := connection.GetDb()
	tran, errMsg := db.Begin()
	if errMsg != nil {
		return error
	}
	defer tran.Rollback()

	tranInfo := DbTransaction{
		transaction: tran,
	}

	// 执行
	var isCommit bool = false
	isCommit, errMsg = executeFun(&tranInfo)
	if isCommit {
		tranInfo.transaction.Commit()
	}

	return nil
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
