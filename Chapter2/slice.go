package main

import "fmt"

func main() {
    // define a array
    var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Create a array slice based on array above
    var mySlice []int = myArray[:5]

    fmt.Println("Elements of myArray: ")

    for _, v := range myArray {
        fmt.Print(v, " ")
    }

    fmt.Println("\nElements of mySlice: ")

    for _, v := range mySlice {
        fmt.Print(v, " ")
    }

    fmt.Println()

}
