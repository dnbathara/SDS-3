package main

import "fmt"

type DataDiri struct {
	Nama, Jurusan, Fakultas string
	Umur, WA                int
}

func main() {
	Ucup := DataDiri{
		Nama:     "Ucup Surucup",
		Jurusan:  "Elektro",
		Fakultas: "Teknik",
		Umur:     19,
		WA:       034334554,
	}

	Jamilah := DataDiri{
		Nama:     "Jamilah",
		Jurusan:  "Keperawatan",
		Fakultas: "Fikes",
		Umur:     19,
		WA:       032112121,
	}

	fmt.Println("Nama saya adalah :", Ucup.Nama)
	fmt.Println("Jurusan saya adalah :", Ucup.Jurusan)
	fmt.Println("Fakultas saya adalah :", Ucup.Fakultas)
	fmt.Println("Umur saya adalah : ", Ucup.Umur, "tahun.")
	fmt.Println("Nomor Whatsapp saya adalah: ", Ucup.WA)
	fmt.Println("")

	fmt.Println("Nama saya adalah :", Jamilah.Nama)
	fmt.Println("Jurusan saya adalah :", Jamilah.Jurusan)
	fmt.Println("Fakultas saya adalah :", Jamilah.Fakultas)
	fmt.Println("Umur saya adalah : ", Jamilah.Umur, "tahun.")
	fmt.Println("Nomor Whatsapp saya adalah: ", Jamilah.WA)
	fmt.Println("")
	fmt.Println("=====Setelah di rubah=====")
	
	Ucup.Umur = 20
	Jamilah.Jurusan = "Sastra Indonesia"
	Jamilah.Fakultas = "FIB"
	
	fmt.Println("Nama saya adalah :", Ucup.Nama)
	fmt.Println("Jurusan saya adalah :", Ucup.Jurusan)
	fmt.Println("Fakultas saya adalah :", Ucup.Fakultas)
	fmt.Println("Umur saya adalah : ", Ucup.Umur, "tahun.")
	fmt.Println("Nomor Whatsapp saya adalah: ", Ucup.WA)
	fmt.Println("")
	
	fmt.Println("Nama saya adalah :", Jamilah.Nama)
	fmt.Println("Jurusan saya adalah :", Jamilah.Jurusan)
	fmt.Println("Fakultas saya adalah :", Jamilah.Fakultas)
	fmt.Println("Umur saya adalah : ", Jamilah.Umur, "tahun.")
	fmt.Println("Nomor Whatsapp saya adalah: ", Jamilah.WA)
}
