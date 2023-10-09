package note

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Uid  string
	Code string
	Data string
}
