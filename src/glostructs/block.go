package glostructs

type Block struct {
	Hash [64]byte //Hash id of the block, formed by hash(n-1)+chain_hash+listofentries
	Entries []Transaction //Listo of entries that will compose this block
}

func (b *Block) New() {

}

//add transaction to block if block is accepted
func (b *Block) AddTransaction(t Transaction) {

	if(t.isValid()) {
		b.Entries = append(b.Entries, t)
	}

}