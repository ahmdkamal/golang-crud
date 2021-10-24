package users

import (
	"awesomeProject/exceptions"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

var validate = validator.New()

type userValidation struct {
	FirstName string `validate:"required,max=250" json:"first_name"`
	LastName string `validate:"required,max=250" json:"last_name"`
	Email    string `validate:"required,email" json:"email"`
}

func createUserValidation(user userValidation) (interface{}, bool) {
	err := validate.Struct(user)

	if err != nil {
		var errors []exceptions.Errors
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, false
		}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, exceptions.Errors{
				Field: err.Field(),
				Message: err.ActualTag(),
			})
		}

		// from here you can create your own error messages in whatever language you wish
		return errors, true
	}

	return nil, false
}

func updateUserValidation(user userValidation) (interface{}, bool) {
	err := validate.Struct(user)
	fmt.Println(err)

	if err != nil {
		var errors []exceptions.Errors
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, false
		}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, exceptions.Errors{
				Field: err.Field(),
				Message: err.ActualTag(),
			})
		}

		// from here you can create your own error messages in whatever language you wish
		return errors, true
	}

	return nil, false
}

func checkUserIdValidation(c *gin.Context) (interface{}, bool) {
	var errors []exceptions.Errors
	userId, err := strconv.Atoi(c.Param("user"))

	var user User
	show(&user, uint(userId))

	if user.ID == 0 || err != nil {
		errors = append(errors, exceptions.Errors{
			Field: "id",
			Message: "Invalid id",
		})
		return errors, true
	}
	return nil, false
}