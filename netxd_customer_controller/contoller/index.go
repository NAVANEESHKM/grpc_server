package controller

import (
	"context"
	"netxd_grpc_mongo/netxd_dal/interfaces"
	"netxd_grpc_mongo/netxd_dal/models"
	pro "github.com/NAVANEESHKM/grpc_proto"
	
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	db := &models.Customer{BankID:req.BankID,Customer_Name:req.Customer_Name,Customer_ID:req.Customer_ID,Balance: req.Balance}
	result, err := CustomerService.CreateCustomer(db)
	if err != nil {
		return nil, err
	}
	responseCustomer := &pro.CustomerResponse{
		Balance: result.Balance,
		
	}
	
	return responseCustomer, nil
}
func (s *RPCServer) UpdateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerDetailsResponse, error) {
	
	
   result,err:=CustomerService.UpdateCustomer(req.From_ID,req.TO_ID,req.Amount)
   if err != nil {
	return nil, err
}
responseCustomer := &pro.CustomerDetailsResponse{
	From_ID: result.From_ID,
	
}

return responseCustomer, nil
}
