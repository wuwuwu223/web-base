package dao

import "web-base/model"

//CreateUser
func CreateUser(user *model.User) error {
	return db.Create(user).Error
}

//GetUser
func GetUser(username string) (*model.User, error) {
	var user model.User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

//GetUserByID
func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	return &user, err
}

//UpdateUser
func UpdateUser(user *model.User) error {
	return db.Save(user).Error
}

//DeleteUser
func DeleteUser(user *model.User) error {
	return db.Delete(user).Error
}
