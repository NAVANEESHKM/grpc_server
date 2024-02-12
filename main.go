package main

import (
	"context"
	"fmt"
	"net"

	pro "github.com/NAVANEESHKM/grpc_proto"
	"netxd_grpc_mongo/netxd_customer_config/config"
	"netxd_grpc_mongo/netxd_customer_config/constants"
	"netxd_grpc_mongo/netxd_dal/services"
	controller "netxd_grpc_mongo/netxd_customer_controller/contoller"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := config.GetCollection(client, "bankdb", "customers")
	TransactionCollection:=config.GetCollection(client,"bankdb","transactions")
	controller.CustomerService = services.InitCustomerService(client,CustomerCollection, TransactionCollection,context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	

	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &controller.RPCServer{})
	reflection.Register(s)

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
