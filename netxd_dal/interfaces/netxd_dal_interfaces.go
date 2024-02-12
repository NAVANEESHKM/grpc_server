// this interface will provide the requried methods
package interfaces

import ("netxd_grpc_mongo/netxd_dal/models"
)



type ICustomer interface{
	CreateCustomer(user *models.Customer) (*models.Customer, error)
	UpdateCustomer(fromid int32, toid int32, amount int32)(*models.Transaction,error)
	
}