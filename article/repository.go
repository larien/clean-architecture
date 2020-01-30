package article

import (
	"log"

	"github.com/larien/clean-architecture/helper/database"
)

// Repository defines the methods to be exposed in Repository layer
type Repository interface {
	// Create saves the article in database
	Create(a *Article) error
}

// repository holds the dependencies for Controller layer
type repository struct {
	DB database.Driver
}

// NewRepository creates a new Repository with access to database
func NewRepository(db database.Driver) Repository {
	db.AutoMigrate(&Article{})
	return &repository{db}
}

func (r *repository) Create(a *Article) error {
	return nil
}
