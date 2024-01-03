package application

import (
	"net/http"
	"strconv"

	"github.com/deamgo/workbench/auth/jwt"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/deamgo/workbench/pkg/types"
	"github.com/deamgo/workbench/service/application"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ApplicationCreate(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name        string `json:"name" validate:"required"`
			WorkspaceID string `json:"workspace_id" validate:"required"`
		}
		req.WorkspaceID = c.Param("workspace_id")
		err := c.ShouldBindJSON(&req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: "Parameter parsing failed"}))
			return
		}
		token := c.GetHeader("Authorization")
		developerID, err := jwt.ExtractIDFromToken(token)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: "Parameter parsing failed"}))
			return
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		dID, _ := strconv.Atoi(developerID)
		app := &application.Application{
			Name:        req.Name,
			WorkspaceID: req.WorkspaceID,
			CreatedBy:   uint64(dID),
		}
		err = ctx.ApplicationService.ApplicationAdd(c, app)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{
			Code: e.Success,
			Msg:  "Success"}))
	}
}

func ApplicationList(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			WorkspaceID string `json:"workspace_id" validate:"required"`
		}
		req.WorkspaceID = c.Param("workspace_id")
		validate := validator.New()
		err := validate.Struct(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		app := &application.Application{
			WorkspaceID: req.WorkspaceID,
		}
		list, err := ctx.ApplicationService.ApplicationList(c, app)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{
			Code: e.Success,
			Msg:  "Success",
			Data: list,
		}))
	}
}

func ApplicationSearch(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			WorkspaceID string `json:"workspace_id" validate:"required"`
			Q           string `json:"q" validate:"required"`
		}
		req.WorkspaceID = c.Param("workspace_id")
		req.Q = c.Query("q")
		validate := validator.New()
		err := validate.Struct(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		app := &application.Application{
			WorkspaceID: req.WorkspaceID,
			Name:        req.Q,
		}
		list, err := ctx.ApplicationService.ApplicationSearchByName(c, app)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{
			Code: e.Success,
			Msg:  "Success",
			Data: list,
		}))
	}
}

func ApplicationDuplicate(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name        string `json:"name" validate:"required"`
			WorkspaceID string `json:"workspace_id" validate:"required"`
			Description string `json:"description" validate:"required"`
		}

		req.WorkspaceID = c.Param("workspace_id")
		err := c.ShouldBindJSON(&req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{
				Code: e.Failed,
				Msg:  "Parameter parsing failed",
			}))
			return
		}
		token := c.GetHeader("Authorization")
		developerID, err := jwt.ExtractIDFromToken(token)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(types.NewValidResponse(&types.Resp{
				Code: e.Failed,
				Msg:  "Parameter parsing failed",
			})))
			return
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		dID, _ := strconv.Atoi(developerID)
		app := &application.Application{
			Name:        req.Name,
			WorkspaceID: req.WorkspaceID,
			Description: req.Description,
			CreatedBy:   uint64(dID),
		}
		err = ctx.ApplicationService.ApplicationDuplicate(c, app)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{Code: e.Failed, Msg: err.Error()}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{Code: e.Success, Msg: "Success"}))
	}
}
