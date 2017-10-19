package main

import(
	glo "glostructs"
	"fmt"
)
var version string = "v2017.0.1.0"

func main(){
	var wallet glo.Wallet
	fmt.Println("GLOBALIUM "+version)
	fmt.Println("Example of empty wallet: ")
	fmt.Println(wallet)

}