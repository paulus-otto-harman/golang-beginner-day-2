package main

func luasLingkaran(jariJari int) float32 {
	const pi = 3.14

	return pi * float32(jariJari) * float32(jariJari)
}

func namaHari(angka uint) string {
	var nama string

	switch angka {
	case 1:
		nama = "Senin"
	case 2:
		nama = "Selasa"
	case 3:
		nama = "Rabu"
	case 4:
		nama = "Kamis"
	case 5:
		nama = "Jumat"
	case 6:
		nama = "Sabtu"
	case 7:
		nama = "Minggu"
	default:
		nama = "-"
	}

	return nama
}

func cariNama(cari string) int {
	daftarNama := []string{"bambang", "dewi", "andi", "budi", "amir"}
	var ketemu int
	for index, nama := range daftarNama {
		if nama == cari {
			ketemu = index + 1
			break
		}
	}

	return ketemu
}

func konversiGrade(nilai int) string {
	var grade string

	switch {
	case nilai <= 50:
		grade = "E"
	case nilai < 60:
		grade = "D"
	case nilai < 70:
		grade = "C"
	case nilai < 80:
		grade = "B"
	default:
		grade = "A"
	}

	return grade
}

func prima(nilai int) bool {
	if nilai <= 1 {
		return false
	}

	for i := 2; i*i <= nilai; i++ {
		if nilai%i == 0 {
			return false
		}
	}

	return true
}

func tampilkanPrima(deret []int) []int {
	deretPrima := []int{}

	for _, angka := range deret {
		if prima(angka) {
			deretPrima = append(deretPrima, angka)
		}
	}

	return deretPrima
}

// function ini mengembalikan lebih dari 1 nilai
func luasDanKelilingLingkaran(jariJari float32) (float32, float32) {
	const pi = 3.14

	luasLingkaran := pi * jariJari * jariJari
	kelilingLingkaran := 2 * pi * jariJari

	return luasLingkaran, kelilingLingkaran
}

// diskon per item
func hargaNet(hargaJual float32, diskon float32) float32 {
	return hargaJual - hargaJual*diskon/100

}

// variadic
func cekHarga(discount float32, daftarHarga ...float32) []float32 {
	hasilDiskon := []float32{}

	for _, harga := range daftarHarga {
		hasilDiskon = append(hasilDiskon, hargaNet(float32(harga), discount))
	}

	return hasilDiskon
}

func main() {
	//fmt.Println(luasLingkaran(10))

	//fmt.Println(namaHari(3))

	//fmt.Println(cariNama("bambang"))

	//const nilai = 40
	//fmt.Printf("nilai %d adalah %s \n", nilai, konversiGrade(nilai))

	//bilangan := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	//fmt.Println(tampilkanPrima(bilangan))

	//const r = 10
	//luas, keliling := luasDanKelilingLingkaran(r)
	//fmt.Printf("Lingkaran dengan jari-jari %d, memiliki luas %f dan keliling %f \n", r, luas, keliling)

	//const hargaBarang = 100000
	const diskon = 10
	//fmt.Printf("harga barang %d didiskon sebesar %d%%, adalah %f \n", hargaBarang, diskon, hargaNet(hargaBarang, diskon))

	//fmt.Println("100000, 200000, 300000", cekHarga(diskon, 100000, 200000, 300000))

}
