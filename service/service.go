package service

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
)
var Host string
var Port int
func PrepareConn() (*thrift.TBinaryProtocolFactory, *thrift.TTransport, error) {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(fmt.Sprintf("%s:%d",Host,Port))
	if err != nil {
		panic(err)
	}

	useTransport := transportFactory.GetTransport(transport)

	if err := useTransport.Open(); err != nil {
		panic(err)
	}
	return protocolFactory, &useTransport, nil
}
