package model

type BookData struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Format      string  `json:"format"`
	Url         string  `json:"url"`
	Genre       string   `json:"genre"`
	ImgUrl      string  `json:"imgUrl"`
    IsClassic   int     `json:"isEssential"`
    IsRecent       int `json:"isRecent"`
	SortTitle   string   `json:"sortTitle"`
}