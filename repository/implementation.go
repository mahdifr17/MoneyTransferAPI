package repository

import (
	"context"

	"github.com/MoneyTransferAPI/entity"
)

func (r *Repository) SaveDisbursementLog(ctx context.Context, request *entity.DisbursementRecord) error {
	if request.ID != 0 { // update
		if err := r.Db.UpdateColumn("StatusDesc", request).Error; err != nil {
			return err
		}
	}

	// insert
	if err := r.Db.Create(request).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) SaveAccountInfo(ctx context.Context, request entity.AccountInfo) error {
	// Create
	if err := r.Db.Create(&request).Error; err != nil {
		return err
	}

	return nil
}
