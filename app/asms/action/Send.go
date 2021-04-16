package action

import (
	"errors"
	config2 "github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/sms"
	"main.go/app/asms/model/IhuyiModel"
	"main.go/app/asms/model/LogErrorModel"
	"main.go/app/asms/model/LogSuccessModel"
	"main.go/app/asms/model/ProjectModel"
	"main.go/app/asms/model/TencentModel"
	"main.go/app/asms/model/Zz253Model"
	ihuyi2 "main.go/extends/ihuyi"
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

func App_253(id interface{}, phone, quhao, text string) error {
	zz := Zz253Model.Api_find(id)
	if len(zz) > 0 {

		sms := sms253.NewSms253(Calc.Any2String(zz["appcode"]), Calc.Any2String(zz["appsecret"]))

		res, err := sms.SendSms(text, phone)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			return err
		}
		if res.Code == "0" {
			if !ProjectModel.Api_dec_amount(zz["pid"]) {
				return errors.New("ProjectModelApi_dec_amount")
			}
			LogSuccessModel.Api_insert(zz["pid"], text)
			return nil
		} else {
			LogErrorModel.Api_insert(zz["pid"], res.ErrorMsg)
			return errors.New(res.ErrorMsg)
		}
	} else {
		return errors.New("未找到项目")

	}
}

func App_ihuyi(id interface{}, phone, quhao, text string) error {
	ihuyi := IhuyiModel.Api_find(id)
	if len(ihuyi) > 0 {
		apiid := Calc.Any2String(ihuyi["apiid"])
		apikey := Calc.Any2String(ihuyi["apikey"])
		str := ""
		if quhao == "86" {
			str, _ = ihuyi2.Ihuyi_send(apiid, apikey, phone, text)
		} else {
			str, _ = ihuyi2.Ihuyi_send_intl(apiid, apikey, phone, text)
		}

		if true {
			if !ProjectModel.Api_dec_amount(ihuyi["pid"]) {
				return errors.New("ProjectModelApi_dec_amount")
			}
			LogSuccessModel.Api_insert(ihuyi["pid"], text)
			return nil
		} else {
			LogErrorModel.Api_insert(ihuyi["pid"], str)
			return errors.New(str)
		}
	} else {
		return errors.New("未找到项目")

	}
}
