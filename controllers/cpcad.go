package controllers

import (
	"github.com/astaxie/beego"
	"beedemo2/models"
	//"encoding/base64"
	//"apiutil/utils"
)

type CpcAdController struct {
	beego.Controller
}

//@router / [get]
func (c *CpcAdController)GetAll() {
	//appId := c.GetString("_appid")
	res,err:=models.GetCpcResult()
	if (err != nil) {
		c.Data["json"] = "err"
	} else {
		c.Data["json"] = res
	}
	c.ServeJSON()
	//base64.StdEncoding("")
	//utils.CreateUUID()
}

