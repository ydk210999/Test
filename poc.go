// sycoin
package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"fmt"
	"encoding/json"
	//	"fmt"
)

type JUANZHENREN struct {
	ID   string
	List []*JUANZHENXINXI
}

type BEIJUANZHENREN struct {
	ID   string
	List []*JUANZHENXINXI
}

type JUANZHENXINXI struct {
	FromId string
	ToId   string
	Fund int64
}

//=================================================================================================================================
//	 Structure Definitions
//=================================================================================================================================
type Chaincode struct {
}

//常量定义

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	return nil, nil
}

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	//	fmt.Printf("****start to invoke******")
	var a JUANZHENXINXI
	err := json.Unmarshal([]byte(args[0]), &a)
	//	fmt.Printf("*****end to unmarshal*****")
	if err != nil {
		//		fmt.Printf("*****debug001****")
		return nil, err
	}
	key1 := a.FromId
	//	fmt.Printf("*****debug002****")
	jasonbytes, err := stub.GetState(key1)
	if err != nil {
		return nil, err
	}
	var b JUANZHENREN
	if jasonbytes != nil {
		//		fmt.Printf("*****debug003****")
		err = json.Unmarshal(jasonbytes, &b)
		if err != nil {
			return nil, err
		}
	}
	b.ID = a.FromId
	b.List = append(b.List, &a)
	jasonbytes, err = json.Marshal(b)
	stub.PutState(key1, jasonbytes)

	//
	//	fmt.Printf("*****debug004****")
	key2 := a.ToId
	jasonbytes, err = stub.GetState(key2)
	if err != nil {
		//		fmt.Printf("not found********")
		return nil, err
	}
	var c BEIJUANZHENREN
	//	fmt.Printf("init check**", c)
	if jasonbytes != nil {
		err = json.Unmarshal(jasonbytes, &c)
		if err != nil {
			return nil, err
		}
	}
	c.ID = a.ToId
	c.List = append(c.List, &a)
	jasonbytes, err = json.Marshal(c)
	stub.PutState(key2, jasonbytes)
	return nil, nil

}

func (t *Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	jsonbytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, err
	}
	return jsonbytes, nil

}

//============================================================================================================
//     Function main函数
//============================================================================================================
func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		panic(err)
	}
}
