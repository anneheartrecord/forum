package controllers

import (
	"golangstudy/jike/awesomeProject/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	//查询到所有的社区  id name 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
func CommunityDetailHandler(c *gin.Context) {
	//获取社区id
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
