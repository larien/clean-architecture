package article

type Controller interface {
	Create(p *Article) error
}

type controller struct {
	// attributes
}

func NewController(r Repository) Controller {
	return &controller{
		// initialized attributes
	}
}

func (c *controller) Create(p *Article) error {
	return nil
}
