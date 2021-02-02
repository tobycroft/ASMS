package action

import (
	"fmt"
	config2 "github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/sms"
	"main.go/app/asms/model/LogErrorModel"
	"main.go/app/asms/model/LogSuccessModel"
	"main.go/app/asms/model/ProjectModel"
	"main.go/app/asms/model/TencentModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

func App_tencent(id interface{}, phone, quhao, text string) {
	tencent := TencentModel.Api_find(id)
	if len(tencent) > 0 {
		config := &config2.Config{
			AppId:  Calc.Any2String(tencent["appid"]),
			AppKey: Calc.Any2String(tencent["appkey"]),
			Sign:   Calc.Any2String(tencent["sign"]),
		}
		s := sms.NewSms(config)
		res, err := s.GetSmsSender().Fetchs(phone, quhao, []string{text}, tencent["861810"].(int64))
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		}
		if res.Result == 0 {
			ProjectModel.Api_dec_amount(tencent["pid"])
			LogSuccessModel.Api_insert(tencent["pid"], text)
		} else {
			LogErrorModel.Api_insert(tencent["pid"], res.Errmsg)
		}
		fmt.Println(res.Errmsg, res.Ext, res.Result, err)
	} else {

	}

}
