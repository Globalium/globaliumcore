package glostructs

import (
	"math/big"
	"crypto/ecdsa"
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
func (t *Transaction)New(origin Direction, destiny Direction, amount uint64, pkPrivate string) bool {
	
	if VerifyAmount(origin, amount) {
		
		t.origin = origin
		t.destiny = destiny
		t.amount = amount

		t.Signed(pkPrivate)

		if t.VerifyTransaction() {
			return true
		} else {
			return false
		}

	} else {
		return false
	}

}

//sign transaction with privatekey
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

//guardar la transaccion en el bloque
func (t *Transaction) Save() {
	//solo guardaríamos:
	// from
	// to
	// amount
}


//test new transaction
func TestNewTransaction() {

	var a,b Direction

	PK := StringToPrivateKey(EXAMPLE_PK)

	a.Charge(CreateIDAccount(PK.PublicKey.X, PK.PublicKey.Y))
	b.New()

	var auxT Transaction
	auxT.New(a,b,0, EXAMPLE_PK)

	if auxT.VerifyTransaction() {
		fmt.Println("TestNewTransaction: Correct")
	} else {
		fmt.Println("TestNewTransaction: Incorrect")
	}

}



