package main

import (
	"github.com/aryo24/task-5-pbi-btpns-aryowahyunugroho/database"
	"github.com/aryo24/task-5-pbi-btpns-aryowahyunugroho/models"
	"github.com/aryo24/task-5-pbi-btpns-aryowahyunugroho/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
