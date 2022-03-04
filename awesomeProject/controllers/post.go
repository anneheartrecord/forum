package controllers

import (
	"golangstudy/jike/awesomeProject/logic"
	"golangstudy/jike/awesomeProject/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c钟取到当前发帖子的用户id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic create post failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//根据id去除帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic GetpostDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回相应
	ResponseSuccess(c, data)
}
func GetPostListHandler(c *gin.Context) {
	//获取分页数
	pageNumStr := c.Query("page")
	pageSizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic get post list failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
