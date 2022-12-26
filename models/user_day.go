package models

type UserDay struct {
	ID        int64 `gorm:"column:id" json:"id"`
	Uid       int64 `gorm:"column:uid" json:"uid"`
	Day       int64 `gorm:"column:day" json:"day"`
	Num       int64 `gorm:"column:num" json:"num"`
	SysCreate int64 `gorm:"column:sys_create" json:"sys_create"`
	SysUpdate int64 `gorm:"column:sys_update" json:"sys_update"`
}

func (UserDay) TableName() string {
	return "user_day"
}
