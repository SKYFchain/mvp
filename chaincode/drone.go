package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"gitlab.qdlt.io/skyf/skyfchain/chaincode/model"
	"math/rand"
	"strconv"
	"time"
)

func droneKey(stub shim.ChaincodeStubInterface, id int64) (string, error) {
	return stub.CreateCompositeKey("drone", []string{string(id)})
}

func initDrones() []model.Drone {
	rand.Seed(42)
	docs := []model.Doc{model.Doc{"General overview"}, model.Doc{"Technical certificate"}, model.Doc{"User's manual"}}
	t1 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	var p1 int64 = 16
	drone1 := model.Drone{
		ObjectType:     "drone",
		Id:             1,
		Name:           "Drone 1",
		Model:          "Heavy UAV",
		Capacity:       400,
		Image:          "media/drone1.png",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: &t1,
		UID:            uuid.New(),
		Version:        &model.Version{"H10", "2.1.1"},
		Point:          &p1}
	t2 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	var p2 int64 = 21
	drone2 := model.Drone{
		ObjectType:     "drone",
		Id:             2,
		Name:           "Drone 2",
		Model:          "Light UAV",
		Capacity:       100,
		Image:          "media/drone2.png",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: &t2,
		UID:            uuid.New(),
		Version:        &model.Version{"L16", "2.1.1"},
		Point:          &p2}
	t3 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	var p3 int64 = 31
	drone3 := model.Drone{
		ObjectType:     "drone",
		Id:             3,
		Name:           "Drone 3",
		Model:          "Heavy HA UAV",
		Capacity:       250,
		Image:          "media/drone3.png",
		Description:    "High-altitude model.",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: &t3,
		UID:            uuid.New(),
		Version:        &model.Version{"HA5", "2.0.5"},
		Point:          &p3}
	t4 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	var p4 int64 = 41
	drone4 := model.Drone{
		ObjectType:     "drone",
		Id:             4,
		Name:           "Drone 4",
		Model:          "Medium UAV",
		Capacity:       250,
		Image:          "media/drone4.png",
		Operator:       "SKYF",
		Docs:           docs,
		Status:         "mission",
		NextInspection: &t4,
		UID:            uuid.New(),
		Version:        &model.Version{"L16", "2.0.5"},
		Point:          &p4}
	t5 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	var p5 int64 = 51
	drone5 := model.Drone{
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
		NextInspection: &t5,
		UID:            uuid.New(),
		Version:        &model.Version{"HR2", "1.2.19"},
		Point:          &p5}
	t6 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	drone6 := model.Drone{
		ObjectType: "drone",
		Id:         6,
		Name:       "Drone 6",
		ETA:        &t6,
		UID:        uuid.New(),
		Notes:      "Waiting for new transmission",
		Status:     "repair"}
	t7 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	drone7 := model.Drone{
		ObjectType: "drone",
		Id:         7,
		Name:       "Drone 7",
		ETA:        &t7,
		UID:        uuid.New(),
		Notes:      "Check chassis",
		Status:     "inspection"}
	t8 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	drone8 := model.Drone{
		ObjectType: "drone",
		Id:         8,
		Name:       "Drone 8",
		ETA:        &t8,
		UID:        uuid.New(),
		Notes:      "Change oil filter",
		Status:     "maintenance"}
	return []model.Drone{drone1, drone2, drone3, drone4, drone5, drone6, drone7, drone8}
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
