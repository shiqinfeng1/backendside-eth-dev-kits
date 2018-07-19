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

// Migrate ...
func migrate() {
	idb := MysqlBegin()
	idb.AutoMigrate(&UserInfo{}, &ProblemInfo{}, &AppendProblemInfo{}, &AssessProblemInfo{})
	idb.MysqlCommit()
}

// InitMysql 初始化数据库
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

	if idb.HasTable(&UserInfo{}) == false {
		idb.CreateTable(&UserInfo{})
	}
	if idb.HasTable(&ProblemInfo{}) == false {
		idb.CreateTable(&ProblemInfo{})
	}
	if idb.HasTable(&AppendProblemInfo{}) == false {
		idb.CreateTable(&AppendProblemInfo{})
	}
	if idb.HasTable(&AssessProblemInfo{}) == false {
		idb.CreateTable(&AssessProblemInfo{})
	}
	mysql = idb
	migrate()
}

// MysqlClose 关闭数据库
func MysqlClose() {
	mysql.Close()
}

// MysqlBegin 打开一个transaction
func MysqlBegin() *GormDB {
	txn := mysql.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	return &GormDB{txn, false}
}

// MysqlCommit 提交并关闭transaction
func (c *GormDB) MysqlCommit() {
	if c.gdbDone {
		return
	}
	tx := c.Commit()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}

// MysqlRollback 回滚并关闭transaction
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
