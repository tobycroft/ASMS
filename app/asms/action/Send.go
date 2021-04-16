package action

import (
	"errors"
	"fmt"
	config2 "github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/sms"
	"main.go/app/asms/model/LogErrorModel"
	"main.go/app/asms/model/LogSuccessModel"
	"main.go/app/asms/model/ProjectModel"
	"main.go/app/asms/model/TencentModel"
	"main.go/app/asms/model/Zz253Model"
	"main.go/extends/sms253"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

func App_tencent(id interface{}, phone, quhao, text string) error {
	tencent := TencentModel.Api_find(id)
	if len(tencent) > 0 {
		config := &config2.Config{
			AppId:  Calc.Any2String(tencent["appid"]),
			AppKey: Calc.Any2String(tencent["appkey"]),
			Sign:   Calc.Any2String(tencent["sign"]),
		}
		s := sms.NewSms(config)
		res, err := s.GetSmsSender().Fetchs(phone, quhao, []string{text}, tencent["tplid"].(int64))
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			return err
		}
		if res.Result == 0 {
			if !ProjectModel.Api_dec_amount(tencent["pid"]) {
				return errors.New("ProjectModelApi_dec_amount")
			}
			LogSuccessModel.Api_insert(tencent["pid"], text)
			return nil
		} else {
			LogErrorModel.Api_insert(tencent["pid"], res.Errmsg)
			return errors.New(res.Errmsg)
		}
		//fmt.Println(res.Errmsg, res.Ext, res.Result, err)
	} else {
		return errors.New("未找到项目")
	}
}

func Aoo_253(id interface{}, phone, quhao, text string) error {
	zz := Zz253Model.Api_find(id)
	if len(zz) > 0 {

	} else {
		return errors.New("未找到项目")

	}

	sms := sms253.NewSms253("N00xxxxxx", "eLENxxxxxx")

	msg := "【253云通讯】" + "测试人员您的验ddd证码为888888请在5分钟内输入。"

	rl, err := sms.SendSms(msg, "13600xxxxx")
	if err != nil {
		return err
	}

	fmt.Println(rl)
}
