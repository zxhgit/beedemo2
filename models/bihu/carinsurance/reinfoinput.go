package carinsurance

import (
	"bytes"
	"errors"
	"github.com/astaxie/beego/config"
	"reflect"
	"strings"
)

type InputModel struct {
	LicenseNo            string `query:"LicenseNo" ignore:""`
	CityCode             int    `query:"CityCode" ignore:"0"`
	CarVin               string `query:"CarVin" ignore:""`
	EngineNo             string `query:"EngineNo" ignore:""`
	Group                int    `query:"Group" ignore:"0"`
	CanShowNo            int    `query:"CanShowNo" ignore:"0"`
	CanShowExhaustScale  int    `query:"CanShowExhaustScale" ignore:"0"`
	ShowXiuLiChangType   int    `query:"ShowXiuLiChangType" ignore:"0"`
	TimeFormat           int    `query:"TimeFormat" ignore:"0"`
	ShowAutoMoldCode     int    `query:"ShowAutoMoldCode" ignore:"0"`
	ShowFybc             int    `query:"ShowFybc" ignore:"0"`
	ShowSheBei           int    `query:"ShowSheBei" ignore:"0"`
	RenewalCarType       int    `query:"RenewalCarType" ignore:"0"`
	ShowOrg              int    `query:"ShowOrg" ignore:"0"`
	SixDigitsAfterIdCard string `query:"SixDigitsAfterIdCard" ignore:""`
	ShowSanZheJieJiaRi   int    `query:"ShowSanZheJieJiaRi" ignore:"0"`
	RenewalSource        int    `query:"RenewalSource" ignore:"0"`
	ShowBaoFei           int    `query:"ShowBaoFei" ignore:"0"`
	keys                 []string
	params               map[string]string
}

func (input *InputModel) CheckInput() error {
	var err error
	if len(input.LicenseNo) <= 0 {
		err = errors.New("LicenseNo 必传")
	}
	if input.CityCode <= 0 {
		err = errors.New("CityCode 必传")
	}
	if input.Group <= 0 {
		err = errors.New("Group 必传")
	}
	return err
}

func (input *InputModel) BuildQueryMap() {
	modelType := reflect.TypeOf(input).Elem()
	keys := make([]string, 0)
	params := make(map[string]string)
	if modelType == nil || modelType.Kind() != reflect.Struct {
		return
	}
	modelValue := reflect.ValueOf(input).Elem()
	num := modelType.NumField()
	for i := 0; i < num; i++ {
		f := modelType.Field(i)
		kStr := f.Tag.Get("query")
		if len(kStr) > 0 {
			vStr := config.ToString(modelValue.Field(i).Interface())
			if len(vStr) > 0 && vStr != f.Tag.Get("ignore") {
				keys = append(keys, kStr)
				params[kStr] = vStr
			}
		}
	}
	input.keys = keys
	input.params = params
	return
}

func (input *InputModel) BuildParamStr() string {
	keys := input.keys
	params := input.params
	var p bytes.Buffer
	i := 0
	for ki := range keys {
		i++
		p.WriteString(keys[ki])
		p.WriteString("=")
		p.WriteString(params[keys[ki]])
		p.WriteString("&")
	}
	return strings.TrimRight(p.String(), "&")
}

func (input *InputModel) GetKeys() *[]string {
	return &(input.keys)
}

func (input *InputModel) SetKeys(keys []string) {
	input.keys = keys
}

func (input *InputModel) GetParams() *map[string]string {
	return &(input.params)
}
