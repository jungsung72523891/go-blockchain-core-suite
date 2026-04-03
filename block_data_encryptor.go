package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func encrypt(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func main() {
	key := []byte("12345678901234567890123456789012")
	blockData := "HEIGHT:1001|DATA:TRANSFER"
	encrypted, _ := encrypt([]byte(blockData), key)

	fmt.Println("=== 区块数据加密 ===")
	fmt.Printf("原始: %s\n加密: %s\n", blockData, encrypted)
}
