package thousandfaces

import (
	"encoding/json"
	"beedemo2/utils"
)

const thousandFacesUrl = "http://data.api.che168.com/rcm/app/"

type RecommendResultModel struct {
	ReturnCode          int `json:"returncode"`
	Message             string `json:"message"`
	Result              *RecommendResultListModel `json:"result"`
}

type RecommendResultListModel struct {
	List   []*RecommendCarItem `json:"list"`
}

type RecommendCarItem struct {
	InfoId int `json:"infoid"`
	UrType string `json:"urtype"`
	Score  int `json:"score"`
}

func GetThousandFacesIds(condition map[string]string)(r *RecommendResultModel,err error)  {
	var httpRes string
	httpRes,err= utils.HttpGet(thousandFacesUrl,condition)
	json.Unmarshal([]byte(httpRes), &r)
	return
}