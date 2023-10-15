package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID         uint            `gorm:"primary_key"`
	Name       string          `gorm:"varchar(255)"`
	Note       string          `gorm:"varchar(255)"`
	IsComplete bool            `gorm:"boolean,default:false"`
	CreatedAt  time.Time       `gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `gorm:"autoUpdateTime"`
	DeletedAt  *gorm.DeletedAt `gorm:"index"`
}
