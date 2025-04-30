package main

import (
    "fmt"
)

func convertToTitle(columnNumber int) string {
    result := ""

    for columnNumber > 0 {
        columnNumber-- 
	    remainder := columnNumber % 26
        char := string('A' + remainder)
        result = char + result
        columnNumber /= 26
    }

    return result
}

func main() {
    var input int
    fmt.Print("Masukkan angka kolom: ")
    fmt.Scan(&input)

    result := convertToTitle(input)
    fmt.Printf("Nama kolom Excel: %s\n", result)
}
