package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"lottery/conf"
	"sync"
	"time"
)

var rdsLock sync.Mutex

var cacheInstance *RedisConn

type RedisConn struct {
	pool      *redis.Pool
	showDebug bool
}

func (r *RedisConn) Do(commandName string, args ...interface{}) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()

	t1 := time.Now().UnixNano()
	reply, err := conn.Do(commandName, args)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Println(err)
		}
	}
	t2 := time.Now().UnixNano()
	if r.showDebug {
		log.Printf("[redis][err][%dus]cmd=%s,err=%s,args=%v,reply=%s\n",
			(t2-t1)/1000, commandName, err, args, reply)
	}
	return reply, err
}

func GetCache() *RedisConn {
	if cacheInstance != nil {
		return cacheInstance
	}
	rdsLock.Lock()
	defer rdsLock.Unlock()
	if cacheInstance != nil {
		return cacheInstance
	}
	return instanceCache()
}

func instanceCache() *RedisConn {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.RedisCache.Host, conf.RedisCache.Port))
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			return dial, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         10000,
		MaxActive:       10000,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}

	cacheInstance = &RedisConn{
		pool:      pool,
		showDebug: true,
	}
	return cacheInstance
}
