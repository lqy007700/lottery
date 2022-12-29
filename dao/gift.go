package dao

import (
	"gorm.io/gorm"
	"lottery/models"
)

type Gift struct {
	db *gorm.DB
}

func NewGiftDao(db *gorm.DB) *Gift {
	return &Gift{
		db: db,
	}
}

func (g *Gift) Get(id int64) *models.Gift {
	gift := &models.Gift{ID: id}
	g.db.First(gift)
	return gift
}

func (g *Gift) GetAll() []models.Gift {
	giftList := make([]models.Gift, 0)
	g.db.Order("display_order asc").Find(&giftList)
	return giftList
}

func (g *Gift) Count() int64 {
	var count int64
	g.db.Model(&models.Gift{}).Count(&count)
	return count
}

func (g *Gift) Delete(id int64) error {
	res := g.db.Delete(&models.Gift{}, id)
	return res.Error
}

func (g *Gift) Update(gift *models.Gift) error {
	res := g.db.Save(gift)
	return res.Error
}

func (g *Gift) Create(gift *models.Gift) error {
	res := g.db.Create(gift)
	return res.Error
}
