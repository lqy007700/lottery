package dao

import (
	"gorm.io/gorm"
	"lottery/models"
)

type UserDay struct {
	db *gorm.DB
}

func NewUserDayDao(db *gorm.DB) *UserDay {
	return &UserDay{
		db: db,
	}
}

func (g *UserDay) Get(id int64) *models.UserDay {
	user := models.UserDay{}
	g.db.Where("id = ?", id).First(&user)
	return &user
}

func (g *UserDay) GetAll(page, size int) []models.UserDay {
	offset := (page - 1) * size

	userList := make([]models.UserDay, 0)
	g.db.Limit(size).Offset(offset).Find(&userList)
	return userList
}

func (g *UserDay) Count() int64 {
	var count int64
	g.db.Model(&models.UserDay{}).Count(&count)
	return count
}

func (g *UserDay) Update(user *models.UserDay) error {
	res := g.db.Save(user)
	return res.Error
}

func (g *UserDay) Create(user *models.UserDay) error {
	res := g.db.Create(user)
	return res.Error
}
