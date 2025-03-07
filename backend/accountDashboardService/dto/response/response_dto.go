package respdto

import "time"

type AccountListPagination struct {
	TotalCount int `json:"totalCount"`
	Page       int `json:"page"`
	Count      int `json:"count"`
}

type AccountListResponse struct {
	AccountListPagination AccountListPagination `json:"accountListPagination"`
	AccountList           []Customer             `json:"accountList"`
}

type GetListAccountResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	ResponseCode string      `json:"responseCode"`
	Data         interface{} `json:"data"`
	Pagination AccountListPagination `json:"pagination"`
}

type Customer struct {
	Name         string        `json:"name"`
	BankAccounts []BankAccount `json:"bank_accounts"`
	Pockets      []Pocket      `json:"pockets"`
	TermDeposits []TermDeposit `json:"term_deposits"`
}

type BankAccount struct {
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency"`
}

type Pocket struct {
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

type TermDeposit struct {
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	InterestRate float64 `json:"interest_rate"`
	MaturityDate string  `json:"maturity_date"`
}

type VerifyTokenResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    VerifyTokenUserDetail `json:"data"`
}

type VerifyTokenUserDetail struct {
	Email  string `json:"email"`
	Role   string `json:"role"`
	ExpiredAt time.Time `json:"expired_at"`
}