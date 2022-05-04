package controllers

import (
	"fmt"

	"github.com/mahirB/user_managment_api/models"
)

func Deploy() {
	models.DropTables()
	models.UserDeploy()
	models.PermissionsDeploy()
	models.AddMockData()

	fmt.Println("Deployed succesfully")
}
