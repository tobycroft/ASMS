package sms253

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	urls string = "http://smssh1.253.com/msg/send/json"
)

// Sms253 模型
type Sms253 struct {
	AppCode   string `json:"appcode"`   // 账号ID
	AppSecret string `json:"appsecret"` // 账号密码

}

// Relust 返回模型
type Relust struct {
	Code     string `json:"code"`
	MsgID    string `json:"msgId"`
	ErrorMsg string `json:"errorMsg"`
	Time     string `json:"time"`
}

// NewSms253 新建模型
func NewSms253(appcode string, appsecret string) *Sms253 {
	return &Sms253{
		AppCode:   appcode,
		AppSecret: appsecret,
	}
}

// SendSms 发送短信
func (a *Sms253) SendSms(template string, recnum string) (*Relust, error) {
	params := make(map[string]interface{})
	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = a.AppCode    //创蓝API账号
	params["password"] = a.AppSecret //创蓝API密码
	params["phone"] = recnum         //手机号码

	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	params["msg"] = url.QueryEscape(template)
	params["report"] = "true"

	bytesData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(urls, "application/json;charset=UTF-8", bytes.NewReader([]byte(bytesData)))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	v := &Relust{}

	if err := json.Unmarshal(body, v); err != nil {
		return nil, err
	}

	return v, nil
}
