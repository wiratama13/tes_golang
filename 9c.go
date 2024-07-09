package main

import (
    "fmt"
    "strings"
)

// Fungsi untuk menghapus nilai tertentu dari array
func removeElements(arr []int, toRemove []int) []int {
    removeMap := make(map[int]bool)
    for _, num := range toRemove {
        removeMap[num] = true
    }

    result := []int{}
    for _, num := range arr {
        if !removeMap[num] {
            result = append(result, num)
        }
    }
    return result
}

// Fungsi untuk mengubah input string menjadi slice int
func parseInput(input string) []int {
    parts := strings.Split(input, ",")
    result := make([]int, len(parts))
    for i, part := range parts {
        fmt.Sscanf(part, "%d", &result[i])
    }
    return result
}

func main() {
    array := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
    input := "9,1,6,4"

    // Parse input
    toRemove := parseInput(input)

    // Hapus elemen-elemen tertentu
    result := removeElements(array, toRemove)

    // Tampilkan hasil
    fmt.Println(result) // Output: [8 3 8 2 3 3 8 2]
}
