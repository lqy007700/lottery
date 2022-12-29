package dao

import (
	"gorm.io/gorm"
	"lottery/models"
)

type User struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (g *User) Get(id int64) *models.User {
	user := models.User{}
	g.db.Where("id = ?", id).First(&user)
	return &user
}

func (g *User) GetAll(page, size int) []models.User {
	offset := (page - 1) * size

	userList := make([]models.User, 0)
	g.db.Limit(size).Offset(offset).Find(&userList)
	return userList
}

func (g *User) Count() int64 {
	var count int64
	g.db.Model(&models.User{}).Count(&count)
	return count
}

func (g *User) Update(user *models.User) error {
	res := g.db.Save(user)
	return res.Error
}

func (g *User) Create(user *models.User) error {
	res := g.db.Create(user)
	return res.Error
}
