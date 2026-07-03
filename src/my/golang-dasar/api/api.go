package api

import (
	"encoding/json"
	"log"
)

type Book struct {
	ID          int
	Nama        string
	TahunTerbit string
	Penulis     string
	Penerbit    int
}

func ApiCalling() string {
	books := []Book{
		{
			ID:          1,
			Nama:        "Laskar Pelangi",
			TahunTerbit: "2005",
			Penulis:     "Andrea Hirata",
			Penerbit:    101,
		},
		{
			ID:          2,
			Nama:        "Bumi",
			TahunTerbit: "2014",
			Penulis:     "Tere Liye",
			Penerbit:    102,
		},
		{
			ID:          3,
			Nama:        "Laut Bercerita",
			TahunTerbit: "2017",
			Penulis:     "Leila S. Chudori",
			Penerbit:    101,
		},
		{
			ID:          4,
			Nama:        "Cantik Itu Luka",
			TahunTerbit: "2002",
			Penulis:     "Eka Kurniawan",
			Penerbit:    103,
		},
		{
			ID:          5,
			Nama:        "Perahu Kertas",
			TahunTerbit: "2009",
			Penulis:     "Dee Lestari",
			Penerbit:    102,
		},
		{
			ID:          6,
			Nama:        "Ronggeng Dukuh Paruk",
			TahunTerbit: "1982",
			Penulis:     "Ahmad Tohari",
			Penerbit:    101,
		},
		{
			ID:          7,
			Nama:        "Gadis Kretek",
			TahunTerbit: "2012",
			Penulis:     "Ratih Kumala",
			Penerbit:    103,
		},
		{
			ID:          8,
			Nama:        "Aroma Karsa",
			TahunTerbit: "2018",
			Penulis:     "Dee Lestari",
			Penerbit:    102,
		},
		{
			ID:          9,
			Nama:        "Filosofi Kopi",
			TahunTerbit: "2006",
			Penulis:     "Dee Lestari",
			Penerbit:    104,
		},
		{
			ID:          10,
			Nama:        "Negeri 5 Menara",
			TahunTerbit: "2009",
			Penulis:     "Ahmad Fuadi",
			Penerbit:    101,
		},
	}

	json, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		log.Fatalf("Gagal convert ke JSON: %s", err)
	}
	return string(json)
}
