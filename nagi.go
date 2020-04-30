package main

import (
	"github.com/gin-gonic/gin"

	"nagi-docker/dk"
	"nagi-docker/controllers/nagi_con"
)

func router() *gin.Engine {
	r := gin.Default()

	cc := nagi_con.Ctl_container{}
	r.GET("", cc.Index)
	r.POST("/container", cc.CreateContainer)
	r.POST("/compose", cc.CreateCompose)

	return r
}


func main() {
	if err := dk.Init(); err != nil {
		panic(err)
	}
	
	r := router()
	r.Run()
}