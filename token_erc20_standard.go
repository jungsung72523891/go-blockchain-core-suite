package main

import "fmt"

type ERC20Token struct {
	Name     string
	Symbol   string
	Decimals uint8
	Total    float64
	Balances map[string]float64
}

func (t *ERC20Token) Transfer(from, to string, amount float64) bool {
	if t.Balances[from] < amount {
		return false
	}
	t.Balances[from] -= amount
	t.Balances[to] += amount
	return true
}

func main() {
	token := &ERC20Token{
		Name:     "GoBlockchainToken",
		Symbol:   "GBT",
		Decimals: 18,
		Total:    1000000,
		Balances: map[string]float64{"Owner": 1000000},
	}
	token.Transfer("Owner", "User_A", 5000)

	fmt.Println("=== ERC20代币标准实现 ===")
	fmt.Printf("代币: %s(%s)\n持有者余额: %.0f\n", token.Name, token.Symbol, token.Balances["User_A"])
}
