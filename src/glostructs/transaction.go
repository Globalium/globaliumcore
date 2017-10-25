package glostructs

type Transaction struct {
	Origin  Direction
	Destiny Direction
	Amount  uint64
}
