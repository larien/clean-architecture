package article

// Controller defines the methods to be exposed in Controller layer
type Controller interface {
	// Create requests the received article to be stored
	Create(a *Article) error
	// List returns all the stored articles
	List() (*[]Article, error)
	// Detail returns an article by its ID
	Detail(ID uint) (*Article, error)
}

// controller holds the dependencies for Controller layer
type controller struct {
	Repository
}

// NewController creates a new Controller with access to Repository methods
func NewController(r Repository) Controller {
	return &controller{
		Repository: r,
	}
}

func (c *controller) Create(a *Article) error {
	return c.Repository.Create(a)
}

func (c *controller) List() (*[]Article, error) {
	return c.Repository.List()
}

func (c *controller) Detail(id uint) (*Article, error) {
	return nil, nil
}
