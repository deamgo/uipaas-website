package user

import (
	"net/http"

	"github.com/deamgo/uipass-waitlist-page/backend/context"
	"github.com/deamgo/uipass-waitlist-page/backend/dao"
	"github.com/deamgo/uipass-waitlist-page/backend/pkg/log"
	"github.com/deamgo/uipass-waitlist-page/backend/pkg/types"
	"github.com/deamgo/uipass-waitlist-page/backend/service/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserGetReq struct {
	ID string
}
type UserGetRep struct {
	user.User
}

func UserGet(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req UserGetReq
			err error
		)
		req.ID = c.Param("id")
		userService := user.User{UserID: req.ID}
		ctx.UserService.UserGet(c, &userService)
		if err != nil {
			switch err {
			case dao.DBError:
				log.Errorw("failed to create api source",
					zap.Error(err),
					zap.Any("user", req),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, types.NewErrorResponse(err.Error()))
			}
			c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(req))
			return

		}
	}
}
