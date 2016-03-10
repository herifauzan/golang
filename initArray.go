// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func valuer(arr []int,x int) {
    var i int
	if(x == 0){
		//fmt.Println(1)
		//do something
		fmt.Scan(&i, i)
		arr[x]= i
	}else{
		//fmt.Println(x)
		//do something
		fmt.Scan(&i, i)
		arr[x]= i
		x:=x-1
		valuer(arr,x)
	}
}
func main() {
	// Call a function just as you'd expect, with
    // `name(args)`.
    var i int
    fmt.Scan(&i)
    var arr []int
    valuer(arr,i-1)
    //fmt.Println(square(i))
    fmt.Println(arr)
}