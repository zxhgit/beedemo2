package utils

import (
	"time"
	"github.com/astaxie/beego/httplib"
	"bytes"
)

func HttpGet(url string,params map[string]string)(content string,err error)  {
	//参数处理
	var joinedParams bytes.Buffer
	joinedParams.WriteString(url)
	joinedParams.WriteString("?")
	for k, v := range params {
		joinedParams.WriteString(k)
		joinedParams.WriteString("=")
		joinedParams.WriteString(v)
		joinedParams.WriteString("&")
	}
	content,err= remoteGet(joinedParams.String())
	return
}

func remoteGet(requestUrl string)(content string,err error) {
	request := httplib.NewBeegoRequest(requestUrl,"GET")
	request.SetTimeout(2 * time.Second, 5 * time.Second)
	content, err = request.String()
	if err != nil {
		content =""+ err.Error()
	}
	return
}
