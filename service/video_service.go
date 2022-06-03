package service

import (
	"fmt"
	"strconv"
	"tiktok-lite/common"
	"tiktok-lite/dao"
	"time"
)

// convertVideoList 将[]dao.Video转换为[]common.Video
func convertVideoList(videos []dao.Video, userId int) []common.Video {
	if dao.DEBUG {
		fmt.Println("service.convertVideoList")
	}

	var videoList []common.Video
	for _, v := range videos {
		user, _ := GetUserInfoById(v.UserId, userId)
		favoriteCount := dao.GetLikedCountByVideoId(v.VideoId)
		commentCount := dao.GetCommentCountByVideoId(v.VideoId)
		isFavorite := dao.GetFavorite(userId, v.VideoId) > 0
		video := common.Video{Id: int64(v.VideoId), Author: user, PlayUrl: v.PlayUrl, CoverUrl: v.CoverUrl, FavoriteCount: favoriteCount, CommentCount: commentCount, IsFavorite: isFavorite, Title: v.Title}
		videoList = append(videoList, video)
	}
	return videoList
}

// GetFeedVideoList 获取feed流
// 返回 视频流,所有视频中最早的时间
func GetFeedVideoList(latestTime string) ([]common.Video, int64) {
	if DEBUG {
		fmt.Println("service.GetFeedVideoList")
	}

	var (
		videoList  []common.Video
		beforeTime time.Time
	)

	if latestTime == "" {
		beforeTime = time.Now()
	} else {
		// 转换时间
		unixTime, err := strconv.Atoi(latestTime)
		if err != nil {
			return videoList, time.Now().Unix()
		}
		beforeTime = time.Unix(int64(unixTime/1000), 0)
	}
	// 获取feed流视频
	videos := dao.GetFeedVideoList(beforeTime)
	fmt.Printf("Feed流：%v\n", videos)
	if len(videos) == 0 {
		return videoList, time.Now().Unix()
	}
	// 将dao.Video转成common.Video
	videoList = convertVideoList(videos, 0)
	// 视频中最早的时间
	earliestTime := videos[len(videos)-1].UpdateTime.Unix() * 1000
	return videoList, earliestTime
}

// GetFeedVideoListLogin 获取feed流
// 返回 视频流,所有视频中最早的时间
func GetFeedVideoListLogin(userId int, latestTime string) ([]common.Video, int64) {
	if DEBUG {
		fmt.Println("service.GetFeedVideoListLogin")
	}

	var (
		videoList  []common.Video
		beforeTime time.Time
	)

	if latestTime == "" {
		beforeTime = time.Now()
	} else {
		// 转换时间
		unixTime, err := strconv.Atoi(latestTime)
		if err != nil {
			return videoList, time.Now().Unix()
		}
		beforeTime = time.Unix(int64(unixTime/1000), 0)
	}
	// 获取登录feed流
	videos := dao.GetFeedVideoListLogin(userId, beforeTime)
	fmt.Printf("登录Feed流：%v\n", videos)
	if len(videos) == 0 {
		return videoList, time.Now().Unix()
	}
	// 将dao.Video转成common.Video
	videoList = convertVideoList(videos, 0)
	// 视频中最早的时间
	earliestTime := videos[len(videos)-1].UpdateTime.Unix() * 1000
	return videoList, earliestTime
}

// AddVideoInfo 添加视频信息
func AddVideoInfo(userId int, playUrl string, coverUrl string, title string) bool {
	if DEBUG {
		fmt.Println("service.AddVideoInfo")
	}

	return dao.AddVideo(userId, playUrl, coverUrl, title)

}

// GetVideoListById 获取用户投稿视频
func GetVideoListById(userId int) []common.Video {
	if DEBUG {
		fmt.Println("service.GetVideoListById")
	}

	var videoList []common.Video
	// 获取登录feed流
	videos := dao.GetVideoListById(userId)
	fmt.Printf("登录Feed流：%v\n", videos)
	if len(videos) == 0 {
		return videoList
	}
	// 将dao.Video转成common.Video
	videoList = convertVideoList(videos, 0)
	return videoList
}