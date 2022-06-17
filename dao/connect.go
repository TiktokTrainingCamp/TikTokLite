package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var DEBUG = false

// ConnectDB 连接mysql数据库
func ConnectDB() (err error) {
	dsn := "root:zjj1028xym@tcp(127.0.0.1:3306)/tiktok?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
	}
	return
}

// FindUser 查找用户
func FindUser(username string) bool {
	if DEBUG {
		fmt.Println("dao.FindUser")
	}
	result := db.Where(&User{Username: username}).Limit(1).Find(&User{})
	if result.Error != nil {
		fmt.Println("dao.FindUser", result.Error)
	}
	return result.RowsAffected > 0
}

// ValidateUser 验证用户
// Discarded method
//func ValidateUser(username string, password string) int {
//	if DEBUG {
//		fmt.Println("dao.ValidateUser")
//	}
//	var userInfo User
//	result := db.Where(&User{Username: username, Password: password}).Limit(1).Find(&userInfo)
//	if result.Error != nil {
//		fmt.Println("dao.ValidateUser", result.Error)
//	}
//	return userInfo.UserId
//}

// GetPasswordByUsername 获取用户存储密码
func GetPasswordByUsername(username string) (int, string) {
	if DEBUG {
		fmt.Println("dao.GetPasswordByUsername")
	}
	var userInfo User
	result := db.Where(&User{Username: username}).Limit(1).Find(&userInfo)
	if result.Error != nil {
		fmt.Println("dao.ValidateUser", result.Error)
	}
	return userInfo.UserId, userInfo.Password
}

// AddUser 添加用户
func AddUser(name string, username string, password string) int {
	if DEBUG {
		fmt.Println("dao.AddUser")
	}
	var user = User{Name: name, Username: username, Password: password}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("dao.AddUser", result.Error)
		return 0
	}
	return user.UserId
}

// GetUserInfo 获取用户信息
func GetUserInfo(userId int) (User, bool) {
	if DEBUG {
		fmt.Println("dao.GetUserInfo")
	}

	//db = db.Debug()
	var user User
	result := db.Where(&User{UserId: userId}).Limit(1).Find(&user)
	if result.Error != nil {
		fmt.Println("dao.GetUserInfo", result.Error)
		return user, false
	}
	return user, true
}

// GetToken 查找token
// Discarded method
//func GetToken(userToken string) (int, Token) {
//	if DEBUG {
//		fmt.Println("dao.GetToken")
//	}
//	var tokenInfo Token
//	now := time.Now()
//	result := db.Where(&Token{UserToken: userToken}).Where("expire_time > ?", now).Limit(1).Find(&tokenInfo)
//	if result.Error != nil {
//		fmt.Println("dao.GetToken", result.Error)
//	}
//	return tokenInfo.UserId, tokenInfo
//}

// AddToken 添加token
// Discarded method
//func AddToken(userId int, userToken string) bool {
//	if DEBUG {
//		fmt.Println("dao.AddToken")
//	}
//	expireTime := time.Now().Add(tokenValidDuration) // token的失效时间是现在的时间后延十分钟
//	var tokenInfo = Token{UserToken: userToken, UserId: userId, ExpireTime: expireTime}
//	result := db.Create(&tokenInfo)
//	if result.Error != nil {
//		fmt.Println("dao.AddToken", result.Error)
//		return false
//	}
//	return true
//}

// RemoveToken 删除指定token
// Discarded method
//func RemoveToken(tokenId int) bool {
//	if DEBUG {
//		fmt.Println("dao.RemoveToken")
//	}
//	result := db.Delete(&Token{}, tokenId)
//	if result.Error != nil {
//		fmt.Println("dao.RemoveToken", result.Error)
//		return false
//	}
//	return true
//}

// ClearToken 删除所有token
// Discarded method
//func ClearToken() bool {
//	if DEBUG {
//		fmt.Println("dao.ClearToken")
//	}
//	result := db.Where("1 = 1").Delete(&Token{})
//	if result.Error != nil {
//		fmt.Println("dao.ClearToken", result.Error)
//		return false
//	}
//	return true
//}

// RefreshToken 更新token
// Discarded method
//func RefreshToken(tokenInfo Token) bool {
//	if DEBUG {
//		fmt.Println("dao.RefreshToken")
//	}
//	tokenInfo.ExpireTime = time.Now().Add(tokenValidDuration) // token的失效时间是现在的时间后延十分钟
//	result := db.Save(&tokenInfo)
//	if result.Error != nil {
//		fmt.Println("dao.RefreshToken", result.Error)
//		return false
//	}
//	return true
//}

// GetFeedVideoList 获取feed流(按时间倒序)
// 上传时间早于给定时间
func GetFeedVideoList(latestTime time.Time) []Video {
	if DEBUG {
		fmt.Println("dao.GetFeedVideoList")
	}
	var videoList []Video
	result := db.Where("update_time < ?", latestTime).Order("update_time desc").Limit(3).Find(&videoList)
	if result.Error != nil {
		fmt.Println("dao.GetFeedVideoList", result.Error)
		return []Video{}
	}
	return videoList
}

// GetFeedVideoListLogin 获取登录后的feed流(按时间倒序)
// 上传时间早于给定时间且不能刷到自己的视频
func GetFeedVideoListLogin(userId int, latestTime time.Time) []Video {
	if DEBUG {
		fmt.Println("dao.GetFeedVideoListLogin")
	}
	var videoList []Video
	result := db.Where("user_id <> ?", userId).Where("update_time < ?", latestTime).Order("update_time desc").Find(&videoList)
	if result.Error != nil {
		fmt.Println("dao.GetFeedVideoListLogin", result.Error)
		return []Video{}
	}
	return videoList
}

// GetVideoListById 获取用户上传的视频
func GetVideoListById(userId int) []Video {
	if DEBUG {
		fmt.Println("dao.GetVideoListById")
	}
	var videoList []Video
	result := db.Where(&Video{UserId: userId}).Order("update_time desc").Find(&videoList)
	if result.Error != nil {
		fmt.Println("dao.GetVideoListById", result.Error)
		return []Video{}
	}
	return videoList
}

// AddVideo 添加视频
func AddVideo(userId int, playUrl string, coverUrl string, title string) bool {
	if DEBUG {
		fmt.Println("dao.AddVideo")
	}
	result := db.Create(&Video{UserId: userId, PlayUrl: playUrl, CoverUrl: coverUrl, Title: title, UpdateTime: time.Now()})
	if result.Error != nil {
		fmt.Println("dao.AddVideo", result.Error)
		return false
	}
	return true
}

// GetFavorite 获取点赞
func GetFavorite(userId int, videoId int) int {
	if DEBUG {
		fmt.Println("dao.GetFavorite")
	}
	var favorite Favorite
	result := db.Where(&Favorite{UserId: userId, VideoId: videoId}).Limit(1).Find(&favorite)
	if result.Error != nil {
		fmt.Println("dao.GetFavorite", result.Error)
	}
	return favorite.Id
}

// AddFavorite 添加点赞
func AddFavorite(userId int, videoId int) bool {
	if DEBUG {
		fmt.Println("dao.AddFavorite")
	}
	result := db.Create(&Favorite{UserId: userId, VideoId: videoId})
	if result.Error != nil {
		fmt.Println("dao.AddFavorite", result.Error)
		return false
	}
	return true
}

// RemoveFavorite 移除点赞
func RemoveFavorite(id int) bool {
	if DEBUG {
		fmt.Println("dao.RemoveFavorite")
	}
	result := db.Delete(&Favorite{}, id)
	if result.Error != nil {
		fmt.Println("dao.RemoveFavorite", result.Error)
		return false
	}
	return true
}

// GetFavoriteVideoList 获取用户点赞的视频列表
func GetFavoriteVideoList(userId int) []Video {
	if DEBUG {
		fmt.Println("dao.GetFavoriteVideoList")
	}
	var videoList []Video
	// 子查询
	result := db.Where("video_id in (?)", db.Table("favorites").Select("video_id").Where(&Favorite{UserId: userId})).Find(&videoList)
	if result.Error != nil {
		fmt.Println("dao.GetFavoriteVideos", result.Error)
		return []Video{}
	}
	return videoList
}

// GetFavoriteCountByVideoId 视频获赞数
func GetFavoriteCountByVideoId(videoId int) int64 {
	if DEBUG {
		fmt.Println("dao.GetFavoriteCountByVideoId")
	}
	var count int64
	result := db.Model(&Favorite{}).Where(&Favorite{VideoId: videoId}).Count(&count)
	if result.Error != nil {
		fmt.Println("dao.GetFavoriteCountByVideoId", result.Error)
	}
	return count
}

// GetFavoriteCountByUserId 用户获赞数
func GetFavoriteCountByUserId(userId int) int64 {
	if DEBUG {
		fmt.Println("dao.GetFavoriteCountByUserId")
	}
	var count int64
	result := db.Model(&Favorite{}).Where("video_id in (?)", db.Table("videos").Select("video_id").Where(&Video{UserId: userId})).Count(&count)
	if result.Error != nil {
		fmt.Println("dao.GetFavoriteCountByUserId", result.Error)
	}
	return count
}

// AddComment 添加评论
func AddComment(userId int, videoId int, content string) (Comment, bool) {
	if DEBUG {
		fmt.Println("dao.AddComment")
	}
	comment := Comment{UserId: userId, VideoId: videoId, Content: content, CreateTime: time.Now()}
	result := db.Create(&comment)
	if result.Error != nil {
		fmt.Println("dao.AddComment", result.Error)
		return comment, false
	}
	return comment, true
}

// RemoveComment 移除评论
func RemoveComment(CommentId int) bool {
	if DEBUG {
		fmt.Println("dao.RemoveComment")
	}
	result := db.Delete(&Comment{}, CommentId)
	if result.Error != nil {
		fmt.Println("dao.RemoveComment", result.Error)
		return false
	}
	return true
}

// GetComment 获取评论
func GetComment(commentId int, videoId int, userId int) int {
	if DEBUG {
		fmt.Println("dao.GetComment")
	}
	var comment Comment
	result := db.Where(&Comment{UserId: userId, VideoId: videoId, CommentId: commentId}).Limit(1).Find(&comment)
	if result.Error != nil {
		fmt.Println("dao.GetComment", result.Error)
	}
	return comment.CommentId
}

// GetVideoCommentList 获取视频评论
func GetVideoCommentList(videoId int) []Comment {
	if DEBUG {
		fmt.Println("dao.GetVideoCommentList")
	}
	var commentList []Comment
	result := db.Where(&Comment{VideoId: videoId}).Order("create_time desc").Find(&commentList)
	if result.Error != nil {
		fmt.Println("dao.GetVideoCommentList", result.Error)
		return []Comment{}
	}
	return commentList
}

// GetCommentCountByVideoId 视频评论数
func GetCommentCountByVideoId(videoId int) int64 {
	if DEBUG {
		fmt.Println("dao.GetCommentCountByVideoId")
	}
	var count int64
	result := db.Model(&Comment{}).Where(&Comment{VideoId: videoId}).Count(&count)
	if result.Error != nil {
		fmt.Println("dao.GetCommentCountByVideoId", result.Error)
	}
	return count
}

// GetRelation 判断是否关注
func GetRelation(followId int, followerId int) int {
	if DEBUG {
		fmt.Println("dao.GetRelation")
	}
	var relation Relation
	result := db.Where(&Relation{FollowId: followId, FollowerId: followerId}).Limit(1).Find(&relation)
	if result.Error != nil {
		fmt.Println("dao.GetRelation", result.Error)
	}
	return relation.Id
}

// AddRelation 添加关注
func AddRelation(followId int, followerId int) bool {
	if DEBUG {
		fmt.Println("dao.AddRelation")
	}
	result := db.Create(&Relation{FollowId: followId, FollowerId: followerId})
	if result.Error != nil {
		fmt.Println("dao.AddRelation", result.Error)
		return false
	}
	return true
}

// RemoveRelation 移除关注
func RemoveRelation(id int) bool {
	if DEBUG {
		fmt.Println("dao.RemoveRelation")
	}
	result := db.Delete(&Relation{}, id)
	if result.Error != nil {
		fmt.Println("dao.RemoveRelation", result.Error)
		return false
	}
	return true
}

// GetFollowCount 查询关注数
func GetFollowCount(userId int) int64 {
	if DEBUG {
		fmt.Println("dao.GetFollowCount")
	}
	var count int64
	result := db.Model(&Relation{}).Where(&Relation{FollowerId: userId}).Count(&count)
	if result.Error != nil {
		fmt.Println("dao.GetFollowCount", result.Error)
	}
	return count
}

// GetFollowerCount 查询关注数
func GetFollowerCount(userId int) int64 {
	if DEBUG {
		fmt.Println("dao.GetFollowerCount")
	}
	var count int64
	result := db.Model(&Relation{}).Where(&Relation{FollowId: userId}).Count(&count)
	if result.Error != nil {
		fmt.Println("dao.GetFollowerCount", result.Error)
	}
	return count
}

// GetFollowListById 获取关注列表
func GetFollowListById(userId int) []User {
	if DEBUG {
		fmt.Println("dao.GetFollowListById")
	}
	var followList []User
	result := db.Where("user_id in (?)", db.Table("relations").Select("follow_id").Where(&Relation{FollowerId: userId})).Find(&followList)
	if result.Error != nil {
		fmt.Println("dao.GetFollowListById", result.Error)
	}
	return followList
}

// GetFollowerListById 获取粉丝列表
func GetFollowerListById(userId int) []User {
	if DEBUG {
		fmt.Println("dao.GetFollowerListById")
	}
	var followerList []User
	result := db.Where("user_id in (?)", db.Table("relations").Select("follower_id").Where(&Relation{FollowId: userId})).Find(&followerList)
	if result.Error != nil {
		fmt.Println("dao.GetFollowerListById", result.Error)
	}
	return followerList
}
