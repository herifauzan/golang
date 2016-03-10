package main

import "fmt"


func main(){
  var i int
  var j int
  fmt.Scan(&i)
  fmt.Println(i)
  s := make([]int, 1)
  fmt.Println("slilce created")
  x := 1
    for x <= i {
        fmt.Scan(&j)
        fmt.Println(j)
        s = append(s, j);
        fmt.Println("slice append")
        x = x + 1
    }

  fmt.Println(s);
}