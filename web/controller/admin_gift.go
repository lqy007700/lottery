package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"lottery/common"
	"lottery/models"
	"lottery/services"
	"lottery/web/viewmodels"
	"time"
)

type AdminGiftController struct {
	Ctx            iris.Context
	ServiceUser    services.User
	ServiceGift    services.Gift
	ServiceBlackIp services.Blackip
	ServiceCode    services.Code
	ServiceUserDay services.UserDay
	ServiceGiftLog services.GiftLog
}

func (c *AdminGiftController) Get() mvc.Result {
	list := c.ServiceGift.GetAll()
	total := len(list)

	for i, gift := range list {
		prizedata := make([][2]int, 0)

		err := json.Unmarshal([]byte(gift.PrizeDate), &prizedata)
		if err != nil || prizedata == nil || len(prizedata) < 1 {
			list[i].PrizeDate = "[]"
		} else {
			newpd := make([]string, len(prizedata))
			for index, pd := range prizedata {
				ct := common.FormatFromUnixTime(int64(pd[0]))
				newpd[index] = fmt.Sprintf("【%s】: %d", ct, pd[1])
			}
			str, err := json.Marshal(newpd)
			if err == nil && len(str) > 0 {
				list[i].PrizeDate = string(str)
			} else {
				list[i].PrizeDate = "[]"
			}
		}
		// 奖品当前的奖品池数量
		//num := utils.GetGiftPoolNum(gift.ID)
		//list[i].Title = fmt.Sprintf("【%d】%s", num, list[i].Title)

	}

	return mvc.View{
		Name: "admin/gift.html",
		Data: iris.Map{
			"Title":    "后台",
			"Channel":  "gift",
			"Datalist": list,
			"Total":    total,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminGiftController) GetEdit() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	giftInfo := viewmodels.ViewGift{}
	if id > 0 {
		data := c.ServiceGift.Get(int64(id))
		if data != nil {
			giftInfo.Id = int(data.ID)
			giftInfo.Title = data.Title
			giftInfo.PrizeNum = int(data.PrizeNum)
			giftInfo.PrizeCode = data.PrizeCode
			giftInfo.PrizeTime = int(data.PrizeTime)
			giftInfo.Img = data.Img
			giftInfo.Displayorder = int(data.DisplayOrder)
			giftInfo.Gtype = int(data.Gtype)
			giftInfo.Gdata = data.Gdata
			giftInfo.TimeBegin = common.FormatFromUnixTime(data.Begin)
			giftInfo.TimeEnd = common.FormatFromUnixTime(data.End)
		}
	}
	return mvc.View{
		Name: "admin/giftEdit.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "gift",
			"info":    giftInfo,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminGiftController) PostSave() mvc.Result {
	data := viewmodels.ViewGift{}
	err := c.Ctx.ReadForm(&data)
	if err != nil {
		fmt.Println("admin_gift.PostSave ReadForm error=", err)
		return mvc.Response{
			Text: fmt.Sprintf("ReadForm转换异常, err=%s", err),
		}
	}
	giftInfo := models.Gift{}
	giftInfo.ID = int64(data.Id)
	giftInfo.Title = data.Title
	giftInfo.PrizeNum = int64(data.PrizeNum)
	giftInfo.PrizeCode = data.PrizeCode
	giftInfo.PrizeTime = int64(data.PrizeTime)
	giftInfo.Img = data.Img
	giftInfo.DisplayOrder = int64(data.Displayorder)
	giftInfo.Gtype = int64(data.Gtype)
	giftInfo.Gdata = data.Gdata
	t1, err1 := common.ParseTime(data.TimeBegin)
	t2, err2 := common.ParseTime(data.TimeEnd)
	if err1 != nil || err2 != nil {
		return mvc.Response{
			Text: fmt.Sprintf("开始时间、结束时间的格式不正确, err1=%s, err2=%s", err1, err2),
		}
	}
	giftInfo.Begin = t1.Unix()
	giftInfo.End = t2.Unix()
	if giftInfo.ID > 0 {
		datainfo := c.ServiceGift.Get(giftInfo.ID)
		if datainfo != nil {
			giftInfo.SysUpdate = int64(time.Now().Unix())
			giftInfo.SysIp = common.ClientIp(c.Ctx.Request())
			// 对比修改的内容项
			if datainfo.PrizeNum != giftInfo.PrizeNum {
				// 奖品总数量发生了改变
				giftInfo.LeftNum = datainfo.LeftNum - (datainfo.PrizeNum - giftInfo.PrizeNum)
				if giftInfo.LeftNum < 0 || giftInfo.PrizeNum <= 0 {
					giftInfo.LeftNum = 0
				}
				giftInfo.SysStatus = datainfo.SysStatus
				//utils.ResetGiftPrizeData(&giftInfo, c.ServiceGift)
			}
			if datainfo.PrizeTime != giftInfo.PrizeTime {
				// 发奖周期发生了变化
				//utils.ResetGiftPrizeData(&giftInfo, c.ServiceGift)
			}
			c.ServiceGift.Update(&giftInfo)
		} else {
			giftInfo.ID = 0
		}
	}
	if int(giftInfo.ID) <= 0 {
		giftInfo.LeftNum = giftInfo.PrizeNum
		giftInfo.SysIp = common.ClientIp(c.Ctx.Request())
		giftInfo.SysCreate = time.Now().Unix()
		c.ServiceGift.Create(&giftInfo)
		// 更新奖品的发奖计划
		//utils.ResetGiftPrizeData(&giftInfo, c.ServiceGift)
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}

func (c *AdminGiftController) GetDelete() mvc.Result {
	return mvc.Response{
		Path: "/admin/gift",
	}
}

func (c *AdminGiftController) GetReset() mvc.Result {
	return mvc.Response{
		Path: "/admin/gift",
	}
}
