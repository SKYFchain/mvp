package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/SKYFchain/mvp/chaincode/model"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
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
		ID:             1,
		Name:           "SKYF-P2.01",
		Model:          "SKYF P2",
		Capacity:       0,
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
		ID:             2,
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
		ID:             3,
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
		ID:             4,
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
		ID:             5,
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
		ID:         6,
		Name:       "Drone 6",
		ETA:        &t6,
		UID:        uuid.New(),
		Notes:      "Waiting for new transmission",
		Status:     "repair"}
	t7 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	drone7 := model.Drone{
		ObjectType: "drone",
		ID:         7,
		Name:       "Drone 7",
		ETA:        &t7,
		UID:        uuid.New(),
		Notes:      "Check chassis",
		Status:     "inspection"}
	t8 := time.Now().Add(time.Hour * time.Duration(rand.Intn(240)))
	drone8 := model.Drone{
		ObjectType: "drone",
		ID:         8,
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

	droneID, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	key, err := droneKey(stub, droneID)
	if err != nil {
		return shim.Error(err.Error())
	}

	droneBytes, err := stub.GetState(key)

	if err != nil {
		return shim.Error(err.Error())
	} else if droneBytes == nil {
		return shim.Error(fmt.Sprintf("{\"Error\":\"Drone does not exist: %v\"}", droneID))
	}

	return shim.Success(droneBytes)
}

func (t *SkyfchainChaincode) saveDrone(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	droneBytes := []byte(args[0])
	var drone model.Drone
	err := json.Unmarshal(droneBytes, &drone)
	if err != nil {
		return shim.Error(err.Error())
	}

	lastDroneID, err := getLastDroneID(stub)

	if err != nil {
		return shim.Error(err.Error())
	}

	drone.ID = lastDroneID + 1
	drone.ObjectType = "drone"
	drone.Status = "idle"
	droneBytes, err = json.Marshal(drone)
	if err != nil {
		return shim.Error(err.Error())
	}
	key, err := droneKey(stub, drone.ID)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(key, droneBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(droneBytes)
}

func getLastDroneID(stub shim.ChaincodeStubInterface) (int64, error) {
	queryString := "{\"selector\":{\"docType\":\"drone\"},\"sort\":[{\"id\":\"desc\"}], \"fields\":[\"id\"], \"limit\": 1}"
	iterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return 0, err
	}
	defer iterator.Close()

	var droneID = droneID{ID: 1}
	if iterator.HasNext() {
		queryResponse, err := iterator.Next()
		if err != nil {
			return 0, err
		}
		err = json.Unmarshal(queryResponse.Value, &droneID)

		if err != nil {
			return 0, err
		}
	}

	return droneID.ID, nil
}

type droneID struct {
	ID int64 `json:"id"`
}
