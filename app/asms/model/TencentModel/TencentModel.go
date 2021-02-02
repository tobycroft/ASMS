package TencentModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "tencent"

func Api_find(pid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"pid": pid,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
