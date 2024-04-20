package loyaltymodel

type Promotion struct {
	Id                     uint32 `json:"id"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Activ                  bool   `json:"activ"`
	Сost                   int    `json:"cost"`
	IdCategoryOrganization int    `json:"id_category_organization"`
}
