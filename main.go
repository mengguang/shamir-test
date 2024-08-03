package main

import (
	"fmt"
	"log"

	"github.com/mengguang/shamir-test/s4"
)

func main() {
	input := "This is the original text."
	fmt.Printf("Original Text: %s\n", input)
	shares, err := s4.DistributeBytes([]byte(input), 3, 2)
	if err != nil {
		log.Fatal(err)
	}
	for n, share := range shares {
		fmt.Printf("share %d: %02X\n", n, share)
	}

	result, err := s4.RecoverBytes(shares[1:3])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recovery: %s\n", string(result))

	result, err = s4.RecoverBytes(shares[0:2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recovery: %s\n", string(result))

	shares, err = s4.DistributeBytesAES([]byte(input), 3, 2)
	if err != nil {
		log.Fatal(err)
	}
	for n, share := range shares {
		fmt.Printf("share %d: %02X\n", n, share)
	}

	result, err = s4.RecoverBytesAES(shares[1:3])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recovery: %s\n", string(result))

	result, err = s4.RecoverBytesAES(shares[0:2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recovery: %s\n", string(result))

}
