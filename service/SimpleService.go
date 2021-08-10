/**
  author: kevin
 */

package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t ServiceSetup) SetInfo(name, num string) (string, error) {

	eventID := "eventSetInfo"
	reg,notifier := RegitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "set", Args: [][]byte{[]byte(name), []byte(num), []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		fmt.Println("err1:",err)
		return string(respone.TransactionID), err
	}

        err = EventResult(notifier, eventID)
	if err != nil {
		fmt.Println("err2:",err)
		return string(respone.TransactionID), err
	}

	return string(respone.TransactionID), nil
}


func (t ServiceSetup) GetInfo(name string) (string, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "get", Args: [][]byte{[]byte(name)}}
	respone, err:= t.Client.Query(req)
	if err != nil {
		fmt.Println("err3:",err)
		return string(respone.Payload), err
	}

	return string(respone.Payload), nil
}

func (t ServiceSetup) PrintInfo() string{
	return "12345"
}
