package user

import (
	"fmt"
	"net/http"

	"github.com/deamgo/uipaas-home/backend/context"
	"github.com/deamgo/uipaas-home/backend/dao"
	"github.com/deamgo/uipaas-home/backend/pkg/log"
	"github.com/deamgo/uipaas-home/backend/pkg/types"
	"github.com/deamgo/uipaas-home/backend/service/user"

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
		fmt.Println("#######", req.ID)
		userService := user.User{UserID: req.ID}
		userInfo, err := ctx.UserService.UserGet(c, &userService)
		//fmt.Println("##########", userInfo.UserID)
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

const (
	LoginSuccess = 0
	LoginFailed  = -1
)

type UserPostReq struct {
	loginUser *user.User
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
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
				Code: LoginFailed,
				Msg:  "login format error",
				Data: nil,
			})
		}
		err = ctx.UserService.UserLogin(c, req.loginUser)

		if err != nil {
			switch err {
			case dao.DBError:
				log.Errorw("failed to get userinfo",
					zap.Error(err),
					zap.Any("userlogin", req),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, &Resp{
					Code: LoginFailed,
					Msg:  "login failed",
					Data: nil,
				})
			}
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: LoginSuccess,
			Msg:  "login failed",
			Data: nil,
		}))
	}
}
