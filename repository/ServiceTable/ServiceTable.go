package ServiceTable

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
)

func (r ServiceTableRepo) AddService(account *models.Services) error {
	tx := db.Data.Create(account)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *ServiceTableRepo) GetService(pagination *models.Pagination) ([]models.Services, error) {
	var Service []models.Services
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("services").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.AccountAgent{}).Find(&Service)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return Service, nil
}

func (r ServiceTableRepo) UpdateService(Service *models.Services) error {

	tx := db.Data.Model(models.Services{}).Table("services").Where("id = ?", Service.ID).
		Updates(models.Services{VendorId: Service.VendorId, Name: Service.Name, Type: Service.Type})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r ServiceTableRepo) DeleteService(Service *models.Services) error {

	query := db.Data.Table("services").Where("id =?", Service.ID).Delete(Service)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r ServiceTableRepo) ServicesStatus(Service *models.Services) error {
	tx := db.Data.Where("id = ?", Service.ID).Table("services").Scan(Service)
	if tx.Error != nil {
		//log.Println(tx.Error)

		return tx.Error
	}
	if Service.Active == true {
		Service.Active = false
	} else {
		Service.Active = true
	}
	tx = db.Data.Where("id = ?", Service.ID).Table("services").Update("active", Service.Active)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r ServiceTableRepo) TotalPageServices(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("services").Count(&length)
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
