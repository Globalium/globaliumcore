package glostructs

import (
	"fmt"
)

type Blockchain struct{
	Blocks []Block	//Blocks that composes the blockchain Globalium
}

//Delete Blockchain if exists
//create genesis block
//create account with MAX_SUPPLY
//Run node Blockchain
func (b *Blockchain)New() {

}

//Sync node with Network of Globalium
//Confirm Transaction or Create transaction
func (b *Blockchain)Run() {

}

//Sync node with other nodes
func (b *Blockchain) Sync() {

}

// create account
// print PrivateKey
func (b *Blockchain) CreateAddress() {

	var a Direction
	
	id, pk := a.New()

	fmt.Println("Your address is: " + id)
	fmt.Println("Your PrivateKey is: " + pk)

}

//create transaction
func (b *Blockchain) CreateTransaction(from string, to string, amount uint64, pk string) bool {

	var t Transaction
	var d, h Direction

	d.Charge(from)
	h.Charge(to)

	return t.New(d, h, amount, pk)

}

//verify transaction
func (b *Blockchain) VerifyTransaction(t Transaction) bool {

	var a Direction

	a.Charge(t.origin.id)

	if VerifyAmount(a,t.amount) {
		return true
	} else {
		return false
	}
}
