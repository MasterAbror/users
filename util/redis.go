package util

import (
	"context"
	"encoding/json"
	"time"

	"github.com/MasterAbror/users/database"
	"github.com/MasterAbror/users/models"
)

func GetRedis(id string) any {
	var rdb = database.NewRedisClient()
	get := rdb.Get(context.Background(), id)
	response, err := get.Result()
	if err != nil {
		msg := "err"
		return msg
	}
	var rau models.UserRedis
	var er = json.Unmarshal([]byte(response), &rau)
	if er != nil {
		msg := "err"
		return msg
	}
	return rau
}

func SetRedis(id string, rau string) string {
	var rdb = database.NewRedisClient()
	key := id
	content := rau
	ttl := time.Duration(1800) * time.Second
	save := rdb.Set(context.Background(), key, content, ttl)
	if err := save.Err(); err != nil {
		msg := "Gagal simpan sesi!"
		return msg
	} else {
		msg := ""
		return msg
	}
}
