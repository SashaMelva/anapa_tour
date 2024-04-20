package loyaltymodel

import "time"

type BalanceBonus struct {
	Id        uint32
	AccountId uint32
	Summ      int
}

type HystoryAppnedBonus struct {
	Id          uint32
	PromotionId uint32
	IdBalance   uint32
	Date        time.Time
	Summ        int
}
