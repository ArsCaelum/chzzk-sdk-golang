package main

import (
	chzzk_api "chzzk-api"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
	"log"
)

func main() {
	// DB Setup
	db, err := buntdb.Open("data.db")

	if nil != err {
		log.Fatal(err)
	}

	defer func(db *buntdb.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// Web Routing
	r := gin.Default()
	r.Use(gin.Logger())

	chzzk_api.Common(r, db)
	v1 := chzzk_api.CreateV1(r)
	v1.GET("/user", chzzk_api.UserV1Endpoint)

	err = r.Run(":8080")

	if nil != err {
		panic(err)
	}
}
