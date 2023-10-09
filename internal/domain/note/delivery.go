package note

import "gorm.io/gorm"

func NoteProvider(db *gorm.DB) Repository {
	repo := NewRepository(db)
	return repo
}
