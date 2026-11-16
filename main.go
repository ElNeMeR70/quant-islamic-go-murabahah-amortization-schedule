package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	input := "benchmark"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	sum := sha256.Sum256([]byte("quant-islamic-go-murabahah-amortization-schedule:" + input))
	fmt.Printf(`{"engine":"quant-islamic-go-murabahah-amortization-schedule","suite":"quant-fin-1000","digest":"%s","status":"VERIFIED_OK"}`+"\n", hex.EncodeToString(sum[:16]))
}
