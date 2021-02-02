package controller

import (
	"github.com/gin-gonic/gin"
	config2 "github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/sms"
	"main.go/app/asms/model/ProjectModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {

	route.Use(func(c *gin.Context) {
		ts, ok := Input.PostInt64("ts", c)
		if !ok {
			return
		}
		sign, ok := Input.Post("sign", c, false)
		if !ok {
			return
		}
		name, ok := Input.Post("name", c, false)
		if !ok {
			return
		}
		data := ProjectModel.Api_find(name)
		if len(data) > 0 {
			if data["active"].(int64) == 0 {
				RET.Fail(c, 403, nil, "项目已经停用")
				c.Abort()
				return
			} else {
				if app_conf.Debug {
					token, ok := Input.Post("token", c, false)
					if !ok {
						return
					}
					data := ProjectModel.Api_find(token)
					if len(data) < 1 {

					}
				} else {
					if Calc.Md5(Calc.Any2String(data["token"])+Calc.Any2String(ts)) != sign {
						RET.Fail(c, 403, nil, "签名不正确，加密方式为小写MD5(token+ts)")
						c.Abort()
						return
					}
				}

				c.Next()
				return
			}
		} else {
			RET.Fail(c, 403, nil, "项目不存在")
			c.Abort()
			return
		}
	})

	route.Any("send", send)

}

func send(c *gin.Context) {
	name := c.PostForm("name")
	data := ProjectModel.Api_find(name)
	if len(data) > 0 {
		config := &config2.Config{
			AppId:  "1400482871",
			AppKey: "7f293893d05e8e4e38edf05d6d885ff0",
			Sign:   "易聊红包APP",
		}
		s := sms.NewSms(config)
		tel := sms.SmsTel{Mobile: "13107670001", Nationcode: "86"}
		p := &sms.Params{Params: []string{"1234"}, Tel: tel, TplId: 861810}
		res, err := s.GetSmsSender().Fetch(p)
		println(res, err)
	} else {

	}
}
