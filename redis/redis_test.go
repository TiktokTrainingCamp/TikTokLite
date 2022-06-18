package redis

import (
	"fmt"
	"testing"
	"tiktok-lite/dao"
	"time"
)

func TestRedis(t *testing.T) {
	ExpireTime = 5 * time.Second
	dao.DEBUG = true

	err := InitRedisClient()
	if err != nil {
		panic(err)
	}

	err = SetToken("user123", 1)
	if err != nil {
		error.Error(err)
	}

	time.Sleep(3 * time.Second)

	userId, err := GetToken("user123")
	if err != nil {
		error.Error(err)
	}
	fmt.Printf("userId: %d\n", userId)

	time.Sleep(3 * time.Second)

	userId, err = GetToken("user123")
	if err != nil {
		error.Error(err)
	}
	fmt.Printf("userId: %d\n", userId)

	time.Sleep(6 * time.Second)

	userId, err = GetToken("user123")
	if err != nil {
		error.Error(err)
	}
	fmt.Printf("userId: %d\n", userId)

	rdb.Set("123", "asd", 0)
	userId, err = GetToken("123")
	if err != nil {
		fmt.Println(err)
	}
}
