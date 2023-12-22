package workspace

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/types"
	"github.com/deamgo/workbench/service/workspace"

	"github.com/gin-gonic/gin"
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

		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, Resp{
				Code: e.Failed,
				Msg:  "upload workspace logo error",
				Data: nil,
			})
			return
		}

		err = c.SaveUploadedFile(file, "./public/"+file.Filename)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, Resp{
				Code: e.Failed,
				Msg:  "upload workspace logo error",
				Data: nil,
			})
			return
		}

		f, err := os.Open("./public/" + file.Filename)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, Resp{
				Code: e.Failed,
				Msg:  "upload workspace logo error",
				Data: nil,
			})
			return
		}
		defer func() {
			err = os.Remove("./public/" + file.Filename)
			if err != nil {
				log.Println(err)
			}
		}()

		path, err := ctx.WorkspaceService.WorkspaceGetFilePath(f)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, Resp{
				Code: e.Failed,
				Msg:  "upload workspace logo file unusual.",
				Data: nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
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
			c.AbortWithStatusJSON(http.StatusBadRequest, Resp{
				Code: e.Failed,
				Msg:  "Get developId data exception",
				Data: nil,
			})
			return
		}
		list, err := ctx.WorkspaceService.WorkspaceGetListById(c, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, Resp{
				Code: e.Failed,
				Msg:  "Get developId data exception",
				Data: nil,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
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
		UpdateBy:    id,
	}
}

func getDevelopId(c *gin.Context) (uint64, error) {
	developIdStr := c.Value("username").(string)
	i, err := strconv.ParseInt(developIdStr, 10, 64)
	if err != nil {
		log.Println("developId data exception")
		return 0, err
	}
	return uint64(i), nil
}
