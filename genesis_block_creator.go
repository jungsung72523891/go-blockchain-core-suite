package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// 区块结构
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// 计算区块哈希
func calculateHash(b Block) string {
	record := fmt.Sprintf("%d%s%s%s", b.Index, b.Timestamp, b.Data, b.PrevHash)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func main() {
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "GO_BLOCKCHAIN_GENESIS_BLOCK_2026",
		PrevHash:  "0",
	}
	genesis.Hash = calculateHash(genesis)

	fmt.Println("=== 创世区块生成 ===")
	fmt.Printf("索引: %d\n哈希: %s\n数据: %s\n", genesis.Index, genesis.Hash, genesis.Data)
}
