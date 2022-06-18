package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"tiktok-lite/dao"
	"time"
)

var rdb *redis.Client

var ExpireTime = 10 * time.Minute

//InitRedisClient 初始化redis连接
func InitRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //Redis服务监听的端口号
		Password: "12345",
		DB:       0, //选择0-15号redis数据库，
		PoolSize: 100,
	})
	result, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

// SetToken 存储/更新token
func SetToken(token string, userId int) (err error) {
	if dao.DEBUG {
		fmt.Println("redis.SetToken")
	}
	err = rdb.Set(token, userId, ExpireTime).Err()
	if err != nil {
		return err
	}
	fmt.Println("添加/更新数据：", token)
	return nil
}

func GetToken(token string) (userId int, err error) {
	if dao.DEBUG {
		fmt.Println("redis.GetToken")
	}
	result, err := rdb.Get(token).Result()
	if err != nil {
		return -1, err
	}
	userId, err = strconv.Atoi(result)
	if err != nil {
		return -1, err
	}
	err = SetToken(token, userId)
	if err != nil {
		error.Error(err)
	}
	return userId, nil
}
