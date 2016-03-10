// You can edit this code!
// Click here and start typing.
package main

import "fmt"
func square(a int) int {

    // Go requires explicit returns, i.e. it won't
    // automatically return the value of the last
    // expression.
    return a*a
}
func loop(x int) {
	var i int
	if(x == 1){
			//fmt.Println(1)
			//do something
			fmt.Scan(&i, i)
			fmt.Println(square(i))
			//loping untuk array fmt.Println(sumArray(i))
		}else{
			//fmt.Println(x)
			//do something
			fmt.Scan(&i, i)
			fmt.Println(square(i))
			//looping untuk array fmt.Println(sumArray(i))
			x:=x-1
			loop(x)
		}
}
func sumArray(x int) int{
	var sum int = 0
	var i int
	if(x == 1){
			//fmt.Println(1)
			//do something
			fmt.Scan(&i, i)
			if(i > 0){
				sum:= sum + square(i)
			}
			// fmt.Println(square(i))
		}else{
			//fmt.Println(x)
			//do something
			fmt.Scan(&i, i)
			if(i > 0){
				sum := square(i)
			}
			//fmt.Println(square(i))
			x:=x-1
			loop(x)
		}
		return sum
}
func main() {
	// Call a function just as you'd expect, with
    // `name(args)`.
    var i int
    fmt.Scan(&i)
    loop(i)
    //fmt.Println(square(i))

}