package main

/*
Goはポインタを扱います。 ポインタは値のメモリアドレスを指します。
変数 T のポインタは、 *T 型で、ゼロ値は nil です。
 */

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}