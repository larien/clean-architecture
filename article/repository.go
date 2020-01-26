package article

type Repository interface {
	// TODO - implement methods
}

type repository struct {
	// attributes
}

func NewRepository() Repository {
	return repository{
		// initialized attributes
	}
}

func (r *repository) Method(){
	// method stuff
}

// methods implemented