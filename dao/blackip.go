package dao

import (
	"gorm.io/gorm"
	"lottery/models"
)

type BlackIp struct {
	db *gorm.DB
}

func NewBlackIpDao(db *gorm.DB) *BlackIp {
	return &BlackIp{
		db: db,
	}
}

func (g *BlackIp) GetByIp(ip string) *models.Blackip {
	blackIp := models.Blackip{}
	g.db.Where("ip = ?", ip).First(&blackIp)
	return &blackIp
}

func (g *BlackIp) GetAll() []models.Blackip {
	blackIpList := make([]models.Blackip, 0)
	g.db.Find(blackIpList)
	return blackIpList
}

func (g *BlackIp) Count() int64 {
	var count int64
	g.db.Model(&models.Blackip{}).Count(&count)
	return count
}

func (g *BlackIp) Delete(id int64) error {
	res := g.db.Delete(&models.Blackip{}, id)
	return res.Error
}

func (g *BlackIp) Update(blackIp *models.Blackip) error {
	res := g.db.Save(blackIp)
	return res.Error
}

func (g *BlackIp) Create(blackIp *models.Blackip) error {
	res := g.db.Create(blackIp)
	return res.Error
}
