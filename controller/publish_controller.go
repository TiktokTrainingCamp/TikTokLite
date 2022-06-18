package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"tiktok-lite/common"
	"tiktok-lite/object"
	"tiktok-lite/service"
	"tiktok-lite/util"
	"time"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

var IP string // 服务器ip,服务器启动时会由main函数修改

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
	videoName := fmt.Sprintf("%d_%s_%s", userId, time.Now().Format("2006-01-02-15-04-05"), filename)
	fmt.Println("视频文件名：", videoName)
	saveFile := filepath.Join("./public/", videoName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "保存视频：" + err.Error(),
		})
		return
	}

	// 保存预览图片
	suffix := ".jpg"
	imageName := strings.Split(videoName, ".")[0] + suffix
	err = util.GeneratePreview(videoName, imageName)
	if err != nil {
		fmt.Println("生成封面：", err.Error())
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "保存视频：" + err.Error(),
		})
		return
	}

	// 上传到云存储服务器
	coverUrl, err := object.UploadImgFile(imageName)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "对象云存储：" + err.Error(),
		})
		return
	}
	playUrl, err := object.UploadVideoFile(videoName)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "对象云存储：" + err.Error(),
		})
		return
	}
	// 本地路径
	//playUrl := "http://" + IP + "/static/" + finalName
	//coverUrl := "http://" + IP + "/static/" + imageName

	// 添加视频
	title := c.PostForm("title")
	success = service.AddVideoInfo(userId, playUrl, coverUrl, title)
	if !success {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  videoName + "Uploaded fail",
		})
		return
	} else {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  "Uploaded successfully",
		})
		return
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
		return
	}

	videoList := service.GetVideoListById(userId)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
