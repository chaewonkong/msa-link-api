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
func (r *Repository) Add(p AddPayload) (res Link, err error) {
	// save link to database

	link := Link{URL: p.URL, Title: p.Title, Description: p.Description}
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

// Update updates a link
func (r *Repository) Update(p UpdatePayload) (res Link, err error) {
	link := Link{
		Model: gorm.Model{ID: p.ID},
	}
	tx := r.db.Model(&link).
		Updates(Link{Title: p.Title, Description: p.Description, ThumbnailImage: p.ThumbnailImage})

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
