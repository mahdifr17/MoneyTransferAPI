package handler

import "github.com/MoneyTransferAPI/usecase"

type Server struct {
	UseCase usecase.UseCaseInterface
}

func NewServer(usecase usecase.UseCaseInterface) *Server {
	return &Server{UseCase: usecase}
}
