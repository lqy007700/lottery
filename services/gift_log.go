package services

import (
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
)

type GiftLog interface {
	GetAll() []models.GiftLog
	Get(id int64) *models.GiftLog
	Count() int64
	Create(log *models.GiftLog) error
}

type giftLog struct {
	dao *dao.GiftLog
}

func NewGiftLogSrv() GiftLog {
	return &giftLog{
		dao: dao.NewGiftLogDao(datasource.New().DB),
	}
}

func (g *giftLog) GetAll() []models.GiftLog {
	return g.dao.GetAll()
}

func (g *giftLog) Get(id int64) *models.GiftLog {
	return g.dao.Get(id)
}

func (g *giftLog) Count() int64 {
	return g.dao.Count()
}

func (g *giftLog) Create(log *models.GiftLog) error {
	return g.dao.Create(log)
}