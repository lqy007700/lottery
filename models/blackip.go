package models

type Blackip struct {
	ID        int64  `gorm:"column:id" json:"id"`
	Ip        string `gorm:"column:ip" json:"ip"`
	Blacktime int64  `gorm:"column:blacktime" json:"blacktime"`
	SysCreate int64  `gorm:"column:sys_create" json:"sys_create"`
}

func (Blackip) TableName() string {
	return "blackip"
}
