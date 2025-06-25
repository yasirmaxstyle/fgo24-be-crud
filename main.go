package main

import (
	"dashboard-backend/routers"
	"dashboard-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	utils.InitDB()

	routers.CombineRouter(r)

	r.Run()
}
