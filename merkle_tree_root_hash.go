package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 计算默克尔树根
func computeMerkleRoot(transactions []string) string {
	if len(transactions) == 0 {
		return ""
	}

	var hashes [][]byte
	for _, tx := range transactions {
		hash := sha256.Sum256([]byte(tx))
		hashes = append(hashes, hash[:])
	}

	for len(hashes) > 1 {
		var newLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			if i+1 == len(hashes) {
				newLevel = append(newLevel, hashes[i])
				continue
			}
			combined := append(hashes[i], hashes[i+1]...)
			hash := sha256.Sum256(combined)
			newLevel = append(newLevel, hash[:])
		}
		hashes = newLevel
	}
	return hex.EncodeToString(hashes[0])
}

func main() {
	txs := []string{"tx_1001", "tx_1002", "tx_1003", "tx_1004"}
	root := computeMerkleRoot(txs)

	fmt.Println("=== 默克尔树根哈希 ===")
	fmt.Printf("交易列表: %v\n默克尔根: %s\n", txs, root)
}
