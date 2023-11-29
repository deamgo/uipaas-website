package company

import (
	"strconv"

	"github.com/deamgo/uipaas-home/backend/context"
	"github.com/deamgo/uipaas-home/backend/dao"
	"github.com/deamgo/uipaas-home/backend/pkg/log"
	"github.com/deamgo/uipaas-home/backend/pkg/types"
	"github.com/deamgo/uipaas-home/backend/service/company"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

const (
	GetInfoListSuccess = 0
	GetInfoListFailed  = -1
)

type companyGetReq struct {
	PageSize string `json:"pageSize"`
	PageNum  string `json:"pageNum"`
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResp struct {
	Items []*company.Company `json:"items"`
	Total int64              `json:"total"`
}

func CompanyGet(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			req companyGetReq
		)

		req.PageSize = c.Query("pageSize")
		req.PageNum = c.Query("pageNum")
		size, _ := strconv.Atoi(req.PageSize)
		num, _ := strconv.Atoi(req.PageNum)

		list, total, err := ctx.CompanyService.CompanyGet(c, size, num)

		if err != nil {
			switch err {
			case dao.DBError:
				log.Errorw("failed to get companyList",
					zap.Error(err),
					zap.Any("companyList", list),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
					Code: GetInfoListFailed,
					Msg:  "failed to get companyList",
					Data: nil,
				}))
			}
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: GetInfoListSuccess,
			Msg:  "get companyList success",
			Data: PageResp{
				Items: list,
				Total: total,
			},
		}))
	}

}
