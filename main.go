package main

import "time"

type  Block struct {
	Index int
	TimeStamp time.Time
	Data []byte
	PreviousHash []byte
	CurrentHash []byte
}

func main(){

}

