package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	TimeStamp    []byte
	Data         []byte
	PreviousHash []byte
	CurrentHash  []byte
}

type Blockchain struct {
	Blocks []*Block
}

func (block *Block) createBlock(Data string, PreviousHash []byte, PreviousIndex int) {
	block.Index = PreviousIndex + 1
	block.TimeStamp = []byte(time.Now().Format(time.RFC3339))
	block.CurrentHash = block.generateHash(block.TimeStamp)
	block.PreviousHash = PreviousHash
	block.Data = []byte(Data)

}

func (block *Block) generateHash(TimeStamp []byte) []byte {
	hash := sha256.New()
	hash.Write(TimeStamp)
	hashBytes := hash.Sum(nil)
	return hashBytes
}

func main() {
	Print("Hi!")
	var blockchain Blockchain
	for i := 0; i < 10; i++ {
		blockchain.Blocks = append(blockchain.Blocks, new(Block))
	}
	//blockchain = append(*blockchain, &block0, &block1)
	fmt.Println(blockchain)
}
