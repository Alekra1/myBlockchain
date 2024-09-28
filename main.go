package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/lukesampson/figlet/figletlib"
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

func (block *Block) generateBlock(Data string) {
	block.TimeStamp = []byte(time.Now().Format(time.RFC3339))
	block.generateHash(block.TimeStamp)
	block.Data = []byte(Data)

}

func (block *Block) generateHash(TimeStamp []byte) []byte {
	hash := sha256.New()
	hash.Write(TimeStamp)
	hashBytes := hash.Sum(nil)
	return hashBytes
}

func Print(message string){
	fontsDirectory := "/home/alekra1/myBlockChain/fonts"
	f, err := figletlib.GetFontByName(fontsDirectory, "smslant")
	if err != nil {
		fmt.Println("Could not find that font!")
		return
	}
	figletlib.PrintMsg(message, f, 80, f.Settings(), "left")
}

func main() {
	block := new(Block)
	block.CurrentHash = block.generateHash([]byte(time.Now().Format(time.RFC3339)))
	Print("Here is your hash:")
	fmt.Printf("%x", block.CurrentHash)
}
