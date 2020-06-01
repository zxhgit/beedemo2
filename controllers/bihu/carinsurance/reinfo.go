package carinsurance

import (
	"beedemo2/models/bihu/carinsurance"
	"github.com/astaxie/beego"
)

type ReInfoController struct {
	beego.Controller
}

// @Title 1获取用户车辆信息和去年投保信息
// @Param LicenseNo 			query string 	true 	"车牌号（字母全部大写）"
// @Param CityCode 				query int 		true 	"参考文档上部城市列表"
// @Param CarVin 				query string 			"车架号（字母全部大写）一般次新车多用这俩参数，此时车牌号非必填 "
// @Param EngineNo 				query string 			"发动机号（字母全部大写）一般次新车多用这俩参数，此时车牌号非必填 "
// @Param Group 				query int 		true 	"为了向下兼容，固定赋值1 "
// @Param CanShowNo 			query int 				"是否展示商业/交强投保单号0：(默认)否 ，1:是"
// @Param CanShowExhaustScale 	query int 				"是否展示排量信息 0：（默认）否  1：是"
// @Param ShowXiuLiChangType 	query int 				"是否展示修理厂类型0（默认）否  1：是"
// @Param TimeFormat 			query int 				"按照实时起保返回到期时间（商业/交强）0：（默认）否 1：是"
// @Param ShowAutoMoldCode 		query int 				"是否展示精友码：0：（默认）否  1：是"
// @Param ShowFybc 				query int 				"修理期间费用补偿险：0：（默认）否  1：是"
// @Param ShowSheBei 			query int 				"新增设备险种：0：（默认）否  1：是"
// @Param RenewalCarType 		query int 				"大小号牌：0小车，1大车，默认0"
// @Param ShowOrg 				query int 				"是否展示机构名称 0：否(默认) 1:是"
// @Param SixDigitsAfterIdCard 	query string 			"车主证件号后六位。非必填。平安系统续保指定字段"
// @Param ShowSanZheJieJiaRi 	query int 				"是否展示三责险附加法定节假日限额翻倍险0：（默认）否  1：是"
// @Param RenewalSource 		query int 				"指定保险公司续保（如果清楚是哪家保司，就将该值赋值，查询速度会有相应的提高）"
// @Param ShowBaoFei 			query int 				"是否展示上年保费，0否（默认）1是。（仅支持人保、太平洋，且人保营销系统无法获取保费）"
// @router / [get]
func (c *ReInfoController) Get() {
	var err error
	inputModel := carinsurance.InputModel{}
	if err = c.ParseForm(&inputModel); err != nil {
		c.Data["json"] = err.Error()
	}
	if err = inputModel.CheckInput(); err != nil {
		c.Data["json"] = err.Error()
	}
	if err != nil {
		c.ServeJSON()
		return
	}
	res1, err := carinsurance.GetReInfo(&inputModel)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = res1
	}
	c.ServeJSON()
}
