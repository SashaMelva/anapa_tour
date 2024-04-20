package loyaltymodel

import "github.com/SashaMelva/anapa_tour/internal/storage/model"

type Pin struct {
	Id          uint32    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Coordinat   []float64 `json:"coordinat"`
	Activ       bool      `json:"activ"`
	OrgId       int       `json:"organization_id"`
}

type PinWhithOrg struct {
	Id          uint32             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Coordinat   []float64          `json:"coordinat"`
	Activ       bool               `json:"activ"`
	Org         model.Organization `json:"organization"`
}

type PinVewFormat struct {
	Org    OrgMin    `json:"organization"`
	Action ActionMin `json:"action"`
	Pin    PinMin    `json:"pin"`
}

type OrgMin struct {
	Id   uint32 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type ActionMin struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type PinMin struct {
	Id        uint32 `json:"id"`
	Coordinat string `json:"coordinat"`
}
