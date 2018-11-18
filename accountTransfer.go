package main

import (
	"fmt"
	"strconv"

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
	function, args := stub.GetFunctionAndParameters()

	if function == "openAccount" {
		return t.openAccount(stub, args)
	} else if function == "transfer" {
		return t.transfer(stub, args)
	} else if function == "inquire" {
		return t.inquire(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"transfer\" \"open\" \"query\"")
}

// openAccount open account
func (t *AccountTransferChaincode) openAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var user string // Entity
	var amount int  // Asset holdings
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	user = args[0]
	// string to int
	amount, err = strconv.Atoi(args[1])

	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	fmt.Printf("user = %s, Aval = %d \n", user, amount)

	// Write the state to the ledger
	err = stub.PutState(user, []byte(strconv.Itoa(amount)))
	if err != nil {
		return shim.Error(err.Error())
	}

	amtBytes, err := stub.GetState(user)
	jsonResp := "{\"Name\":\"" + user + "\",\"Amount\":\"" + string(amtBytes) + "\"}"

	fmt.Printf("Query Response:%s\n", jsonResp)

	return shim.Success(amtBytes)
}

// transfer transfer user's amount to b
func (t *AccountTransferChaincode) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var userA, userB string
	var amtA, amtB, amt int
	var err error

	userA = args[0]
	// args[1] parameter
	amt, _ = strconv.Atoi(args[1])
	userB = args[2]

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// userA
	amtABytes, err := stub.GetState(userA)

	if err != nil {
		return shim.Error("Failed to get state")
	}
	// convert
	amtA, _ = strconv.Atoi(string(amtABytes))

	// userB
	amtBBytes, err := stub.GetState(userB)

	if err != nil {
		return shim.Error("Failed to get state")
	}

	amtB, _ = strconv.Atoi(string(amtBBytes))

	amtA = amtA - amt
	amtB = amtB + amt

	fmt.Printf("user %s amount = %d, user %s amount = %d\n", userA, amtA, userB, amtB)

	// Write the state back to the ledger
	err = stub.PutState(userA, []byte(strconv.Itoa(amtA)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(userB, []byte(strconv.Itoa(amtB)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *AccountTransferChaincode) inquire(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var userA string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	userA = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(userA)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + userA + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + userA + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + userA + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	return shim.Success(Avalbytes)
}

// main call main func
func main() {
	err := shim.Start(new(AccountTransferChaincode))

	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
