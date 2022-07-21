package main

import (
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.SetupJoinTable(&User{}, "Languages", &UserLanguage{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	english := Language{
		Name: "English",
	}
	DB.Save(&english)

	user := User{
		Name: "jinzhu",
		Languages: []*Language{
			&english,
		},
	}

	DB.Omit("Languages.*").Save(&user)

	var result UserLanguage
	if err := DB.Limit(1).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.ID == "" {
		t.Error("Failed, BeforeCreate() is not working")
	} else {
		log.Printf("BeforeCreate() is working, %+v", result)
	}
}
