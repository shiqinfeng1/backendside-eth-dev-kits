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

// UserInfo 首次提问的问题信息
type UserInfo struct {
	Num int `gorm:"AUTO_INCREMENT"` // 自增
	Model
	UserID   string `gorm:"not null;unique;type:varchar(32) index:idx_uinfo_userid"`
	IsSynced bool
}

// ProblemInfo 首次提问的问题信息
type ProblemInfo struct {
	Num int `gorm:"AUTO_INCREMENT"` // 自增
	Model
	UserID         string `gorm:"type:varchar(32) index:idx_pinfo_userid"`
	PartnerOrderID string `gorm:"type:varchar(256)"`
	Content        string `gorm:"type:text"`
	ProblemID      int32  `gorm:"index:idx_pinfo_pid"`
	Price          int32
	DoctorID       string `gorm:"type:varchar(256)"`
	PaidType       string `gorm:"type:varchar(32)"`
	IsReply        bool   `gorm:"index"`
	IsRefund       bool   `gorm:"index"`
	IsDeleted      bool
	IsClosed       bool
}

// AppendProblemInfo 追加的问题信息
type AppendProblemInfo struct {
	Num int `gorm:"AUTO_INCREMENT"` // 自增
	Model
	Content   string `gorm:"type:text"`
	ProblemID string `gorm:"type:varchar(32) index:idx_appendpinfo_pid"`
	IsReply   bool   `gorm:"index"`
	IsDeleted bool
	IsClosed  bool
}

// AssessProblemInfo 问题评价
type AssessProblemInfo struct {
	Num int `gorm:"AUTO_INCREMENT"` // 自增
	Model
	UserID        string `gorm:"type:varchar(32) index:idx_assesspinfo_userid"`
	ProblemID     string `gorm:"type:varchar(32) index:idx_assesspinfo_pid"`
	AssessLevel   string `gorm:"type:text"`
	AssessTag     string `gorm:"type:text"`
	AssessContent string `gorm:"type:text"`
}
