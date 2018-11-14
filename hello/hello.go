/*
 これはGo言語テスト用プログラム
  package main
  runしたい場合にはmainパッケージを定義して func mainを定義するのが必須
*/
package main

//import fmt -> fmtという名前のpackageを読み込んでいる。ダブルクォーテーション必須
import "fmt"

//func -> function の意味となるが、functionって記載するとエラーになります
func main() {
	// Println は JavaのSystem.out.printlnに酷似している。ダブルクォーテーションではなくシングルにするとエラー
	fmt.Println("Hello, World!")
}
