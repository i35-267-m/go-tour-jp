package main

/*
定数( constant )は、 const キーワードを使って変数と同じように宣言します。
定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使えます。
なお、定数は := を使って宣言できません。
 */

import "fmt"

const world  = "Hello World"

func main() {
	fmt.Println("こんにちわ！" + world)
}
