package main

import "fmt"

func main() {
//variable a adalah string bernilai "initial"
    var a string = "initial"
    fmt.Println(a)
//variable b dan c adalah integer bernilai 1 dan 2
    var b, c int = 1, 2
//dua nilai berbeda pada satu variable selalu dipisahkan spasi
    fmt.Println(b,c)
//variable d bernilai true
    var d = true
    fmt.Println(d)
	
//variable e adalah integer bernilai default(default integer = 0)
    var e int
    fmt.Println(e)
//cara pintas untuk mendeklarasikan variable
//f adalah string bernilai short
    f := "short"
    fmt.Println(f)

    g:= 700
    fmt.Println(g)
}
