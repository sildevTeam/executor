package service

import (
	"sync"
	"../testservice"
)

var instance *testservice.SilGatewayTestServiceClient
var once sync.Once

func GetIns() *testservice.SilGatewayTestServiceClient {
	once.Do(func() {
		protocolFactory, useTransport, err := PrepareConn()
		if err != nil {
			panic(err)
		}
		var client = testservice.NewSilGatewayTestServiceClientFactory(*useTransport, protocolFactory)
		instance = client
	})
	return instance
}
