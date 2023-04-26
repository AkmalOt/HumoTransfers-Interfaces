package SysMessage

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
)

func (r SysMessageRepo) AddSysMessage(message *models.SysMessage) error {
	tx := db.Data.Create(message)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *SysMessageRepo) GetSysMessage(pagination *models.Pagination) ([]models.SysMessage, error) {
	var SysMessage []models.SysMessage
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("sys_messages").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.SysMessage{}).Find(&SysMessage)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return SysMessage, nil
}

func (r SysMessageRepo) UpdateSysMessage(message *models.SysMessage) error {

	tx := db.Data.Model(models.SysMessage{}).Where("id = ?", message.ID).Updates(models.SysMessage{Name: message.Name})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (r SysMessageRepo) DeleteSysMessage(message *models.SysMessage) error {

	query := db.Data.Table("sys_messages").Where("id =?", message.ID).Delete(message)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r SysMessageRepo) SysMessageStatus(message *models.SysMessage) error {
	tx := db.Data.Where("id = ?", message.ID).Table("sys_messages").Scan(message)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	if message.Active == true {
		message.Active = false
	} else {
		message.Active = true
	}
	tx = db.Data.Where("id = ?", message.ID).Table("sys_messages").Update("active", message.Active)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r SysMessageRepo) TotalPageSysMessage(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("sys_messages").Count(&length)
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
