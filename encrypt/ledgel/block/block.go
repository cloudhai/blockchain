package block

import (
	"blockchain/entity"
	"blockchain/encrypt/ed25519"
)

func NewBlock(txs []*entity.Transaction,key *ed25519.EdKeys) entity.Block{
	var block entity.Block
	block.Trans = txs

}


func GetCurBlockHeight() int{
	r
}
