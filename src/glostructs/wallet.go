package glostructs

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

type WalletECDSA struct {
	Pk     *ecdsa.PrivateKey
	Amount uint64 //amount of wallet
}

func (w *WalletECDSA) New() {

	priAndPub, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)

	//validate generate keys of account
	if err != nil {
		panic(err)
	}

	w.Pk = priAndPub
	w.Amount = 0
}

func TestNewWallet() {

	var wallet WalletECDSA
	wallet.New()
	f, err := os.Create("extra/pruebaAccount.txt")

	if err != nil {
		panic(err)
	}

	//f.WriteString()
	defer f.Close()

	fmt.Println("\nYour private key is:")
	fmt.Println(wallet.Pk.D.String())

}
