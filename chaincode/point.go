package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type point struct {
	ObjectType string  `json:"docType"`
	Id         int64   `json:"id"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
}

func pointKey(stub shim.ChaincodeStubInterface, id int64) (string, error) {
	return stub.CreateCompositeKey("point", []string{string(id)})
}

func initPoints() []point {
	return []point{point{ObjectType: "point", Id: 11, Lat: 1.2583861960946845, Lng: 103.83418327546315},
		point{ObjectType: "point", Id: 12, Lat: 1.1175038553177987, Lng: 104.08365980834049},
		point{ObjectType: "point", Id: 13, Lat: 1.1098663583218364, Lng: 104.08511893596244},
		point{ObjectType: "point", Id: 14, Lat: 1.1025721182785055, Lng: 104.08803717345279},
		point{ObjectType: "point", Id: 15, Lat: 1.0962647220902284, Lng: 104.09299390376918},
		point{ObjectType: "point", Id: 16, Lat: 1.0904722074210613, Lng: 104.10181299895328},
		point{ObjectType: "point", Id: 17, Lat: 1.0878119367906187, Lng: 104.1136147186188},
		point{ObjectType: "point", Id: 18, Lat: 1.0896140703587003, Lng: 104.12751927880254},
		point{ObjectType: "point", Id: 19, Lat: 1.0934757359520424, Lng: 104.13721815794986},
		point{ObjectType: "point", Id: 20, Lat: 0.9190947424427305, Lng: 104.48672072142642},
		point{ObjectType: "point", Id: 21, Lat: 21.03526879821853, Lng: 105.78183575277865},
		point{ObjectType: "point", Id: 22, Lat: 21.402481071672483, Lng: 103.02095184418033},
		point{ObjectType: "point", Id: 31, Lat: 35.957529729150075, Lng: 129.41059513693392},
		point{ObjectType: "point", Id: 32, Lat: 33.605544414920466, Lng: 130.439817405665743},
		point{ObjectType: "point", Id: 41, Lat: 33.24848943576627, Lng: 131.7428456128174},
		point{ObjectType: "point", Id: 42, Lat: 33.56249487191811, Lng: 133.57320616337142},
		point{ObjectType: "point", Id: 51, Lat: 60.9717854408869, Lng: 76.5673948705728},
		point{ObjectType: "point", Id: 52, Lat: 61.739460834342516, Lng: 75.59443662258468},
		point{ObjectType: "point", Id: 61, Lat: 35.115308892606386, Lng: 126.78893699055266},
		point{ObjectType: "point", Id: 62, Lat: 33.523279771082116, Lng: 126.5469619535304},
		point{ObjectType: "point", Id: 71, Lat: 52.384802933061934, Lng: 4.929993587785475},
		point{ObjectType: "point", Id: 72, Lat: 53.59804425094406, Lng: 6.700785746464817}}
}

func (t *SkyfchainChaincode) points(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) > 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	queryString := "{\"selector\":{\"docType\":\"points\"}}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SkyfchainChaincode) point(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	pointId, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	key, err := pointKey(stub, pointId)
	if err != nil {
		return shim.Error(err.Error())
	}

	pointBytes, err := stub.GetState(key)

	if err != nil {
		return shim.Error(err.Error())
	} else if pointBytes == nil {
		return shim.Error(fmt.Sprintf("{\"Error\":\"Point does not exist: %v\"}", pointId))
	}

	return shim.Success(pointBytes)
}
