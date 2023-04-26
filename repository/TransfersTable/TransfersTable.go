package TransferTable

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
)

func (r TransferTableRepo) AddTransfer(transfer *models.Transfers) error {
	tx := db.Data.Table("transfers").Create(transfer)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *TransferTableRepo) GetTransfer(pagination *models.Pagination) ([]models.Transfers, error) {
	var Transfer []models.Transfers
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("transfers").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Transfers{}).Find(&Transfer)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return Transfer, nil
}

func (r TransferTableRepo) UpdateTransfers(transfer *models.Transfers) error {

	tx := db.Data.Table("transfers").Model(models.Transfers{}).Where("id = ?", transfer.ID).Updates(models.Transfers{Entity: transfer.Entity, EntityId: transfer.EntityId, LangId: transfer.LangId, Value: transfer.Value})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r TransferTableRepo) DeleteTransfers(transfers *models.Transfers) error {

	query := db.Data.Table("transfers").Where("id =?", transfers.ID).Delete(transfers)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r TransferTableRepo) TotalPageTransfer(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("transfers").Count(&length)
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
