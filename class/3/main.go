package main

import "fmt"

func daftarMerekSepatu(daftar []string, merekYangDibeli string) []string {
	// catat merek sepatu yang dibeli ke dalam daftar dan tidak boleh ada duplikasi (unique)
	flag := false
	for _, merek := range daftar {
		if merek == merekYangDibeli {
			flag = true
		}
	}

	if !flag {
		daftar = append(daftar, merekYangDibeli)
	}

	return daftar
}

func hitungBelanja(keranjangBelanja []string) (int, int) {
	var totalBelanja int = 0

	merekSepatu := []string{}

	const hargaAdidas = 200000
	const hargaPuma = 15000
	const hargaKappa = 600000

	// hitung total belanja
	for _, item := range keranjangBelanja {

		switch item {
		case "adidas":
			totalBelanja += hargaAdidas
		case "puma":
			totalBelanja += hargaPuma
		case "kappa":
			totalBelanja += hargaKappa
		}

		// buat daftar merek sepatu yang dibeli
		merekSepatu = daftarMerekSepatu(merekSepatu, item)
	}

	// hitung diskon dengan cara concat/join merek-merek sepatu
	var semuaMerek string
	for _, merek := range merekSepatu {
		semuaMerek += merek
	}

	var totalDiskon int = 0
	switch semuaMerek {
	case "adidaspuma", "pumaadidas":
		totalDiskon = 50000
	case "pumakappa", "kappapuma":
		totalDiskon = 150000
	case "adidaskappa", "kappaadidas":
		totalDiskon = 75000
	}

	return totalBelanja, totalDiskon
}

func main() {
	//contoh 1
	keranjang := []string{"kappa", "adidas"}
	total, diskonSepatu := hitungBelanja(keranjang)
	fmt.Printf("%s -> Harga total %d, diskon %d, total bayar = %d\n", keranjang, total, diskonSepatu, total-diskonSepatu)

	//contoh 2
	keranjang = []string{"kappa", "kappa"}
	total, diskonSepatu = hitungBelanja(keranjang)
	fmt.Printf("%s -> Harga total %d, diskon %d, total bayar = %d\n", keranjang, total, diskonSepatu, total-diskonSepatu)
}
