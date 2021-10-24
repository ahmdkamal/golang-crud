package users

import (
	"awesomeProject/shared"
)

func store(user *User) {
	shared.DBConnection.Create(&user)
}

func index(users *[]User) {
	shared.DBConnection.Find(&users)
}

func show(user *User, userId uint) {
	shared.DBConnection.Find(&user, userId)
}

func update(user *User) {
	shared.DBConnection.Save(&user)
}

func destroy(userId uint) int64 {
	deleted := shared.DBConnection.Delete(&User{}, userId)
	return deleted.RowsAffected
}
