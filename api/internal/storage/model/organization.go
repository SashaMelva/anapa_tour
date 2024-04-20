package model

type Organization struct {
	Id           uint32 `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	FullName     string `json:"full_name" db:"full_name"`
	Description  string `json:"description" db:"description"`
	CategoryId   int    `json:"category_id" db:"category_id"`
	AdminId      int    `json:"admin_id" db:"admin_id"`
	LegalAddress string `json:"legal_address" db:"legal_address"`
}

type OrganizationFull struct {
	Id           uint32                `json:"id" db:"id"`
	Name         string                `json:"name" db:"name"`
	FullName     string                `json:"full_name" db:"full_name"`
	Description  string                `json:"description" db:"description"`
	CategoryData *CategoryOrganization `json:"category_organization" db:"category_organization"`
	AdminId      int                   `json:"admin_id" db:"admin_id"`
	LegalAddress string                `json:"legal_address" db:"legal_address"`
}

type CategoryOrganization struct {
	Id   uint32 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
