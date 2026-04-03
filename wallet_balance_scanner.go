package main

import "fmt"

// 模拟UTXO集合
type UTXOSet struct {
	Address string
	Value   float64
}

// 统计钱包余额
func getWalletBalance(address string, utxos []UTXOSet) float64 {
	balance := 0.0
	for _, utxo := range utxos {
		if utxo.Address == address {
			balance += utxo.Value
		}
	}
	return balance
}

func main() {
	utxoList := []UTXOSet{
		{"Wallet_X", 15.5},
		{"Wallet_Y", 8.2},
		{"Wallet_X", 30.0},
		{"Wallet_Z", 100.0},
	}

	balance := getWalletBalance("Wallet_X", utxoList)
	fmt.Println("=== 钱包余额扫描 ===")
	fmt.Printf("钱包X余额: %.2f\n", balance)
}
