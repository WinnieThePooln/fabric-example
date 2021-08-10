/**
  author: hanxiaodong
 */
package main

import (
    "os"
    "fmt"
    "github.com/kongyixueyuan.com/education/sdkInit"
    "github.com/kongyixueyuan.com/education/service"
    "github.com/kongyixueyuan.com/education/web"
    "github.com/kongyixueyuan.com/education/web/controller"
    "reflect"
//    "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

const (
    configFile = "config.yaml"
    initialized = false
    SimpleCC = "simplecc"
)


func main() {

    initInfo := &sdkInit.InitInfo{

        ChannelID: "kevinkongyixueyuan",
        ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/kongyixueyuan.com/education/fixtures/artifacts/channel.tx",

        OrgAdmin:"Admin",
        OrgName:"Org1",
        OrdererOrgName: "orderer.kevin.kongyixueyuan.com",
	ChaincodeID: SimpleCC,
	ChaincodeGoPath: os.Getenv("GOPATH"),
	ChaincodePath: "github.com/kongyixueyuan.com/education/chaincode/",
	UserName:"User1",
    }

    sdk, err := sdkInit.SetupSDK(configFile, initialized)
    if err != nil {
        fmt.Printf(err.Error())
        return
    }

    defer sdk.Close()

    err = sdkInit.CreateChannel(sdk, initInfo)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
    if err != nil {
        fmt.Println(err.Error())
	return
    }
    fmt.Println(channelClient)

    svs := service.ServiceSetup{
	ChaincodeID:SimpleCC,
	Client:channelClient,
    }
    fmt.Println("start service")
    fmt.Println("type name",reflect.TypeOf(svs.ChaincodeID).Name())
    fmt.Println("type kind",reflect.TypeOf(svs.ChaincodeID).Name())
    fmt.Println(" type name", reflect.TypeOf(svs.Client).Name())
    fmt.Println(" type kind", reflect.TypeOf(svs.Client).Kind())

    svs.PrintInfo()
    fmt.Println("test")
/*
    eventID := "eventSetInfo"
    reg, _ := service.RegitserEvent(svs.Client, svs.ChaincodeID, eventID)
    defer svs.Client.UnregisterChaincodeEvent(reg)

    req := channel.Request{ChaincodeID: svs.ChaincodeID, Fcn: "set", Args: [][]byte{[]byte("123"), []byte("456"), []byte(eventID)}}
    fmt.Println(" type name", reflect.TypeOf(req).Name())
    fmt.Println(" type kind", reflect.TypeOf(req).Kind())

    respone, _ := svs.Client.Execute(req)
    
    fmt.Println("respone id",string(respone.TransactionID))
*/
	//err = service.EventResult(notifier, eventID)
	//if err != nil {
	//	fmt.Println(err)
	//}
    msg, err := svs.SetInfo("Hanxiaodong", "Kongyixueyuan")
    if err != nil {
	fmt.Println(err)
    } else {
	fmt.Println(msg)
    }
    msg, err = svs.GetInfo("Hanxiaodong")
    if err != nil {
	fmt.Println(err)
    } else {
	fmt.Println(msg)
    }
    app := controller.Application{
	Fabric: &svs,
    }
    web.WebStart(&app)
}
