package main

import (
	"fmt"
)

func BitsToByte(bits float64) float64{
	bytes := bits/8
	return bytes
}

func BitsToKiloByte(bits float64) float64{
	kiloBytes := BitsToByte(bits)/1024
	return kiloBytes
}

func BitsToMegaByte(bits float64) float64{
	megaBytes := BitsToByte(bits)/1024
	return megaBytes
}

func BitsToGigaByte(bits float64) float64{
	gigaBytes := BitsToMByte(bits)/1024
	return gigaBytes
}

func BitsToTeraByte(bits float64) float64{
	teraBytes := BitsToGigaByte(bits)/1024
	return teraBytes
}

func main(){
	var sizeInBits float64
	print("Enter data size in bits: ")
	fmt.Scan(&sizeInBits)

	var sizeInBytes float64 = BitsToByte(sizeInBits)

	fmt.Println("Original Size(bits):", int(sizeInBits), "bits")
	fmt.Println("Converted Size(bytes):", int(sizeInBytes), "bytes")
}