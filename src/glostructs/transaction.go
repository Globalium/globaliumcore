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
}

//object transaction
type TransactionEncripted struct {
	tra Transaction
	r,s *big.Int
	pb ecdsa.PublicKey
}

//constructor Transaction
//Comprueba en la blockchain si la direccion tiene suficientes monedas
func (t *Transaction)New(origin Direction, destiny Direction, amount uint64, pkPrivate string) bool {
	
	if VerifyAmount(origin, amount) {

		t.origin = origin
		t.destiny = destiny
		t.amount = amount

		tr := t.Signed(pkPrivate)

		if tr.VerifyTransaction() {
			return true
		} else {
			return false
		}

	} else {
		return false
	}

}

//sign transaction with privatekey
func (t *Transaction)Signed(pk string) TransactionEncripted {
	
	var tr TransactionEncripted
	PK := StringToPrivateKey(pk)

	tr.tra.origin = t.origin
	tr.tra.destiny = t.destiny
	tr.tra.amount = t.amount
	
	r, s, err := ecdsa.Sign(rand.Reader, &PK, []byte(fmt.Sprintf("%v",tr.tra)))

	if err != nil {
		panic(err)
	}
	
	//copy transaction
	tr.r = r
	tr.s = s
	tr.pb = PK.PublicKey

	return tr
}

//return verify transaction
func (t *TransactionEncripted)VerifyTransaction() bool {
	return ecdsa.Verify(&t.pb,[]byte(fmt.Sprintf("%v",t.tra)),t.r,t.s) && (CreateIDAccount(t.pb.X, t.pb.Y) == t.tra.origin.id)
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

	if auxT.New(a,b,0, EXAMPLE_PK) {
		fmt.Println("TestNewTransaction: Correct")
	} else {
		fmt.Println("TestNewTransaction: Incorrect")
	}

}



