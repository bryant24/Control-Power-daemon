package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os/exec"
)

var VERSION = "1.0.0.703"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"errcode": 0})
	})
	r.POST("/shutdown", func(c *gin.Context) {
		ShutDown()
		c.JSON(200, gin.H{"errcode": 0})
	})
	r.Run(":1234")

	log.Println("成功监听端口:1234")
	log.Println("节点http启动成功,版本号:", VERSION)

}

func ShutDown() {
	log.Println("触发关机")
	_ = exec.Command("shutdown", `-s`, `-t`, `0`).Run()
}
