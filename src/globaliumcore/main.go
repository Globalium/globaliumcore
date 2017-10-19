package main

import(
	glo "glostructs"
	"fmt"
	"bufio"
	"os"
)
var version string = "v2017.0.1.0"

func menu() (string, error){
	fmt.Println("What do you want to do?")
	fmt.Println("\t1.- Test new wallet")
	in := bufio.NewReader(os.Stdin)
	opcion, err := in.ReadString('\n')
	if err != nil {
		return "",err
	}
	opcion = opcion[:len(opcion)-1]
	return opcion,nil
}

func main(){
	fmt.Println("GLOBALIUM "+version+"\n")
	opcion, err := menu()
	if err != nil {
		panic(err)
	}
	switch opcion{
	case "1":
		fmt.Println("Testing: Create new wallet...")
		glo.TestNewWallet()
	default:
		fmt.Println("Option don't supported.")
	}
}