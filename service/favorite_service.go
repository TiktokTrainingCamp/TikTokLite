package service

import (
	"fmt"
	"tiktok-lite/common"
	"tiktok-lite/dao"
)

// UserLikeVideo 给视频点赞
func UserLikeVideo(userId int, videoId int) bool {
	if DEBUG {
		fmt.Println("service.UserLikeVideo")
	}

	// 查询是否点赞
	isLike := dao.GetFavorite(userId, videoId) > 0
	if isLike {
		return true
	}
	// 点赞
	return dao.AddFavorite(userId, videoId)
}

// UserUnlikeVideo 取消视频点赞
func UserUnlikeVideo(userId int, videoId int) bool {
	if DEBUG {
		fmt.Println("service.UserUnlikeVideo")
	}

	// 查询是否点赞
	id := dao.GetFavorite(userId, videoId)
	if id == 0 {
		return true
	}
	// 取消点赞
	return dao.RemoveFavorite(id)
}

// GetFavoriteVideoListById 获取用户的点赞列表
func GetFavoriteVideoListById(userId int) []common.Video {
	if DEBUG {
		fmt.Println("service.UserUnlikeVideo")
	}

	videos := dao.GetFavoriteVideoList(userId)
	videoList := convertVideoList(videos, userId)
	return videoList
}
