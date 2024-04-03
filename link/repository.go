package link

import "gorm.io/gorm"

// Repository represents the repository for link
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new link repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

// Add adds a new link
func (r *Repository) Add(l AddPayload) (res Link, err error) {
	// save link to database

	link := Link{URL: l.URL, Title: l.Title, Description: l.Description}
	tx := r.db.Create(&link)

	if tx.Error != nil {
		err = tx.Error
		return
	}

	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
		return
	}

	res = link
	return
}
