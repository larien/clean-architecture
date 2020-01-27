package article

// Repository defines the methods to be exposed in Repository layer
type Repository interface {
	// Create saves the article in database
	Create(a *Article) error
}

type repository struct {
	// attributes
}

// NewRepository creates a new Repository with access to database
func NewRepository() Repository {
	// create database
	return &repository{
		// initialized attributes
	}
}

func (r *repository) Create(a *Article) error {
	return nil
}
