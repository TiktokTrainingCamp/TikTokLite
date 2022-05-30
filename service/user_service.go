package service

import (
	"fmt"
	"strconv"
	"tiktok-lite/common"
	"tiktok-lite/dao"
)

var DEBUG = true

// generateToken 生成token
func generateToken(username string, password string) string {
	if dao.DEBUG {
		fmt.Println("service.generateToken")
	}

	//TODO 实现token算法
	return username + strconv.Itoa(len(username)) + password + strconv.Itoa(len(password))
}

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

// ValidateToken 验证token
func ValidateToken(token string) (int, bool) {
	if DEBUG {
		fmt.Println("service.ValidateToken")
	}

	// 查询token
	userId, tokenInfo := dao.GetToken(token)
	if userId == 0 {
		fmt.Println("token已过期")
		return 0, false
	}
	// 更新token
	success := dao.RefreshToken(tokenInfo)
	if !success {
		return 0, false
	}
	return userId, true
}

// ClearAllToken 清除token信息
func ClearAllToken() bool {
	if DEBUG {
		fmt.Println("service.ClearAllToken")
	}

	return dao.ClearToken()
}

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
		userId int
		token  string
	)
	// 添加用户
	userId = dao.AddUser(username, username, password)
	if userId == 0 {
		return userId, token, false
	}
	// 生成token
	token = generateToken(username, password)
	// 添加token
	success := dao.AddToken(userId, token)
	if !success {
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
		userId int
		token  string
	)
	// 验证账号密码
	userId = dao.ValidateUser(username, password)
	if userId == 0 {
		return userId, token, false
	}
	// 生成token
	token = generateToken(username, password)
	// 添加token
	success := dao.AddToken(userId, token)
	if !success {
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
