package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/MoneyTransferAPI/entity"
	"github.com/labstack/echo"
)

func (s *Server) Check(ctx echo.Context) error {
	accInfo := entity.AccountInfo{
		AccountNumber: "12938019283",
		AccountName:   "Mahdi",
		BankID:        "1092",
		BankName:      "BCA",
	}

	err := s.UseCase.SaveAccountInfo(ctx.Request().Context(), accInfo)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) ValidateAccount(ctx echo.Context) error {
	// validate input
	bData, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	reqAccountInfo := entity.ReqInquiryAccountInfo{}
	if err := json.Unmarshal(bData, &reqAccountInfo); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	accInfo, err := s.UseCase.ValidateAccount(ctx.Request().Context(), reqAccountInfo)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, accInfo)
}

func (s *Server) Disbursement(ctx echo.Context) error {
	// validate input
	bData, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	reqDisbursement := entity.ReqDisbursement{}
	if err := json.Unmarshal(bData, &reqDisbursement); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// lock process, prevent duplicate trx

	// store disbursement log
	disburseLog, err := s.UseCase.StoreDisbursementLog(ctx.Request().Context(), reqDisbursement, entity.RespDisbursement{StatusDesc: entity.DisbursementStatusPending})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	// do disbursement
	respDisbursement, err := s.UseCase.Disbursement(ctx.Request().Context(), reqDisbursement)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	disburseLog.StatusDesc = respDisbursement.StatusDesc

	// update disbursement log
	err = s.UseCase.UpdateDisbursementLog(ctx.Request().Context(), disburseLog)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, respDisbursement)
}

func (s *Server) DisbursementCallback(ctx echo.Context) error {
	return nil
}
