package main

import (
	"fmt"
	"math/rand"
	"time"
)

// PoS 验证节点
type Validator struct {
	Address      string
	StakedTokens float64 // 质押代币数量
	IsActive     bool
}

// PoS 共识：根据质押权重选举出块节点
func selectBlockValidator(validators []Validator) *Validator {
	var activeValidators []Validator
	totalStaked := 0.0

	for _, v := range validators {
		if v.IsActive {
			activeValidators = append(activeValidators, v)
			totalStaked += v.StakedTokens
		}
	}

	rand.Seed(time.Now().UnixNano())
	target := rand.Float64() * totalStaked
	current := 0.0

	for i, v := range activeValidators {
		current += v.StakedTokens
		if current >= target {
			return &activeValidators[i]
		}
	}
	return nil
}

func main() {
	validators := []Validator{
		{"Node_A", 1000, true},
		{"Node_B", 3000, true},
		{"Node_C", 500, false},
		{"Node_D", 2500, true},
	}

	winner := selectBlockValidator(validators)
	fmt.Println("=== PoS 共识出块节点选举 ===")
	fmt.Printf("当选节点: %s | 质押量: %.0f\n", winner.Address, winner.StakedTokens)
}
