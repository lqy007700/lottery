package models

type User struct {
	ID        int64  `gorm:"column:id" json:"id"`
	Username  string `gorm:"column:username" json:"username"`
	Blacktime int64  `gorm:"column:blacktime" json:"blacktime"`
	Name      string `gorm:"column:name" json:"name"`
	Mobile    int64  `gorm:"column:mobile" json:"mobile"`
	SysCreate int64  `gorm:"column:sys_create" json:"sys_create"`
	SysUpdate int64  `gorm:"column:sys_update" json:"sys_update"`
	SysIp     string `gorm:"column:sys_ip" json:"sys_ip"`
}

func (User) TableName() string {
	return "user"
}
