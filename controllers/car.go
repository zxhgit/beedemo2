package controllers

import (
	"github.com/astaxie/beego"
	//"beedemo2/models/solr"
	"time"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"beedemo2/models"
)

type CarController struct {
	beego.Controller
}

// @Title Get
// @Description find car by carId
// @Param	carId		path 	string	true		"the carId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :carId is empty
// @router /:carId [get]
func (c *CarController)Get(){
	carId := c.Ctx.Input.Param(":carId")
	c.Data["json"]=	"getcar"+carId
	c.ServeJSON()
}

//@router / [get]
func (c *CarController)GetAll() {
	appId := c.GetString("_appid")
	udId := c.GetString("udid")
	seriesId, _ := c.GetInt("seriesid")
	areaId, _ := c.GetInt("areaid")
	pId, _ := c.GetInt("pid")
	cId, _ := c.GetInt("cid")
	carIds := c.GetString("carids")
	carId, _ := c.GetInt("carid")
	price, _ := c.GetFloat("price", 0)
	//res, err := models.GetSameSeriesRecommendSolr(appId,udId,seriesId,areaId,pId,cId)
	//fmt.Println(time.Now())
	//res, err := models.GetRecommendCar(appId, udId, seriesId, areaId, pId, cId, 12.45)
	//fmt.Println(time.Now())
	//res, err := models.GetRecommendCarAsync(appId, udId, seriesId, areaId, pId, cId, 12.45)
	//fmt.Println(time.Now())
	//res, err :=models.GetDefaultSolr(appId, udId, areaId, pId, cId, 10)
	//res, err := models.GetThousandFacesRecommend(appId, udId, carIds, areaId, pId, cId)
	res, err := models.GetOfferRecommendList(appId, udId, carIds, carId, seriesId, areaId, pId, cId, price)
	if (err != nil) {
		c.Data["json"] = "err"
	} else {
		c.Data["json"] = res
	}
	c.ServeJSON()
}






type RemoteResult struct {
	Url string
	Result string
}

func RemoteGet(requestUrl string, resultChan chan RemoteResult) {

	request := httplib.NewBeegoRequest(requestUrl,"GET")
	request.SetTimeout(2 * time.Second, 5 * time.Second)
	//request.String()
	content, err := request.String()
	if err != nil {
		content =""+ err.Error()
	}
	resultChan <- RemoteResult{Url:requestUrl, Result:content}
}
func MultiGet(urls []string) []RemoteResult {
	fmt.Println(time.Now())
	resultChan := make(chan RemoteResult, len(urls))
	defer close(resultChan)
	var result []RemoteResult
	//fmt.Println(result)
	for _, url := range urls {
		go RemoteGet(url, resultChan)
	}
	for i:= 0; i < len(urls); i++ {
		res := <-resultChan
		result = append(result, res)
	}
	fmt.Println(time.Now())
	return result
}

func main() {
	urls := []string{
		"http://127.0.0.1/test.php?i=13",
		"http://127.0.0.1/test.php?i=14",
		"http://127.0.0.1/test.php?i=15",
		"http://127.0.0.1/test.php?i=16",
		"http://127.0.0.1/test.php?i=17",
		"http://127.0.0.1/test.php?i=18",
		"http://127.0.0.1/test.php?i=19",
		"http://127.0.0.1/test.php?i=20"}
	content := MultiGet(urls)
	fmt.Println(content)
}