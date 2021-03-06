package foodbusiness

import (
	"context"
	"errors"
	"food-delivery/common"
	"food-delivery/module/food/foodmodel"
)

type GetFoodStore interface {
	GetAllFoods(paging common.Paging) ([]foodmodel.GetFood, error)
	GetFoodById(id int) (foodmodel.Food, error)
}

type getFoodBusiness struct {
	store GetFoodStore
}

func NewGetFoodBusiness(store GetFoodStore) *getFoodBusiness {
	return &getFoodBusiness{
		store: store,
	}
}

func (getFood *getFoodBusiness) GetAll(ctx context.Context, paging common.Paging) ([]foodmodel.GetFood, error) {
	result, err := getFood.store.GetAllFoods(paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (getFood *getFoodBusiness) GetById(ctx context.Context, id int) (foodmodel.Food, error) {
	result, err := getFood.store.GetFoodById(id)
	if err != nil {
		return result, err
	}

	if result == (foodmodel.Food{}) {
		return result, errors.New("Không tìm thấy note")
	}

	return result, nil
}