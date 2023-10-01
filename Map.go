package main

import "fmt"

func main() {
	/*person := map[string]string{
		"name":      "tara",
		"education": "S1",
		"adress":    "Purbalingga",
		"age":       "19",
	}
	testing := map[string]int{
		"wa":   01234567,
		"usia": 19,
	}
	fmt.Println(person)
	fmt.Println(testing)

	fmt.Println("Jumlah data yang ada di map person adalah : ", len(person))
	fmt.Println("Jumlah data yang ada di map testing adalah : ", len(testing))

	fmt.Println("nama aku adalah :", person["name"], ", aku berasal dari", person["adress"])

	person["name"] = "Bathara"
	person["adress"] = "Purwokerto"

	fmt.Println("Setelah diubah")
	fmt.Println("nama aku adalah :", person["name"], ", aku berasal dari", person["adress"])

	delete(person, "age")
	fmt.Println("jumlah data map person setelah di hapus", len(person))
	*/

	Data := map[string]string{
		"nama":       "Ucup Surucup",
		"pendidikan": "S1 Teknik Elektro",
		"alamat":     "Bandung",
		"umur":       "19",
	}

	Organisasi := map[string]string{
		"smp":    "osis",
		"sma":    "pramuka",
		"kuliah": "pmr",
	}

	Nilai := map[string]float32{
		"semester_1": 3.5,
		"semester_2": 3.6,
		"semester_3": 3.7,
	}

	fmt.Println("Perkenalkan nama saya", Data["nama"], ". Saya sekarang kuliah", Data["pendidikan"], "di Purwokerto.")
	fmt.Println("Asal saya dari", Data["alamat"], "dan umur saya", Data["umur"], "tahun.")
	fmt.Println("Saat SMP saya aktif di", Organisasi["smp"], " ,lalu saat SMA saya aktif di", Organisasi["sma"], "dan saat kuliah saya bergabung di", Organisasi["kuliah"], ".")
	fmt.Println("Nilai kuliah saya semasa semester 1 adalah ", Nilai["semester_1"], ", semester 2 adalah", Nilai["semester_2"], "dan ketika semester 3 adalah", Nilai["semester_3"])
}
