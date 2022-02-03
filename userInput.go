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
	gigaBytes := BitsToMegaByte(bits)/1024
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
	var sizeInKiloBytes float64 = BitsToKiloByte(sizeInBits)
	var sizeInMegaBytes float64 = BitsToMegaByte(sizeInBits)
	var sizeInGigaBytes float64 = BitsToGigaByte(sizeInBits)
	var sizeInTeraBytes float64 = BitsToTeraByte(sizeInBits)

	fmt.Println("Original Size(bits):", int(sizeInBits), "bits")
	fmt.Println("Converted Size(bytes):", int(sizeInBytes), "bytes")
	fmt.Println("Converted Size(kilobytes):", int(sizeInKiloBytes), "kilobytes")
	fmt.Println("Converted Size(megabytes):", int(sizeInMegaBytes), "megabytes")
	fmt.Println("Converted Size(gigabytes):", int(sizeInGigaBytes), "gigabytes")
	fmt.Println("Converted Size(terabytes):", int(sizeInTeraBytes), "terabytes")
}