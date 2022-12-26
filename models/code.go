package models

type Code struct {
	ID        int64  `gorm:"column:id" json:"id"`
	GiftId    int64  `gorm:"column:gift_id" json:"gift_id"`
	Code      string `gorm:"column:code" json:"code"`             //  编码
	SysStatus int64  `gorm:"column:sys_status" json:"sys_status"` //  0正常 1删除 2已发放
	SysCreate int64  `gorm:"column:sys_create" json:"sys_create"`
	SysUpdate int64  `gorm:"column:sys_update" json:"sys_update"`
}

func (Code) TableName() string {
	return "code"
}
