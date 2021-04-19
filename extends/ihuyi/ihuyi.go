package ihuyi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"main.go/tuuz/Jsong"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Ihuyi_send(APIID, APIKEY, phone, content string) (string, error) {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := APIID   //查看用户名 登录用户中心->验证码通知短信>产品总览->API接口信息->APIID
	_password := APIKEY //查看密码 登录用户中心->验证码通知短信>产品总览->API接口信息->APIKEY
	_mobile := phone
	_content := content
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	v.Set("format", "json")
	body := strings.NewReader(v.Encode()) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	jobject, err := Jsong.JObject(string(data))
	fmt.Println(jobject, err)

	return string(data), err
}

func Ihuyi_send_intl(APIID, APIKEY, quhao, phone, content string) (string, error) {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := APIID   //查看用户名 登录用户中心->验证码通知短信>产品总览->API接口信息->APIID
	_password := APIKEY //查看密码 登录用户中心->验证码通知短信>产品总览->API接口信息->APIKEY
	_mobile := phone
	_content := content
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	v.Set("format", "json")
	body := strings.NewReader(v.Encode()) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://api.isms.ihuyi.com/webservice/isms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	return string(data), err
}
