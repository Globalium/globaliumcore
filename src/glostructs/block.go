package glostructs

import (
	"bufio"
	"os"
	"fmt"
)

type Block struct {
	Hash [64]byte //Hash id of the block, formed by hash(n-1)+chain_hash+listofentries
	Entries []Transaction //Listo of entries that will compose this block
}

func (b *Block) New() {

}

//add transaction to block if block is accepted
func (b *Block) AddTransaction(t Transaction) {
	b.Entries = append(b.Entries, t)
}

func (b *Block) Save() {

}

func TestNewBlock() {

	var a,b Direction
	var bl Block
	
	a.New()
	b.New()

	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Input PrivateKey: ")
	pKString, err := in.ReadString('\n')

	if err != nil {
		panic(err)
	}

	fmt.Printf("Create a block with 10 transactions...")
	for x := 0; x < 10; x++ {
		var auxTransaction Transaction
		auxTransaction.New(a,b,1,pKString)
		bl.AddTransaction(auxTransaction)
	}

	bl.Save()

	fmt.Println("TestNewBlock: Succed!")
}