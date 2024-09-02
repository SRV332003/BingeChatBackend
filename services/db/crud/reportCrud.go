package crud

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/db"
)




func CreateReport(college *models.Report) error {
	result := db.Instance.Create(college)
	if result.Error != nil {
		return result.Error
	}
	return nil
}