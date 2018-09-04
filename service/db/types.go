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
	Success         bool   `gorm:"not null"`
	MinedBlock      uint64 `gorm:"not null"`
	Comfired        int    `gorm:"not null"`
	ListenTimeout   bool   `gorm:"not null"`
	ListenTimeoutAt time.Time
	Description     string `gorm:"type:varchar(256)"`
}

// PointsInfo 积分交易信息
type PointsInfo struct {
	Model
	ChainType      string `gorm:"not null;type:varchar(32)"`
	UserID         string `gorm:"not null;type:varchar(32);index:idx_ptinfo_userid"`
	UserAddress    string `gorm:"not null;type:varchar(42)"`
	TxnType        string `gorm:"not null;type:varchar(32)"` // buy consume refund
	TxnHash        string `gorm:"not null;unique;type:varchar(66)"`
	PreBalance     uint64 `gorm:"not null"`
	ExpectBalance  uint64 `gorm:"not null"`
	IncurredAmount uint64 `gorm:"not null"` //本次交易发生的金额
	CurrentStatus  string `gorm:"not null"` //交易执行的状态
}
