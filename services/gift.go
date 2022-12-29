package services

import (
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
)

type Gift interface {
	GetAll() []models.Gift
	Count() int64
	Get(id int64) *models.Gift
	Delete(id int64) error
	Update(gift *models.Gift) error
	Create(gift *models.Gift) error
}

type gift struct {
	dao *dao.Gift
}

func NewGift() Gift {
	return &gift{
		dao: dao.NewGiftDao(datasource.New().DB),
	}
}

func (g *gift) Get(id int64) *models.Gift {
	return g.dao.Get(id)
}

func (g *gift) GetAll() []models.Gift {
	return g.dao.GetAll()
}

func (g *gift) Count() int64 {
	return g.dao.Count()
}

func (g *gift) Delete(id int64) error {
	return g.dao.Delete(id)
}

func (g *gift) Update(gift *models.Gift) error {
	return g.dao.Update(gift)
}

func (g *gift) Create(gift *models.Gift) error {
	return g.dao.Create(gift)
}
