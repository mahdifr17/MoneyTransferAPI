package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MoneyTransferAPI/entity"
	"github.com/MoneyTransferAPI/repository"
)

const (
	mockUsername    = "mahdi"
	mockPassword    = "firdaus"
	mockMerchantKey = "3JE1FjrpaL0BI96P"
	mockUrl         = "https://c7b0adcc-1a0f-4fc6-977f-9c37c54ed3ba.mock.pstmn.io"
)

type UseCase struct {
	Repository repository.RepositoryInterface
}

func NewUseCase(repository repository.RepositoryInterface) *UseCase {
	return &UseCase{Repository: repository}
}

func (uc *UseCase) prepareRequest(url, method string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil
	}
	req.SetBasicAuth(mockUsername, mockPassword)
	req.Header.Add("Merchant-Key", mockMerchantKey)
	req.Header.Add("Content-Type", "application/json")

	return req
}

func (uc *UseCase) ValidateAccount(ctx context.Context, reqData entity.ReqInquiryAccountInfo) (accInfo entity.AccountInfo, err error) {
	// prepare payload
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(reqData)
	if err != nil {
		return entity.AccountInfo{}, err
	}

	// prepare request
	reqClient := http.Client{}
	req := uc.prepareRequest(fmt.Sprintf("%s/inquiry-account", mockUrl), http.MethodPost, &buf)

	// do request
	resp, err := reqClient.Do(req)
	if err != nil {
		return entity.AccountInfo{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return entity.AccountInfo{}, err
	}

	// read request
	bRespBody, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(bRespBody, &accInfo); err != nil {
		return entity.AccountInfo{}, err
	}

	return accInfo, nil
}

func (uc *UseCase) Disbursement(ctx context.Context, reqData entity.ReqDisbursement) (respDisburse entity.RespDisbursement, err error) {
	// prepare payload
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(reqData)
	if err != nil {
		return entity.RespDisbursement{}, err
	}

	// prepare request
	reqClient := http.Client{}
	req := uc.prepareRequest(fmt.Sprintf("%s/disbursement", mockUrl), http.MethodPost, &buf)

	// do request
	resp, err := reqClient.Do(req)
	if err != nil {
		return entity.RespDisbursement{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return entity.RespDisbursement{}, err
	}

	// read request
	bRespBody, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(bRespBody, &respDisburse); err != nil {
		return entity.RespDisbursement{}, err
	}

	return respDisburse, nil
}

func (uc *UseCase) DisbursementCallback(ctx context.Context, reqData entity.RespDisbursement) error {
	return nil
}

func (uc *UseCase) StoreDisbursementLog(ctx context.Context, reqData entity.ReqDisbursement, respData entity.RespDisbursement) (disburseLog *entity.DisbursementRecord, err error) {
	disburseLog = &entity.DisbursementRecord{
		FromAccount:       reqData.FromAccount,
		ToAccount:         reqData.ToAccount,
		ToAccountFullname: reqData.ToAccountFullName,
		ToBankID:          reqData.ToBankID,
		ToBankName:        reqData.ToBankName,
		StatusDesc:        respData.StatusDesc,
	}
	err = uc.Repository.SaveDisbursementLog(ctx, disburseLog)
	return disburseLog, err
}

func (uc *UseCase) UpdateDisbursementLog(ctx context.Context, recordData *entity.DisbursementRecord) error {
	return uc.Repository.SaveDisbursementLog(ctx, recordData)
}

func (uc *UseCase) SaveAccountInfo(ctx context.Context, reqData entity.AccountInfo) error {
	return uc.Repository.SaveAccountInfo(ctx, reqData)
}
