package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok-lite/common"
	"tiktok-lite/service"
)

type CommentListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

// CommentAction 评论/删除评论
func CommentAction(c *gin.Context) {
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "video_id parse failed"},
		})
		return
	}

	token := c.Query("token")
	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"},
		})
		return
	}

	var comment common.Comment

	actionType := c.Query("action_type")
	if actionType == "1" { // 评论
		content := c.Query("comment_text")
		comment, success = service.UserCommentVideo(userId, videoId, content)
		if !success {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "Comment failed"},
			})
			return
		}
	} else { // 删除评论
		commentId, err := strconv.Atoi(c.Query("comment_id"))
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "comment_id parse failed"},
			})
			return
		}
		// 判断是否有删除权限
		isLegal := service.CheckComment(commentId, videoId, userId)
		if !isLegal {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "Permission denied"},
			})
			return
		}
		// 删除评论
		success := service.DeleteCommentById(commentId)
		if !success {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "Comment delete failed"},
			})
			return
		}
	}

	c.JSON(http.StatusOK, CommentActionResponse{
		Response: common.Response{StatusCode: 0},
		Comment:  comment,
	})

}

// CommentList 获取视频的评论列表
func CommentList(c *gin.Context) {
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "video_id parse failed"},
		})
		return
	}
	token := c.Query("token")
	// 验证token
	userId, success := service.ValidateToken(token)
	if !success {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "Token expired, please login again"},
		})
		return
	}

	commentList := service.GetCommentListById(videoId, userId)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: commentList,
	})
}
