package article

import (
	"testing"
)

func TestRepository_Create(t *testing.T) {
	t.Run("when repository fails to create article", func(t *testing.T) {
		// is := assert.New(t)

		// db, err := database.New("localhost", "larien", "clean_architecture", "")
		// if err != nil {
		// 	t.Fatalf("failed to connect to database: %v", err)
		// }

		// repository := NewRepository(db)
		// is.NotNil(repository)

		// db.Close()
		// article := &Article{}
		// is.Nil(faker.FakeData(article))
		// is.NotNil(repository.Create(article))
	})
	t.Run("when repository creates the article", func(t *testing.T) {
		// is := assert.New(t)

		// db, err := database.New("localhost", "larien", "clean_architecture", "")
		// if err != nil {
		// 	t.Fatalf("failed to connect to database: %v", err)
		// }

		// repository := NewRepository(db)
		// is.NotNil(repository)

		// article := &Article{}
		// is.Nil(faker.FakeData(article))
		// article.ID = 0
		// is.Nil(repository.Create(article))
		// time.Sleep(time.Second)
		// article2 := &Article{}
		// res := db.First(article2, 7)
		// fmt.Printf("1: %+v\n", article)
		// fmt.Printf("2: %+v\n", article2)
		// is.Nil(res.Error)
		// is.Equal(article.ID, article2.ID)
	})
}
