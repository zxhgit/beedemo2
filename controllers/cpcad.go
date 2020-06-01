package controllers

import (
	"beedemo2/models"
	"github.com/astaxie/beego"
	//"encoding/base64"
	//"apiutil/utils"
)

type CpcAdController struct {
	beego.Controller
}

//@router / [get]
func (c *CpcAdController) GetAll() {
	//appId := c.GetString("_appid")
	res, err := models.GetCpcResult()
	if err != nil {
		c.Data["json"] = "err"
	} else {
		res.Message += "02"
		c.Data["json"] = res
	}
	c.ServeJSON()
	//base64.StdEncoding("")
	//utils.CreateUUID()
}
