package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/asms/action"
	"main.go/app/asms/model/ProjectModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"main.go/tuuz/Vali"
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
	phone, ok := Input.Post("phone", c, false)
	if !ok {
		return
	}
	err := Vali.Length(phone, 8, 11)
	if err != nil {
		RET.Fail(c, 400, nil, err.Error())
		return
	}
	quhao, ok := Input.Post("quhao", c, false)
	if !ok {
		return
	}
	err = Vali.Length(quhao, 1, 3)
	if err != nil {
		RET.Fail(c, 400, nil, err.Error())
		return
	}
	text, ok := Input.Post("text", c, false)
	if !ok {
		return
	}
	if data["amount"].(int64) < 1 {
		RET.Fail(c, 400, "没有数量了", "没有可以用于扣除的数量了")
		return
	}
	if len(data) > 0 {
		switch data["type"].(string) {
		case "tencent":
			if !ProjectModel.Api_dec_amount(data["id"]) {
				RET.Fail(c, 400, "没有数量了", "没有可以用于扣除的数量了")
				return
			}
			ret, err := action.App_tencent(data["id"], phone, quhao, text)
			if err != nil {
				RET.Fail(c, 300, ret, err.Error())
			} else {
				RET.Success(c, 0, ret, nil)
			}
			break

		case "zz253":
			if !ProjectModel.Api_dec_amount(data["id"]) {
				RET.Fail(c, 400, "没有数量了", "没有可以用于扣除的数量了")
				return
			}
			err := action.App_253(data["id"], phone, quhao, text)
			if err != nil {
				RET.Fail(c, 300, err, err.Error())
			} else {
				RET.Success(c, 0, err, nil)
			}
			break

		case "ihuyi":
			if !ProjectModel.Api_dec_amount(data["id"]) {
				RET.Fail(c, 400, "没有数量了", "没有可以用于扣除的数量了")
				return
			}
			str, err := action.App_ihuyi(data["id"], phone, quhao, text)
			if err != nil {
				RET.Fail(c, 300, err.Error(), err.Error())
			} else {
				RET.Success(c, 0, err, str)
			}
			break

		case "aliyun":
			if !ProjectModel.Api_dec_amount(data["id"]) {
				RET.Fail(c, 400, "没有数量了", "没有可以用于扣除的数量了")
				return
			}
			ret, err := action.App_aliyun(data["id"], phone, quhao, text)
			if err != nil {
				RET.Fail(c, 300, err, err.Error())
			} else {
				RET.Success(c, 0, ret, nil)
			}
			break

		default:
			RET.Fail(c, 404, nil, "没有执行方法")
			break
		}

	} else {
		RET.Fail(c, 404, nil, "项目未找到")
	}
}
