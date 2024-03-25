package main

import "C"

import (
	"fmt"
	"unsafe"

	"github.com/Unique-Divine/gonibi"
)

// C representation of the NibiruClient struct
type C_NibiruClient struct {
	ChainId          *C.char
	Keyring          *C.char        // Assuming Keyring can be serialized
	EncCfg           *C.char        // Assuming EncodingConfig can be serialized
	Querier          *C.char        // Assuming Querier can be serialized
	CometRPC         *C.char        // Assuming cmtrpcclient.Client can be serialized
	AccountRetriever *C.char        // Assuming authtypes.AccountRetriever can be serialized
	GrpcClient       unsafe.Pointer // (See explanation below)
}

func main() {}

const (
	Success = 0
	Fail    = 1
)

//export NewNibiruClientDefault
func NewNibiruClientDefault() C.int {
	fmt.Println("Call [NewNibiruClientDefault]")
	grpcConn, err := gonibi.GetGRPCConnection(gonibi.DefaultNetworkInfo.GrpcEndpoint, true, 2)
	if err != nil {
		fmt.Println("[NewNibiruClientDefault] GetGRPCConnection error: " + err.Error())
		return Fail
	}
	gosdk, err := gonibi.NewNibiruClient("nibiru-localnet-0", grpcConn, gonibi.DefaultNetworkInfo.TmRpcEndpoint)
	fmt.Println("[NewNibiruClientDefault] Connected to " + gosdk.ChainId)
	if err != nil {
		fmt.Println("[NewNibiruClientDefault] Connect to network error: " + err.Error())
		return Fail
	}

	return Success
}

//export NewNibiruClient
func NewNibiruClient(chainId *C.char, grpcEndpoint *C.char, rpcEndpoint *C.char) C.int {
	fmt.Println("Call [NewNibiruClient]")
	grpcConn, err := gonibi.GetGRPCConnection(C.GoString(grpcEndpoint), true, 2)
	if err != nil {
		fmt.Println("[NewNibiruClient] GetGRPCConnection error: " + err.Error())
		return Fail
	}
	gosdk, err := gonibi.NewNibiruClient(C.GoString(chainId), grpcConn, C.GoString(rpcEndpoint))
	fmt.Println("[NewNibiruClient] Connected to " + gosdk.ChainId)
	if err != nil {
		fmt.Println("[NewNibiruClient] Connect to network error: " + err.Error())
		return Fail
	}

	return Success
}
