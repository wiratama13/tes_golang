package main

import (
    "fmt"
    "sort"
)

// Fungsi untuk mengurutkan array dan menghilangkan duplikat
func sortAndRemoveDuplicates(arr []int) []int {
    // Gunakan map untuk menghapus duplikat
    uniqueMap := make(map[int]bool)
    for _, num := range arr {
        uniqueMap[num] = true
    }

    // Pindahkan elemen unik ke dalam slice
    uniqueSlice := []int{}
    for num := range uniqueMap {
        uniqueSlice = append(uniqueSlice, num)
    }

    // Urutkan slice
    sort.Ints(uniqueSlice)

    return uniqueSlice
}

func main() {
    array := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
    sortedArray := sortAndRemoveDuplicates(array)
    fmt.Println(sortedArray)  // Output: [1 2 3 4 6 8 9]
}
