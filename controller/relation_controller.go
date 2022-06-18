package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok-lite/common"
	"tiktok-lite/service"
)

type UserListResponse struct {
	common.Response
	UserList []common.User `json:"user_list"`
}

// RelationAction 关注/取消关注
func RelationAction(c *gin.Context) {
	followId, err := strconv.Atoi(c.Query("to_user_id"))
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "Parse to_user_id failed"},
		})
		return
	}
	token := c.Query("token")
	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"},
		})
		return
	}
	actionType := c.Query("action_type")
	if actionType == "1" { //关注
		success = service.FollowUser(userId, followId)
	} else { // 取消关注
		success = service.UnFollowUser(userId, followId)
	}
	if !success {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Action failed"})
	} else {
		fmt.Println("1 关注/2 取消关注：", actionType, " 成功")
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	}
}

// FollowList 关注列表
func FollowList(c *gin.Context) {
	token := c.Query("token")
	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"})
		return
	}
	// 获取user_id
	targetUserId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil { // 解析user_id失败
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "user_id parse failed"})
		return
	}
	// 获取关注列表
	userList := service.GetFollowListById(targetUserId, userId)

	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{StatusCode: 0},
		UserList: userList,
	})
}

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"})
		return
	}
	// 获取user_id
	targetUserId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil { // 解析user_id失败
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "user_id parse failed"})
		return
	}
	// 获取粉丝列表
	userList := service.GetFollowerListById(targetUserId, userId)

	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{StatusCode: 0},
		UserList: userList,
	})
}
