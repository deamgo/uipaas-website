package router

import (
	"github.com/deamgo/workbench/api/workspace"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/deamgo/workbench/api/user"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/middleware"
)

func NewRouter(ctx context.ApplicationContext) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(middleware.Cors())
	e.MaxMultipartMemory = 8 << 20 // 8 MiB
	mountAPIs(e, ctx)
	return e
}
func mountAPIs(e *gin.Engine, ctx context.ApplicationContext) {
	api := e.Group("v1")
	{
		api.POST("/signup", user.SignUp(ctx))
		api.POST("/signup_verify", user.SignUpVerify(ctx))
		api.POST("/signin", user.SignIn(ctx))
		api.POST("/workspace/create", workspace.WorkspaceCreate(ctx))
	}

	//workspaceApi := api.Group("workspace")
	//{
	//	workspaceApi.POST("/create", func(c *gin.Context) {
	//		logoFile, err := c.FormFile("logo")
	//		if err != nil {
	//			c.JSON(http.StatusBadRequest, gin.H{
	//				"error": err.Error(),
	//			})
	//			return
	//		}
	//		logoFileName := logoFile.Filename
	//		if strings.HasSuffix(logoFileName, ".jpg'") ||
	//			strings.HasSuffix(logoFileName, ".png'") ||
	//			strings.HasSuffix(logoFileName, ".jpeg'") {
	//			c.JSON(http.StatusOK, gin.H{
	//				"message": "请使用正确图像文件",
	//			})
	//		}
	//
	//		name := c.PostForm("name")
	//
	//		err = c.SaveUploadedFile(logoFile, "/Users/wenxing/uipaas-website/apps/workbench/backend/public/"+logoFile.Filename)
	//		if err != nil {
	//			c.JSON(http.StatusBadRequest, gin.H{
	//				"error": err.Error(),
	//			})
	//			return
	//		}
	//
	//		w := workspace.Workspace{
	//			Name:        name,
	//			Logo:        "/public/" + logoFile.Filename,
	//			Id:          "ZHANWQ",
	//			Lable:       "短内容测试",
	//			Description: "长内容测试",
	//		}
	//
	//		fmt.Print(db.DB)
	//		result := db.DB.Table("workspace").Create(&w)
	//
	//		c.JSON(http.StatusOK, gin.H{
	//			"message": fmt.Sprint(result),
	//		})
	//	})
	//}

}
