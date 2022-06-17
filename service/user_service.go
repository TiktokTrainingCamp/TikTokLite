package service

import (
	"fmt"
	"tiktok-lite/common"
	"tiktok-lite/dao"
	"tiktok-lite/redis"
	"tiktok-lite/util"
)

var DEBUG = true
var key = "ThisIsPrivateKey"

// convertUser 将dao.User转为common.User
func convertUser(userInfo dao.User, userId int) common.User {
	if dao.DEBUG {
		fmt.Println("service.convertUser")
	}

	//TODO 获赞数,app里有但是接口里没有
	followCount := getFollowCount(userInfo.UserId)
	followerCount := getFollowerCount(userInfo.UserId)
	isFollow := dao.GetRelation(userInfo.UserId, userId) > 0
	user := common.User{Id: int64(userInfo.UserId), Name: userInfo.Name, FollowCount: followCount, FollowerCount: followerCount, IsFollow: isFollow}
	return user
}

// convertUserList 将[]dao.User转为[]common.User
func convertUserList(users []dao.User, userId int) []common.User {
	if dao.DEBUG {
		fmt.Println("service.convertUserList")
	}

	var userList []common.User
	for _, u := range users {
		user := convertUser(u, userId)
		userList = append(userList, user)
	}
	return userList
}

// generateToken 生成token
func generateToken(userId int) (string, bool) {
	if dao.DEBUG {
		fmt.Println("service.generateToken")
	}

	//实现token算法
	token, success := util.GenerateToken(userId, key)
	return token, success
}

// ValidateToken 验证token
func ValidateToken(token string) (int, bool) {
	if DEBUG {
		fmt.Println("service.ValidateToken")
	}

	var (
		userId  int
		success bool
	)
	// 解析token
	_, success = util.ParseToken(token, key)
	if !success {
		fmt.Println("token解析失败")
		return 0, false
	}
	// 查询token
	userId, err := redis.GetToken(token)
	if userId <= 0 || err != nil {
		fmt.Println("token已过期")
		return 0, false
	}
	// 更新token
	err = redis.SetToken(token, userId)
	if err != nil {
		return 0, false
	}
	return userId, true
}

// ClearAllToken 清除token信息
// Discarded method
//func ClearAllToken() bool {
//	if DEBUG {
//		fmt.Println("service.ClearAllToken")
//	}
//
//	return dao.ClearToken()
//}

// UserExist 判断用户是否存在
func UserExist(username string) bool {
	if DEBUG {
		fmt.Println("service.UserExist")
	}

	success := dao.FindUser(username)
	return success
}

// UserRegister 注册用户
func UserRegister(username string, password string) (int, string, bool) {
	if DEBUG {
		fmt.Println("service.UserRegister")
	}

	var (
		userId  int
		token   string
		success bool
	)
	// 加密密码
	password, err := util.HashAndSalt(password)
	if err != nil {
		return userId, token, false
	}
	// 添加用户
	userId = dao.AddUser(username, username, password)
	if userId == 0 {
		return userId, token, false
	}
	// 生成token
	token, success = generateToken(userId)
	if !success {
		return userId, token, false
	}
	// 添加token
	err = redis.SetToken(token, userId)
	if err != nil {
		return userId, token, false
	}
	return userId, token, true
}

// UserLogin 用户登录
func UserLogin(username string, password string) (int, string, bool) {
	if DEBUG {
		fmt.Println("service.UserLogin")
	}

	var (
		userId  int
		token   string
		success bool
	)
	// 验证账号密码
	userId, encrpPassword := dao.GetPasswordByUsername(username)
	if userId == 0 {
		return userId, token, false
	}
	success = util.ComparePasswords(encrpPassword, password)
	if !success {
		return userId, token, false
	}
	//userId = dao.ValidateUser(username, password)
	// 生成token
	token, success = generateToken(userId)
	if !success {
		return userId, token, false
	}
	// 添加token
	err := redis.SetToken(token, userId)
	if err != nil {
		return userId, token, false
	}
	return userId, token, true
}

// GetUserInfo 通过token获取用户信息
//func GetUserInfo(token string) (common.User, bool) {
//	if DEBUG {
//		fmt.Println("service.GetUserInfo")
//	}
//	var user common.User
//
//	user, success = GetUserInfoById(userId)
//	if !success {
//		return user, false
//	}
//	return user, true
//}

// GetUserInfoById 通过user_id获取用户信息
func GetUserInfoById(targetUserId int, userId int) (common.User, bool) {
	if DEBUG {
		fmt.Println("service.GetUserInfoById")
	}

	var user common.User
	// 查询用户
	userInfo, success := dao.GetUserInfo(targetUserId)
	if !success {
		return user, false
	}
	// 将dao.User转为common.User
	user = convertUser(userInfo, userId)
	return user, true
}
