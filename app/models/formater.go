package models

import (
	"github.com/shopspring/decimal"
)

type SumFormat struct {
	ID           int             `json:"id"`
	MerchantCode int             `json:"merchant_code"`
	Currency     string          `json:"currency"`
	Product      string          `json:"bet_count"`
	Amount       decimal.Decimal `json:"bet_amount"`
	Category     decimal.Decimal `json:"bet_valid"`
	Description  decimal.Decimal `json:"bet_winlose"`
}

type CreateSummaryInput struct {
	RangeDate    string `json:"range_date" binding:"required"`
	ProviderCode string `json:"provider_code" binding:"required"`
}
