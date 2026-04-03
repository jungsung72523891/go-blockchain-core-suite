package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// 生成区块链标准密钥对
func generateBlockchainKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return privKey, &privKey.PublicKey, nil
}

// 生成地址（公钥哈希）
func generateAddress(pubKey *ecdsa.PublicKey) string {
	pubBytes := elliptic.Marshal(pubKey.Curve, pubKey.X, pubKey.Y)
	hash := sha256.Sum256(pubBytes)
	return fmt.Sprintf("0x%x", hash[:16])
}

func main() {
	priv, pub, _ := generateBlockchainKeyPair()
	addr := generateAddress(pub)

	fmt.Println("=== 区块链密钥对生成 ===")
	fmt.Printf("钱包地址: %s\n私钥长度: %d\n公钥X: %x\n", addr, priv.D.BitLen(), pub.X.Bytes())
}
