package glostructs

type Block struct {
	Hash [64]byte //Hash id of the block, formed by hash(n-1)+chain_hash+listofentries
	Entries []Transaction //Listo of entries that will compose this block
}