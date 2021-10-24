package users

func indexUserService() []User {
	var users []User
	index(&users)
	return users
}

func createUserService(user userValidation) User {
	modelUser := User{FirstName: user.FirstName, LastName: user.LastName, Email:user.Email}
	store(&modelUser)
	return modelUser
}

func updateUserService(user userValidation, userId uint) User {
	modelUser := User{FirstName: user.FirstName, LastName: user.LastName, Email:user.Email}
	modelUser.ID = userId

	update(&modelUser)

	return modelUser
}

func showUserService(userId uint) User {
	var user User
	show(&user, userId)
	return user
}

func deleteUserService(userId uint) int64 {
	return destroy(userId)
}