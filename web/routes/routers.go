package routes

import (
	"github.com/kataras/iris/v12/_examples/mvc/login/web/middleware"
	"github.com/kataras/iris/v12/mvc"
	"lottery/bootstrap"
	"lottery/services"
	"lottery/web/controller"
)

func Configure(b *bootstrap.Bootstrapper) {
	userSrv := services.NewUserSrv()
	userDaySrv := services.NewUserDaySrv()
	giftLogSrv := services.NewGiftLogSrv()
	giftSrv := services.NewGift()
	codeSrv := services.NewCodeSrv()

	index := mvc.New(b.Party("/"))
	index.Register(userSrv, userDaySrv, giftSrv, giftLogSrv, codeSrv)
	index.Handle(new(controller.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(userSrv, userDaySrv, giftSrv, giftLogSrv, codeSrv)
	admin.Handle(new(controller.AdminController))

	adminGift := admin.Party("/gift")
	adminGift.Router.Use(middleware.BasicAuth)
	adminGift.Register(giftSrv)
	adminGift.Handle(new(controller.AdminGiftController))
}