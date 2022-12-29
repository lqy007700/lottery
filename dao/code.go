package dao

import (
	"gorm.io/gorm"
	"lottery/models"
)

type Code struct {
	db *gorm.DB
}

func NewCodeDao(db *gorm.DB) *Code {
	return &Code{
		db: db,
	}
}

func (g *Code) Get(id int64) *models.Code {
	code := &models.Code{ID: id}
	g.db.First(code)
	return code
}

func (g *Code) GetAll() []models.Code {
	codeList := make([]models.Code, 0)
	g.db.Find(&codeList)
	return codeList
}

func (g *Code) Count() int64 {
	var count int64
	g.db.Model(&models.Code{}).Count(&count)
	return count
}

func (g *Code) Delete(id int64) error {
	res := g.db.Delete(&models.Code{}, id)
	return res.Error
}

func (g *Code) Update(code *models.Code) error {
	res := g.db.Save(code)
	return res.Error
}

func (g *Code) Create(code *models.Code) error {
	res := g.db.Create(code)
	return res.Error
}
