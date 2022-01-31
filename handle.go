package zdpgo_grpc

import (
	"github.com/zhangdapeng520/zdpgo_code"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangdapeng520/zdpgo_gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleErrorGrpcToHttp 将GRPC的错误转换为HTTP错误
// @param err grpc的错误消息
// @param c gin上下文对象
func (g *Grpc) HandleErrorGrpcToHttp(err error, c *gin.Context) {
	rsp := zdpgo_gin.NewResponse()
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				rsp.Code = zdpgo_code.CODE_NOT_FOUND
				rsp.Message = zdpgo_code.MESSAGE_NOT_FOUND
			case codes.Internal:
				rsp.Code = zdpgo_code.CODE_SERVER_ERROR
				rsp.Message = zdpgo_code.MESSAGE_SERVER_ERROR
			case codes.InvalidArgument:
				rsp.Code = zdpgo_code.CODE_PARAM_ERROR
				rsp.Message = zdpgo_code.MESSAGE_PARAM_ERROR
			case codes.Unavailable:
				rsp.Code = zdpgo_code.CODE_GRPC_NOT_USE
				rsp.Message = zdpgo_code.MESSAGE_GRPC_NOT_USE
			case codes.AlreadyExists:
				rsp.Code = zdpgo_code.CODE_EXISTS_ERROR
				rsp.Message = zdpgo_code.MESSAGE_EXISTS_ERROR
			default:
				rsp.Code = zdpgo_code.CODE_SERVER_ERROR
				rsp.Message = zdpgo_code.MESSAGE_SERVER_ERROR
			}
			c.JSON(http.StatusOK, rsp)
			return
		}
	}
	c.JSON(http.StatusOK, rsp)
}
