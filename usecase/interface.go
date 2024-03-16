package usecase

import (
	"context"

	"github.com/MoneyTransferAPI/entity"
)

type UseCaseInterface interface {
	ValidateAccount(ctx context.Context, reqData entity.ReqInquiryAccountInfo) (accInfo entity.AccountInfo, err error)
	Disbursement(ctx context.Context, reqData entity.ReqDisbursement) (resp entity.RespDisbursement, err error)
	DisbursementCallback(ctx context.Context, reqData entity.RespDisbursement) error
	StoreDisbursementLog(ctx context.Context, reqData entity.ReqDisbursement, respData entity.RespDisbursement) (*entity.DisbursementRecord, error)
	UpdateDisbursementLog(ctx context.Context, recordData *entity.DisbursementRecord) error

	SaveAccountInfo(ctx context.Context, request entity.AccountInfo) error
}
