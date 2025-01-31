package chzzk_api

import "github.com/gin-gonic/gin"

const V1 = "v1"

func CreateV1(r *gin.Engine) *gin.RouterGroup {
	return r.Group(V1)
}
