package main

/*
var ステートメントは変数( variable )を宣言します。 関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言できます。
var ステートメントはパッケージ、または、関数で利用できます。例のコードで示します。
 */

import "fmt"

func main() {
	var i, j int = 1,2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}


/*
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

 */
