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

// PendingTransactionInfo 交易信息
type PendingTransactionInfo struct {
	Model
	UserID          string `gorm:"not null;type:varchar(32);index:idx_ptinfo_userid"`
	From            string `gorm:"not null;type:varchar(42)"`
	To              string `gorm:"not null;type:varchar(42)"`
	TxHash          string `gorm:"not null;unique;type:varchar(66)"`
	ChainType       string `gorm:"not null;type:varchar(32)"`
	Nonce           uint64 `gorm:"not null"`
	Mined           bool   `gorm:"not null"`
	ListenTimeout   bool   `gorm:"not null"`
	ListenTimeoutAt time.Time
}
