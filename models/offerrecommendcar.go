package models

import (
	"beedemo2/models/solr"
	"strconv"
	//"sort"
	//"fmt"
	"strings"
	"beedemo2/models/thousandfaces"
        ."github.com/ahmetb/go-linq"
	"beedemo2/utils"
)

type RecommendModel struct {
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
	IsLoan           int `json:"isloan"`
	DealerId         int `json:"dealerid"`
	//DownPayment      string `json:"downpayment"`
}

type PSortedModel struct {
	PRn int
	Price float64
	SolrModel *RecommendModel
}

type CrossSortedModel struct {
	CrossRn int
	SolrModel *RecommendModel
}

func GetSameSeriesRecommendSolr(appId,udId string,seriesId,areaId,pId,cId int)(r []*RecommendModel,err error) {
	para := make(map[string]string)
	para["_appid"] = appId
	para["udid"] = udId
	para["pageindex"] = "1"
	para["pagesize"] = "100"
	para["seriesid"] = strconv.Itoa(seriesId)
	para["areaid"] = strconv.Itoa(areaId)
	para["pid"] = strconv.Itoa(pId)
	para["cid"] = strconv.Itoa(cId)
	para["orderby"] = "4"
	para["dealertype"] = "9"
	r = []*RecommendModel{}
	var res *solr.SolrResultModel
	res, err = solr.GetDataBySearchInternal(para)
	if err != nil {
		return
	}
	if res == nil || res.Result == nil || res.Result.CarList == nil {
		return
	}
	for _, val := range res.Result.CarList {
		r = append(r, convThousandFacesModel(val))
	}
	return
}

func GetOtherSeriesRecommendSolr(appId,udId string,areaId,pId,cId int)(r []*RecommendModel,err error) {
	leveId := "16,17,18,19,20"
	para := make(map[string]string)
	para["_appid"] = appId
	para["udid"] = udId
	para["pageindex"] = "1"
	para["pagesize"] = "100"
	para["levelid"] = leveId
	para["areaid"] = strconv.Itoa(areaId)
	para["pid"] = strconv.Itoa(pId)
	para["cid"] = strconv.Itoa(cId)
	para["orderby"] = "4"
	para["dealertype"] = "9"
	r = []*RecommendModel{}
	var res *solr.SolrResultModel
	res, err = solr.GetDataBySearchInternal(para)
	if err != nil {
		return
	}
	if res == nil || res.Result == nil || res.Result.CarList == nil {
		return
	}
	for _, val := range res.Result.CarList {
		r = append(r, convThousandFacesModel(val))
	}
	return
}

func GetDefaultSolr(appId,udId string,areaId,pId,cId,count int)(r []*RecommendModel,err error) {
	para := make(map[string]string)
	para["_appid"] = appId
	para["udid"] = udId
	para["pageindex"] = "1"
	para["pagesize"] = strconv.Itoa(count)
	para["areaid"] = strconv.Itoa(areaId)
	para["pid"] = strconv.Itoa(pId)
	para["cid"] = strconv.Itoa(cId)
	para["dealertype"] = "9"
	r = []*RecommendModel{}
	var res *solr.SolrResultModel
	res, err = solr.GetDataBySearchInternal(para)
	if err != nil {
		return
	}
	if res == nil || res.Result == nil || res.Result.CarList == nil {
		return
	}
	for _, val := range res.Result.CarList {
		r = append(r, convThousandFacesModel(val))
	}
	return
}

func buildPSortedSlicesLinq(sSlices []*RecommendModel,price float64,maxCount int) (smSlices []*PSortedModel,err error) {
	slices := []*PSortedModel{}
	err = nil
	if sSlices == nil || len(sSlices) == 0 {
		return
	}
	pRn := 1
	for _, val := range sSlices {
		var f float64
		f, err = strconv.ParseFloat(val.Price, 8)
		sm := &PSortedModel{
			PRn:       pRn,
			Price:     CalcAbs(price - f),
			SolrModel: val,
		}
		slices = append(slices, sm)
		pRn++
	}
	From(slices).OrderByT(func(sm *PSortedModel) float64 { return sm.Price }).
		ThenByT(func(sm *PSortedModel) int { return sm.PRn }).
		Take(maxCount).ToSlice(&smSlices)
	return
}

func buildCrossSortedSlices(smSlices []*PSortedModel,crossStartRn int)(cssSlices []*CrossSortedModel,err error) {
	cssSlices = []*CrossSortedModel{}
	err = nil
	if smSlices == nil || len(smSlices) == 0 {
		return
	}
	for _, val := range smSlices {
		css := &CrossSortedModel{
			CrossRn:   crossStartRn,
			SolrModel: val.SolrModel,
		}
		cssSlices = append(cssSlices, css)
		crossStartRn = crossStartRn + 2
	}
	return
}

func getSameSeriesRecommend(appId,udId string,seriesId,areaId,pId,cId,maxCount int,price float64)(cssRes []*CrossSortedModel,err error) {
	cssRes= []*CrossSortedModel{}
	solrRes, err := GetSameSeriesRecommendSolr(appId, udId, seriesId, areaId, pId, cId)
	if err!=nil{
		return
	}
	pSortedRes, err := buildPSortedSlicesLinq(solrRes, price, maxCount)
	if err!=nil{
		return
	}
	cssRes, err = buildCrossSortedSlices(pSortedRes, 1)
	return
}

func getOtherSeriesRecommend(appId,udId string,areaId,pId,cId,maxCount int,price float64)(cssRes []*CrossSortedModel,err error) {
	cssRes= []*CrossSortedModel{}
	solrRes, err := GetOtherSeriesRecommendSolr(appId, udId, areaId, pId, cId)
	if err!=nil{
		return
	}
	pSortedRes, err := buildPSortedSlicesLinq(solrRes, price, maxCount)
	if err!=nil{
		return
	}
	cssRes, err = buildCrossSortedSlices(pSortedRes, 2)
	return
}

func GetThousandFacesRecommend1(appId,udId,carIds string,areaId,pId,cId int)(res []*RecommendModel,err error) {
	para := make(map[string]string)
	para["uid"] = "test"
	para["cookieid"] = udId
	para["platform"] = "app"
	para["recentinfos"] = carIds
	para["cityid"] = strconv.Itoa(cId)
	para["provinceid"] = strconv.Itoa(pId)
	para["bucket"] = "2"
	res = []*RecommendModel{}
	var tRes *thousandfaces.RecommendResultModel
	tRes, err = thousandfaces.GetThousandFacesIds(para)
	if err != nil {
		return
	}
	if tRes == nil || tRes.Result == nil || tRes.Result.List == nil {
		return
	}
	var arrId []int
	for _, rc := range tRes.Result.List {
		arrId = append(arrId, rc.InfoId)
	}
	var cMap map[int]*RecommendModel
	cMap, err = buildCarsInfo(appId, udId, arrId)
	if err != nil {
		return
	}
	var ok bool
	var car *RecommendModel
	for _, rc := range tRes.Result.List {
		car, ok = cMap[rc.InfoId]
		if ok {
			car.Score = rc.Score
			car.UrType = rc.UrType
			res = append(res, car)
		}
	}
	return
}

func GetThousandFacesRecommend(udId,carIds string,pId,cId int)(res []*thousandfaces.RecommendCarItem,err error) {
	para := make(map[string]string)
	para["uid"] = "test"
	para["cookieid"] = udId
	para["platform"] = "app"
	para["recentinfos"] = carIds
	para["cityid"] = strconv.Itoa(cId)
	para["provinceid"] = strconv.Itoa(pId)
	para["bucket"] = "2"
	res = []*thousandfaces.RecommendCarItem{}
	var tRes *thousandfaces.RecommendResultModel
	tRes, err = thousandfaces.GetThousandFacesIds(para)
	if err != nil {
		return
	}
	if tRes == nil || tRes.Result == nil || tRes.Result.List == nil {
		return
	}
	res=append(res,tRes.Result.List...)
	return
}

func GetRecommendCarAsync(appId,udId string,seriesId,areaId,pId,cId int,price float64)(res []*RecommendModel,err error) {
	maxCount := 24
	res = []*RecommendModel{}
	var sameRes, otherRes []*CrossSortedModel
	combinedRes := []*CrossSortedModel{}
	resultChan := make(chan []*CrossSortedModel, 2)
	defer close(resultChan)
	go func() {
		sameRes, err = getSameSeriesRecommend(appId, udId, seriesId, areaId, pId, cId, maxCount, price)
		resultChan <- sameRes
	}()
	go func() {
		otherRes, err = getOtherSeriesRecommend(appId, udId, areaId, pId, cId, maxCount, price)
		resultChan <- otherRes
	}()
	for i := 0; i < 2; i++ {
		resItem := <-resultChan
		combinedRes = append(combinedRes, resItem...)
	}
	if err != nil {
		return
	}
	From(combinedRes).OrderByT(func(csm *CrossSortedModel)int {return csm.CrossRn}).
		SelectT(func(csm *CrossSortedModel) *RecommendModel {return csm.SolrModel}).
		ToSlice(&res)
	return
}

func GetOfferRecommendList(appId,udId,carIds string,carId, seriesId,areaId,pId,cId int,price float64) (res []*RecommendModel,err error) {
	size := 48
	res = []*RecommendModel{}
	//推荐车源
	var seriesR []*RecommendModel
	seriesR, err = GetRecommendCarAsync(appId, udId, seriesId, areaId, pId, cId, price)
	if err != nil {
		return
	}
	From(seriesR).WhereT(func(c *RecommendModel) bool { return c.CarId != carId }).
		Take(size).ToSlice(&res) //过滤当前车
	rCount := len(res)
	tfR:=[]*thousandfaces.RecommendCarItem{}
	if rCount < size { //补充
		//千人千面
		tfR, err = GetThousandFacesRecommend(udId, carIds, pId, cId) //改为先过滤再fill车源信息
		//过滤车源:上一步结果+当前询价车
		var resCarId []int
		From(res).SelectT(func(c *RecommendModel) int { return c.CarId }).ToSlice(&resCarId)
		var filtedR []*thousandfaces.RecommendCarItem
		From(tfR).WhereT(func(c *thousandfaces.RecommendCarItem) bool {
			res, _ := utils.Contain(c.InfoId, resCarId)
			return res && c.InfoId != carId
		}).OrderByDescendingT(func(c *thousandfaces.RecommendCarItem) int { return c.Score }).
			Take(size - rCount).ToSlice(&filtedR)
		//填充千人千面车源信息
		var arrId []int
		From(filtedR).SelectT(func(c *thousandfaces.RecommendCarItem) int { return c.InfoId }).ToSlice(&arrId)
		var cMap map[int]*RecommendModel
		cMap, err = buildCarsInfo(appId, udId, arrId)
		if err != nil {
			return
		}
		var tRes []*RecommendModel
		From(cMap).OrderByDescendingT(func(m *RecommendModel) int { return m.Score }).ToSlice(&tRes)
		res = append(res, tRes...)
		//默认列表
		rCount = len(res)
		if rCount < size {
			var dsR []*RecommendModel
			dsR, err = GetDefaultSolr(appId, udId, areaId, pId, cId, size-rCount+10)
			if err != nil {
				return
			}
			var resCarId []int
			From(res).SelectT(func(c *RecommendModel) int { return c.CarId }).ToSlice(&resCarId)
			//过滤车源:上一步结果+当前询价车
			From(dsR).WhereT(func(c *RecommendModel) bool {
				res, _ := utils.Contain(c.CarId, resCarId)
				return res && c.CarId != carId
			}).Take(size - rCount).ToSlice(&dsR)
			res = append(res, dsR...)
		}
	}
	From(res).ForEachT(func(c *RecommendModel) {
		c.UrType="99"
		tfCar:= From(tfR).FirstWith(func(tfc interface{}) bool {
			return tfc.(*thousandfaces.RecommendCarItem).InfoId==c.CarId
		})
		if tfCar!=nil{
			c.Score= tfCar.(*thousandfaces.RecommendCarItem).Score
			c.UrType=tfCar.(*thousandfaces.RecommendCarItem).UrType
		}
	})
	return
}

func buildCarsInfo(appId,udId string,ids []int)(res map[int]*RecommendModel,err error) {
	res=make(map[int]*RecommendModel)
	if ids == nil || len(ids) == 0 {
		return
	}
	arrStrIds := []string{}
	for _, val := range ids {
		arrStrIds = append(arrStrIds, strconv.Itoa(val))
	}
	para := make(map[string]string)
	para["_appid"] = appId
	para["udid"] = udId
	para["pageindex"] = "1"
	para["pagesize"] = "100"
	para["dealertype"] = "9"
	para["carids"] = strings.Join(arrStrIds, ",")
	var carMap map[int]*solr.SearchModel
	carMap, err = getCarMapBySearchInternal(para)
	if err != nil {
		return
	}
	var car *solr.SearchModel
	var ok bool
	for _, val := range ids {
		car, ok = carMap[val]
		if ok {
			res[val]=convThousandFacesModel(car)
		}
	}
	return
}

func getCarMapBySearchInternal(condition map[string]string)(res map[int]*solr.SearchModel, err error) {
	res = make(map[int]*solr.SearchModel)
	var m *solr.SolrResultModel
	m, err = solr.GetDataBySearchInternal(condition)
	if err != nil {
		return
	}
	if m == nil || m.Result == nil || m.Result.CarList == nil {
		return
	}
	for _, val := range m.Result.CarList {
		res[val.CarId] = val
	}
	return
}

func convThousandFacesModel(sm *solr.SearchModel)(m *RecommendModel)  {
	m=&RecommendModel{
		sm.CarId,
		sm.CarName,
		sm.Price,
		sm.Image,
		sm.Image330x220,
		sm.Mileage,
		sm.RegistrationDate,
		sm.SourceId,
		sm.Pid,
		sm.Cid,
		sm.CName,
		sm.PublishDate,
		sm.NewCarPrice,
		sm.IsTopConfig,
		sm.IsNewCar,
		sm.UrType,
		sm.Score,
		sm.IsLoan,
		sm.DealerId,
	}
	if sm.InterestFree==8{
		m.IsLoan=1
	}
	return
}

func CalcAbs(x float64) (ret float64) {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}