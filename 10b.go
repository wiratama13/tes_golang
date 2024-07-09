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

// Fungsi untuk mengurutkan angka dari terbesar ke terkecil dan huruf dari terkecil ke terbesar
func sortString(input string) (string, string) {
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

	// Menghapus duplikasi dan mengurutkan angka
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	numbers = removeDuplicates(numbers)

	// Menghapus duplikasi dan mengurutkan huruf
	sort.Strings(letters)
	letters = removeStringDuplicates(letters)

	// Mengonversi kembali ke string
	sortedNumbers := intSliceToString(numbers)
	sortedLetters := strings.Join(letters, "")

	return sortedNumbers, sortedLetters
}

// Fungsi untuk menghapus duplikasi dari slice int
func removeDuplicates(arr []int) []int {
	uniqueMap := make(map[int]bool)
	for _, num := range arr {
		uniqueMap[num] = true
	}
	uniqueSlice := []int{}
	for num := range uniqueMap {
		uniqueSlice = append(uniqueSlice, num)
	}
	return uniqueSlice
}

// Fungsi untuk menghapus duplikasi dari slice string
func removeStringDuplicates(arr []string) []string {
	uniqueMap := make(map[string]bool)
	for _, str := range arr {
		uniqueMap[str] = true
	}
	uniqueSlice := []string{}
	for str := range uniqueMap {
		uniqueSlice = append(uniqueSlice, str)
	}
	return uniqueSlice
}

// Fungsi untuk mengonversi slice int ke string
func intSliceToString(arr []int) string {
	var result []string
	for _, num := range arr {
		result = append(result, fmt.Sprintf("%d", num))
	}
	return strings.Join(result, ",")
}

func main() {
	// Generate random string
	randomString := generateRandomString()
	fmt.Println("Generated String:", randomString)

	// Sort string
	sortedNumbers, sortedLetters := sortString(randomString)

	// Print sorted results
	fmt.Println("Sorted Numbers:", sortedNumbers)
	fmt.Println("Sorted Letters:", sortedLetters)
}
