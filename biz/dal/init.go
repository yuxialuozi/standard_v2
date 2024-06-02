package dal

import (
	"cwgo_test/biz/dal/mysql"
	"cwgo_test/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
