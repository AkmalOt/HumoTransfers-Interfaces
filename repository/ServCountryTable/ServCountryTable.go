package ServCountry

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
)

func (r ServCountryRepo) AddServCountry(servCountry *models.ServicesCountry) error {
	tx := db.Data.Table("services_country").Create(servCountry)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *ServCountryRepo) GetServCountry(pagination *models.Pagination) ([]models.ServicesCountry, error) {
	var ServCountry []models.ServicesCountry
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("services_country").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.ServicesCountry{}).Find(&ServCountry)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return ServCountry, nil
}

func (r ServCountryRepo) DeleteServCountry(servCountry *models.ServicesCountry) error {

	query := db.Data.Table("services_country").Where("service_id =?", servCountry.ServiceId).Where("country_id = ?", servCountry.CountryId).Delete(servCountry)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r ServCountryRepo) ServCountryStatus(servCountry *models.ServicesCountry) error {
	tx := db.Data.Table("services_country").Where("service_id =?", servCountry.ServiceId).Where("country_id = ?", servCountry.CountryId).Scan(servCountry)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	if servCountry.Active == true {
		servCountry.Active = false
	} else {
		servCountry.Active = true
	}
	tx = db.Data.Table("services_country").Where("service_id =?", servCountry.ServiceId).Where("country_id = ?", servCountry.CountryId).Update("active", servCountry.Active)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r ServCountryRepo) TotalPageServCountry(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("services_country").Count(&length)
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
