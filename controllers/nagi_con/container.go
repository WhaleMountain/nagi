package nagi_con

import (
	"github.com/gin-gonic/gin"

	//"nagi-docker/services/nagi_con"
	"nagi-docker/services/nagi_con"
)

type Ctl_container struct{}

func (cc Ctl_container) Index(c *gin.Context) {
	c.AbortWithStatus(404)
}

func (cc Ctl_container) CreateContainer(c *gin.Context) {
	sc := nagi_con.Service_container{}

	if err := sc.CreateNewContainer(c); err != nil {
		c.AbortWithStatus(500)
	} else {
		c.JSON(201, gin.H{
			"kind": "Container",
			"status": "Success",
		})
	}
}

func (cc Ctl_container) CreateCompose(c *gin.Context) {
	sc := nagi_con.Service_container{}

	if err := sc.CreateNewCompose(c); err != nil {
		c.AbortWithStatus(500)
	} else {
		c.JSON(201, gin.H{
			"kind": "Compose",
			"status": "Success",
		})
	}
}