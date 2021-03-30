package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/qlboob/bsngo/bsngo"
)

func main() {
	err := shim.Start(new(bsngo.BsnGo))
	if err != nil {
		fmt.Printf("Error starting BsnChainCode: %s", err)
	}
}
