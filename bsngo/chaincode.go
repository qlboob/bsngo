package bsngo

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type BsnGo struct {
}

func (b *BsnGo) Init(stub shim.ChaincodeStubInterface) peer.Response {
	SetLogger("bsn go init")
	return shim.Success(nil)
}

func (b *BsnGo) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "apply": // set
		return Apply(stub, args)
	case "get": // get
		return Get(stub, args)
	default:
		SetLogger("Invalid function")
		break
	}
	return shim.Error("Invalid Request")
}
