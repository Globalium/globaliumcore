package glostructs

import (
	"crypto/ecdsa"
	"fmt"
)

//transaction format in plane trext
type Transaction struct {
	Origin  Direction
	Destiny Direction
	Amount  uint64
}

//format transactionEncripted
type TransactionEncripted struct {
	publicKey ecdsa.PublicKey //public key
	textEncripted string //text encripted
}


//new transaction constructor
func (t *Transaction) New(dest Direction, amount uint64) {
	
	if amount <= t.Origin.Amount {
		
	} else {
		fmt.Println("ErrorTransaction: id " + t.Origin.IdAccount + "not enough globaliums.")
	}

}

// comprueba que una transacción es válida según la blockchain local
func (t *Transaction) isValid() (bool) {
	return true
}

//encripta la transaccion con la privatekey pasada por parametro
func (t *Transaction) encripted(pk string) TransactionEncripted {
	var aux TransactionEncripted
	return aux
}

// constructor el cual encripta la transacción a realizar
// para codificar la informacion
// string es la clave privada que se da cuando se crea la cuenta
func (te *TransactionEncripted) New(pk string, t Transaction) {

}

//desencripta la transaccion a partir de la key publica
func (te *TransactionEncripted) desencripted(pk ecdsa.PublicKey) (Transaction, bool) {
	var aux Transaction
	return aux, true
}
