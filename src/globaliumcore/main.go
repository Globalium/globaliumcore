package main

import(
	glo "glostructs"
	"fmt"
)
var version string = "2017.0.1.0"

func main(){
	var wallet glo.Wallet
	fmt.Println("GLOBALIUM 0.1.0")
	fmt.Println("Example of empty wallet: ")
	fmt.Println(wallet)
}