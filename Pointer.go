package main

import "fmt"

func main() {
	var Nama = "Ucup Surucup"
	var NIM = "K1A033483"
	var Alamat = "Banyumas"

	fmt.Println(Nama)
	fmt.Println(NIM)
	fmt.Println(Alamat)
	fmt.Println(&Alamat)
	
	fmt.Println("")
	fmt.Println("=====Setelah perubahan=====")
	fmt.Println("")
	
	
	Nama = "Jajang"
	NIM = "J1A099834"
	Alamat = "Rembang"
	
	fmt.Println(Nama)
	fmt.Println(NIM)
	fmt.Println(Alamat)
	fmt.Println(&Alamat)

}
