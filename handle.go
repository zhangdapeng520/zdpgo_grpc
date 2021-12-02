package zgo_grpc

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	api "zgo_api"
)

// 将GRPC的错误转换为HTTP错误
func HandleErrorGrpcToHttp(err error, c *api.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, api.H{
					"msg": "要访问的资源不存在",
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, api.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, api.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, api.H{
					"msg": "服务不可用",
				})
			case codes.AlreadyExists:
				c.JSON(http.StatusInternalServerError, api.H{
					"msg": "数据已存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, api.H{
					"msg": "服务器错误",
				})
			}
			return
		}
	}
}
