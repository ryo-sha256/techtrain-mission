package Models

import (
	"fmt"

	cnfg "techtrain-mission/config"
)

// NOTE: This file contains all of the functions that
//       interact directly with database.
//       There should be no need for additional service code

// Create User: writes user data to database
func CreateUser(u *User) (err error) {
	if err = cnfg.DB_CONN.Create(u).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Get User By Token: finds user data by a given token
// @param: userModel, token
func GetUserByToken(u *User, t string) (err error) {
	if err = cnfg.DB_CONN.Where("token = ?", t).First(u).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Update User: updates user's name
// @param: userModel, name, token
func UpdateUser(u *User, n string, t string) (err error) {
	if err = cnfg.DB_CONN.Where("token = ?", t).Update("name", t).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Delete User: deletes user by token
// @param: userModel, token
// NOTE: not required to implement delete operation for now
func DeleteUser(u *User, t string) (err error) {
	return nil
}
