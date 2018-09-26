package main

/*
struct (構造体)は、フィールド( field )の集まりです。
structのフィールドは、ドット( . )を用いてアクセスします。
 */

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
}
