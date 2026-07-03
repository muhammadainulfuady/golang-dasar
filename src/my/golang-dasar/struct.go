package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	Jurusan string
	Nilai   float64
}

func main() {
	mhs := []Mahasiswa{
		{"Budi", "Informatika", 80},
		{"Andi", "Sistem Informasi", 60},
		{"Siti", "Informatika", 95},
		{"Ainul", "Teknik Komputer", 88},
		{"Muhammad", "Informatika", 72},
	}
	showData(mhs)
	logicPassed(mhs)
	bestMark(mhs)
	lowestMark(mhs)
	totalNilai := sumMark(mhs)
	fmt.Println("Total Nilai Seluruh Mahasiswa:", totalNilai)
	fmt.Println("Rata-rata Nilai Seluruh Mahasiswa:", averageMark(mhs))
}

func showData(mhs []Mahasiswa) {
	for i, m := range mhs {
		fmt.Println("Data Mahasiswa ke-", i+1)
		fmt.Println("Nama : ", m.Nama, "\nJurusan : ", m.Jurusan, "\nNilai : ", m.Nilai)
		fmt.Println("")
	}
}

func logicPassed(mhs []Mahasiswa) {
	for _, m := range mhs {
		if m.Nilai >= 75 {
			fmt.Println("Mahasiswa yang lulus:", m.Nama)
		} else {
			fmt.Println("Mahasiswa yang tidak lulus:", m.Nama)
		}
	}
}

func sumMark(mhs []Mahasiswa) float64 {
	total := 0.0
	for _, m := range mhs {
		total += m.Nilai
	}
	return total
}

func averageMark(mhs []Mahasiswa) float64 {
	total := sumMark(mhs)
	return total / float64(len(mhs))
}

func bestMark(mhs []Mahasiswa) {
	best := mhs[0]
	for _, m := range mhs {
		if m.Nilai > best.Nilai {
			best = m
		}
	}
	fmt.Println("Mahasiswa dengan nilai tertinggi:", best.Nama, "\nNilai", best.Nilai)
}

func lowestMark(mhs []Mahasiswa) {
	lowest := mhs[0]
	for _, m := range mhs {
		if m.Nilai < lowest.Nilai {
			lowest = m
		}
	}
	fmt.Println("Mahasiswa dengan nilai terendah:", lowest.Nama, "\nNilai", lowest.Nilai)
}

func (m Mahasiswa) sayHello() {
	fmt.Println("Hello, my name is", m.Nama)
}
