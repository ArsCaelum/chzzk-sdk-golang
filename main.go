package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
	"log"
)

const OPEN_API_URL = "https://openapi.chzzk.naver.com"

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

	r.GET("/api/auth", func(c *gin.Context) {
		code := c.Request.URL.Query().Get("code")
		state := c.Request.URL.Query().Get("state")

		err := db.Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set("code", code, nil)
			if nil != err {
				return err
			}

			_, _, err = tx.Set("state", state, nil)
			if nil != err {
				return err
			}

			return nil
		})

		if nil != err {
			log.Fatal(err)
		}
	})

	r.GET("/api/auth/view", func(c *gin.Context) {
		var data = make(map[string]interface{})

		err := db.View(func(tx *buntdb.Tx) error {
			val, err := tx.Get("code")

			if nil != err {
				return err
			}
			data["code"] = val

			val, err = tx.Get("state")
			if nil != err {
				return err
			}

			data["state"] = val

			return nil
		})

		if nil != err {
			log.Print(err.Error())
			return
		}

		c.JSON(200, data)
	})

	err = r.Run(":8080")

	if nil != err {
		panic(err)
	}
}
