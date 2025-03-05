package main

import "fmt"

func Tiket() {
	fmt.Println("=====================")
	fmt.Println("=======PARKING=======")
	fmt.Println("=====================")
	fmt.Println("===PILIH KENDARAAN===")
	fmt.Println("       1. Mobil      ")
	fmt.Println("       2. Motor      ")
	fmt.Println("       3. Bus        ")
	fmt.Println("       4. Truk       ")
	fmt.Println("=====================")
}

func mobil(x int) int {
	var jammobi int
	y := 3000
	jammobi = x * y
	return jammobi
}

func motor(x int) int {
	var jam int
	y := 1500
	jam = x * y
	return jam
}

func Bus(x int) int {
	var jam int
	y := 5000
	jam = x * y
	return jam
}

func truk(x int) int {
	var jam int
	y := 4500
	jam = x * y
	return jam
}

func main() {
	var pilih, parkir int
	var jamMobil, jamMtor, jamBus, jamTruk int

	Tiket()
	fmt.Println("Selamat datang di E-parking")
	fmt.Println("Silahkan pilih kendaraan anda")
	fmt.Scanln(&pilih)

	for pilih != 5 {
		if pilih == 1 {
			fmt.Println("Masukkan berapa jam anda parkir")
			fmt.Scanln(&parkir)
			jamMobil = mobil(parkir)
			fmt.Printf("Total bayar parkir mobil anda RP%d\n", jamMobil)
		} else if pilih == 2 {
			fmt.Println("Masukkan berapa jam anda parkir")
			fmt.Scanln(&parkir)
			jamMtor = motor(parkir)
			fmt.Printf("Total bayar parkir motor anda RP%d\n", jamMtor)
		} else if pilih == 3 {
			fmt.Println("Masukkan berapa jam anda parkir")
			fmt.Scanln(&parkir)
			jamBus = Bus(parkir)
			fmt.Printf("Total bayar parkir BUS RP%d\n", jamBus)
		} else if pilih == 4 {
			fmt.Println("Masukkan berapa jam anda parkir")
			fmt.Scanln(&parkir)
			jamTruk = truk(parkir)
			fmt.Printf("Total bayar parkir Truk RP%d\n", jamTruk)
		} else {
			fmt.Println("Selamat datang pegawai tercinta")
		}

		fmt.Println("=====================")
		fmt.Println("Silahkan pilih kendaraan lagi atau tekan 5 untuk keluar")
		fmt.Scanln(&pilih)
	}

	fmt.Println("Terima kasih telah menggunakan layanan E-parking")
}
