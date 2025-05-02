package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("\n--- MINI OYUNLAR ---")
		fmt.Println("1. Sayı Tahmin Oyunu")
		fmt.Println("2. Taş-Kağıt-Makas")
		fmt.Println("3. Kelime Tahmini")
		fmt.Println("4. Basit Matematik Testi")
		fmt.Println("5. Zar Atma Oyunu")
		fmt.Println("0. Çıkış")
		fmt.Print("Seçiminiz: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		secim := scanner.Text()

		switch secim {
		case "1":
			sayiTahminOyunu()
		case "2":
			tasKagitMakas()
		case "3":
			kelimeTahmini()
		case "4":
			matematikTesti()
		case "5":
			zarAtma()
		case "0":
			fmt.Println("Çıkış yapılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

func sayiTahminOyunu() {
	hedef := rand.Intn(100) + 1
	hak := 5

	fmt.Println("\n1-100 arası bir sayı tuttum. 5 hakkın var!")

	for hak > 0 {
		fmt.Printf("Kalan hak: %d | Tahmin: ", hak)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		tahmin, _ := strconv.Atoi(scanner.Text())

		if tahmin == hedef {
			fmt.Println("Doğru bildin!")
			return
		} else if tahmin < hedef {
			fmt.Println("Daha büyük")
		} else {
			fmt.Println("Daha küçük")
		}
		hak--
	}

	fmt.Printf("Kaybettin! Sayı: %d\n", hedef)
}

func tasKagitMakas() {
	secenekler := []string{"taş", "kağıt", "makas"}
	bilgisayar := secenekler[rand.Intn(3)]

	fmt.Print("\nTaş mı, kağıt mı, makas mı? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	oyuncu := strings.ToLower(scanner.Text())

	if oyuncu != "taş" && oyuncu != "kağıt" && oyuncu != "makas" {
		fmt.Println("Geçersiz seçim!")
		return
	}

	fmt.Printf("Bilgisayar: %s\n", bilgisayar)

	if oyuncu == bilgisayar {
		fmt.Println("Berabere!")
	} else if (oyuncu == "taş" && bilgisayar == "makas") ||
		(oyuncu == "kağıt" && bilgisayar == "taş") ||
		(oyuncu == "makas" && bilgisayar == "kağıt") {
		fmt.Println("Kazandın!")
	} else {
		fmt.Println("Kaybettin!")
	}
}

func kelimeTahmini() {
	kelimeler := []string{"go", "programlama", "bilgisayar", "algoritma", "yazılım"}
	hedef := kelimeler[rand.Intn(len(kelimeler))]
	tahminler := make(map[string]bool)
	can := 5

	fmt.Println("\nHarf tahmin et! (5 hakkın var)")

	for can > 0 {
		fmt.Printf("Kalan can: %d | Kelime: ", can)
		for _, harf := range hedef {
			if tahminler[string(harf)] {
				fmt.Printf("%c ", harf)
			} else {
				fmt.Print("_ ")
			}
		}

		fmt.Print("\nTahmin: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		tahmin := strings.ToLower(scanner.Text())

		if len(tahmin) != 1 {
			fmt.Println("Sadece 1 harf girin!")
			continue
		}

		if tahminler[tahmin] {
			fmt.Println("Bu harfi zaten denedin!")
			continue
		}

		tahminler[tahmin] = true

		if strings.Contains(hedef, tahmin) {
			fmt.Println("Doğru harf!")
		} else {
			fmt.Println("Yanlış harf!")
			can--
		}

		// Kelimeyi tamamen tahmin etti mi kontrolü
		tamaminiBildi := true
		for _, harf := range hedef {
			if !tahminler[string(harf)] {
				tamaminiBildi = false
				break
			}
		}

		if tamaminiBildi {
			fmt.Printf("Tebrikler! Kelime: %s\n", hedef)
			return
		}
	}

	fmt.Printf("Kaybettin! Kelime: %s\n", hedef)
}

func matematikTesti() {
	sorular := 5
	puan := 0

	fmt.Println("\n5 soruluk matematik testi başlıyor!")

	for i := 1; i <= sorular; i++ {
		a := rand.Intn(10) + 1
		b := rand.Intn(10) + 1
		islem := rand.Intn(3) // 0:+, 1:-, 2:*

		var cevap int
		var islemStr string

		switch islem {
		case 0:
			cevap = a + b
			islemStr = "+"
		case 1:
			cevap = a - b
			islemStr = "-"
		case 2:
			cevap = a * b
			islemStr = "*"
		}

		fmt.Printf("%d) %d %s %d = ", i, a, islemStr, b)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		kullaniciCevap, _ := strconv.Atoi(scanner.Text())

		if kullaniciCevap == cevap {
			fmt.Println("Doğru!")
			puan++
		} else {
			fmt.Printf("Yanlış! Doğru cevap: %d\n", cevap)
		}
	}

	fmt.Printf("Test bitti! Puan: %d/%d\n", puan, sorular)
}

func zarAtma() {
	fmt.Println("\nZar atma oyunu! Bilgisayarla yarış.")
	fmt.Print("Zarını at (ENTER'a bas): ")
	bufio.NewScanner(os.Stdin).Scan()

	oyuncuZar := rand.Intn(6) + 1
	bilgisayarZar := rand.Intn(6) + 1

	fmt.Printf("Sen: %d | Bilgisayar: %d\n", oyuncuZar, bilgisayarZar)

	if oyuncuZar > bilgisayarZar {
		fmt.Println("Kazandın!")
	} else if oyuncuZar < bilgisayarZar {
		fmt.Println("Kaybettin!")
	} else {
		fmt.Println("Berabere!")
	}
}
