package models

type Gift struct {
	ID           int64  `gorm:"column:id" json:"id"`
	Title        string `gorm:"column:title" json:"title"`
	PrizeNum     int64  `gorm:"column:prize_num" json:"prize_num"`         //  奖品数量 0 无限, >0 限量, <0 无奖品
	LeftNum      int64  `gorm:"column:left_num" json:"left_num"`           //  剩余奖品数量
	PrizeCode    string `gorm:"column:prize_code" json:"prize_code"`       //  0-9999 表示100%, 0-0表示万分之一
	PrizeTime    int64  `gorm:"column:prize_time" json:"prize_time"`       //  发奖周期 天
	Img          string `gorm:"column:img" json:"img"`                     //
	DisplayOrder int64  `gorm:"column:display_order" json:"display_order"` //  排序
	Gtype        int64  `gorm:"column:gtype" json:"gtype"`                 //  奖品类型 1 虚拟币，2虚拟券，3 实物小奖，4 实物大奖
	Gdata        string `gorm:"column:gdata" json:"gdata"`                 //  拓展数据
	Begin        int64  `gorm:"column:begin" json:"begin"`                 //
	End          int64  `gorm:"column:end" json:"end"`                     //
	PrizeDate    string `gorm:"column:prize_date" json:"prize_date"`       //  发奖计划
	PrizeBegin   int64  `gorm:"column:prize_begin" json:"prize_begin"`     //  发奖周期开始
	PrizeEnd     int64  `gorm:"column:prize_end" json:"prize_end"`         //  周期结束
	SysStatus    int64  `gorm:"column:sys_status" json:"sys_status"`       //  0正常 1删除
	SysCreate    int64  `gorm:"column:sys_create" json:"sys_create"`
	SysUpdate    int64  `gorm:"column:sys_update" json:"sys_update"`
	SysIp        string `gorm:"column:sys_ip" json:"sys_ip"`
}

func (Gift) TableName() string {
	return "gift"
}
