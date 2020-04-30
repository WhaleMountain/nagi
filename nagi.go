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

	//environment := [][]string{
	//	{"-e", "MYSQL_ROOT_PASSWORD=mysql"},
	//	{"-e", "WORDPRESS_DB_HOST=nagi-db", "WORDPRESS_DB_USER=root", "WORDPRESS_DB_PASSWORD=mysql"},
	//}

	//NewCompose(cli, ctx, []string{"nagi-db", "nagi-word"},[]string{"mysql:5.7", "wordpress:latest"}, environment, []string{"3306", "80"}, []string{"3306", "8080"}, netResp)
}