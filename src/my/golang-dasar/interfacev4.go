package main

import "fmt"

type Pengiriman interface {
	MengirimBarang() string
	MenghitungBiaya() int
	MenampilkanInformasi() string
}

type KurirMotor struct {
	Name           string
	NomorKendaraan string
	Kapasitas      int
}
type KurirMobil struct {
	Name           string
	NomorKendaraan string
	Kapasitas      int
}
type KurirDrone struct {
	Name           string
	NomorKendaraan string
	Kapasitas      int
}

func (k KurirMotor) MenghitungBiaya() int {
	return k.Kapasitas * 1000
}

func (k KurirMotor) MengirimBarang() string {
	status := "Kurir Motor Mengirim Paket"
	return status
}

func (k KurirMotor) MenampilkanInformasi() string {
	return showInformation("Kurir Motor", k.Name, k.NomorKendaraan, k.Kapasitas, k)
}

func (k KurirMobil) MenghitungBiaya() int {
	return k.Kapasitas * 1000
}

func (k KurirMobil) MengirimBarang() string {
	status := "Kurir Mobil Mengirim Paket"
	return status
}

func (k KurirMobil) MenampilkanInformasi() string {
	return showInformation("Kurir Mobil", k.Name, k.NomorKendaraan, k.Kapasitas, k)
}

func (k KurirDrone) MenghitungBiaya() int {
	return k.Kapasitas * 1000
}

func (k KurirDrone) MengirimBarang() string {
	status := "Kurir Drone Mengirim Paket"
	return status
}

func (k KurirDrone) MenampilkanInformasi() string {
	return showInformation("Kurir Drone", k.Name, k.NomorKendaraan, k.Kapasitas, k)
}

func TampilkanPengiriman(p Pengiriman) {
	fmt.Println(p.MenampilkanInformasi())
}

func main() {
	kurirMotor := KurirMotor{
		Name:           "Budi",
		NomorKendaraan: "AB1234CD",
		Kapasitas:      10,
	}

	kurirMobil := KurirMobil{
		Name:           "Andi",
		NomorKendaraan: "AB5678EF",
		Kapasitas:      20,
	}

	kurirDrone := KurirDrone{
		Name:           "Citra",
		NomorKendaraan: "DR1234",
		Kapasitas:      5,
	}

	TampilkanPengiriman(kurirMotor)
	TampilkanPengiriman(kurirMobil)
	TampilkanPengiriman(kurirDrone)

}

func showInformation(jenisKurir string, name string, nomorKendaraan string, kapasitas int, p Pengiriman) string {
	result := "\n"
	informasi := []string{
		"Jenis Kurir: " + jenisKurir,
		"Nama Kurir: " + name,
		"Nomor Kendaraan: " + nomorKendaraan,
		"Kapasitas: " + fmt.Sprint(kapasitas),
		"Biaya Pengiriman: " + fmt.Sprint(p.MenghitungBiaya()),
		"Status Pengiriman: " + p.MengirimBarang(),
	}
	for _, info := range informasi {
		result += info + "\n"
	}
	if kapasitas > 5 {
		result += "Muatan Berat"
	} else {
		result += "Muatan Ringan"
	}
	return result
}
