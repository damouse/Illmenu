package illmenu

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Base model implementation
type Model struct {
	Id        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// Representation of a dish on a menu
// In the future we can also include information about the query from the user, including geolocation
type Dish struct {
	Model
	Name   string `json:"name"`
	Images string `json:"images"` // ordered, comma-seperated list of urls to images

	// Metadata
	Hits         int       `json:"-"`
	Scrapes      int       `json:"-"`
	LastSearched time.Time `json:"lastSearched"`
}

// Initialize GORM. connect to the postgres db, and initialize the schema
func InitORM(auth string, shouldLog bool) *gorm.DB {
	db, err := gorm.Open("postgres", auth)
	panicOnError(err)
	db.LogMode(shouldLog)

	db.AutoMigrate(&Dish{})

	return db
}
