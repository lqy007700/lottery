package services

import (
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
)

type Code interface {
	Get(id int64) *models.Code
	GetAll() []models.Code
	Count() int64
	Create(user *models.Code) error
	Update(user *models.Code) error
}

type code struct {
	dao *dao.Code
}

func NewCodeSrv() Code {
	return &code{
		dao: dao.NewCodeDao(datasource.New().DB),
	}
}

func (u *code) Get(id int64) *models.Code {
	return u.dao.Get(id)
}

func (u *code) GetAll() []models.Code {
	return u.dao.GetAll()
}

func (u *code) Count() int64 {
	return u.dao.Count()
}

func (u *code) Create(user *models.Code) error {
	return u.dao.Create(user)
}

func (u *code) Update(user *models.Code) error {
	return u.dao.Update(user)
}