package main

import "fmt"
func Recursive(slice []int, number int) int {
    if number == 0 {
        return slice[0]
    }
    return slice[number] + Recursive(slice,number-1)
}
func Process() int{
  var i int
  var j int
  fmt.Scan(&i)
  fmt.Println(i)
  s := make([]int, 1)
  fmt.Println("slice created")
  n:=0
  x := 0
    for x < i {
        fmt.Scan(&j)
        fmt.Println(j)
        if(j > 0){
          s = append(s, j*j);
          fmt.Println("slice append")
          n++
        }else{
          fmt.Println("zero or negatif not append")
        }
        
        x = x + 1
    }
  //sum(s);
  fmt.Println(s);
  sum:= Recursive(s,n)
  return sum
}
func Loop(i int){
  if(i==1){
    fmt.Println(Process())
  }else{
    fmt.Println(Process())
    i--
    Loop(i)
  }
}
func main(){
  var x int
  fmt.Scan(&x)
  Loop(x-1)
}