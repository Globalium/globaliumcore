package glostructs

type Wallet struct{
	PublicKey [64]byte 	//Public key and account id of the wallet
	Amount uint64		//Amount of money holded by the account, unsigned int of 8 digits as minimum
}

type Entry struct{
	Origin Wallet
	Destiny Wallet
	Amount uint
}

type Block struct {
	Hash [64]byte //Hash id of the block, formed by hash(n-1)+chain_hash+listofentries
	Entries []Entry //Listo of entries that will compose this block
}

type Blockchain struct{
	Blocks []Block	//Blocks that composes the blockchain Globalium
}