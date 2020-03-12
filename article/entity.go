package article

import (
	"encoding/json"
	"net/http"
	"time"
)

// Article represents the article entity's attributes.
type Article struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}

func (a *Article) Decode(r *http.Request) error {
	json.NewDecoder(r.Body).Decode(&a)
	return nil
}
