package crud

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/db"
)

type College struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

func CreateCollege(college *models.College) error {
	result := db.Instance.Create(college)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCollegeById(id uint) (*models.College, error) {
	var college models.College
	result := db.Instance.First(&college, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &college, nil
}

func GetAllCollegesNames() ([]College, error) {
	var colleges []College
	// select only name and id
	result := db.Instance.Select("name", "id").Find(&colleges)
	if result.Error != nil {
		return nil, result.Error
	}

	return colleges, nil
}

func GetVerifiedColleges() ([]models.College, error) {
	var colleges []models.College
	result := db.Instance.Where("verified = ?", true).Find(&colleges)
	if result.Error != nil {
		return nil, result.Error
	}
	return colleges, nil
}

func UpdateCollegeStatus(college *models.College) error {
	tx := db.Instance.Begin()
	result := tx.Model(&models.College{}).Where("id = ?", college.ID).Updates(college)

	if result.RowsAffected > 1 {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CheckCollegeExists(emailFormat string) (bool, models.College, error) {
	var college models.College
	result := db.Instance.Where("email_format = ?", emailFormat).First(&college)
	if result.Error != nil {
		return false, college, result.Error
	}
	return true, college, nil
}
