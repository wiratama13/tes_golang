package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Fungsi untuk menghasilkan string acak dengan jumlah huruf dan angka yang ditentukan
func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numbers = []rune("0123456789")

	result := make([]rune, 100)
	for i := 0; i < 50; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	for i := 50; i < 100; i++ {
		result[i] = numbers[rand.Intn(len(numbers))]
	}

	return string(result)
}

// Fungsi untuk mengurutkan hasil string random sesuai aturan yang diminta
func sortString(input string) string {
	var numbers []int
	var letters []string

	// Memisahkan angka dan huruf
	for _, char := range input {
		if char >= '0' && char <= '9' {
			numbers = append(numbers, int(char-'0'))
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			letters = append(letters, string(char))
		}
	}

	// Mengurutkan angka dari terbesar ke terkecil
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	// Mengurutkan huruf dari terkecil ke terbesar
	sort.Strings(letters)

	// Menggabungkan hasil sesuai format yang diinginkan
	var result []string
	for _, num := range numbers {
		result = append(result, fmt.Sprintf("%d", num))
	}
	for _, letter := range letters {
		result = append(result, letter)
	}

	return strings.Join(result, ",")
}

func main() {
	// Generate random string
	randomString := generateRandomString()
	fmt.Println("Generated String:", randomString)

	// Sort string
	sortedString := sortString(randomString)

	// Print sorted result
	fmt.Println("Sorted String:", sortedString)
}
