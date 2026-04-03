package main

import "fmt"

const initialReward = 50.0 // 初始奖励
const halvingInterval = 210000 // 减半高度

// 计算当前高度区块奖励
func calculateBlockReward(height int) float64 {
	reward := initialReward
	halvings := height / halvingInterval
	for i := 0; i < halvings; i++ {
		reward /= 2
	}
	return reward
}

func main() {
	heights := []int{0, 210000, 420000, 630000, 840000}
	fmt.Println("=== 区块链区块奖励减半算法 ===")
	for _, h := range heights {
		fmt.Printf("区块高度 %d → 奖励: %.6f\n", h, calculateBlockReward(h))
	}
}
