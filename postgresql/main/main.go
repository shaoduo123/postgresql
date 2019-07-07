package main

import (
	"fmt"
	"go_dev/postgresql/pseudo"
)

func main() {
	var inputInt int32 = 123456789
	outputInt := pseudo.PseudoEncrypt(inputInt)
	selfInverse := pseudo.PseudoEncrypt(pseudo.PseudoEncrypt(inputInt))
	fmt.Printf("input is int %d and pseudoEncrypt output is %d, self-inverse is: pseudo_encrypt(pseudo_encrypt(X)) =%d\n", inputInt, outputInt, selfInverse)
	fmt.Printf("---------------\n")
	var inputUint64 uint64 = 12345678912345
	outputInt64 := pseudo.PseudoEncrypt64(inputUint64, 12) //figure out the length of the number
	fmt.Printf("inputUint64 is int %d and pseudoEncrypt outputInt64 is %d  self-inverse is: pseudo_encrypt(pseudo_encrypt(X)) =%d\n", inputUint64, outputInt64, selfInverse64)
}
