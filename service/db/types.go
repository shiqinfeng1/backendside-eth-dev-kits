package db

import "time"

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, which could be embedded in your models
//    type User struct {
//      common.Model
//    }
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AccountInfo 账户信息
type AccountInfo struct {
	Model
	UserID  string `gorm:"not null;type:varchar(32);index:idx_ainfo_userid"`
	Path    string `gorm:"type:varchar(256)"`
	Address string `gorm:"not null;unique;type:varchar(256)"`
}
