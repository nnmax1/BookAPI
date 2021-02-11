package datastore

import (
	"log"
	"os"
	"github.com/nnmax1/BookAPI/loader"
	"github.com/nnmax1/BookAPI/model"
)

type Books struct {
	Store *[]*model.BookData `json:"data"`
}


func (b *Books) Initialize() {
	filename := "./assets/book_data.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b.Store = loader.LoadData(file)
}