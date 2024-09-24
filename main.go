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

func (block *Block)generateBlock(Data string){
	block.TimeStamp = []byte(time.Now().Format(time.RFC3339))
	block.generateHash(block.TimeStamp)
	block.Data = []byte(Data)

}

func (block *Block) generateHash(TimeStamp []byte) []byte {
	h := sha256.New()
	h.Write(TimeStamp)
	hashBytes := h.Sum(nil)
	return hashBytes
}

func main() {
	block := new(Block)
	block.CurrentHash = block.generateHash())
	fmt.Printf("%x", block.CurrentHash)
}
