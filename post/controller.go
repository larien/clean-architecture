package post

type Controller interface {
	Create(p *Post) error
}

type controller struct {
	// attributes
}

func NewController(r Repository) Controller {
	return &controller{
		// initialized attributes
	}
}

func (c *controller) Create(p *Post) error {
	return nil
}
