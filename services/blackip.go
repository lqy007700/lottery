package services

import (
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
)

type Blackip interface {
	GetByIp(id string) *models.Blackip
	GetAll() []models.Blackip
	Count() int64
	Create(user *models.Blackip) error
	Update(user *models.Blackip) error
}

type blackip struct {
	dao *dao.BlackIp
}

func NewBlackIpSrv() Blackip {
	return &blackip{
		dao: dao.NewBlackIpDao(datasource.New().DB),
	}
}

func (u *blackip) GetByIp(id string) *models.Blackip {
	return u.dao.GetByIp(id)
}

func (u *blackip) GetAll() []models.Blackip {
	return u.dao.GetAll()
}

func (u *blackip) Count() int64 {
	return u.dao.Count()
}

func (u *blackip) Create(user *models.Blackip) error {
	return u.dao.Create(user)
}

func (u *blackip) Update(user *models.Blackip) error {
	return u.dao.Update(user)
}
