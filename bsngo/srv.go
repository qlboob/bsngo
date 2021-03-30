package bsngo

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	preKey = "ming_yuan_"
)

func Apply(stubInterface shim.ChaincodeStubInterface, str []string) peer.Response {
	defer func() {
		SetLogger("apply end")
	}()
	SetLogger("apply start")
	if len(str) != 1 {
		return shim.Error("param error")
	}
	var data BaseModel
	err := json.Unmarshal([]byte(str[0]), &data)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}
	if err := data.Check(); err != nil {
		return shim.Error(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}

	jsonBt, err := json.Marshal(&data)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}
	mainKey := GetMainKey(data.BaseKey)
	err = stubInterface.PutState(mainKey, jsonBt)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to add data:%s", err.Error()))
	}

	SetLogger("finish saving data", mainKey)

	return shim.Success([]byte("SUCCESS"))
}

func Get(stubInterface shim.ChaincodeStubInterface, str []string) peer.Response {
	if len(str) != 1 {
		return shim.Error("param error")
	}

	key := str[0]
	mainKey := GetMainKey(key)

	result, err := stubInterface.GetState(mainKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) == 0 {
		SetLogger(fmt.Sprintf("data with key value of not found ", mainKey))
		return shim.Error("data not found")
	}
	var dbModel BaseModel
	err = json.Unmarshal(result, &dbModel)
	if err != nil {
		return shim.Error(fmt.Sprintf("data formatting error:%s", err.Error()))
	}
	return shim.Success([]byte(dbModel.BaseValue))
}
