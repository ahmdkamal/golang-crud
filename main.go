package main

import (
	"awesomeProject/shared"
	"awesomeProject/users"
)

func main() {

	shared.Connect()
	users.UserInitialization()

	err := shared.Router.Run("localhost:3030")

	if err != nil {
		panic(err)
	}
}
