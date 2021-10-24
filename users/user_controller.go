package users

import (
	"awesomeProject/shared"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func indexUsers(c *gin.Context) {
	shared.Success(c, 200, "Successfully", indexUserService())
}

func createUser(c *gin.Context) {
	var user userValidation

	data, _ := c.GetRawData()
	if err := json.Unmarshal(data, &user); err != nil {
		panic(err)
	}

	var validation interface{}
	var hasError bool
	validation, hasError = createUserValidation(user)

	if hasError {
		shared.ErrorResponse(c, 422, "Hello", validation)
		return
	}

	shared.Success(c, 200, "Hello", createUserService(user))
}

func showUser(c *gin.Context) {
	var validation interface{}
	var hasError bool
	validation, hasError = checkUserIdValidation(c)
	if hasError {
		shared.ErrorResponse(c, 422, "Error", validation)
		return
	}
	userId, _ := strconv.Atoi(c.Param("user"))

	shared.Success(c, 200, "Hello", showUserService(uint(userId)))
}

func updateUser(c *gin.Context) {
	var user userValidation

	data, _ := c.GetRawData()
	if err := json.Unmarshal(data, &user); err != nil {
		panic(err)
	}

	var validation interface{}
	var hasError bool
	validation, hasError = updateUserValidation(user)

	if hasError {
		shared.ErrorResponse(c, 422, "Hello", validation)
		return
	}
	validation, hasError = checkUserIdValidation(c)
	if hasError {
		shared.ErrorResponse(c, 422, "Hello", validation)
		return
	}

	userId, _ := strconv.Atoi(c.Param("user"))

	shared.Success(c, 200, "Hello", updateUserService(user, uint(userId)))
}

func deleteUser(c *gin.Context) {
	var validation interface{}
	var hasError bool
	validation, hasError = checkUserIdValidation(c)
	if hasError {
		shared.ErrorResponse(c, 422, "Error", validation)
		return
	}

	userId, _ := strconv.Atoi(c.Param("user"))
	deleted := deleteUserService(uint(userId))

	shared.Success(c, 204, fmt.Sprintf("%s %d %s %d", "Deleted Successfully and ", deleted, " rows affected ", userId), nil)
}
