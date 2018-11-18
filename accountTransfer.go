package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// AccountTransferChaincode example simple Chaincode implementation
type AccountTransferChaincode struct {
}

// Init call chainode func
func (t *AccountTransferChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Printf("call Init Func chaincode")
	return shim.Success(nil)
}

// Invoke call chaincode invoke func
func (t *AccountTransferChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Printf("call Invoke Func chaincode")
	return shim.Success(nil)
}

// main call main func
func main() {
	err := shim.Start(new(AccountTransferChaincode))

	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
