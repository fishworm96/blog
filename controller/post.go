package controller

import (
	"blog/dao/mysql"
	"blog/logic"
	"blog/models"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
// @Summary 创建帖子接口
// @Description 创建帖子接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param PostParam body models.Post false "社区ID"
// @Security ApiKeyAuth
// @Success 200
// @Router /posts [post]
func CreatePostHandler(c *gin.Context) {
	// 1.获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) err", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 从c取到当前发送请求的用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3.响应返回
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 根据帖子id获取帖子信息
// @Summary 帖子信息接口
// @Description 根据帖子id获取帖子信息的接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param id path int true "查询帖子id"
// @Security ApiKeyAuth
// @Success 200
// @Router /post/{id} [get]
func GetPostDetailHandler(c *gin.Context) {
	// 获取参数 (从url中获取帖子的id)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Int64("pid", pid), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id获取数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorPostNotExist) {
			ResponseError(c, CodePostNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表的函数
// @Summary 获取帖子列表接口
// @Description 获取帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200
// @Router /post [get]
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList(page, size) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// UpdatePostHandler 更新帖子
// @Summary 更新帖子接口
// @Description 根据文章id来接收标题和内容修改帖子接口
// @Tags 帖子相关接口
// Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param id path int true "文章id"
// @Param ParamPost body models.ParamPost false "修改帖子内容"
// @Security ApiKeyAuth
// @Success 200
// @Router /post/edit/{id} [put]
func UpdatePostHandler(c *gin.Context) {
	// 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("update post with param", zap.Int64("pid", pid), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamPost)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) err", zap.Any("err", err))
		zap.L().Error("update post with invalid parma")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 修改帖子信息
	if err := logic.UpdatePost(pid, p); err != nil {
		zap.L().Error("logic.UpdatePost(p) failed", zap.Any("post", p), zap.Error(err))
		if errors.Is(err, mysql.ErrorPostNotExist) {
			ResponseError(c, CodePostNotExist)
			return
		}
		ResponseError(c, CodeUpdateFailed)
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
}

// DeletePostHandler 删除帖子接口
// @Summary 删除帖子接口
// @Description 根据帖子id删除帖子的接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param id path int true "删除文章帖子参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /post/delete/{id} [delete]
func DeletePostHandler(c *gin.Context) {
	// 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("delete post with invalid param", zap.Int64("pid", pid), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id删除文章
	if err := logic.DeletePostById(pid); err != nil {
		zap.L().Error("logic.DeletePostById(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		if errors.Is(err, mysql.ErrorPostNotExist) {
			ResponseError(c, CodePostNotExist)
			return
		}
		ResponseError(c, CodeDeleteFailed)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// 定义默认参数，如果没有传参数就用默认参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p)
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostListNew(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, data)
}
