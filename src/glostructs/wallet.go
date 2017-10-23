package glostructs

import (
	"golang.org/x/crypto/ripemd160"
	"crypto/sha256"
	"encoding/hex"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

type WalletECDSA struct {
	idAccount     string
	Amount uint64 //amount of wallet
}

func (w *WalletECDSA) New() {

	priAndPub, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	//validate generate keys of account
	if err != nil {
		panic(err)
	}

	h1 := sha256.New()
	h2 := sha256.New()
	
	h1.Write([]byte(priAndPub.X.String()))
	h2.Write([]byte(priAndPub.Y.String()))

	var strAux string = hex.EncodeToString(h1.Sum(nil)) + hex.EncodeToString(h2.Sum(nil))
	r := ripemd160.New()
	r.Write([]byte(strAux))

	fmt.Println(hex.EncodeToString(r.Sum(nil)))
	

	//imprimimos la privateKey
	fmt.Println("Your private Key:")
	fmt.Println(hex.EncodeToString(priAndPub.D.Bytes()))
	
	//guardamos la public key
	w.Save(priAndPub.PublicKey, w.idAccount)
}

func (w *WalletECDSA) Save(k ecdsa.PublicKey, id string) {

	f, err := os.Create("extra/accounts.txt")

	if err != nil {
		panic(err)
	}
	
	f.WriteString(id)
	f.WriteString(k.X.String())
	f.WriteString(k.Y.String())

	defer f.Close()
}

func TestNewWallet() {

	var wallet WalletECDSA
	
	wallet.New()

	f, err := os.Open("extra/account1PrivateKey.txt")
	

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Println("\nEsta firmado por el cliente de la wallet")

}
