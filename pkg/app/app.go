package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linxbin/cron-service/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (r *Response) ToResponse(data interface{}) {
	response := ResponseSuccess{
		Code:    errcode.Success.Code,
		Message: errcode.Success.Message,
		Result:  data,
	}
	r.Ctx.JSON(http.StatusOK, response)
}

type ResponseList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  gin.H  `json:"result"`
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	response := ResponseList{
		Code:    errcode.Success.Code,
		Message: errcode.Success.Message,
		Result: gin.H{
			"items": list,
			"pager": Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	}
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code, "message": err.Message}
	details := err.Details
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
