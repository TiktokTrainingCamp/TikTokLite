package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"strings"
	"tiktok-lite/controller"
	"tiktok-lite/dao"
	"tiktok-lite/redis"
	"tiktok-lite/service"
	"time"
)

// GetOutBoundIP 获取服务器运行ip
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func main() {
	// 显示ip地址
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("server is running on http://%s:8080\n", ip)
	controller.IP = ip + ":8080"

	// 调试信息控制
	service.DEBUG = true
	dao.DEBUG = false
	redis.ExpireTime = time.Minute * 10

	// 连接mysql数据库
	err = dao.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql连接成功")

	// 连接redis
	err = redis.InitRedisClient()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis初始化成功")

	// 清除token
	//if success := service.ClearAllToken(); success {
	//	fmt.Println("清除token")
	//}

	//dao.TestConnect()
	r := gin.Default()
	initRouter(r)

	err = r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}

}
