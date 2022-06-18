package service

import (
	"fmt"
	"tiktok-lite/common"
	"tiktok-lite/dao"
)

// GetFollowCount 获取关注数
func getFollowCount(userId int) int64 {
	if dao.DEBUG {
		fmt.Println("service.GetFollowCount")
	}
	return dao.GetFollowCount(userId)
}

// GetFollowerCount 获取粉丝数
func getFollowerCount(userId int) int64 {
	if dao.DEBUG {
		fmt.Println("service.GetFollowerCount")
	}
	return dao.GetFollowerCount(userId)
}

// FollowUser 关注用户
func FollowUser(userId int, followId int) bool {
	if DEBUG {
		fmt.Println("service.FollowUser")
	}

	// 查询是否关注
	isFollow := dao.GetRelation(followId, userId) > 0
	if isFollow {
		return true
	}
	// 关注
	return dao.AddRelation(followId, userId)
}

// UnFollowUser 取关用户
func UnFollowUser(userId int, followId int) bool {
	if DEBUG {
		fmt.Println("service.UnFollowUser")
	}

	// 查询是否关注
	id := dao.GetRelation(followId, userId)
	if id == 0 {
		return true
	}
	// 取消关注
	return dao.RemoveRelation(id)
}

// GetFollowListById 获取用户的关注列表
func GetFollowListById(targetUserId int, userId int) []common.User {
	if DEBUG {
		fmt.Println("service.GetFollowListById")
	}

	// 获取用户关注列表
	users := dao.GetFollowListById(targetUserId)
	// 转换格式
	userList := convertUserList(users, userId)
	return userList
}

// GetFollowerListById 获取用户的粉丝列表
func GetFollowerListById(targetUserId int, userId int) []common.User {
	if DEBUG {
		fmt.Println("service.GetFollowerListById")
	}

	// 获取用户粉丝列表
	users := dao.GetFollowerListById(targetUserId)
	// 转换格式
	userList := convertUserList(users, userId)
	return userList
}
