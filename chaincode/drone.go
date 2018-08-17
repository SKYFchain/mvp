package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"math/rand"
	"strconv"
	"time"
)

type drone struct {
	ObjectType     string    `json:"docType"`
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Model          string    `json:"model"`
	Capacity       int       `json:"capacity"`
	Image          string    `json:"image"`
	Description    string    `json:"description"`
	Operator       string    `json:"operator"`
	Docs           []doc     `json:"docs"`
	Status         string    `json:"status"`
	NextInspection time.Time `json:"nextInspection"`
	UID            uuid.UUID `json:"uid"`
	Version        version   `json:"version"`
	Point          int64     `json:"point"`
}

type version struct {
	Hardware string `json:"hardware"`
	Software string `json:"software"`
}

type doc struct {
	Name string `json:"name"`
}

func droneKey(stub shim.ChaincodeStubInterface, id int64) (string, error) {
	return stub.CreateCompositeKey("drone", []string{string(id)})
}

func initDrones() []drone {
	rand.Seed(42)
	docs := []doc{doc{"General overview"}, doc{"Technical certificate"}, doc{"User's manual"}}
	drone1 := drone{
		ObjectType:     "drone",
		Id:             1,
		Name:           "Drone 1",
		Model:          "Heavy UAV",
		Capacity:       400,
		Image:          "media/drone1.png",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: time.Now().Add(time.Hour * time.Duration(rand.Intn(240))),
		UID:            uuid.New(),
		Version:        version{"H10", "2.1.1"},
		Point:          16}
	drone2 := drone{
		ObjectType:     "drone",
		Id:             2,
		Name:           "Drone 2",
		Model:          "Light UAV",
		Capacity:       100,
		Image:          "media/drone2.png",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: time.Now().Add(time.Hour * time.Duration(rand.Intn(240))),
		UID:            uuid.New(),
		Version:        version{"L16", "2.1.1"},
		Point:          21}
	drone3 := drone{
		ObjectType:     "drone",
		Id:             3,
		Name:           "Drone 3",
		Model:          "Heavy HA UAV",
		Capacity:       250,
		Image:          "media/drone3.png",
		Description:    "High-altitude drone.",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: time.Now().Add(time.Hour * time.Duration(rand.Intn(240))),
		UID:            uuid.New(),
		Version:        version{"HA5", "2.0.5"},
		Point:          31}
	drone4 := drone{
		ObjectType:     "drone",
		Id:             4,
		Name:           "Drone 4",
		Model:          "Medium UAV",
		Capacity:       250,
		Image:          "media/drone4.png",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: time.Now().Add(time.Hour * time.Duration(rand.Intn(240))),
		UID:            uuid.New(),
		Version:        version{"L16", "2.0.5"},
		Point:          41}
	drone5 := drone{
		ObjectType:     "drone",
		Id:             5,
		Name:           "Drone 5",
		Model:          "Heavy HR UAV",
		Capacity:       200,
		Image:          "media/drone5.png",
		Description:    "High-reliability drone",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: time.Now().Add(time.Hour * time.Duration(rand.Intn(240))),
		UID:            uuid.New(),
		Version:        version{"HR2", "1.2.19"},
		Point:          51}
	return []drone{drone1, drone2, drone3, drone4, drone5}
}

func (t *SkyfchainChaincode) drones(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) > 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	queryString := "{\"selector\":{\"docType\":\"drone\"}}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SkyfchainChaincode) drone(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	droneId, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	key, err := droneKey(stub, droneId)
	if err != nil {
		return shim.Error(err.Error())
	}

	droneBytes, err := stub.GetState(key)

	if err != nil {
		return shim.Error(err.Error())
	} else if droneBytes == nil {
		return shim.Error(fmt.Sprintf("{\"Error\":\"Drone does not exist: %v\"}", droneId))
	}

	return shim.Success(droneBytes)
}
