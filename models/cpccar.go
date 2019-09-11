package models

import (
	"strconv"
)

type CpcResultModel struct {
	ReturnCode          int `json:"returncode"`
	Message             string `json:"message"`
	Result              *CpcResultListModel `json:"result"`
}

type CpcResultListModel struct {
	List_10033 []*CpcCarModel `json:"list_10033"`
	List_999 []*CpcCarModel `json:"list_999"`
}

type CpcCarModel struct {
	Version           string `json:"version"`
	ImgPath           string `json:"imgpath"`
	Mileage           string `json:"mileage"`
	CarLicense        string `json:"carLicense"`
	LocatedType       string `json:"locatedType"`
	SellPrice         string `json:"sellPrice"`
	ProductId         string `json:"productId"`
	Pubtime           string `json:"pubtime"`
	Type              string `json:"type"`
	Title             string `json:"title"`
	PvId              string `json:"pvid"`
	Ah_json_module_id string `json:"ah_json_module_id"`
	RdPostUrl         string `json:"rdposturl"`
	IsHavead          string `json:"ishavead"`
	Url               string `json:"url"`
	ThirdAdUrl        string `json:"thirdadurl"`
	ThirdClickUrl     []string `json:"thirdclickurl"`
}

func GetCpcResult() (r *CpcResultModel,err error) {
	r = &CpcResultModel{
		ReturnCode: 0,
		Message:    "ok",
	}
	li := &CpcResultListModel{}
	li.List_10033, li.List_999, err = getCpcCars()
	r.Result = li
	err = nil
	return
}

func getCpcCars()(list10033 []*CpcCarModel,list999 []*CpcCarModel,err error) {
	list10033 = []*CpcCarModel{}
	list999 = []*CpcCarModel{}
	typeId := 3831
	for i := 0; i < 12; i++ {
		var car *CpcCarModel
		if (i == 3 || i == 5 || i == 7 || i == 9) {
			car = &CpcCarModel{
				PvId:              "300935323335616138642D396461622D343764322D393661332D63393061653063363137306509312E3009363865393466333534373139323466326365323635373366316663333639313034396264393665630909323037340930313961663438312D356464612D346439302D626535662D6537646564353134626534660931330932313239373309320949500931",
				Ah_json_module_id: "999",
				RdPostUrl:         "https://rdx.autohome.com.cn/app/realdeliver?",
				IsHavead:          "0",
			}
			car.Type = strconv.Itoa(typeId)
			list999 = append(list999, car)
		} else {
			car = &CpcCarModel{
				Version:           "1.0.0",
				ImgPath:           "http://2sc2.autoimg.cn/escimg/g14/M04/23/38/330x220_0_q87_autohomecar__wKjByVo3dj-ABl6_AAIfXs_x5Uc879.jpg",
				Mileage:           "5.1",
				CarLicense:        "2014",
				LocatedType:       "2",
				SellPrice:         "22.50",
				ProductId:         "3164047",
				Pubtime:           "2017-11-02 14:55:27",
				Title:             "奥迪Q53(进口) 2012款 35 TFSI quattro 舒适型",
				PvId:              "300935323335616138642D396461622D343764322D393661332D63393061653063363137306509312E3009363865393466333534373139323466326365323635373366316663333639313034396264393665630909323037340930313961663438312D356464612D346439302D626535662D6537646564353134626534660931330932313239373309320949500931",
				Ah_json_module_id: "10033",
				RdPostUrl:         "https://rdx.autohome.com.cn/app/realdeliver?",
				IsHavead:          "1",
				Url:               "https://clickx.autohome.com.cn/click?t=MS4wCTUyMzVhYThkLTlkYWItNDdkMi05NmEzLWM5MGFlMGM2MTcwZQkwMTlhZjQ4MS01ZGRhLTRkOTAtYmU1Zi1lN2RlZDUxNGJlNGYJMTMJMjA3NAkxMjE2CTIxMjk3MwkJJSVDTElDS19VUkxfVU5FU0MlJSZ1cmw9aHR0cHMlM0ElMkYlMkZkc3BtbnQuYXV0b2hvbWUuY29tLmNuJTJGbW9uaXRvciUzRnJ0eXBlJTNEMiUyNmRwbWslM0Q0RGY1TUdVRkxERWJKNDA5RGFzUllsVmFLYVdyRkljdlN5YzJzJTI1MkZoYTMzY2ZTV1AyNmxqZWNCbEtiT3o0RUpnclNscGg0dTlTM1hkUEdXT2s5MW1LSTB4VmIlMjUyRmUlMjUyQlVzTjdHQmxvN2JsWjNpTkxTRGoyNjFmZUp3eFVlYUt6QkljbURFSjU4T3RaanlRYVFHcnQ3d1NLSXdOTVAlMjUyRm5xVFl3bkd4NTJwZTBFaXlZYlNXJTI1MkJpdjFTSVlBSmFPcVNpQ1lwZ0ZGcHE0dlpDanlaZUZ5aXB2a0xVWUJ4SWJQVDRUTXczWFIwcHFiNUMxR0FZUUQ3NTdnYmRkeHBQYXZub1ZJaHdUUjFwOXU5WDNTUWZIamp6N0ZuZmNocEJPYVRqVm9zaERGUjVzTFlCbWlSQkNqYmk0RUxmY2g5SmElMjUyRkxxVU41eERGUjVyYlVDaHk1TEp6U3olMjUyQkZyTWN4NUphdkR1VU41eUgxcDM0cThPaHpaSEhIbjY2MVRkZEJkVWVhT29CWTgyUnc0JTI1MkJxYjVDMUhVZFRHTHo5a0tQSmxnZEtiU3pCTXg0SDB0czl1dFZ3bUJER1NPZnFoS0hJVXRhWWZIclVONXVEQm95cElVUW5DdE5IWG42NlZEY2NRSmFPSyUyNTJCcEZMRXZRUnclMjUyQjR1QlJ3bUJORnlpMGhRYVBJVm9YS2VMZ1VNQjBHMHhqOGVKVTMzY1lUR3IzNGxUWmJnd2JONm01QzdFM1hCUjUlMjUyQnZnSW1qWmVDMkh2OVFyQUkxc01OS2kxRFl0c1RSYzI3cmtPd1RKQ0dTJTI1MkJtdFJLRGJWUWRLYSUyNTJCS0VvRWtSd3gwb2I0RXJ5RmFFUzJwcmhuUk1sZ1pLYVc3Q1lwJTI1MkZIRWhxOSUyNTJCMVMybUFDV2lpbHVRJTI1MkJBSm5FYU1xVDRXdDk3RzBwMzRya0JtaWRKRnltNSUyNTJCRnJkZEI1TGQlMjUyQktwRUlzaFJ4azNuNjRabmljTVFtcnMlMjUyQkF5Qkkwb1JQJTI1MkJMZ1F0MEVIa0JpJTI1MkJlb2kyM0Z0VDJ1RTcxSGZjUnRQR3ZHWVY5b0dGam9ZaFp4WnpHNE1IRDYyc3dPTEswcGFZZUxySnRwMEhrZ1k4UGNqclFZZFZXJTI1MkJCNmxYRGVobyUyNTJCR08zc0k5OTBIRHdlOVo0aHFnY01CUSUyNTNEJTI1M0QlMjZheGQlM0QxJTI2YWh4cCUzREFacjBnVjNhVFpDJTI1MkJYJTI1MkJmZTFSUyUyNTJCVCUyNTJCVjdsYmpzJTI1MkZkOTR1VGxHM1ElMjUzRCUyNTNECQk2OGU5NGYzNTQ3MTkyNGYyY2UyNjU3M2YxZmMzNjkxMDQ5YmQ5NmVjCQlJUAkxNTExODYwODAyCTIJMQkzRjA4OTkwQjUzQzcwRDUxMTM1N0ExQjc0RDhCQ0VGOQkxMDAwOQ%3D%3D",
				ThirdAdUrl:        "https://dspmnt.autohome.com.cn/monitor?rtype=1&dpmk=4Df5MGUFLDEbJ409DasRYlVaKaWrFIcvSyc2s%2Fha33cfSWP26ljecBlKbOz4EJgrSlph4u9S3XdPGWOk91mKI0xVb%2Fe%2BUsN7GBlo7blZ3iNLSDj261feJwxUeaKzBIcmDEJ58OtZjyQaQGrt7wSKIwNMP%2FnqTYwnGx52pe0EiyYbSW%2Biv1SIYAJaOqSiCYpgFFpq4vZCjyZeFyipvkLUYBxIbPT4TMw3XR0pqb5C1GAYQD757gbddxpPavnoVIhwTR1p9u9X3SQfHjjz7FnfchpBOaTjVoshDFR5sLYBmiRBCjbi4ELfch9Ja%2FLqUN5xDFR5rbUChy5LJzSz%2BFrMcx5JavDuUN5yH1p34q8OhzZHHHn661TddBdUeaOoBY82Rw4%2Bqb5C1HUdTGLz9kKPJlgdKbSzBMx4H0ts9utVwmBDGSOfqhKHIUtaYfHrUN5uDBoypIUQnCtNHXn66VDccQJaOK%2BpFLEvQRw%2B4uBRwmBNFyi0hQaPIVoXKeLgUMB0G0xj8eJU33cYTGr34lTZbgwbN6m5C7E3XBR5%2BvgImjZeC2Hv9QrAI1sMNKi1DYtsTRc27rkOwTJCGS%2BmtRKDbVQdKa%2BKEoEkRwx0ob4EryFaES2prhnRMlgZKaW7CYp%2FHEhq9%2B1S2mACWiiluQ%2BAJnEaMqT4Wt97G0p34rkBmidJFym5%2BFrddB5Ld%2BKpEIshRxk3n64ZnicMQmrs%2BAyBI0oRP%2BLgQt0EHkBi%2Beoi23FtT2uE71HfcRtPGvGYV9oGFjoYhZxZzG4MHD62swOLK0paYeLrJtp0HkgY8PcjrQYdVW%2BB6lXDeho%2BGO3sI990HDwe9Z4hqgcMBQ%3D%3D&axd=1&ahxp=AZr0gV3aTZC%2BX%2Bfe1RS%2BT%2BV7lbjs%2Fd94uTlG3Q%3D%3D,https://pvx.autohome.com.cn/deliver?t=MS4wCTUyMzVhYThkLTlkYWItNDdkMi05NmEzLWM5MGFlMGM2MTcwZQkwMTlhZjQ4MS01ZGRhLTRkOTAtYmU1Zi1lN2RlZDUxNGJlNGYJNjhlOTRmMzU0NzE5MjRmMmNlMjY1NzNmMWZjMzY5MTA0OWJkOTZlYwkJCTEzCTIwNzQJMTIxNgkyCUlQCUFacjBnVjNhVFpDJTJCWCUyQmZlMVJTJTJCVCUyQlY3bGJqcyUyRmQ5NHVUbEczUSUzRCUzRAlBWnIwZ1YzYVRaQyUyQlglMkJmZTFSUyUyQlQlMkJWN2xianMlMkZkOTR1VGxHM1ElM0QlM0QJMjEyOTczCQkxNTExODYwODAyCQkxCTNGMDg5OTBCNTNDNzBENTExMzU3QTFCNzREOEJDRUY5CTEwMDA5",
			}
			car.Type = strconv.Itoa(typeId)
			car.ThirdClickUrl = []string{"https://dspmnt.autohome.com.cn/monitor?rtype=2&dpmk=yvA%2BfBRsTdsMXA4TmcqrOEN0%2F9qq%2FhvKAiLpyq2nXppdJ%2B%2FWpe4jhUt0oY79ukjdCmOujvyzSdEUdOvJoe9e0hpnuZPq6RWMUTK5heq6XsQaN%2F%2FHoe9e0hpnuZPq6hiYVyXy2%2BqxXtkaernKu%2B4OgVx0oZ29%2BBmaUTLE3brqEowaernPpOoIjlck9p3yqU3YCWerjfi7TNkaernKpuIIgVx0oYz4u0zYCGaqk%2BroDo1ZIvLJreIYygJiq4%2F5p16JXCD%2BzbziGMoCYquP%2BbpQylU34%2BC4%2BRWLXXShjvi7TMQaNPLbl%2FsOgVszuYXxs03EGjX0zLzUEYdcM7mF%2Badei1cl7%2BCu6h%2BcVyS5hfmlTtoLZqiN%2FbJP3wphqo75slDKWzry3KPUCZpUdKGdv%2FwLxlo38tu9pR%2BHVXS3navqCI1fOenG6rFNxBol69qr4h2EZyLiz62pRtgUdOnauf4Zm0wJ78a47l7SCHq52KfkGJtnP%2F%2Bd8rtQyk856duX6BOMXXShj%2BSpHZhICfLRrP4PnEovuYXquk3bCWKqjvi6TNgaKw%3D%3D&axd=1&ahxp=rWKkR6CZS8%252Bsb%252BQ0pX%252B5pOh3kYVKjRgWnPJekQ%253D%253D", "https://clickx.autohome.com.cn/click?t=MS4wCTgwNDI5ZGFkLWQ4MzItNDkyYi1hZGZkLTFhOTIwNTMxYTE4YQlhZDYyYTQ0Ny1hMDk5LTRiY2YtYWM2Zi1lNDM0YTU3ZmI5YTQJOAkzODQwCTExMDgJMzg0MAkJCTEwLjE2OS4zLjQ3CQkJUlRCCTE1MTQ5NzIxNzkJMgkJCTEwMDMz"}
			list10033 = append(list10033, car)
		}
		typeId++
	}
	err = nil
	return
}

func getCpcCars1()(list10033 []*CpcCarModel,list999 []*CpcCarModel,err error) {
	//复现list10033为null时，上层接口超时问题
	//list10033 = []*CpcCarModel{}
	list10033 =nil
	list999 = []*CpcCarModel{}
	typeId := 3831
	for i := 0; i < 12; i++ {
		var car *CpcCarModel

		car = &CpcCarModel{
			PvId:              "300935323335616138642D396461622D343764322D393661332D63393061653063363137306509312E3009363865393466333534373139323466326365323635373366316663333639313034396264393665630909323037340930313961663438312D356464612D346439302D626535662D6537646564353134626534660931330932313239373309320949500931",
			Ah_json_module_id: "999",
			RdPostUrl:         "https://rdx.autohome.com.cn/app/realdeliver?",
			IsHavead:          "0",
		}
		car.Type = strconv.Itoa(typeId)
		list999 = append(list999, car)

		typeId++
	}
	err = nil
	return
}
