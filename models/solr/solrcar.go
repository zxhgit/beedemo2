package solr

import (
	"beedemo2/utils"
	"encoding/json"
)

const searchUrl = "http://2sc6.api.che168.com/v56/car/SearchInternal.ashx"

type SolrResultModel struct {
	ReturnCode          int `json:"returncode"`
	Message             string `json:"message"`
	Result              *SolrResultListModel `json:"result"`
}

type SolrResultListModel struct {
	PageCount int `json:"pagecount"`
	RowCount  int `json:"rowcount"`
	PageIndex int `json:"pageindex"`
	PageSize  int `json:"pagesize"`
	CpcNum    int `json:"cpcnum"`
	CarList   []*SearchModel `json:"carlist"`
}

//对应search接口字段
type SearchModel struct {
	CarId            int `json:"carid"`
	CarName          string `json:"carname"`
	Price            string `json:"price"`
	Image            string `json:"image"`
	Image330x220     string `json:"image330x220"`
	Mileage          string `json:"mileage"`
	RegistrationDate string `json:"registrationdate"`
	SourceId         int `json:"sourceid"`
	Pid              int `json:"pid"`
	Cid              int `json:"cid"`
	CName            string `json:"cname"`
	PublishDate      string `json:"publishdate"`
	NewCarPrice      string `json:"newcarprice"`
	IsTopConfig      int `json:"istopconfig"`
	IsNewCar         int `json:"isnewcar"`
	UrType           string `json:"urtype"`
	Score            int `json:"score"`
	InterestFree     int `json:"interestfree"`
	IsLoan           int `json:"isloan"`
	SeriesId         int `json:"seriesid"`
	DealerId         int `json:"dealerid"`
}

func GetDataBySearchInternal(condition map[string]string)(m *SolrResultModel,err error) {
	var httpRes string
	httpRes, err = utils.HttpGet(searchUrl, condition)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(httpRes), &m)
	return
}

