package Country

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
	"log"
)

func (r CountryRepo) AddCountry(country *models.Countries) error {

	tx := db.Data.Create(country)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *CountryRepo) GetCountries(pagination *models.Pagination) ([]models.Countries, error) {
	var countries []models.Countries
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("countries").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Countries{}).Find(&countries)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return countries, nil
}

func (r CountryRepo) UpdateCountries(country *models.Countries) error {

	//log.Println(1, country, 2, countries)
	tx := db.Data.Model(models.Countries{}).Where("id = ?", country.ID).Updates(models.Countries{Name: country.Name, Icon: country.Icon, Active: country.Active})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (r CountryRepo) DeleteCountries(country *models.Countries) error {
	query := db.Data.Table("countries").Where("id =?", country.ID).Delete(country)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r CountryRepo) CountryStatus(country *models.Countries) error {
	tx := db.Data.Where("id = ?", country.ID).Table("countries").Scan(country)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if country.Active == true {
		country.Active = false
	} else {
		country.Active = true
	}
	tx = db.Data.Where("id = ?", country.ID).Table("countries").Update("active", country.Active)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r CountryRepo) TotalPageCountry(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}

	query := db.Data.Table("countries").Count(&length)
	if query.Error != nil {
		//log.Println(query.Error, "error in TotalPageCountry")
		return 0, query.Error
	}

	totalPage := length / limit
	if length%limit != 0 {
		totalPage++
	}

	return totalPage, nil
}
