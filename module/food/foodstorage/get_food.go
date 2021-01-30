package foodstorage

import (
	"errors"
	"fmt"
	"food-delivery/common"
	"food-delivery/module/food/foodmodel"
)

func (s *store) GetAllFoods(paging common.Paging) ([]foodmodel.GetFood, error) {
	sqlQuery := fmt.Sprintf(`
		SELECT id, name, price, short_description
		FROM foods
		ORDER BY created_at DESC
		LIMIT %d
		OFFSET %d
	`, paging.Limit, (paging.Page - 1) * paging.Limit)
	
	rows, err := s.db.Raw(sqlQuery).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []foodmodel.GetFood
	var item foodmodel.GetFood
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
