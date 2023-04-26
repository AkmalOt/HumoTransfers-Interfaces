package ServiceRules

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
)

func (r ServiceRulesRepo) AddServiceRules(ServiceRules *models.ServicesRules) error {
	tx := db.Data.Table("services_rules").Create(ServiceRules)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *ServiceRulesRepo) GetServiceRules(pagination *models.Pagination) ([]models.ServicesRules, error) {
	var ServicesRules []models.ServicesRules
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("services_rules").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.ServicesRules{}).Find(&ServicesRules)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return ServicesRules, nil
}

func (r ServiceRulesRepo) UpdateServiceRules(serviceRules *models.ServicesRules) error {

	tx := db.Data.Model(models.ServicesRules{}).Table("services_rules").Where("id = ?", serviceRules.ID).Updates(models.ServicesRules{Name: serviceRules.Name, Type: serviceRules.Type})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r ServiceRulesRepo) DeleteServiceRules(ServiceRules *models.ServicesRules) error {

	query := db.Data.Table("services_rules").Where("id =?", ServiceRules.ID).Delete(ServiceRules)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r ServiceRulesRepo) TotalPageServiceRules(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("services_rules").Count(&length)
	if query.Error != nil {
		//log.Println(query.Error)
		return 0, query.Error
	}
	totalPage := length / limit
	if length%limit != 0 {
		totalPage++
	}
	return totalPage, nil
}
