package main

import (
    "fmt"
    "strings"
)

// Fungsi untuk mengecek apakah elemen ada di dalam array dan memberikan saran kota jika tidak ada
func isPresentWithSuggestion(array []string, input string) (bool, []string) {
    suggestions := []string{}

    // Cek apakah input ada dalam array
    for _, item := range array {
        if item == input {
            return true, suggestions
        }
    }

    // Jika tidak ditemukan, cari saran berdasarkan abjad pertama dan terakhir
    inputFirstLetter := strings.ToUpper(string(input[0]))
    inputLastLetter := strings.ToUpper(string(input[len(input)-1]))
    for _, item := range array {
        itemFirstLetter := strings.ToUpper(string(item[0]))
        itemLastLetter := strings.ToUpper(string(item[len(item)-1]))

        if itemFirstLetter == inputFirstLetter || itemLastLetter == inputLastLetter {
            suggestions = append(suggestions, item)
        }
    }

    return false, suggestions
}

func main() {
    array := []string{"Bandung", "Cimahi", "Ambon", "Jayapura", "Makasar"}

    // Contoh penggunaan fungsi isPresentWithSuggestion
    input := "Bogor"
    present, suggestions := isPresentWithSuggestion(array, input)
    if present {
        fmt.Printf("Kota %s ditemukan dalam array.\n", input)
    } else {
        if len(suggestions) > 0 {
            fmt.Printf("Kota %s tidak ditemukan dalam array. Saran Kota: %v?\n", input, suggestions)
        } else {
            fmt.Printf("Kota %s tidak ditemukan dalam array dan tidak ada saran kota dengan abjad pertama atau terakhir yang sama.\n", input)
        }
    }   
}
