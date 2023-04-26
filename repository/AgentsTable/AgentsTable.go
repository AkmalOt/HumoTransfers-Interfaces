package AgentsTable

import (
	"awesomeProject2/db"
	"awesomeProject2/models"
	"log"
)

func (r AgentsTableRepo) AddAgent(agents *models.Agents) error {
	tx := db.Data.Create(agents)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *AgentsTableRepo) GetAgent(pagination *models.Pagination) ([]models.Agents, error) {
	var Agents []models.Agents
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("agents").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Agents{}).Find(&Agents)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return Agents, nil
}

func (r AgentsTableRepo) UpdateAgents(agent *models.Agents) error {

	tx := db.Data.Model(models.Agents{}).Where("id = ?", agent.ID).Updates(models.Agents{Name: agent.Name, LegalName: agent.LegalName})
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r AgentsTableRepo) DeleteAgents(agent *models.Agents) error {
	query := db.Data.Table("agents").Where("id =?", agent.ID).Delete(agent)
	if query.Error != nil {
		//log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r AgentsTableRepo) AgentStatus(agent *models.Agents) error {
	tx := db.Data.Where("id = ?", agent.ID).Table("agents").Scan(agent)
	if tx.Error != nil {
		//log.Println(tx.Error)
		return tx.Error
	}
	if agent.Active == true {
		agent.Active = false
	} else {
		agent.Active = true
	}
	tx = db.Data.Where("id = ?", agent.ID).Table("agents").Update("active", agent.Active)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r AgentsTableRepo) TotalPageAgents(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("agents").Count(&length)
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
