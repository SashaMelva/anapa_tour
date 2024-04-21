package loyaltymodel

import "time"

type BalanceBonus struct {
	Id        uint32 `json:"id"`
	AccountId uint32 `json:"account_id"`
	Summ      int    `json:"summ"`
}

type HystoryAppnedBonus struct {
	Id          uint32
	PromotionId uint32
	IdBalance   uint32
	Date        time.Time
	Summ        int
}

type ChangeBalance struct {
	AccountId uint32 `json:"account_id"`
	Summ      int    `json:"summ"`
	Plus      bool   `json:"plus"`
}

type ChangeBalanceForOrganization struct {
	AccountId     uint32
	OrgCategoryId uint32
}
