package main

import (
    "fmt"
)

// Fungsi untuk mengecek apakah elemen ada di dalam array
func isPresent(array []string, input string) bool {
    for _, item := range array {
        if item == input {
            return true
        }
    }
    return false
}

func main() {
    array := []string{"Bandung", "Cimahi", "Ambon", "Jayapura", "Makasar"}

    // Contoh penggunaan fungsi isPresent
    fmt.Println(isPresent(array, "Bandung"))  // Output: true
    fmt.Println(isPresent(array, "Jakarta"))  // Output: false
}
