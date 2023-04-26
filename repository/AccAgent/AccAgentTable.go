package AccAgent

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
	"log"
)

func (r AccAgentRepo) AddAccount(account *models.AccountAgent) error {
	tx := db.Data.Create(account)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *AccAgentRepo) GetAccountAgent(pagination *models.Pagination) ([]models.AccountAgent, error) {
	var AccAgents []models.AccountAgent
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("account_agents").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.AccountAgent{}).Find(&AccAgents)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return AccAgents, nil
}

func (r AccAgentRepo) UpdateAccountAgent(AccAgent *models.AccountAgent) error {

	tx := db.Data.Model(models.AccountAgent{}).Where("id = ?", AccAgent.ID).Updates(models.AccountAgent{AgentId: AccAgent.AgentId, CurrencyId: AccAgent.CurrencyId, IsDefault: AccAgent.IsDefault, Type: AccAgent.Type})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (r AccAgentRepo) DeleteAccountAgent(AccAgent *models.AccountAgent) error {

	query := db.Data.Table("account_agents").Where("id =?", AccAgent.ID).Delete(AccAgent)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r AccAgentRepo) UpdateAccountDefault(AccAgent *models.AccountAgent) error {
	tx := db.Data.Where("id = ?", AccAgent.ID).Table("account_agents").Scan(AccAgent)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	if AccAgent.IsDefault == true {
		AccAgent.IsDefault = false
	} else {
		AccAgent.IsDefault = true
	}
	tx = db.Data.Where("id = ?", AccAgent.ID).Table("account_agents").Update("is_default", AccAgent.IsDefault)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r AccAgentRepo) AccountAgentStatus(AccAgent *models.AccountAgent) error {
	tx := db.Data.Where("id = ?", AccAgent.ID).Table("account_agents").Scan(AccAgent)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	if AccAgent.Active == true {
		AccAgent.Active = false
	} else {
		AccAgent.Active = true
	}
	tx = db.Data.Where("id = ?", AccAgent.ID).Table("account_agents").Update("active", AccAgent.Active)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r AccAgentRepo) TotalPageAccount(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("account_agents").Count(&length)
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
