package main

import "fmt"

// Fungsi sederhana tanpa parameter dan tanpa return
func sayHello() {
    fmt.Println("Hello, Golang!")
}

// Fungsi dengan parameter
func greet(name string) {
    fmt.Println("Hello,", name)
}

// Fungsi dengan return value
func add(a int, b int) int {
    return a + b
}

func displayInfo() {
    var nama string = "Yuki"
    umur := 21
    var tinggi float64 = 165.5

    fmt.Println("Nama:", nama)
    fmt.Println("Umur:", umur)
    fmt.Println("Tinggi:", tinggi)
}
