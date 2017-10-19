package glostructs

type Wallet struct{
	PublicKey [64]byte 	//Public key and account id of the wallet
	Amount uint64		//Amount of money holded by the account, unsigned int of 8 digits as minimum
}