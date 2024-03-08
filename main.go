package main

import (
	"fmt"
)

func main() {
	var sisi, alas, tinggi, jari, depan, samping, miring float64
	var s, t, v0, a, m, h, v, f, n, Vl, F, A float64
	var celcius, reamur, fahrenheit, kelvin float64
	var pilihan1, pilihan2 string

	fmt.Println("Pilihan Program yang akan digunakan dengan memasukkan input huruf proram")
	fmt.Println("A. Menghitung luas Persegi")
	fmt.Println("B. Menghitung luas Segitiga")
	fmt.Println("C. Menghitung luas Lingkaran")
	fmt.Println("D. Menghitung sudut sinus, cosinus, tangen")
	fmt.Println("E. Menghitung gerak lurus beraturan")
	fmt.Println("F. Menghitung gerak lurus berubah beraturan")
	fmt.Println("G. Menghitung energi potensial, kinetik")
	fmt.Println("H. Menghitung Frekuensi atau Getaran")
	fmt.Println("I. Menghitung masa jenis")
	fmt.Println("J. Menghitung daya, tekanan, usaha dan gaya")
	fmt.Println("K. Konversi untuk semua satuan suhu")
	fmt.Scanln(&pilihan1)

	switch pilihan1 {
	case "A":
		fmt.Println("Masukkan nilai panjang sisi Persegi")
		fmt.Scanln(&sisi)
		fmt.Printf("%.2f * %.2f = %.2f", alas, tinggi, alas*tinggi)

	case "B":
		fmt.Println("Masukkan nilai Alas")
		fmt.Scanln(&alas)
		fmt.Println("Masukkan nilai Tinggi")
		fmt.Scanln(&tinggi)
		fmt.Printf("%.2f * %.2f = %.2f", alas, tinggi, alas*tinggi)

	case "C":
		fmt.Println("Masukkan nilai panjang jari - jari lingkaran")
		fmt.Scanln(&jari)
		fmt.Printf("phi * %.2f * %.2f - %.2f", jari, jari, 3.14*jari*jari)

	case "D":
		fmt.Println("Pilih nilai yang ingin dicari dengan memasukkan input huruf")
		fmt.Println("A. Nilai Sinus")
		fmt.Println("B. Nilai Cosinus")
		fmt.Println("C. Nilai Tangen")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "A":
			fmt.Println("Masukkan nilai panjang sisi depan")
			fmt.Scanln(&depan)
			fmt.Println("Masukkan nilai panjang sisi miring")
			fmt.Scanln(&miring)
			fmt.Printf("Sin %.2f/%.2f - %.2f", depan, miring, depan/miring)

		case "B":
			fmt.Println("Masukkan nilai panjang sisi samping")
			fmt.Scanln(&samping)
			fmt.Println("Masukkan nilai panjang sisi miring")
			fmt.Scanln(&miring)
			fmt.Printf("Cas %.2f/%.2f = %.2f", samping, miring, samping/miring)

		case "C":
			fmt.Println("Masukkan nilai panjang sisi depan")
			fmt.Scanln(&depan)
			fmt.Println("Masukkan nilai panjang sisi samping")
			fmt.Scanln(&samping)
			fmt.Printf("Tan %.2f/%.2f = %.2f", depan, samping, depan/samping)

		}

	case "E":
		fmt.Println("Masukkan nilai jarak")
		fmt.Scanln(&s)
		fmt.Println("Masukkan nilai waktu")
		fmt.Scanln(&t)
		fmt.Printf("V - %.2f/%.2f - %.2f", s, t, s*t)

	case "F":
		fmt.Println("Masukkan nilai kecepatan awal")
		fmt.Scanln(&v0)
		fmt.Println("Masukkan nilai percepatan")
		fmt.Scanln(&a)
		fmt.Println("Masukkan nilai waktu")
		fmt.Scanln(&t)
		fmt.Printf("Vt = %.2f + %.2f*%.2f = %.2f", v0, a, t, v0+(a*t))

	case "G":
		fmt.Println("Pilih nilai yang ingin dicari dengan memasukkan input huruf")
		fmt.Println("A. Energi Pontensial")
		fmt.Println("B. Energi Kinetik")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "A":
			fmt.Println("Masukkan nilai Massa")
			fmt.Scanln(&m)
			fmt.Println("Masukkan nilai Ketinggian")
			fmt.Scanln(&h)
			fmt.Printf("Ep = m * g * h = %.2f", m*10*h)

		case "B":
			fmt.Println("Masukkan nilai Massa")
			fmt.Scanln(&m)
			fmt.Println("Masukkan nilai Kecepatan")
			fmt.Scanln(&v)
			fmt.Printf("Et = 1/2 * m * v^2 = %.2f", 0.5*m*v*v)
		}

	case "H":
		fmt.Println("Pilih nilai yang ingin dicari dengan memasukkan input huruf")
		fmt.Println("A. Frekuensi")
		fmt.Println("B. Getaran")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "A":
			fmt.Println("Masukkan nilai getaran")
			fmt.Scanln(&n)
			fmt.Println("Masukkan nilai waktu")
			fmt.Scanln(&t)
			fmt.Printf("f = %.2f/%.2f = %.2f", n, t, n/t)

		case "B":
			fmt.Println("Masukkan nilai Frekuensi")
			fmt.Scanln(&f)
			fmt.Println("Masukkan nalai waktu")
			fmt.Scanln(&t)
			fmt.Printf("n = %.2f * %.2f = %.2f", f, t, f*t)
		}

	case "I":
		fmt.Println("Masukkkan nilai Massa")
		fmt.Scanln(&m)
		fmt.Println("Masukkan nilai Volume")
		fmt.Scanln(&Vl)
		fmt.Printf("rho = %.2f/%.2f = %.2f", m, Vl, m/Vl)

	case "J":
		fmt.Println("Pilih nilai yang ingin dicari dengan memasukkan input huruf")
		fmt.Println("A. Daya")
		fmt.Println("B. Tekanan")
		fmt.Println("C. Usaha")
		fmt.Println("D. Gaya")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "A":
			fmt.Println("Masukkan nilai gaya")
			fmt.Scanln(&F)
			fmt.Println("Masukkan nilai kecepatan")
			fmt.Scanln(&v)
			fmt.Printf("P = %.2f/%.2f = %.2f", F, v, F/v)

		case "B":
			fmt.Println("Masukkan nilai gaya")
			fmt.Scanln(&F)
			fmt.Println("Masukkan nilai Luas Permukaan (m^2)")
			fmt.Scanln(&A)
			fmt.Printf("p = %.2f * %.2f = %.2f", F, A, F*A)

		case "C":
			fmt.Println("Masukkan nilai gaya")
			fmt.Scanln(&F)
			fmt.Println("Masukkan nilai Perpindahan")
			fmt.Scanln(&s)
			fmt.Printf("W = %.2f * %.2f = %.2f", F, s, F*s)

		case "D":
			fmt.Println("Masukkan nilai Massa")
			fmt.Scanln(&m)
			fmt.Println("Masukkan nilai Percapatan")
			fmt.Scanln(&a)
			fmt.Printf("F = %.2f * %.2f = %.2f", m, a, m*a)
		}

	case "K":
		fmt.Println("Pilih nilai suhu yang diketahui")
		fmt.Println("A. Celcius")
		fmt.Println("B. Reamur")
		fmt.Println("C. Fahreneit")
		fmt.Println("D. Kelvin")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "A":
			fmt.Println("Pilih suhu dalam °C")
			fmt.Scanln(&celcius)
			fmt.Printf("%.2f °R\n", celcius*4/5)
			fmt.Printf("%.2f °F\n", (celcius*9/5)+32)
			fmt.Printf("%.2f K\n", celcius+273)

		case "B":
			fmt.Println("Pilih nilai suhu dalam °R")
			fmt.Scanln(&reamur)
			fmt.Printf("%.2f °C\n", reamur*5/4)
			fmt.Printf("%.2f °F\n", (reamur*9/4)+32)
			fmt.Printf("%.2f K\n", (reamur*5/4)+273)

		case "C":
			fmt.Printf("Pilih nilai suhu dalam °F")
			fmt.Scanln(&fahrenheit)
			fmt.Printf("%.2f °C\n", 5/9*(fahrenheit-32))
			fmt.Printf("%.2f °R\n", 4/9*(fahrenheit+32))
			fmt.Printf("%.2f K\n", (5/9*(fahrenheit-32))+273)

		case "D":
			fmt.Println("Pilih nilai suhu dalam satuan kelvin")
			fmt.Scanln(&kelvin)
			fmt.Printf("%.2f °C\n", kelvin-273)
			fmt.Printf("%.2f °R\n", (kelvin-273)*4/5)
			fmt.Printf("%.2f °F\n", ((kelvin-273)*9/5)+32)
		}

	default:
		fmt.Println("Pilihan tidak valid")
	}
}
