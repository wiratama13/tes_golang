package main

import (
    "fmt"
)

// Fungsi untuk menghitung jumlah duplikat dari setiap elemen
func countDuplicates(arr []int) map[int]int {
    countMap := make(map[int]int)
    for _, num := range arr {
        countMap[num]++
    }
    return countMap
}

// Fungsi untuk menampilkan hasil dalam format yang diinginkan
func formatCountMap(countMap map[int]int) string {
    result := ""
    first := true
    for num, count := range countMap {
        if !first {
            result += ","
        }
        result += fmt.Sprintf("%d[%d]", num, count)
        first = false
    }
    return result
}

func main() {
    array := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}

    // Hitung duplikat
    countMap := countDuplicates(array)

    // Tampilkan hasil
    result := formatCountMap(countMap)
    fmt.Println(result) // Output: 9[1],1[2],6[3],4[2],8[3],3[3],2[2]
}
