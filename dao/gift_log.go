package dao

import (
	"gorm.io/gorm"
	"lottery/models"
)

type GiftLog struct {
	db *gorm.DB
}

func NewGiftLogDao(db *gorm.DB) *GiftLog {
	return &GiftLog{
		db: db,
	}
}

func (g *GiftLog) Get(id int64) *models.GiftLog {
	gift := &models.GiftLog{ID: id}
	g.db.First(gift)
	return gift
}

func (g *GiftLog) GetAll() []models.GiftLog {
	giftList := make([]models.GiftLog, 0)
	g.db.Order("display_order asc").Find(&giftList)
	return giftList
}

func (g *GiftLog) Count() int64 {
	var count int64
	g.db.Model(&models.Gift{}).Count(&count)
	return count
}

func (g *GiftLog) Create(gift *models.GiftLog) error {
	res := g.db.Create(gift)
	return res.Error
}