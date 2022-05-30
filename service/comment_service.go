package service

import (
	"fmt"
	"tiktok-lite/common"
	"tiktok-lite/dao"
)

// convertCommentList 将[]dao.Comment转换为[]common.Comment
func convertCommentList(comments []dao.Comment, userId int) []common.Comment {
	if dao.DEBUG {
		fmt.Println("service.convertCommentList")
	}

	var commentList []common.Comment
	for _, c := range comments {
		user, _ := GetUserInfoById(c.UserId, userId)
		comment := common.Comment{Id: int64(c.CommentId), User: user, Content: c.Content, CreateDate: c.CreateTime.Format("01-02")}
		commentList = append(commentList, comment)
	}
	return commentList
}

// CheckComment 检验评论是否由用户发表
func CheckComment(commentId int, videoId int, userId int) bool {
	if DEBUG {
		fmt.Println("service.CheckComment")
	}

	return dao.GetComment(commentId, videoId, userId) > 0
}

// UserCommentVideo 给视频评论
func UserCommentVideo(userId int, videoId int, content string) (common.Comment, bool) {
	if DEBUG {
		fmt.Println("service.UserCommentVideo")
	}

	c, success := dao.AddComment(userId, videoId, content)
	user, _ := GetUserInfoById(c.UserId, userId)
	// 转换格式
	comment := common.Comment{Id: int64(c.CommentId), User: user, Content: c.Content, CreateDate: c.CreateTime.Format("01-02")}
	return comment, success
}

// DeleteCommentById 删除评论
func DeleteCommentById(commentId int) bool {
	if DEBUG {
		fmt.Println("service.DeleteCommentById")
	}

	// 移除评论
	return dao.RemoveComment(commentId)
}

// GetCommentListById 获取视频评论
func GetCommentListById(videoId int, userId int) []common.Comment {
	if DEBUG {
		fmt.Println("service.GetCommentListById")
	}

	comments := dao.GetVideoCommentList(videoId)
	// 将dao.Comment转换为common.Comment
	commentList := convertCommentList(comments, userId)
	return commentList
}
