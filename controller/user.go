package controller

import (
	"blog/dao/mysql"
	"blog/logic"
	"blog/models"
	"blog/pkg/tools"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
// @Summary 注册用户接口
// @Description 用户注册接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param SignUp body models.ParamSignUp true "注册参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /signUp [post]
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 其你去参数有误，直接返回相应
		zap.L().Error("SignUp with invalid  param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		}
		return
	}
	// 3。返回相应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录函数
// @Summary 登录接口
// @Description 登录接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Login body models.ParamLogin true "登录参数"
// @Success 200
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	// 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回相应
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	data := make(map[string]string)
	data["token"] = token
	// 返回响应
	ResponseSuccess(c, data)
}

func EmailLoginHandler(c *gin.Context) {
	p := new(models.EmailLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("EmailLogin with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	token, err := logic.EmailLogin(p)
	if err != nil {
		zap.L().Error("logic.EmailLogin failed", zap.String("email", p.Email), zap.Error(err))
		ResponseError(c, CodeIncorrect)
		return
	}
	data := make(map[string]string)
	data["token"] = token

	// 返回响应
	ResponseSuccess(c, data)
}

// GetUserInfoHandler 获取用户信息
func GetUserInfoHandler(c *gin.Context) {
	pidStr := c.Param("id")
	id, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get role with param", zap.Int64("id", id), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetUserInfo(id)
	if err != nil {
		zap.L().Error("logic.GetUserInfo(userID) failed", zap.Int64("userID", id), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func GetUserInfoListHandler(c *gin.Context) {
	data, err := logic.GetUserInfoList()
	if err != nil {
		zap.L().Error("logic.GetUserInfoList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// UploadImage 上传头像
func UploadImage(c *gin.Context) {
	pidStr := c.Param("id")
	id, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get role with param", zap.Int64("id", id), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		zap.L().Error("upload Image with invalid param", zap.Any("file", file), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	extName, ok := tools.SuffixName(file)
	if !ok {
		zap.L().Error("tools.SuffixName(file) failed", zap.Any("file", file), zap.Error(err))
		ResponseError(c, CodeFileSuffixNotLegal)
		return
	}
	dst, err := logic.UploadImage(file, extName, id)
	if err != nil {
		zap.L().Error("logic.UploadImage(file) failed", zap.Any("file", file), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeUpdateFailed)
		return
	}
	c.SaveUploadedFile(file, dst)
	ResponseSuccess(c, nil)
}

func GetEmailCode(c *gin.Context) {
	e := new(models.Email)
	if err := c.ShouldBindJSON(e); err != nil {
		zap.L().Error("get Code with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	if err := logic.SendCode(e.Email); err != nil {
		zap.L().Error("err", zap.Error(err))
		return
	}

	ResponseSuccess(c, nil)
}