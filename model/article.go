package model

type Article struct {
	IdArtikel   int    `json:"id_artikel" gorm:"primaryKey"`
	PostingDate int64  `json:"posting_date"`
	Kategori    string `json:"kategori"`
	Judul       string `json:"judul"`
	IsiArtikel  string `json:"isi_artikel"`
	IdUser      int    `json:"id_user"`
}
