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
	Model
	Num      uint   `gorm:"AUTO_INCREMENT"`
	UserID   string `gorm:"not null;unique;type:varchar(32);index:idx_uinfo_userid"`
	IsSynced bool
}

// ProblemInfo 首次提问的问题信息
type ProblemInfo struct {
	Model
	UserID         string `gorm:"type:varchar(32);index:idx_pinfo_userid"`
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
	Model
	UserID    string `gorm:"type:varchar(32)"`
	Content   string `gorm:"type:text"`
	ProblemID int    `gorm:"index:idx_appendpinfo_pid"`
}

// AssessProblemInfo 问题评价
type AssessProblemInfo struct {
	Model
	UserID        string `gorm:"type:varchar(32);index:idx_assesspinfo_userid"`
	ProblemID     int    `gorm:"index:idx_assesspinfo_pid"`
	AssessInfo    string `gorm:"type:text"`
	AssessContent string `gorm:"type:text"`
}
