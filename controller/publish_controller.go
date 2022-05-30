package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"tiktok-lite/common"
	"tiktok-lite/service"
	"tiktok-lite/util"
	"time"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

var IP = "10.11.31.123:8080" // 服务器ip,服务器启动时会由main函数修改

// Publish 投稿视频
func Publish(c *gin.Context) {
	fmt.Println("投稿信息：", c)
	token := c.PostForm("token")

	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"},
		})
		return
	}

	// 获取视频
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "获取视频：" + err.Error(),
		})
		return
	}

	// 将视频保存到服务器
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s_%s", userId, time.Now().Format("2006-01-02-15-04-05"), filename)
	fmt.Println("filename", finalName)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "保存视频：" + err.Error(),
		})
		return
	}

	// 保存预览图片
	imageName, err := util.GeneratePreview(finalName)
	if err != nil {
		fmt.Println("生成封面：", err.Error())
	}

	// 添加视频
	title := c.PostForm("title")
	playUrl := "http://" + IP + "/static/" + finalName
	coverUrl := "http://" + IP + "/static/" + imageName
	success = service.AddVideoInfo(userId, playUrl, coverUrl, title)
	if !success {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  finalName + "Uploaded fail",
		})
	} else {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  "Uploaded successfully",
		})
	}
}

// PublishList 获取投稿列表
func PublishList(c *gin.Context) {
	token := c.Query("token")
	// 验证token
	_, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"},
		})
		return
	}

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "user_id parse failed"},
		})
	}

	videoList := service.GetVideoListById(userId)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
