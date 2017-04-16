package main

import (
	"github.com/go-pg/pg"
	"gopkg.in/gin-gonic/gin.v1"
  "github.com/go-pg/pg/orm"
  "github.com/tinchi/gin-react/models"
)

var dbConn *pg.DB
var router *gin.Engine

func main() {
	router = gin.Default()

	dbConn = pg.Connect(&pg.Options{
		User:     "saint",
		Database: "deposit_manager",
	})

	err := createSchema(dbConn)
	if err != nil {
		panic(err)
	}

	initializeRoutes()

	router.Run(":3001")
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&models.User{}, &models.Deposit{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}
	return nil
}
