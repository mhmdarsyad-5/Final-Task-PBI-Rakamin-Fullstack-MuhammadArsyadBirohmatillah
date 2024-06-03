package main

import (
	"final-task-pbi-rakamin-fullstack-MuhammadArsyadBirohmatillah/database"
	"final-task-pbi-rakamin-fullstack-MuhammadArsyadBirohmatillah/helpers"
	"final-task-pbi-rakamin-fullstack-MuhammadArsyadBirohmatillah/router"

	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnvVariables()
	database.Connect()
	database.Migration()
}

func main() {
	r := gin.Default()
	router.Routers(r)
	r.Run()
}
