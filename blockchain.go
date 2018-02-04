package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "time"
)

var Blockchain []Block

type Block struct {
    Index      int
    Hash       string
    PrevHash   string
    Version      int
    Timestamp  string
    Zone       string
    Data       []Slot
}

type Slot struct {
    Hash       string
    Index      int
    Data       string
    Script     string
    Owner      string
    Sign       string
    FeeOutput  Output
    Input      []Input
    Output     []Output
}

type Input struct {
    SlotHash   string
    Sign       string
}

type Output struct {
    Value   int
    To      string
}

func generateBlock(oldBlock Block, zone string) (Block, error) {
    var newBlock Block

    t := time.Now()

    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = t.String()
    newBlock.Zone = zone
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Hash = calculateHash(newBlock)

    return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
    if oldBlock.Index+1 != newBlock.Index {
        return false
    }

    if oldBlock.Hash != newBlock.PrevHash {
        return false
    }

    if calculateHash(newBlock) != newBlock.Hash {
        return false
    }

    return true
}

func replaceChain(newBlocks []Block) {
    if len(newBlocks) > len(Blockchain) {
        Blockchain = newBlocks
    }
}

func calculateHash(block Block) string {
    record := string(block.Index) + block.Timestamp + string(block.Zone) + block.PrevHash

    for index,slot := range block.Data{
        record += slot.Hash;
    }

    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}
