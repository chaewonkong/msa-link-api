package link

import "gorm.io/gorm"

// Link gorm Model
type Link struct {
	gorm.Model
	URL            string `gorm:"not null"`
	Title          string
	Description    string
	ThumbnailImage string
}
