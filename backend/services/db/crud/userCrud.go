package crud

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/db"
	"errors"
)

func UpdateUserInfo(user *models.UserInfo) error {
	tx := db.Instance.Begin()
	result := tx.Model(&models.UserInfo{}).Where("id = ?", user.ID).Updates(user)

	if result.RowsAffected > 1 {
		tx.Rollback()
		return errors.New("more than one user is updating, invalid sql query")
	}
	tx.Commit()

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateUserInfo(userInfo *models.UserInfo) error {
	result := db.Instance.Create(userInfo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
