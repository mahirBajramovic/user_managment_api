package models

import (
	"strconv"

	"github.com/mahirB/user_managment_api/services"
)

var UserTable = `CREATE TABLE 
				user
				( 
					id INT(11) NOT NULL AUTO_INCREMENT , 
					first_name VARCHAR(50) NOT NULL , 
					last_name VARCHAR(50) NOT NULL , 
					username VARCHAR(25) NOT NULL , 
					password CHAR(25) NOT NULL , 
					email VARCHAR(100) NOT NULL , 
					status INT NOT NULL DEFAULT '0' , 
					PRIMARY KEY (id),
					UNIQUE KEY (email)
				) ENGINE = InnoDB;`

var PermsTable = `CREATE TABLE 
				permissions
				( 
					id INT(11) NOT NULL AUTO_INCREMENT , 
					code VARCHAR(50) NOT NULL , 
					description VARCHAR(100) NOT NULL , 
					PRIMARY KEY (id)
				) ENGINE = InnoDB;`

var UserPermsTable = `CREATE TABLE 
				user_permissions
				( 
					id INT(11) NOT NULL AUTO_INCREMENT , 
					user_id INT NOT NULL, 
					perms_id INT NOT NULL,
					PRIMARY KEY (id),
					FOREIGN KEY (user_id) REFERENCES user(id),
					FOREIGN KEY (perms_id) REFERENCES permissions(id)
				) ENGINE = InnoDB;`

func DropTables() {
	db := services.Access.GetDB()
	db.Exec("DROP TABLE user_permissions")
	db.Exec("DROP TABLE permissions")
	db.Exec("DROP TABLE user")

}

func UserDeploy() {
	db := services.Access.GetDB()

	_, err := db.Exec(UserTable)

	if err != nil {
		panic("Could not create User table " + err.Error())
	}
}

func PermissionsDeploy() {
	db := services.Access.GetDB()

	_, err := db.Exec(PermsTable)

	if err != nil {
		panic("Could not create Perms table " + err.Error())
	}

	_, err = db.Exec(UserPermsTable)

	if err != nil {
		panic("Could not create UserPerms table " + err.Error())
	}
}

func AddMockData() {
	db := services.Access.GetDB()

	db.Exec(`insert into permissions (code,description) values ('Add','You can add user with this permission')`)
	db.Exec(`insert into permissions (code,description) values ('Edit','You can edit user with this permission')`)
	db.Exec(`insert into permissions (code,description) values ('Delete','You can delete user with this permission')`)
	db.Exec(`insert into permissions (code,description) values ('View','You can view user with this permission')`)

	for i := 1; i < 25; i++ {
		db.Exec(`INSERT INTO user (first_name,last_name,username,password,email,status) VALUES ('Mahir','Bajramovic','Mahamaha','nekiPassword','mahir` + strconv.Itoa(i) + `@gmail.com',1)`)
	}
}
