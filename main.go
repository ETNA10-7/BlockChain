package main

import (
	"fmt"

	blockchain "github.com/ETNA10-7/blockchain/block"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {

		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Index : %d\n", block.Index)



	}

	fmt.Println("\nBefore Tampering - Blockchain Validation:", chain.ValidateChain())

	chain.Blocks[2].Data = []byte("Alice sends 10 BTC to Bob") // Modify a transaction

	fmt.Println("\nAfter Tampering - Blockchain Validation:", chain.ValidateChain())


}
