package controller

import (
	"github.com/kataras/iris/v12"
	"lottery/common"
	"lottery/models"
	"lottery/services"
)

type IndexController struct {
	Ctx            iris.Context
	ServiceUser    services.User
	ServiceGift    services.Gift
	ServiceBlackIp services.Blackip
	ServiceCode    services.Code
	ServiceUserDay services.UserDay
	ServiceGiftLog services.GiftLog
}

func (c *IndexController) Get() string {
	c.Ctx.Header("Content-type", "text/html")
	return "<a href='/public/index.html'>开始抽奖</a>"
}

func (c *IndexController) GetGifts() map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = 0
	res["msg"] = ""

	all := c.ServiceGift.GetAll()
	list := make([]models.Gift, 0)

	for _, gift := range all {
		if gift.SysStatus == 0 {
			list = append(list, gift)
		}
	}

	res["gifts"] = list
	return res
}

func (c *IndexController) GetNewPrize() map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = 0
	res["msg"] = ""

	res["list"] = c.ServiceGiftLog.GetAll()
	return res
}

func (c *IndexController) GetLogin() {
	uid := common.Random(10000)
	loginUser := &models.LoginUser{
		Uid:      int64(uid),
		Username: "zhang",
		Now:      int64(common.NowUnix()),
		Ip:       common.ClientIp(c.Ctx.Request()),
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html?from=login"
	}

	common.SetLoginUser(c.Ctx.ResponseWriter(), loginUser)
	common.Redirect(c.Ctx.ResponseWriter(), refer)
}

func (c *IndexController) GetLogout() {
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html?from=logout"
	}
	common.SetLoginUser(c.Ctx.ResponseWriter(), nil)
	common.Redirect(c.Ctx.ResponseWriter(), refer)
}
