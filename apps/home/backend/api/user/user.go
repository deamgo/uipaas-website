package user

import (
	"github.com/deamgo/uipaas-home/backend/context"
	"github.com/deamgo/uipaas-home/backend/dao"
	"github.com/deamgo/uipaas-home/backend/pkg/e"
	"github.com/deamgo/uipaas-home/backend/pkg/log"
	"github.com/deamgo/uipaas-home/backend/pkg/types"
	"github.com/deamgo/uipaas-home/backend/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserGetReq struct {
	ID string `json:"id"`
}
type UserGetResp struct {
	*user.User
}

func UserGet(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req UserGetReq
			err error
		)
		req.ID = c.Param("id")
		userService := user.User{UserID: req.ID}
		userInfo, err := ctx.UserService.UserGet(c, &userService)
		if err != nil {
			switch err {
			case dao.DBError:
				log.Errorw("failed to get user",
					zap.Error(err),
					zap.Any("user", req),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, types.NewErrorResponse(err.Error()))
			}
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(UserGetResp{userInfo}))
	}
}

type UserPostReq struct {
	loginUser *user.User
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

func UserLogin(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			req UserPostReq
		)
		err = c.ShouldBind(&req.loginUser)
		if err != nil {
			log.Errorw("login format error",
				zap.Error(err),
				zap.Any("userlogin", req),
			)

			c.AbortWithStatusJSON(http.StatusBadRequest, &Resp{
				Code: e.Failed,
				Msg:  e.LoginFormatError,
				Data: nil,
			})
		}
		token, err := ctx.UserService.UserLogin(c, req.loginUser)

		if err != nil {
			switch err {
			case dao.DBError:
				log.Errorw("failed to get userinfo",
					zap.Error(err),
					zap.Any("userlogin", req),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
					Code: e.Failed,
					Msg:  e.LoginFailed,
					Data: nil,
				}))
			}
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  e.LoginSuccess,
			Data: UserLoginResp{Token: token},
		}))
	}
}
