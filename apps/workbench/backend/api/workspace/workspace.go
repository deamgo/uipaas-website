package workspace

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/deamgo/workbench/auth/jwt"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/types"
	"github.com/deamgo/workbench/service/workspace"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type WorkspaceCreateReq struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type WorkspaceCreateResp struct {
	*workspace.Workspace
}

func WorkspaceGetLogoPath(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		header, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &types.Resp{
				Code: e.Failed,
				Msg:  "upload workspace logo error",
				Data: nil,
			})
			return
		}

		path, err := ctx.WorkspaceService.WorkspaceGetFilePath(header)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &types.Resp{
				Code: e.Failed,
				Msg:  "upload workspace logo file unusual.",
				Data: nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{
			Code: e.Success,
			Msg:  "upload workspace logo succeed",
			Data: path,
		}))

	}
}

func WorkspaceGetListById(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		id, err := getDevelopId(c)
		if id == 0 || err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &types.Resp{
				Code: e.Failed,
				Msg:  "Get developId data exception",
				Data: nil,
			})
			return
		}
		list, err := ctx.WorkspaceService.WorkspaceGetListById(c, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &types.Resp{
				Code: e.Failed,
				Msg:  "Get developId data exception",
				Data: nil,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{
			Code: e.Success,
			Msg:  "get workspace succeed",
			Data: list,
		}))
	}
}

func WorkspaceCreate(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req WorkspaceCreateReq
			err error
		)
		err = c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := ctx.WorkspaceService.WorkspaceCreate(c, convertWorkspace(req, c))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &types.Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly." + err.Error(),
				Data: nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{
			Code: e.Success,
			Msg:  "create workspace succeed",
			Data: WorkspaceCreateResp{data},
		}))

	}
}

func WorkspaceDel(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Id string `json:"id" validate:"required"`
		}
		req.Id = c.Param("workspace_id")

		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&types.Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
				Data: nil,
			}))
			return
		}
		var workspace = &workspace.Workspace{
			Id: req.Id,
		}
		developerID, err := jwt.ExtractIDFromToken(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		err = ctx.WorkspaceService.WorkspaceDel(c, workspace, developerID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&types.Resp{
				Code: e.Failed,
				Msg:  err.Error(),
			}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&types.Resp{
			Code: e.Success,
			Msg:  "delete workspace succeed",
		}))
	}
}

func convertWorkspace(req WorkspaceCreateReq, c *gin.Context) *workspace.Workspace {
	id, err := getDevelopId(c)
	if id == 0 || err != nil {
		log.Println("Get developId data exception")
	}
	return &workspace.Workspace{
		Name:        req.Name,
		Label:       req.Label,
		Description: req.Description,
		Logo:        req.Logo,
		CreatedBy:   id,
		UpdatedBy:   id,
	}
}

// getDevelopId Get the developer id in the context
func getDevelopId(c *gin.Context) (uint64, error) {
	developIdStr, err := jwt.ExtractIDFromToken(c.GetHeader("Authorization"))
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(developIdStr, 10, 64)
	if err != nil {
		log.Println("developId data exception")
		return 0, err
	}
	return uint64(i), nil
}
