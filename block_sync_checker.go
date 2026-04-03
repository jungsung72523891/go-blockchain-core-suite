package main

import "fmt"

// 检查节点是否同步
func isNodeSynced(localHeight int, networkHeight int) bool {
	return localHeight >= networkHeight
}

func main() {
	local := 15200
	network := 15203

	fmt.Println("=== 区块同步状态检查 ===")
	fmt.Printf("本地高度: %d\n网络高度: %d\n已同步: %t\n", local, network, isNodeSynced(local, network))
}
