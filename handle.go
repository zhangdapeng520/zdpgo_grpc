package zdpgo_grpc

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gin "zdpgo_gin"
)

// 将GRPC的错误转换为HTTP错误
func HandleErrorGrpcToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "要访问的资源不存在",
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务不可用",
				})
			case codes.AlreadyExists:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "数据已存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务器错误",
				})
			}
			return
		}
	}
}
