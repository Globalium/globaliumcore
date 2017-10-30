package glostructs

import (
	"bufio"
	"math/big"
	"crypto/ecdsa"
	"os"
	"fmt"
	"crypto/rand"
)

//transaction format in plane trext
// r, s sign of ecdsa
// public for verify sign
type Transaction struct {
	origin  Direction
	destiny Direction
	amount  uint64
	r,s 	*big.Int
	pb		ecdsa.PublicKey
}

//constructor Transaction
//Comprueba en la blockchain si la direccion tiene suficientes monedas
func (t *Transaction)New(origin Direction, destiny Direction, amount uint64, pkPrivate string) {
	
	if VerifyAmountDirection(origin) {
		
		t.origin = origin
		t.destiny = destiny
		t.amount = amount

		t.Signed(pkPrivate)

		if t.VerifyTransaction() {
			fmt.Println("Transaction created.")
		} else {
			fmt.Println("Transaction NOT created, failed!")
		}

	} else {
		//transaction not valid, in our blockchain origin not enough coins
		fmt.Println("ErrorTransaction: Direction not enough coin in blockchain")
	}

}

//
func (t *Transaction)Signed(pk string) {
	
	PK := StringToPrivateKey(pk)
	
	r, s, err := ecdsa.Sign(rand.Reader, &PK, []byte(fmt.Sprintf("%v",t)))

	if err != nil {
		panic(err)
	}
	
	t.r = r
	t.s = s
	t.pb = PK.PublicKey
}

//return verify transaction
func (t *Transaction)VerifyTransaction() bool {
	return ecdsa.Verify(&t.pb,[]byte(fmt.Sprintf("%v",t)),t.r,t.s) && CreateIDAccount(t.pb.X, t.pb.Y) == t.origin.GetDirection()
}


//test new transaction
func TestNewTransaction() {

	var a,b Direction
	a.New()
	b.New()

	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Input PrivateKey: ")
	pKString, err := in.ReadString('\n')

	if err != nil {
		panic(err)
	}

	var auxT Transaction
	auxT.New(a,b,0, pKString)

	if auxT.VerifyTransaction() {
		fmt.Println("TestNewTransaction: Correct")
	} else {
		fmt.Println("TestNewTransaction: Incorrect")
	}
	

}



