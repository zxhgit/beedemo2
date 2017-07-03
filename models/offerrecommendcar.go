package models

import (
	"beedemo2/models/solr"
	"strconv"
	"sort"
	//"fmt"
	"strings"
	"beedemo2/models/thousandfaces"
        ."github.com/ahmetb/go-linq"
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

//创建该类型，以便实现sort.Interface
type PSortedModelSlices []*PSortedModel

//创建该类型，以便实现sort.Interface
type CrossSortedModelSlices []*CrossSortedModel

func (c PSortedModelSlices) Len() int {
	return len(c)
}

func (c PSortedModelSlices) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c PSortedModelSlices) Less(i, j int) bool {
	if c[i].Price == c[j].Price {
		return c[i].PRn < c[j].PRn
	}
	return c[i].Price < c[j].Price
}

func (c CrossSortedModelSlices) Len() int {
	return len(c)
}
func (c CrossSortedModelSlices) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c CrossSortedModelSlices) Less(i, j int) bool {
	return c[i].CrossRn < c[j].CrossRn
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

func buildPSortedSlices(sSlices []*RecommendModel,price float64,maxCount int) (smSlices PSortedModelSlices,err error) {
	smSlices = PSortedModelSlices{}
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
		smSlices = append(smSlices, sm)
		pRn++
	}
	if !sort.IsSorted(smSlices) {
		sort.Sort(smSlices)
	}
	if maxCount < len(smSlices) {
		smSlices = smSlices[0:maxCount]
	}
	return
}

func buildPSortedSlicesLinq(sSlices []*RecommendModel,price float64,maxCount int) (smSlices PSortedModelSlices,err error) {
	smSlices = PSortedModelSlices{}
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
		smSlices = append(smSlices, sm)
		pRn++
	}
	a:= From(smSlices)

	if !sort.IsSorted(smSlices) {
		sort.Sort(smSlices)
	}
	if maxCount < len(smSlices) {
		smSlices = smSlices[0:maxCount]
	}
	return
}

func buildCrossSortedSlices(smSlices PSortedModelSlices,crossStartRn int)(cssSlices CrossSortedModelSlices,err error) {
	cssSlices = CrossSortedModelSlices{}
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

func getSameSeriesRecommend(appId,udId string,seriesId,areaId,pId,cId,maxCount int,price float64)(cssRes CrossSortedModelSlices,err error) {
	cssRes= CrossSortedModelSlices{}
	solrRes, err := GetSameSeriesRecommendSolr(appId, udId, seriesId, areaId, pId, cId)
	if err!=nil{
		return
	}
	pSortedRes, err := buildPSortedSlices(solrRes, price, maxCount)
	if err!=nil{
		return
	}
	cssRes, err = buildCrossSortedSlices(pSortedRes, 1)
	return
}

func getOtherSeriesRecommend(appId,udId string,areaId,pId,cId,maxCount int,price float64)(cssRes CrossSortedModelSlices,err error) {
	cssRes= CrossSortedModelSlices{}
	solrRes, err := GetOtherSeriesRecommendSolr(appId, udId, areaId, pId, cId)
	if err!=nil{
		return
	}
	pSortedRes, err := buildPSortedSlices(solrRes, price, maxCount)
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

func GetRecommendCar(appId,udId string,seriesId,areaId,pId,cId int,price float64) (res []*RecommendModel,err error) {
	maxCount := 24
	var sameRes, otherRes CrossSortedModelSlices
	combinedRes := CrossSortedModelSlices{}
	res = []*RecommendModel{}

	sameRes, err = getSameSeriesRecommend(appId, udId, seriesId, areaId, pId, cId, maxCount, price)
	if err != nil {
		return
	}
	combinedRes = append(combinedRes,sameRes...)

	otherRes, err = getOtherSeriesRecommend(appId, udId, areaId, pId, cId, maxCount, price)
	if err != nil {
		return
	}
	combinedRes = append(combinedRes,otherRes...)

	if !sort.IsSorted(combinedRes) {
		sort.Sort(combinedRes)
	}
	for _, val := range combinedRes {
		res = append(res, val.SolrModel)
	}
	return
}

func GetRecommendCarAsync(appId,udId string,seriesId,areaId,pId,cId int,price float64)(res []*RecommendModel,err error) {
	maxCount := 24
	res = []*RecommendModel{}
	var sameRes, otherRes CrossSortedModelSlices
	combinedRes := CrossSortedModelSlices{}
	resultChan := make(chan CrossSortedModelSlices, 2)
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
	if !sort.IsSorted(combinedRes) {
		sort.Sort(combinedRes)
	}
	for _, val := range combinedRes {
		res = append(res, val.SolrModel)
	}
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
	for _, c := range seriesR { //过滤当前车
		if c != nil && c.CarId != carId {
			res = append(res, c)
		}
	}
	rCount := len(res)
	if rCount < size { //补充
		//千人千面
		var tfR []*thousandfaces.RecommendCarItem
		tfR, err = GetThousandFacesRecommend(udId, carIds, pId, cId) //改为先过滤再fill车源信息
		//过滤车源:上一步结果+当前询价车
		filtedR:=[]*thousandfaces.RecommendCarItem{}
		for _, c := range tfR {
			var hit bool
			for _, r := range res {
				if r.CarId == c.InfoId {
					hit = true
					break
				}
			}
			if !hit && c.InfoId != carId {
				filtedR = append(filtedR, c)
			}
		}
		//填充千人千面车源信息
		var arrId []int
		for _, rc := range filtedR {
			arrId = append(arrId, rc.InfoId)
		}
		var cMap map[int]*RecommendModel
		cMap, err = buildCarsInfo(appId, udId, arrId)
		if err != nil {
			return
		}
		var ok bool
		var car *RecommendModel
		tfFiled := []*RecommendModel{}//排序？？
		for _, rc := range filtedR {
			car, ok = cMap[rc.InfoId]
			if ok {
				car.Score = rc.Score
				car.UrType = rc.UrType
				tfFiled = append(tfFiled, car)
			}
		}
		//默认列表
		rCount = len(res)
		if rCount<size{
			var dsR []*RecommendModel
			dsR,err= GetDefaultSolr(appId,udId,areaId,pId,cId,size - rCount + 10)
			if err!=nil{
				return
			}
			//过滤车源:上一步结果+当前询价车
			filtedR:=[]*RecommendModel{}
			for _, c := range dsR {
				var hit bool
				for _, r := range res {
					if r.CarId == c.CarId {
						hit = true
						break
					}
				}
				if !hit && c.CarId != carId {
					filtedR = append(filtedR, c)
				}
			}
		}

	}
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