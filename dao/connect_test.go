package dao

import (
	"testing"
)

// TestConnect 测试dao接口
func TestConnect(t *testing.T) {
	err := ConnectDB()
	if err != nil {
		panic(err)
	}
	//fmt.Println(GetUser("ColorOfNight1"))
	//fmt.Println(ValidateUser("admin", "123456"))
	//fmt.Println(AddUser("test", "test", "123456"))
	//fmt.Println(AddToken(1, "admin123456"))
	//fmt.Println(RemoveToken(2))
	//userId, token := GetToken("admin123456")
	//fmt.Println(userId, token)
	//fmt.Println(RefreshToken(token))
	//fmt.Println(GetFeedVideoList())
	//mm, _ := time.ParseDuration("-35h")
	//fmt.Println(GetFeedVideoListLogin(1, time.Unix(int64(1653660143), 0)))
	//fmt.Println(GetVideoListById(1))
	//fmt.Println(AddVideo(1, "test url", "test url"))
	//fmt.Println(GetFavorite(2, 8))
	//fmt.Println(AddFavorite(3, 1))
	//fmt.Println(RemoveFavorite(4))
	//fmt.Println(AddComment(2, 1, "测试评论"))
	//fmt.Println(GetVideoCommentList(1))
	//fmt.Println(GetFavoriteVideoList(4))
	//fmt.Println(AddRelation(3, 4))
	//fmt.Println(GetRelation(3, 4))
	//fmt.Println(RemoveRelation(4))
	//fmt.Println(GetFollowCount(2))
	//fmt.Println(GetFollowerCount(1))
	//fmt.Println(GetLikedCountByVideoId(8))
	//fmt.Println(GetCommentCountByVideoId(8))
	//fmt.Println(GetLikedCountByUserId(1))
	//fmt.Println(GetFollowListById(2))
	//fmt.Println(GetFollowerListById(1))

}
