package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SkyfchainChaincode struct {
}

func main() {
	err := shim.Start(new(SkyfchainChaincode))
	if err != nil {
		fmt.Printf("Error starting Skyfchain chaincode: %s", err)
	}
}

func (t *SkyfchainChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	drones := initDrones()

	for _, drone := range drones {
		key, err := droneKey(stub, drone.Id)

		if err != nil {
			return shim.Error(err.Error())
		}
		droneBytes, err := json.Marshal(drone)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = stub.PutState(key, droneBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	routes := initRoutes()

	for _, route := range routes {
		key, err := routeKey(stub, route.Id)

		if err != nil {
			return shim.Error(err.Error())
		}
		routeBytes, err := json.Marshal(routes)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = stub.PutState(key, routeBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
}

func (t *SkyfchainChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "drones" {
		return t.drones(stub, args)
	}
	if function == "drone" {
		return t.drone(stub, args)
	}
	if function == "routes" {
		return t.routes(stub, args)
	}
	if function == "route" {
		return t.route(stub, args)
	}
	if function == "missions" {
		return t.missions(stub, args)
	}
	if function == "mission" {
		return t.mission(stub, args)
	}
	if function == "points" {
		return t.missions(stub, args)
	}
	if function == "point" {
		return t.mission(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation?")
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}
