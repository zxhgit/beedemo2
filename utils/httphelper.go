package utils

import (
	"bytes"
	"github.com/astaxie/beego/httplib"
	"net/http"
	"net/url"
	"time"
)

func HttpGet(url string, params map[string]string) (content string, err error) {
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
	content, err = remoteGet(joinedParams.String())
	return
}

func HttpGetUrl(url string) (content string, err error) {
	content, err = remoteGet(url)
	return
}

func remoteGet(requestUrl string) (content string, err error) {
	request := httplib.NewBeegoRequest(requestUrl, "GET")
	request.SetTimeout(60*time.Second, 60*time.Second)
	request.SetProxy(func(req *http.Request) (*url.URL, error) {
		u, _ := url.ParseRequestURI("http://127.0.0.1:8888")
		return u, nil
	})
	content, err = request.String()
	if err != nil {
		content = "" + err.Error()
	}
	return
}
