package mysql

type CasbinRule struct {
	PType  string `json:"p_type"  gorm:"column:p_type"`
	RoleId string `json:"role_id" gorm:"column:v0"`
	Api    string `json:"api"     gorm:"column:v1"`
	Method string `json:"method"  gorm:"column:v2"`
}

func (c *CasbinRule) TableName() string {
	return "casbin_rule"
}
