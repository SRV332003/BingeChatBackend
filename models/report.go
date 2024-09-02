package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	ReporterID uint   `json:"reporterId"`
	ReportedID uint   `json:"reportedId"`
	ReportText string `json:"ReportText"`
}
