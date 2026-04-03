package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func signTransaction(priv *ecdsa.PrivateKey, data string) (r, s *big.Int, err error) {
	hash := sha256.Sum256([]byte(data))
	return ecdsa.Sign(rand.Reader, priv, hash[:])
}

func verifyTransaction(pub *ecdsa.PublicKey, data string, r, s *big.Int) bool {
	hash := sha256.Sum256([]byte(data))
	return ecdsa.Verify(pub, hash[:], r, s)
}

func main() {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txData := "FROM:A TO:B AMOUNT:10"
	r, s, _ := signTransaction(priv, txData)
	valid := verifyTransaction(&priv.PublicKey, txData, r, s)

	fmt.Println("=== 交易签名与验签 ===")
	fmt.Printf("交易内容: %s\n验签结果: %t\n", txData, valid)
}
