package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok-lite/common"
	"tiktok-lite/service"
)

type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if exist := service.UserExist(username); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		userId, token, success := service.UserRegister(username, password)
		fmt.Println("注册")
		if success {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 0},
				UserId:   int64(userId),
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "User register failed"},
			})
		}
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userId, token, success := service.UserLogin(username, password)
	if success {
		fmt.Println("登录成功")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   int64(userId),
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"},
		})
		return
	}

	targetUserId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "user_id parse failed"},
		})
		return
	}
	if user, success := service.GetUserInfoById(targetUserId, userId); success {
		fmt.Printf("获取用户信息：%#v\n", user)
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
