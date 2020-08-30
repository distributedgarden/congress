package adapters

import (
	"github.com/gomodule/redigo/redis"
	"os"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(os.Getenv("REDIS_NETWORK"), os.Getenv("REDIS_ADDRESS"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
