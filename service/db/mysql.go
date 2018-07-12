package db

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	cmn "github.com/shiqinfeng1/chunyuyisheng/service/common"
)

// DB gorm数据库实例
var mysql *gorm.DB

// GormDB 封装后的gorm数据库实例
type GormDB struct {
	*gorm.DB
	gdbDone bool
}

// InitDB 初始化数据库
func InitMysql() {
	// var connstring string
	idb, err := gorm.Open("mysql",
		cmn.Config().GetString("mysql.URL")+"/"+
			cmn.Config().GetString("mysql.dbName")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// Then you could invoke `*sql.DB`'s functions with it
	idb.DB().SetMaxIdleConns(cmn.Config().GetInt("mysql.idle"))
	idb.DB().SetMaxOpenConns(cmn.Config().GetInt("mysql.maxOpen"))
	idb.LogMode(cmn.Config().GetBool("mysql.debug"))

	mysql = idb
}

// DBClose 关闭数据库
func MysqlClose() {
	mysql.Close()
}

// DBBegin 打开一个transaction
func MysqlBegin() *GormDB {
	txn := mysql.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	return &GormDB{txn, false}
}

// DBCommit 提交并关闭transaction
func (c *GormDB) MYsqlCommit() {
	if c.gdbDone {
		return
	}
	tx := c.Commit()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}

// DBRollback 回滚并关闭transaction
func (c *GormDB) MysqlRollback() {
	if c.gdbDone {
		return
	}
	tx := c.Rollback()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}
