package users

import "awesomeProject/shared"

func userRoutes()  {
	shared.Router.GET("/users", indexUsers)
	shared.Router.POST("/users", createUser)
	shared.Router.GET("/users/:user", showUser)
	shared.Router.PUT("/users/:user", updateUser)
	shared.Router.PATCH("/users/:user", updateUser)
	shared.Router.DELETE("/users/:user", deleteUser)
}