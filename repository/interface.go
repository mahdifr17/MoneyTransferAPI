package repository

import (
	"context"

	"github.com/MoneyTransferAPI/entity"
)

type RepositoryInterface interface {
	SaveDisbursementLog(ctx context.Context, request *entity.DisbursementRecord) error

	SaveAccountInfo(ctx context.Context, request entity.AccountInfo) error
}
