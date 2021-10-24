package users

import "awesomeProject/shared"

func UserInitialization()  {
	err := shared.DBConnection.AutoMigrate(&User{})

	if err != nil {
		return
	}

	userRoutes()
}