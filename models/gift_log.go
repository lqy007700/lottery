package models

type GiftLog struct {
	ID        int64  `gorm:"column:id" json:"id"`
	GiftId    int64  `gorm:"column:gift_id" json:"gift_id"`
	GiftName  string `gorm:"column:gift_name" json:"gift_name"`
	GiftType  int64  `gorm:"column:gift_type" json:"gift_type"`
	Uid       int64  `gorm:"column:uid" json:"uid"`
	Username  string `gorm:"column:username" json:"username"`
	PrizeCode int64  `gorm:"column:prize_code" json:"prize_code"`
	GiftData  string `gorm:"column:gift_data" json:"gift_data"`
	SysStatus int64  `gorm:"column:sys_status" json:"sys_status"` //  0正常 1删除 2作废
	SysCreate int64  `gorm:"column:sys_create" json:"sys_create"`
	SysIp     string `gorm:"column:sys_ip" json:"sys_ip"`
}

func (GiftLog) TableName() string {
	return "gift_log"
}
