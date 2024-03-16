package entity

type AccountInfo struct {
	ID            uint   `json:"id" orm:"auto;pk;column(id)"`
	AccountNumber string `json:"account_number" orm:"column(account_number);index"`
	AccountName   string `json:"account_name" orm:"column(account_name)"`
	BankID        string `json:"bank_id" orm:"column(bank_id)"` /* clearing code */
	BankName      string `json:"bank_name" orm:"column(bank_name)"`
}

func (acc *AccountInfo) TableName() string {
	return "account_info"
}

type ReqInquiryAccountInfo struct {
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
}
