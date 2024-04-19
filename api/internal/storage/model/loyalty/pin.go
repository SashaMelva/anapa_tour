package loyaltymodel

type Pin struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Coordinat   string `json:"coordinat"`
	Activ       bool   `json:"activ"`
}
