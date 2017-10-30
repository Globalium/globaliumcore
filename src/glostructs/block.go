package glostructs

import (
	"encoding/hex"
	"encoding/json"
	"crypto/sha512"
	"fmt"
)
//account example for test
const (
	EXAMPLE_PK = "MTQ3ODYwOTUxNzE5NDMxOTc0MzQ0NTcwMjMyNTQ4ODg0OTkwMDMxMjQwOTcyNDE0MTAyMTAwMzAzNTkyODUxNTYxOTQwMzEzNjgxMTA6NjU4Mzc4MDEyMzc3NjYxNTMwNzcxNDQwNTAzNjUxMTk5NzUwNzMwOTc2Njc0NzE5NDE3OTg3MDYyMjE0ODk2ODA5ODk3NDI2NTg3Mjg6MTM2MzkyOTUzODU3NzU4NTYxMzUzNDUxMDQwNjQ2MTIyMjgwOTM2OTkzNDI3NTAzMTI4MzAzOTMyMTE1MjEzOTU2MzQ4OTExNjgyNzA="
)

//block of blockchain
type Block struct {
	Hash string //Hash id of the block, formed by hash(n-1)+chain_hash+listofentries
	Entries []Transaction //Listo of entries that will compose this block
}

func (b *Block) New() {

}

//add transaction to block if block is accepted
func (b *Block) AddTransaction(t Transaction) {
	b.Entries = append(b.Entries, t)
}

func (b *Block) Save() {
	
	h1 := sha512.New()
	
	a, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}

	//doble sha
	h1.Write(a)
	h1.Write([]byte(h1.Sum(nil)))

	b.Hash = hex.EncodeToString(h1.Sum(nil))

}

func TestNewBlock() {

	var a,b Direction
	var bl Block

	b.New()

	fmt.Printf("Create a block with 10 transactions...")
	for x := 0; x <= 10; x++ {
		var auxTransaction Transaction
		auxTransaction.New(a,b,1,EXAMPLE_PK)
		bl.AddTransaction(auxTransaction)
	}

	bl.Save()
	fmt.Println(bl.Hash)

	fmt.Println("TestNewBlock: Succed!")
}