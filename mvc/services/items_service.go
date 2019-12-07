package services

import "github.com/mgw2007/golang-microservices/mvc/domain"

import "github.com/mgw2007/golang-microservices/mvc/utils"

type itemsService struct {
}

var (
	//ItemsService for user service operations
	ItemsService itemsService
)

//GetItem for return items
func (i *itemsService) GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {

	return nil, nil
}
