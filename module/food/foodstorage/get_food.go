package foodstorage

import (
	"errors"
	"food-delivery/module/food/foodmodel"
)

func (s *store) GetAllFoods() ([]foodmodel.Food, error) {
	rows, err := s.db.Raw("select * from notes").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []foodmodel.Food
	var item foodmodel.Food
	for rows.Next() {
		s.db.ScanRows(rows, &item)
		result = append(result, item)
	}

	return result, nil
}

func (s *store) GetFoodById(id int) (foodmodel.Food, error) {
	var result foodmodel.Food
	err := s.db.Raw("select * from foods where id = ?", id).Scan(&result).Error
	if err != nil {
		return result, err
	}
	if result == (foodmodel.Food{}) {
		return result, errors.New("Không tìm thấy note")
	}

	return result, nil
}
