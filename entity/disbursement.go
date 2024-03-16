package entity

type DisbursementRecord struct {
	ID                uint   `json:"id" orm:"auto;pk;column(id)"`
	TransactionID     string `json:"transaction_id"`
	FromAccount       string `json:"from_account"`
	ToAccount         string `json:"to_account"`          // Benef Account Number
	ToAccountFullname string `json:"to_account_fullname"` // Benef Name
	ToBankID          string `json:"to_bank_id"`          // Bank ID
	ToBankName        string `json:"to_bank_name"`        // Bank Name
	StatusDesc        string `json:"status_desc"`         // Status description
}

type ReqDisbursement struct {
	FromAccount       string `json:"from_account"`
	ToAccount         string `json:"to_account"`          // Benef Account number
	ToAccountFullName string `json:"to_account_fullname"` // Benef Name
	ToBankID          string `json:"to_bank_id"`          // Benef Bank ID
	ToBankName        string `json:"to_bank_name"`        // Benef Bank Name
	TrxDesc           string `json:"trx_desc"`            // Transaction Description
}

const (
	DisbursementStatusPending    string = "pending"
	DisbursementStatusOnProgress string = "on_progress"
	DisbursementStatusSuccess    string = "success"
	DisbursementStatusFailed     string = "failed"
)

type RespDisbursement struct {
	TransactionID string `json:"transaction_id"`
	StatusDesc    string `json:"status_desc"` // Status description
	Error         string `json:"error"`
}
