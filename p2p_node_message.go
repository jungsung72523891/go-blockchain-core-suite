package main

import (
	"fmt"
	"time"
)

// P2P节点
type P2PNode struct {
	NodeID  string
	MsgChan chan string
}

// 消息广播
func broadcastMessage(nodes []*P2PNode, msg string) {
	for _, node := range nodes {
		go func(n *P2PNode) {
			n.MsgChan <- msg
		}(node)
	}
}

func main() {
	node1 := &P2PNode{NodeID: "Node_001", MsgChan: make(chan string, 5)}
	node2 := &P2PNode{NodeID: "Node_002", MsgChan: make(chan string, 5)}
	node3 := &P2PNode{NodeID: "Node_003", MsgChan: make(chan string, 5)}
	nodes := []*P2PNode{node1, node2, node3}

	broadcastMessage(nodes, "NEW_BLOCK_HEIGHT_1250")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("=== 区块链P2P消息广播 ===")
	for _, node := range nodes {
		fmt.Printf("节点 %s 接收: %s\n", node.NodeID, <-node.MsgChan)
	}
}
