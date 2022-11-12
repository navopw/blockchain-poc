package main

import (
	"fmt"
	"github.com/navopw/blockchain-poc/block"
)

func main() {
	// create blockchain with mining difficulty of 2
	blockchain := block.CreateBlockchain(2)

	// transactions
	blockchain.AddBlock("Stan", "Bob", 5)
	blockchain.AddBlock("Bob", "Stan", 2)

	// check if the blockchain is valid
	fmt.Println(blockchain.IsValid())
}
