package main

import(
	glo "glostructs"
	"fmt"
	"crypto/rand"
    "crypto/rsa"
    "crypto/x509"
	"encoding/pem"
	"errors"
)
var version string = "2017.0.1.0"

func PrivateKeyToEncryptedPEM(bits int, pwd string) ([]byte, error) {
	if pwd == ""{
		err := errors.New("Cannot use an empty password.")
		return nil,err
	}
    // Generate the key of length bits
    key, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, err
	}
	fmt.Println(key)

    // Convert it to pem
    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	

	fmt.Println("OK--------------------")
    // Encrypt the pem
	block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(pwd), x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	der, err := x509.DecryptPEMBlock(block, []byte(pwd))
	if err != nil {
		panic("Error decrypt pem block")
	}

	oldkey, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil{
		panic(err)
	}
	fmt.Println(oldkey)

    return pem.EncodeToMemory(block), nil
}

func main(){
	var wallet glo.Wallet
	fmt.Println("GLOBALIUM 0.1.0")
	fmt.Println("Example of empty wallet: ")
	fmt.Println(wallet)
	clave := "123456789"
	pem,err := PrivateKeyToEncryptedPEM(1024,clave)
	if err != nil{
		fmt.Println(err)
		panic("Error generating rsa private key.")
	}
	fmt.Println(pem)
}