package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"gitlab.qdlt.io/skyf/skyfchain/chaincode/model"
	"strconv"
	"time"
)

func missionKey(stub shim.ChaincodeStubInterface, id int64) (string, error) {
	return stub.CreateCompositeKey("mission", []string{string(id)})
}

func initMissions() []model.Mission {
	certs := []model.Cert{model.Cert{Name: "Operator's license"}, model.Cert{Name: "Certificate of conformity"}}
	legal := []model.Legal{model.Legal{Name: "Flight list"}, model.Legal{Name: "Cargo declaration"}}
	mission1 := model.Mission{
		ObjectType: "mission",
		Id:         1,
		Name:       "Spare parts and instruments",
		Point:      14,
		Cargo:      "1",
		Customer:   "customer",
		Drone:      1,
		Price:      1600,
		Route:      1,
		Status:     "progress",
		ETA:        time.Date(2018, 06, 21, 17, 0, 0, 0, time.Local),
		ETD:        time.Date(2018, 06, 21, 12, 0, 0, 0, time.Local),
		Certs:      certs,
		Legal:      legal}
	mission2 := model.Mission{
		ObjectType: "mission",
		Id:         2,
		Name:       "Mail",
		Point:      21,
		Cargo:      "1",
		Customer:   "customer",
		Drone:      2,
		Price:      800,
		Route:      2,
		Status:     "pending",
		ETA:        time.Date(2018, 06, 22, 17, 0, 0, 0, time.Local),
		ETD:        time.Date(2018, 06, 22, 12, 0, 0, 0, time.Local),
		Certs:      certs,
		Legal:      legal}
	mission3 := model.Mission{
		ObjectType: "mission",
		Id:         3,
		Name:       "Water",
		Point:      31,
		Cargo:      "1",
		Customer:   "customer",
		Drone:      3,
		Price:      1400,
		Route:      3,
		Status:     "progress",
		ETA:        time.Date(2018, 06, 22, 17, 0, 0, 0, time.Local),
		ETD:        time.Date(2018, 06, 22, 12, 0, 0, 0, time.Local),
		Certs:      certs,
		Legal:      legal}
	mission4 := model.Mission{
		ObjectType: "mission",
		Id:         4,
		Name:       "Spare parts",
		Point:      41,
		Cargo:      "1",
		Customer:   "customer",
		Drone:      4,
		Price:      2210,
		Route:      4,
		Status:     "trouble",
		ETA:        time.Date(2018, 06, 22, 17, 0, 0, 0, time.Local),
		ETD:        time.Date(2018, 06, 22, 12, 0, 0, 0, time.Local),
		Certs:      certs,
		Legal:      legal}
	mission5 := model.Mission{
		ObjectType: "mission",
		Id:         5,
		Name:       "Measuring instruments",
		Point:      51,
		Cargo:      "1",
		Customer:   "customer",
		Drone:      5,
		Price:      1850,
		Route:      5,
		Status:     "done",
		ETA:        time.Date(2018, 06, 22, 17, 0, 0, 0, time.Local),
		ETD:        time.Date(2018, 06, 22, 12, 0, 0, 0, time.Local),
		Certs:      certs,
		Legal:      legal}
	return []model.Mission{mission1, mission2, mission3, mission4, mission5}
}

func (t *SkyfchainChaincode) missions(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) > 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	queryString := "{\"selector\":{\"docType\":\"mission\"}}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SkyfchainChaincode) mission(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	missionId, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	key, err := missionKey(stub, missionId)
	if err != nil {
		return shim.Error(err.Error())
	}

	missionBytes, err := stub.GetState(key)

	if err != nil {
		return shim.Error(err.Error())
	} else if missionBytes == nil {
		return shim.Error(fmt.Sprintf("{\"Error\":\"Mission does not exist: %v\"}", missionId))
	}

	return shim.Success(missionBytes)
}
