package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok-lite/common"
	"tiktok-lite/service"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTimeString := c.Query("latest_time")
	var (
		latestTime int
		err        error
	)
	if latestTimeString == "" {
		latestTime = int(time.Now().Unix() * 1000)
	} else {
		latestTime, err = strconv.Atoi(latestTimeString)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "latest_time parse failed"},
			})
			return
		}
	}
	fmt.Printf("lasest_time:%d\n", latestTime)
	token := c.Query("token")

	var videoList []common.Video
	var beforeTime = time.Now().Unix()

	if token != "" { // 登录状态
		// 验证token
		userId, success := service.ValidateToken(token)
		if !success {
			c.JSON(http.StatusOK, UserResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"},
			})
			return
		}
		videoList, beforeTime = service.GetFeedVideoListLogin(userId, latestTime)
	} else { // 非登录状态
		videoList, beforeTime = service.GetFeedVideoList(latestTime)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  int64(beforeTime),
	})
}
