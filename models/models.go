package models

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	TotalPages int64       `json:"total_pages"`
	Records    interface{} `json:"records"`
}

type DbData struct {
	DSN string `yaml:"DSN"`
}

type Countries struct {
	ID     int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name   string `gorm:"column:name"`
	Icon   string `gorm:"column:icon"`
	Active bool   `gorm:"column:active;"`
}

type Languages struct {
	ID     int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name   string `gorm:"column:name"`
	Icon   string `gorm:"column:icon"`
	Active bool   `gorm:"column:active;default true"`
}

type SysMessage struct {
	ID     int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name   string `gorm:"column:name"`
	Active bool   `gorm:"column:active;default true"`
}

type Currency struct {
	ID   int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name string `gorm:"column:name;"`
	Icon string `gorm:"column:icon"`
}

type Transfers struct {
	ID       int64  `gorm:"column:id;primary_key;autoIncrement"`
	Entity   string `gorm:"column:entity"`
	EntityId int64  `gorm:"column:entity_id"`
	LangId   int64  `gorm:"column:lang_id;references:languages"`
	KeyField int64  `gorm:"column:keyfield"`
	Value    string `gorm:"column:value"`
}

type Agents struct {
	ID        int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name      string `gorm:"column:name"`
	LegalName string `gorm:"column:legal_name"`
	Active    bool   `gorm:"column:active;default true"`
}

type AccountAgent struct {
	ID         int64 `gorm:"column:id;primary_key;autoIncrement"`
	AgentId    int64 `gorm:"column:agent_id;references_agents"`
	CurrencyId int64 `gorm:"column:curr_id;references_currency"`
	Active     bool  `gorm:"column:active;default true"`
	IsDefault  bool  `gorm:"column:is_default;default true"`
	Type       int   `gorm:"column:type"`
}

type UserInfo struct {
	ID     int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name   string `gorm:"column:name"`
	Icon   string `gorm:"column:icon"`
	Active bool   `gorm:"column:active;"`
	Sort   int    `gorm:"column:sort"`
}

type Vendor struct {
	ID     int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name   string `gorm:"column:name"`
	Active bool   `gorm:"column:active;"`
}

type Services struct {
	ID       int64  `gorm:"column:id;primary_key;autoIncrement"`
	VendorId int64  `gorm:"column:vendor_id;references_vendor"`
	Name     string `gorm:"column:name"`
	Active   bool   `gorm:"column:active;"`
	Type     string `gorm:"column:type"`
}

type ServicesCountry struct {
	ServiceId int64 `gorm:"column:service_id;references_services"`
	CountryId int64 `gorm:"column:country_id;references_countries"`
	Active    bool  `gorm:"column:active;"`
}

type ServicesRules struct {
	ID   int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name string `gorm:"column:name"`
	Type string `gorm:"column:type;"`
}

type PaymentType struct {
	ID   int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name string `gorm:"column:name"`
}

/* type Register struct {
	ID         int64  `gorm:"column:id;primary_key;autoIncrement"`
	Request    string `gorm:"column:request"`
	Name       string `gorm:"column:name"`
	Icon       string `gorm:"column:icon"`
	Active     bool   `gorm:"column:active;"`
	Sort       int    `gorm:"column:sort"`
	AgentId    int64  `gorm:"column:agent_id;references_agents"`
	CurrencyId int64  `gorm:"column:curr_id;references_currency"`
	IsDefault  bool   `gorm:"column:is_default;default true"`
	Type       string `gorm:"column:type"`
	LegalName  string `gorm:"column:legal_name"`
	Entity     string `gorm:"column:entity"`
	EntityId   int64  `gorm:"column:entity_id"`
	LangId     int64  `gorm:"column:lang_id;references:languages"`
	KeyField   int64  `gorm:"column:keyfield"`
	Value      string `gorm:"column:value"`
	Status     string `gorm:"column:status"`
}
*/
// *bool - null bool
