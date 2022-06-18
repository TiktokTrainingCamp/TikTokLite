package dao

import (
	"fmt"
	"testing"
	"time"
)

// TestConnect 测试dao接口
func TestConnect(t *testing.T) {
	DEBUG = true

	err := ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("存在用户 admin：", FindUser("admin"))
	fmt.Println("不存在用户 administer：", FindUser("administer"))

	userId, encrptPassword := GetPasswordByUsername("admin")
	fmt.Println("id为", userId, "的用户 admin 的加密存储密码：", encrptPassword)
	userId, encrptPassword = GetPasswordByUsername("administer")
	fmt.Println("id为", userId, "的不存在用户 administer 的加密存储密码：", encrptPassword)

	fmt.Println("删除用户：", removeUser("testtest"))
	fmt.Println("创建用户的id为：", AddUser("testtest", "testtest", "123456"))
	fmt.Println("重复用户创建：", 0 == AddUser("testtest", "testtest", "123456"))

	user, _ := GetUserInfo(1)
	fmt.Println("获取用户信息：", user)

	fmt.Println("获取feed流：", GetFeedVideoList(time.Now()))
	fmt.Println("获取登录feed流：", GetFeedVideoListLogin(3, time.Now()))

	fmt.Println("获取用户1的投稿列表：", GetVideoListById(1))
	fmt.Println("添加测试视频：", AddVideo(0, "test url", "test url", ""))

	fmt.Println("用户0给视频0点赞：", AddFavorite(0, 0))
	favorId := GetFavorite(0, 0)
	fmt.Println("用户0喜欢视频0：", favorId > 0)
	fmt.Println("用户0取消给视频0的点赞：", RemoveFavorite(favorId))
	fmt.Println("获取用户4的点赞列表：", GetFavoriteVideoList(4))
	fmt.Println("获取视频27的获赞数：", GetFavoriteCountByVideoId(27))
	fmt.Println("获取用户4的获赞数：", GetFavoriteCountByUserId(4))

	comment, success := AddComment(0, -1, "test")
	fmt.Println("用户0给视频-1评论：", success)
	fmt.Println("获取用户0给视频-1的评论：", GetComment(comment.CommentId, -1, 0) > 0)
	fmt.Println("获取视频-1的评论：", GetVideoCommentList(-1))
	fmt.Println("获取视频-1的评论数：", GetCommentCountByVideoId(-1))
	fmt.Println("用户0删除对视频0的评论：", RemoveComment(comment.CommentId))

	fmt.Println("用户-1关注用户0：", AddRelation(0, -1))
	relationId := GetRelation(0, -1)
	fmt.Println("用户-1是否关注用户0：", relationId > 0)
	fmt.Println("用户-1的关注数：", GetFollowCount(-1))
	fmt.Println("用户0的粉丝数：", GetFollowerCount(0))
	fmt.Println("用户-1的关注列表：", GetFollowListById(-1))
	fmt.Println("用户0的粉丝列表：", GetFollowerListById(0))
	fmt.Println("用户-1取关用户0：", RemoveRelation(relationId))

}
