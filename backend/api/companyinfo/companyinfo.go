package companyinfo

import (
	"github.com/deamgo/uipaas-home/backend/context"
	"github.com/deamgo/uipaas-home/backend/dao"
	"github.com/deamgo/uipaas-home/backend/pkg/log"
	"github.com/deamgo/uipaas-home/backend/pkg/types"
	"github.com/deamgo/uipaas-home/backend/service/companyinfo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

const (
	GetInfoListSuccess = 0
	GetInfoListFailed  = -1
)

type CompanyInfoGetReq struct {
	PageSize string `json:"pageSize"`
	PageNum  string `json:"pageNum"`
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResp struct {
	Items []*companyinfo.CompanyInfo `json:"items"`
	Total int64                      `json:"total"`
}

func CompanyInfoGet(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			req CompanyInfoGetReq
		)

		req.PageSize = c.Query("pageSize")
		req.PageNum = c.Query("pageNum")
		size, _ := strconv.Atoi(req.PageSize)
		num, _ := strconv.Atoi(req.PageNum)

		list, total, err := ctx.CompanyInfoService.CompanyInfoGet(c, size, num)

		if err != nil {
			switch err {
			case dao.DBError:
				log.Errorw("failed to get companyInfoList",
					zap.Error(err),
					zap.Any("companyInfoList", list),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, &Resp{
					Code: GetInfoListFailed,
					Msg:  "failed to get companyInfoList",
					Data: nil,
				})
			}
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: GetInfoListSuccess,
			Msg:  "get companyInfoList success",
			Data: PageResp{
				Items: list,
				Total: total,
			},
		}))
	}

}
