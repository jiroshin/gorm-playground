package main

import (
	"github.com/rs/xid"
	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages"`
}

type Language struct {
	gorm.Model
	Name string
}

// Customize JoinTable
type UserLanguage struct {
	OriginalModel
	UserID     uint `gorm:"uniqIndex:idx_user_languages_user_id_language_id"`
	LanguageID uint `gorm:"uniqIndex:idx_user_languages_user_id_language_id"`
}

type OriginalModel struct {
	ID string `gorm:"primaryKey"`
}

func (o *OriginalModel) BeforeCreate(db *gorm.DB) error {
	if o.ID == "" {
		o.ID = xid.New().String()
	}
	return nil
}
