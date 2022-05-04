package models

import "github.com/mahirB/user_managment_api/services"

type Permission struct {
	ID          int    `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Description string `json:"description" db:"description"`
}

type UserPermission struct {
	ID           int `json:"id" db:"id"`
	UserID       int `json:"user_id" db:"user_id"`
	PermissionID int `json:"perms_id" db:"perms_id"`
}

func PermissionsList() (perms []Permission, err error) {
	query := `SELECT 
				id,code,description
			  FROM permissions`

	db := services.Access.GetDB()

	err = db.Select(&perms, query)

	return perms, err
}

func UserPermissions(userID int) (userPerms []UserPermission, err error) {
	query := `SELECT 
				id,user_id,perms_id
			  FROM 
			  	user_permissions 
			  WHERE 
			  	user_id = ?`

	db := services.Access.GetDB()

	err = db.Select(&userPerms, query, userID)

	return userPerms, err
}

func CreateUserPermissions(userID int, userPerms []UserPermission) error {
	db := services.Access.GetDB()

	db.Exec("DELETE FROM user_permissions WHERE user_id=?", userID)

	query := `INSERT INTO user_permissions
					(user_id,perms_id)
				VALUES
					(:user_id,:perms_id)`

	for _, userPermission := range userPerms {
		_, err := db.NamedExec(query, userPermission)
		if err != nil {
			return err

		}
	}

	return nil
}
