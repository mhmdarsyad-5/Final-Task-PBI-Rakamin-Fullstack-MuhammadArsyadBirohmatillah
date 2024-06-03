package database

import (
	"final-task-pbi-rakamin-fullstack-MuhammadArsyadBirohmatillah/models"
)

func Migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Photo{})
}
