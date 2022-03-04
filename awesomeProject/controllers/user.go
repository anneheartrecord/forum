package controllers

import (
	"errors"
	"golangstudy/jike/awesomeProject/dao/mysql"
	"golangstudy/jike/awesomeProject/logic"
	"golangstudy/jike/awesomeProject/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		//	请求参数有误
		zap.L().Error("signup with invalid param", zap.Error(err))
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(c, CodeInvalidPassword, removeTopStruct(err.Translate(trans)))
		}
		return
	}

	// 业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		//	请求参数有误
		zap.L().Error("login with invalid param", zap.Error(err))
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(c, CodeInvalidPassword, removeTopStruct(err.Translate(trans)))
		}
		return
	}
	token, err := logic.Login(p)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, token)
}
