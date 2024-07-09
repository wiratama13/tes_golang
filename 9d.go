package main

import (
    "fmt"
)

// Fungsi untuk menambahkan nilai pada array dengan maksimal setiap elemen adalah 10
func addValueToArray(arr []int, value int) []int {
    result := []int{}
    carry := value

    for _, num := range arr {
        sum := num + carry
        if sum > 10 {
            result = append(result, 10)
            carry = sum - 10
        } else {
            result = append(result, sum)
            carry = 0
        }
    }

    // Jika ada sisa carry, tambahkan ke elemen-elemen berikutnya
    for carry > 0 {
        if carry > 10 {
            result = append(result, 10)
            carry -= 10
        } else {
            result = append(result, carry)
            carry = 0
        }
    }

    return result
}

func main() {
    array := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
    input := 15

    // Tambahkan nilai ke dalam array
    result := addValueToArray(array, input)

    // Tampilkan hasil
    fmt.Println(result) // Output: [10 10 10 5 8 6 6 3 8 2 3 3 4 1 8 2]
}
