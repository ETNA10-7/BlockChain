package types

import "time"

type Block struct {
	Hash      []byte //Current Hash Block
	Data      []byte //Previous Hash Block
	PrevHash  []byte
	TimeStamp time.Time //Creation of Block
	Index     int       // Block Index Number
	Nonce     int
}
