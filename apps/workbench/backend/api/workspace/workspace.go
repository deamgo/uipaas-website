package workspace

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/types"
	"github.com/deamgo/workbench/service/workspace"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type WorkspaceCreateReq struct {
	Name        string `json:"name"`
	Lable       string `json:"lable"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type WorkspaceCreateResp struct {
	*workspace.Workspace
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

		data, err := ctx.WorkspaceService.WorkspaceCreate(c, convertWorkspace(req))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly." + err.Error(),
				Data: nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "create workspace succeed",
			Data: WorkspaceCreateResp{data},
		}))

	}
}

func convertWorkspace(req WorkspaceCreateReq) *workspace.Workspace {
	return &workspace.Workspace{
		Name:        req.Name,
		Lable:       req.Lable,
		Description: req.Description,
		Logo:        req.Logo,
	}
}
