package loader

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"github.com/nnmax1/BookAPI/model"
)


func LoadData(r io.Reader) *[]*model.BookData  {
	reader := csv.NewReader(r)

	ret := make([]*model.BookData, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		classic, _ := strconv.Atoi(row[6])
		isRecent, _ := strconv.Atoi(row[7])

		if err != nil {
			log.Println(err)
		}
		book := &model.BookData{
			Title:    row[0], 
			Author:   row[1],
			Format:   row[2],
			Url:      row[3], 
			Genre: 	  row[4],
			ImgUrl:   row[5],
			IsClassic: classic,
			IsRecent: isRecent,
			SortTitle: row[8],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, book)
	}
	return &ret
}