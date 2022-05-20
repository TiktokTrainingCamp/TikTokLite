package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"simple-demo/util"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var id = int64(1)
var ip = "10.11.31.230:8080"

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		fmt.Println("user not exist")
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	// save video file
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// save preview image
	imageName, err := util.GeneratePreview(finalName)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// add to database
	id += 1
	DemoVideos = append(DemoVideos, Video{
		Id:            id,
		Author:        DemoUser,
		PlayUrl:       "http://" + ip + "/static/" + finalName,
		CoverUrl:      "http://" + ip + "/static/" + imageName,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	})

	// response
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
