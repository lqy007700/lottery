package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"lottery/services"
)

type AdminController struct {
	Ctx            iris.Context
	ServiceUser    services.User
	ServiceGift    services.Gift
	ServiceBlackIp services.Blackip
	ServiceCode    services.Code
	ServiceUserDay services.UserDay
	ServiceGiftLog services.GiftLog
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":   "后台",
			"Channel": "",
		},
		Layout: "admin/layout.html",
	}
}