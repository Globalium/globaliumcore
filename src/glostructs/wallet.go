package glostructs

import (
	"encoding/hex"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

type WalletECDSA struct {
	Pk     ecdsa.PublicKey
	Amount uint64 //amount of wallet
}

func (w *WalletECDSA) New() {

	priAndPub, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	//validate generate keys of account
	if err != nil {
		panic(err)
	}

	w.Pk = priAndPub.PublicKey
	w.Amount = 0

	f, err := os.Create("extra/account1PrivateKey.txt")
	
	if err != nil {
		panic(err)
	}

	//pasamos la privateKey
	f.WriteString(hex.EncodeToString(priAndPub.D.Bytes()))
	
	//guardamos la public key
	w.Save()

	defer f.Close()
}

func (w *WalletECDSA) Save() {

	f, err := os.Create("extra/accounts.txt")

	if err != nil {
		panic(err)
	}

	f.WriteString(w.Pk.X.String() + w.Pk.Y.String())

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
