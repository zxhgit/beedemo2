package carinsurance

import (
	"beedemo2/utils"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
)

type ReInfoModel struct {
	BusinessStatus int        `json:"businessstatus"`
	StatusMessage  string     `json:"statusmessage"`
	UserInfo       *UserInfo  `json:"UserInfo"`
	SaveQuote      *SaveQuote `json:"SaveQuote"`
	CustKey        string     `json:"CustKey"`
}

type UserInfo struct {
	CarUsedType           int     `json:"CarUsedType"`
	LicenseNo             string  `json:"LicenseNo"`
	LicenseOwner          string  `json:"LicenseOwner"`
	PostedName            string  `json:"PostedName"`
	InsuredName           string  `json:"InsuredName"`
	PurchasePrice         float64 `json:"PurchasePrice"`
	IdType                int     `json:"IdType"`
	CredentislasNum       string  `json:"CredentislasNum"`
	CityCode              int     `json:"CityCode"`
	EngineNo              string  `json:"EngineNo"`
	ModleName             string  `json:"ModleName"`
	RegisterDate          string  `json:"RegisterDate"`
	CarVin                string  `json:"CarVin"`
	ForceExpireDate       string  `json:"ForceExpireDate"`
	BusinessExpireDate    string  `json:"BusinessExpireDate"`
	NextForceStartDate    string  `json:"NextForceStartDate"`
	NextBusinessStartDate string  `json:"NextBusinessStartDate"`
	SeatCount             int     `json:"SeatCount"`
	InsuredIdCard         string  `json:"InsuredIdCard"`
	InsuredIdType         int     `json:"InsuredIdType"`
	InsuredMobile         string  `json:"InsuredMobile"`
	HolderIdCard          string  `json:"HolderIdCard"`
	HolderIdType          int     `json:"HolderIdType"`
	HolderMobile          string  `json:"HolderMobile"`
	RateFactor1           float64 `json:"RateFactor1"`
	RateFactor2           float64 `json:"RateFactor2"`
	RateFactor3           float64 `json:"RateFactor3"`
	RateFactor4           float64 `json:"RateFactor4"`
	FuelType              int     `json:"FuelType"`
	ProofType             int     `json:"ProofType"`
	ClauseType            int     `json:"ClauseType"`
	LicenseColor          int     `json:"LicenseColor"`
	RunRegion             int     `json:"RunRegion"`
	IsPublic              int     `json:"IsPublic"`
	BizNo                 string  `json:"BizNo"`
	ForceNo               string  `json:"ForceNo"`
	ExhaustScale          string  `json:"ExhaustScale"`
	AutoMoldCode          string  `json:"AutoMoldCode"`
	Organization          string  `json:"Organization"`
}

type SaveQuote struct {
	Source                 int64   `json:"Source"`
	CheSun                 float64 `json:"CheSun"`
	SanZhe                 float64 `json:"SanZhe"`
	DaoQiang               float64 `json:"DaoQiang"`
	SiJi                   float64 `json:"SiJi"`
	ChengKe                float64 `json:"ChengKe"`
	BoLi                   float64 `json:"BoLi"`
	HuaHen                 float64 `json:"HuaHen"`
	BuJiMianCheSun         float64 `json:"BuJiMianCheSun"`
	BuJiMianSanZhe         float64 `json:"BuJiMianSanZhe"`
	BuJiMianDaoQiang       float64 `json:"BuJiMianDaoQiang"`
	SheShui                float64 `json:"SheShui"`
	ZiRan                  float64 `json:"ZiRan"`
	BuJiMianChengKe        float64 `json:"BuJiMianChengKe"`
	BuJiMianSiJi           float64 `json:"BuJiMianSiJi"`
	BuJiMianHuaHen         float64 `json:"BuJiMianHuaHen"`
	BuJiMianSheShui        float64 `json:"BuJiMianSheShui"`
	BuJiMianZiRan          float64 `json:"BuJiMianZiRan"`
	BuJiMianJingShenSunShi float64 `json:"BuJiMianJingShenSunShi"`
	HcSanFangTeYue         float64 `json:"HcSanFangTeYue"`
	HcJingShenSunShi       float64 `json:"HcJingShenSunShi"`
	HcXiuLiChangType       string  `json:"HcXiuLiChangType"`
	HcXiuLiChang           string  `json:"HcXiuLiChang"`
	SanZheJieJiaRi         string  `json:"SanZheJieJiaRi"`
	Fybc                   string  `json:"Fybc"`
	FybcDays               string  `json:"FybcDays"`
	SheBeiSunShi           string  `json:"SheBeiSunShi"`
	BjmSheBeiSunShi        string  `json:"BjmSheBeiSunShi"`
}

func GetReInfo(inputModel *InputModel) (res *ReInfoModel, err error) {
	urlStr := "http://iu.91bihu.com/api/CarInsurance/getreinfo"
	agentId := "161303"
	agentSecret := "728891080a7"
	custKey := "custKey_capadmin"
	inputModel.BuildQueryMap()
	keys := *(inputModel.GetKeys())
	para := *(inputModel.GetParams())
	keys = append(keys, "Agent", "CustKey")
	para["Agent"] = agentId
	para["CustKey"] = custKey
	inputModel.SetKeys(keys)
	urlParamStr := inputModel.BuildParamStr()
	secCode := getSecCode(urlParamStr, agentSecret)
	keys = append(keys, "SecCode")
	para["SecCode"] = secCode
	para["LicenseNo"] = url.QueryEscape(para["LicenseNo"])
	inputModel.SetKeys(keys)
	urlParamStr = inputModel.BuildParamStr()
	var httpRes string
	httpRes, err = utils.HttpGetUrl(urlStr + "?" + urlParamStr)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(httpRes), &res)
	return
}

func getSecCode(paramStr string, agentSecret string) string {
	urlStr := paramStr + agentSecret
	bUrlStr := []byte(urlStr)
	h := md5.Sum(bUrlStr)
	md5str := fmt.Sprintf("%x", h)
	return md5str
}
