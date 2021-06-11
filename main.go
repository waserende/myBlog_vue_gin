package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"waserd.ren/myBlog/common"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	fmt.Println("hello word!")

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
