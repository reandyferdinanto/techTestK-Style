package entity

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	ID       int8            `gorm:"primaryKey;autoIncrement" json:"id"`
    Username        string          `gorm:"type:varchar(50)" json:"username"`
    Gender          string          `gorm:"type:enum('Male', 'Female')" json:"gender"`
    SkinType        string          `gorm:"type:enum('Oily', 'Dry', 'Combination', 'Sensitive', 'Normal')" json:"skintype"`
    SkinColor       string          `gorm:"type:enum('Fair', 'Medium', 'Olive', 'Brown', 'Dark')" json:"skincolor"`
    CreatedAt		time.Time		`json:"createdat"`
	UpdatedAt		time.Time		`json:"updatedat"`
	DeletedAt		gorm.DeletedAt 	`json:"-" gorm:"index"`
}