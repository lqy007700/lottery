package services

import (
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
)

type UserDay interface {
	Get(id int64) *models.UserDay
	GetAll(page, size int) []models.UserDay
	Count() int64
	Create(user *models.UserDay) error
	Update(user *models.UserDay) error
}

type userDay struct {
	dao *dao.UserDay
}

func NewUserDaySrv() UserDay {
	return &userDay{
		dao: dao.NewUserDayDao(datasource.New().DB),
	}
}

func (u *userDay) Get(id int64) *models.UserDay {
	return u.dao.Get(id)
}

func (u *userDay) GetAll(page, size int) []models.UserDay {
	return u.dao.GetAll(page, size)
}

func (u *userDay) Count() int64 {
	return u.dao.Count()
}

func (u *userDay) Create(user *models.UserDay) error {
	return u.dao.Create(user)
}

func (u *userDay) Update(user *models.UserDay) error {
	return u.dao.Update(user)
}