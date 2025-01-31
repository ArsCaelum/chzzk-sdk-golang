package chzzk_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserV1 struct {
	ChannelId   string `json:"channelId"`
	ChannelName string `json:"channelName"`
}

func UserV1Endpoint(c *gin.Context) {
	url := OPEN_API_URL + "/open/v1/users/me"
	err, str := sendGet(url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, str)
}
