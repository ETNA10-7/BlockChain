package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/ETNA10-7/blockchain/types"
)

type BlockChain struct {
	Blocks []*types.Block
}

// Generates Hash for the Block
func MakeHash(block *types.Block) []byte {
	info := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	return hash[:]
}

// Create Block Function
func CreateBlock(data string, index int, prevHash []byte) *types.Block {
	block := &types.Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Index: index, Nonce: 0}
	block.Hash = MakeHash(block)
	return block

}



// Add Block to Chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Index+1, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

//First Block Creation
func Genesis() *types.Block {
	return CreateBlock("Genesis", 0, []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*types.Block{Genesis()}}
}

// Validate Blocks Hash
func (chain *BlockChain) ValidateChain() bool {
	for i := 1; i < len(chain.Blocks); i++ {
		prevBlock := chain.Blocks[i-1]
		currentBlock := chain.Blocks[i]

		if !bytes.Equal(currentBlock.PrevHash, prevBlock.Hash) {
			fmt.Printf("Tampering detected! Block %d prevHash does not match Block %d hash.\n", i, i-1)
			return false
		}

		if !bytes.Equal(MakeHash(currentBlock), currentBlock.Hash) {
			fmt.Printf("Tampering detected! Block %d hash does not match calculated hash.\n", i)
			return false
		}
	}
	return true
}



