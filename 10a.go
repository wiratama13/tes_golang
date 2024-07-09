package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
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

// Fungsi untuk menghitung statistik dari string yang dihasilkan
func countStatistics(input string) (int, int, int, int) {
	totalLetters := 0
	totalVowels := 0
	totalNumbers := 0
	totalEvenNumbers := 0

	vowels := "aeiouAEIOU"
	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			totalLetters++
			if strings.ContainsRune(vowels, char) {
				totalVowels++
			}
		} else if char >= '0' && char <= '9' {
			totalNumbers++
			if char%2 == 0 {
				totalEvenNumbers++
			}
		}
	}

	return totalLetters, totalVowels, totalNumbers, totalEvenNumbers
}

func main() {
	// Generate random string
	randomString := generateRandomString()
	fmt.Println("Generated String:", randomString)

	// Hitung statistik
	totalLetters, totalVowels, totalNumbers, totalEvenNumbers := countStatistics(randomString)

	// Tampilkan hasil
	fmt.Printf("Total Huruf: %d\n", totalLetters)
	fmt.Printf("Total Huruf Vokal: %d\n", totalVowels)
	fmt.Printf("Total Angka: %d\n", totalNumbers)
	fmt.Printf("Total Angka Genap: %d\n", totalEvenNumbers)
}
