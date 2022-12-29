package services

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"lottery/common"
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
	"sync"
)

var cacheUserLise = make(map[int]*models.User)
var cacheUserLock = sync.Mutex{}

const (
	cacheUserKey = "info:user:%d"
)

type User interface {
	Get(id int64) *models.User
	GetAll(page, size int) []models.User
	Count() int64
	Create(user *models.User) error
	Update(user *models.User) error
}

type user struct {
	dao *dao.User
}

func NewUserSrv() User {
	return &user{
		dao: dao.NewUserDao(datasource.New().DB),
	}
}

func (u *user) Get(id int64) *models.User {
	data := u.getByUserCache(id)
	if data == nil || data.ID <= 0 {
		data = u.dao.Get(id)
		u.setByUserCache(data)
	}
	return data
}

func (u *user) GetAll(page, size int) []models.User {
	return u.dao.GetAll(page, size)
}

func (u *user) Count() int64 {
	return u.dao.Count()
}

func (u *user) Create(user *models.User) error {
	return u.dao.Create(user)
}

func (u *user) Update(user *models.User) error {
	return u.dao.Update(user)
}

func (u *user) getByUserCache(id int64) *models.User {
	key := fmt.Sprintf(cacheUserKey, id)
	rds := datasource.GetCache()
	dataMap, err := redis.StringMap(rds.Do("HGETALL", key))
	if err != nil {
		return nil
	}

	dataId := common.GetInt64FromStringMap(dataMap, "id", 0)
	if dataId <= 0 {
		return nil
	}

	return &models.User{
		ID:        dataId,
		Username:  common.GetStringFromStringMap(dataMap, "username", ""),
		Blacktime: common.GetInt64FromStringMap(dataMap, "blacktime", 0),
		Name:      common.GetStringFromStringMap(dataMap, "name", ""),
		Mobile:    common.GetInt64FromStringMap(dataMap, "mobile", 0),
		SysCreate: common.GetInt64FromStringMap(dataMap, "syscreate", 0),
		SysUpdate: common.GetInt64FromStringMap(dataMap, "sysupdate", 0),
		SysIp:     common.GetStringFromStringMap(dataMap, "sysip", ""),
	}
}

func (u *user) setByUserCache(data *models.User) {
	if data == nil || data.ID <= 0 {
		return
	}
	key := fmt.Sprintf(cacheUserKey, data.ID)
	rds := datasource.GetCache()

	params := []interface{}{key}
	params = append(params, "id", data.ID)
	params = append(params, "username", data.Username)
	params = append(params, "blacktime", data.Blacktime)
	params = append(params, "name", data.Name)
	params = append(params, "mobile", data.Mobile)
	params = append(params, "syscreate", data.SysCreate)
	params = append(params, "sysupdate", data.SysUpdate)
	params = append(params, "sysip", data.SysIp)

	_, err := rds.Do("HMSET", params...)
	if err != nil {
		log.Println("user_service.setByCache HMSET params=", params, ", error=", err)
		return
	}
}
