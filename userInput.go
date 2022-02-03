package main

import (
	"fmt"
)

func BitsToByte(bits float64) float64{
	bytes := bits/8
	return bytes
}

func main(){
	var sizeInBits float64
	print("Enter data size in bits: ")
	fmt.Scan(&sizeInBits)

	var sizeInBytes float64 = BitsToByte(sizeInBits)

	fmt.Println("Original Size(bits):", int(sizeInBits), "bits")
	fmt.Println("Converted Size(bytes):", int(sizeInBytes), "bytes")
}