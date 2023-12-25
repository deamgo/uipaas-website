package devdepot

import (
	"net/http"
	"strconv"

	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DevdepotList(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		workspace_id := c.Param("workspace_id")
		if len(workspace_id) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": e.Failed,
				"msg":  "The parameters are not formatted correctly",
			})
		}
		pageNum, err := strconv.Atoi(c.Query("pageNum"))
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Parameter parsing failed"})
		}

		list, err := ctx.DevDepotService.DevItemList(c, workspace_id, pageNum)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": e.Success,
			"data": list,
		})
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
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Parameter parsing failed"})
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
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": e.Failed, "msg": err.Error()})
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
		}
		err := c.ShouldBind(&req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Parameter parsing failed"})
			return
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": e.Failed, "msg": err.Error()})
			return
		}
		err = ctx.DevDepotService.DevDepotDel(c, req.DeveloperID)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": e.Failed, "msg": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": e.Success, "msg": "Delete successfully"})
	}
}
