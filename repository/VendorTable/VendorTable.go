package VendorTable

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
)

func (r VendorRepo) AddVendor(account *models.Vendor) error {
	tx := db.Data.Table("vendor").Create(account)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *VendorRepo) GetVendor(pagination *models.Pagination) ([]models.Vendor, error) {
	var Vendor []models.Vendor
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("vendor").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.AccountAgent{}).Find(&Vendor)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return Vendor, nil
}

func (r VendorRepo) UpdateVendor(Vendor *models.Vendor) error {

	tx := db.Data.Model(models.Vendor{}).Table("vendor").Where("id = ?", Vendor.ID).Updates(models.Vendor{Name: Vendor.Name})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r VendorRepo) DeleteVendor(Vendor *models.Vendor) error {

	query := db.Data.Table("vendor").Where("id =?", Vendor.ID).Delete(Vendor)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r VendorRepo) VendorStatus(Vendor *models.Vendor) error {
	tx := db.Data.Where("id = ?", Vendor.ID).Table("vendor").Scan(Vendor)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	if Vendor.Active == true {
		Vendor.Active = false
	} else {
		Vendor.Active = true
	}
	tx = db.Data.Where("id = ?", Vendor.ID).Table("vendor").Update("active", Vendor.Active)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r VendorRepo) TotalPageVendor(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("vendor").Count(&length)
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
