package devdepot

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/deamgo/workbench/api/ws"
	"github.com/deamgo/workbench/auth/jwt"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/dao/devdepot"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/deamgo/workbench/pkg/types"
	"github.com/deamgo/workbench/service/developer"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func DevdepotList(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		workspace_id := c.Param("workspace_id")
		if len(workspace_id) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&types.Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
			}))
		}
		pageNum, err := strconv.Atoi(c.Query("pageNum"))
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: "Parameter parsing failed"}))
		}

		list, err := ctx.DevDepotService.DevItemList(c, workspace_id, pageNum)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{
			Code: e.Success,
			Data: list,
		}))
	}
}

func DevdepotSearch(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Q       string `validate:"required"`
			PageNum int    `validate:"required"`
		}
		var err error
		req.Q = c.Query("q")
		req.PageNum, err = strconv.Atoi(c.Query("pageNum"))
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{
				Msg: "Parameter parsing failed",
			}))
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": e.Failed,
				"msg":  err.Error(),
			})
			return
		}
		workspace_id := c.Param("workspace_id")
		if len(workspace_id) == 0 {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": e.Failed,
				"msg":  "The parameters are not formatted correctly",
			})
		}
		devInfoSearch, err := ctx.DevDepotService.DevInfoSearch(c, workspace_id, req.Q, req.PageNum)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": e.Success,
			"data": devInfoSearch,
		})
	}
}

func DevdepotDel(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			DeveloperID string `json:"developer_id" validate:"required"`
			WorkspaceID string `json:"workspace_id" validate:"required"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Parameter parsing failed"})
			return
		}
		req.WorkspaceID = c.Param("workspace_id")
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": e.Failed, "msg": "The parameters are not formatted correctly"})
			return
		}
		err = ctx.DevDepotService.DevDepotDel(c, req.WorkspaceID, req.DeveloperID)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		// Send a message to the deleted developer
		msg := []byte("3000")
		err = ws.SendMsgToDeveloper(req.DeveloperID, msg)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": e.Success, "msg": "Delete successfully"})
	}
}

func DevdepotRoleModify(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			DeveloperID string `json:"developer_id" validate:"required"`
			WorkspaceID string `json:"workspace_id" validate:"required"`
			Role        string `json:"role" validate:"required"`
		}
		err := c.ShouldBind(&req)
		req.WorkspaceID = c.Param("workspace_id")
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Parameter parsing failed"})
			return
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": e.Failed, "msg": "The parameters are not formatted correctly"})
			return
		}
		devDepotItem := &devdepot.DevDepotItem{
			DeveloperId: req.DeveloperID,
			WorkspaceId: req.WorkspaceID,
			Role:        req.Role,
		}
		err = ctx.DevDepotService.DevDepotRoleModify(c, devDepotItem)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": e.Success, "msg": "Modify successfully"})
	}
}

func DevDepotInvite(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email       string `json:"email" validate:"required,email"`
			WorkspaceID string `json:"workspace_id" validate:"required"`
			DeveloperID string `json:"developer_id"`
			Role        string `json:"role" validate:"required"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Parameter parsing failed"})
			return
		}
		req.WorkspaceID = c.Param("workspace_id")
		token := c.GetHeader("Authorization")
		req.DeveloperID, err = jwt.ExtractIDFromToken(token)

		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Parameter parsing failed"})
			return
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": e.Failed, "msg": "The parameters are not formatted correctly"})
			return
		}
		developer, err := ctx.UserService.DeveloperGetByEmail(c, &developer.Developer{Email: req.Email})
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				logger.Error(err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
				return
			}
		}
		if developer != nil {
			req.DeveloperID = developer.ID
		}
		devDepotItem := &devdepot.DevDepotItem{
			Email:       req.Email,
			WorkspaceId: req.WorkspaceID,
			DeveloperId: req.DeveloperID,
			Role:        req.Role,
		}
		err = ctx.DevDepotService.DevDepotInvite(c, devDepotItem)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": e.Success, "msg": "Invitation email sent."})
	}
}
