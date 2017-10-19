package glotesting

import(
	"fmt"
	"testing"
	"encoding/base64"
	"encoding/json"
	glo "glostructs"	
)
func TestNewWallet(t *testing.T){
	var wallet glo.Wallet
	passphrase := "unafrasedeprueba"
	block, err := wallet.New(passphrase,4096)
	if err != nil{
		panic(err)
	}

	blockjson,err := json.Marshal(block)
	if err != nil{
		panic(err)
	}

	pembase64 := base64.StdEncoding.EncodeToString(blockjson)

	fmt.Println("\nYour wallet is:")
	fmt.Println(wallet)

	fmt.Println("\nYour private key is:")
	fmt.Println(pembase64)
}