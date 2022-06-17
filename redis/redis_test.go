package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	expireTime = 10 * time.Second

	err := InitRedisClient()
	if err != nil {
		panic(err)
	}

	err = SetToken("user123", 1)
	if err != nil {
		error.Error(err)
	}

	time.Sleep(5 * time.Second)

	userId, err := GetToken("user123")
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

	time.Sleep(11 * time.Second)

	userId, err = GetToken("user123")
	if err != nil {
		error.Error(err)
	}
	fmt.Printf("userId: %d\n", userId)
}
