package LogErrorModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "log_error"

func Api_insert(pid, error interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"pid":   pid,
		"error": error,
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
