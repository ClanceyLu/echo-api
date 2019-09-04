package models

import (
	"time"

	"github.com/ClanceyLu/echo-api/pkg/mysql"
	"github.com/jinzhu/gorm"
)

// Model definition
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

var db *gorm.DB

func init() {
	db = mysql.Connect()
}
