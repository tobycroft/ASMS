package LogSuccessModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "log_success"

func Api_insert(pid, text interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"pid":  pid,
		"text": text,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
