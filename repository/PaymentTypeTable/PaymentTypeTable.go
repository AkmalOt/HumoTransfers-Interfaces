package PaymentType

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
)

func (r PaymentTypeRepo) AddPaymentType(paymentType *models.PaymentType) error {
	tx := db.Data.Table("payment_type").Create(paymentType)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *PaymentTypeRepo) GetPaymentType(pagination *models.Pagination) ([]models.PaymentType, error) {
	var paymentType []models.PaymentType
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("payment_type").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.PaymentType{}).Find(&paymentType)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return paymentType, nil
}

func (r PaymentTypeRepo) UpdatePaymentType(PaymentType *models.PaymentType) error {

	tx := db.Data.Model(models.Vendor{}).Table("payment_type").Where("id = ?", PaymentType.ID).Updates(models.Vendor{Name: PaymentType.Name})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r PaymentTypeRepo) DeletePaymentType(PaymentType *models.PaymentType) error {

	query := db.Data.Table("payment_type").Where("id =?", PaymentType.ID).Delete(PaymentType)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r PaymentTypeRepo) TotalPagePaymentType(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("payment_type").Count(&length)
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
