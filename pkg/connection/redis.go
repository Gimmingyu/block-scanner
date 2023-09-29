package connection

import (
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

func opt(addr, pass, db string) *redis.Options {
	option := new(redis.Options)
	option.Addr = addr
	option.Password = pass
	option.DB, _ = strconv.Atoi(db)
	if os.Getenv("APPMODE") == "prod" {
		option.TLSConfig = &tls.Config{
			InsecureSkipVerify: false,
		}
	}

	return option
}

func NewRedisConnection(addr, pass, db string) *redis.Client {
	return redis.NewClient(opt(addr, pass, db))
}
