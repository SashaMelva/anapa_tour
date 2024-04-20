package loyaltymodel

import "time"

type Action struct {
	Id          uint32     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CategoryId  int        `json:"category_id"`
	Activ       bool       `json:"activ"`
	DateStart   *time.Time `json:"date_start"`
	DateEnd     *time.Time `json:"date_end"`
	OrgId       int        `json:"organization_id"`
}
