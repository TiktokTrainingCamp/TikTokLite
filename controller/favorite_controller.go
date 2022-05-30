package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok-lite/common"
	"tiktok-lite/service"
)

// FavoriteAction 点赞/取消
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil { // 解析video_id失败
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Video_id parse failed"})
		return
	}
	actionType := c.Query("action_type")

	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"})
		return
	}
	if actionType == "1" { // 点赞
		success = service.UserLikeVideo(userId, videoId)
	} else { // 取消点赞
		success = service.UserUnlikeVideo(userId, videoId)
	}
	if !success {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Action failed"})
	} else {
		fmt.Println("1 点赞/2 取消点赞：", actionType, " 成功")
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	}

}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"})
		return
	}
	videoList := service.GetFavoriteVideoListById(userId)

	c.JSON(http.StatusOK, VideoListResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: videoList,
	})
}
