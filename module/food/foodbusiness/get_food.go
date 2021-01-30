package foodbusiness

import (
	"errors"
	"food-delivery/module/food/foodmodel"
)

type GetFoodStore interface {
	GetAllFoods() ([]foodmodel.Food, error)
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

func (getFood *getFoodBusiness) GetAll() ([]foodmodel.Food, error) {
	result, err := getFood.store.GetAllFoods()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (getFood *getFoodBusiness) GetById(id int) (foodmodel.Food, error) {
	result, err := getFood.store.GetFoodById(id)
	if err != nil {
		return result, err
	}

	if result == (foodmodel.Food{}) {
		return result, errors.New("Không tìm thấy note")
	}

	return result, nil
}