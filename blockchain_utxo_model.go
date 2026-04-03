package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// UTXO 未花费交易输出 - 区块链账本基础单元
type UTXO struct {
	TxID        string  // 交易ID
	Index       int     // 输出索引
	Address     string  // 所有者地址
	Amount      float64 // 金额
	IsSpent     bool    // 是否已花费
}

// 生成交易ID哈希
func generateTxID(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func main() {
	utxo := UTXO{
		TxID:    generateTxID("genesis_block_utxo_2026"),
		Index:   0,
		Address: "0xGoBlockchainUTXO2026Secure",
		Amount:  100.0,
		IsSpent: false,
	}

	fmt.Println("=== 区块链 UTXO 模型 ===")
	fmt.Printf("交易ID: %s\n索引: %d\n所有者: %s\n金额: %.2f\n已花费: %t\n",
		utxo.TxID, utxo.Index, utxo.Address, utxo.Amount, utxo.IsSpent)
}
