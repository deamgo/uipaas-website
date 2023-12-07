package company

import (
	"github.com/go-playground/validator/v10"
	"strconv"

	"github.com/deamgo/uipaas-home/backend/context"
	"github.com/deamgo/uipaas-home/backend/dao"
	"github.com/deamgo/uipaas-home/backend/pkg/e"
	"github.com/deamgo/uipaas-home/backend/pkg/log"
	"github.com/deamgo/uipaas-home/backend/pkg/types"
	"github.com/deamgo/uipaas-home/backend/service/company"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
					Code: e.Failed,
					Msg:  e.GetListFailed,
					Data: nil,
				}))
			}
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  e.GetListSuccess,
			Data: PageResp{
				Items: list,
				Total: total,
			},
		}))
	}

}

type CompanyPostReq struct {
	*company.Company
}

func CompanyAdd(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req CompanyPostReq
			err error
		)

		err = c.ShouldBind(&req.Company)
		if err != nil {
			log.Errorw("add info format error",
				zap.Error(err),
				zap.Any("companyinfo", req),
			)
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		validate := validator.New()
		err = validate.Struct(req.Company)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}

		err = ctx.CompanyService.CompanyAdd(c, req.Company)
		if err != nil {
			switch err {
			case dao.DBError:
				log.Errorw("failed to add info")
				c.AbortWithStatus(http.StatusInternalServerError)
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
					Code: e.Failed,
					Msg:  err.Error(),
					Data: nil,
				}))
			}
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  e.AddMsgSuccess,
			Data: nil,
		}))
	}
}
