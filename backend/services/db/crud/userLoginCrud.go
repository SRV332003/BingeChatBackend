package crud

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/db"
	"errors"
)

func GetUserLoginById(id uint) (*models.UserLogin, error) {
	var userLogin models.UserLogin
	result := db.Instance.First(&userLogin, id)
	if result.RowsAffected > 1 {
		db.Instance.Logger.Warn(db.Instance.Statement.Context, "More than one user found with the same ID", "id", id)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &userLogin, nil
}

func GetUserLoginByEmail(email string) (*models.UserLogin, error) {
	var user models.UserLogin
	result := db.Instance.Where("email = ?", email).First(&user)
	if result.RowsAffected > 1 {
		db.Instance.Logger.Warn(db.Instance.Statement.Context, "More than one user found with the same email", "email", email)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByRefreshToken(refreshToken string) (*models.UserLogin, error) {
	var user models.UserLogin
	result := db.Instance.Where("refresh_token = ?", refreshToken).First(&user)
	if result.RowsAffected > 1 {
		db.Instance.Logger.Warn(db.Instance.Statement.Context, "More than one user found with the same refresh token", "refresh_token", refreshToken)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUserLogin(user *models.UserLogin) error {
	tx := db.Instance.Begin()
	result := tx.Model(&models.UserLogin{}).Where("id = ?", user.ID).Updates(user)

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

func CreateUserLogin(userLogin *models.UserLogin) error {
	result := db.Instance.Create(userLogin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ValidateLogin(email string, password string) (*models.UserLogin, error) {
	var userLogin models.UserLogin
	result := db.Instance.Where("email = ? AND password = ?", email, password).First(&userLogin)
	if result.RowsAffected > 1 {
		db.Instance.Logger.Warn(db.Instance.Statement.Context, "More than one user found with the same email", "email", email)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &userLogin, nil
}

func CheckUserLoginExists(email string) (bool, models.UserLogin, error) {
	var user models.UserLogin
	result := db.Instance.Where("email = ?", email).First(&user)
	if result.RowsAffected > 1 {
		db.Instance.Logger.Warn(db.Instance.Statement.Context, "More than one user found with the same email", "email", email)
	}
	if result.Error != nil {
		return false, user, result.Error
	}
	return result.RowsAffected > 0, user, nil
}

func CheckUserInfoExists(id uint) (bool, error) {
	var user models.UserLogin
	result := db.Instance.Where("id = ?", id).First(&user)
	if result.RowsAffected > 1 {
		db.Instance.Logger.Warn(db.Instance.Statement.Context, "More than one user found with the same ID", "id", id)
	}
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected < 1 {
		return false, nil
	}
	if user.UserInfoID == 0 {
		return false, nil
	}
	return true, nil
}
