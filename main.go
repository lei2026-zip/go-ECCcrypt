package main

import (
	"blockverfity/BitcoinAddress"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	////第一步、生成私钥 ECC，公钥，并进行hash计算
	curve := elliptic.P256()
	pri, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//ecdsa.GenerateKey(curve,rand.Reader)
	fmt.Println(pri)
	fmt.Println(len(pri))

	//拼接x和y，生成公钥
	pubKey := append(x.Bytes(), y.Bytes()...)
	fmt.Println(hex.EncodeToString(pubKey))
	//pubKey := []byte("ef22241075696721a11ea753fc26603e2d38da385cd56e3555153ed6be082eaa0792bd39a0daed48b7e4e879f6b1cc2c4ba16809ccb91ddc4f28a559f60f0a2f")
	//addr := []byte(" 15Gs4vFKMkNtZuAhefHMnmedCtor6ZnWuZ")
    addr := BitcoinAddress.NewBitcoinAddress(0x00,pubKey)
    fmt.Println(BitcoinAddress.CheckAddress(string(addr)))

}
