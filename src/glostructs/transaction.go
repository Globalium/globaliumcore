package glostructs

type Transaction struct {
	Origin  WalletECDSA
	Destiny WalletECDSA
	Amount  uint
}
