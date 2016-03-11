package main

import "fmt"
func Recursive(slice []int, number int) int {

    if number == 0 {

        return slice[0]
    }

    return slice[number] + Recursive(slice,number-1)
}

func main() {
    s := make([]int, 3)

    // We can set and get just like with arrays.
    s[0] = 5
    s[1] = 76
    s[2] = 1
    
    answer := Recursive(s,2)
    fmt.Printf("Recursive: %d\n", answer)
} 