package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type route struct {
	ObjectType string  `json:"docType"`
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Active     bool    `json:"active"`
	Distance   int     `json:"distance"`
	Altitude   []int   `json:"altitude"`
	Points     []int64 `json:"points"`
}

func routeKey(stub shim.ChaincodeStubInterface, id int64) (string, error) {
	return stub.CreateCompositeKey("route", []string{string(id)})
}

func initRoutes() []route {
	route1 := route{
		ObjectType: "route",
		Id:         1,
		Name:       "Singapore - Pinang",
		Active:     true,
		Distance:   87000,
		Points:     []int64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		Altitude:   []int{100, 200}}
	route2 := route{
		ObjectType: "route",
		Id:         2,
		Name:       "Hanoi — Dien Bien Phu",
		Active:     true,
		Distance:   286000,
		Points:     []int64{21, 22},
		Altitude:   []int{200, 250}}
	route3 := route{
		ObjectType: "route",
		Id:         3,
		Name:       "Pohang — Fukuoka",
		Active:     true,
		Distance:   216000,
		Points:     []int64{31, 32},
		Altitude:   []int{200, 300}}
	route4 := route{
		ObjectType: "route",
		Id:         4,
		Name:       "Oita — Kochi",
		Active:     true,
		Distance:   196000,
		Points:     []int64{41, 42},
		Altitude:   []int{150, 200}}
	route5 := route{
		ObjectType: "route",
		Id:         5,
		Name: "Nizhnevartovsk — Pokachi	",
		Active:   true,
		Distance: 98000,
		Points:   []int64{51, 52},
		Altitude: []int{100, 300}}

	route6 := route{
		ObjectType: "route",
		Id:         6,
		Name:       "Gwangju — Jeju",
		Active:     true,
		Distance:   183000,
		Points:     []int64{61, 62},
		Altitude:   []int{100, 350}}

	route7 := route{
		ObjectType: "route",
		Id:         7,
		Name:       "Amsterdam — Borkum",
		Active:     false,
		Distance:   137000,
		Points:     []int64{71, 72},
		Altitude:   []int{200, 220}}

	return []route{route1, route2, route3, route4, route5, route6, route7}
}
func (t *SkyfchainChaincode) routes(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) > 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	queryString := "{\"selector\":{\"docType\":\"route\"}}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SkyfchainChaincode) route(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	routeId, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	key, err := routeKey(stub, routeId)
	if err != nil {
		return shim.Error(err.Error())
	}

	routeBytes, err := stub.GetState(key)

	if err != nil {
		return shim.Error(err.Error())
	} else if routeBytes == nil {
		return shim.Error(fmt.Sprintf("{\"Error\":\"Route does not exist: %v\"}", routeId))
	}

	return shim.Success(routeBytes)
}
