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
	block0 := new(Block)
	block1 := new(Block)
	fmt.Println(block0)
	fmt.Println(block1)
	blocks := []*Block{block0, block1}
	blockchain := Blockchain{blocks}
	//blockchain = append(*blockchain, &block0, &block1)
	fmt.Println(blockchain)
}
