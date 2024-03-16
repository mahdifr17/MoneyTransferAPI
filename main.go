package main

import (
	"github.com/MoneyTransferAPI/entity"
	"github.com/MoneyTransferAPI/handler"
	"github.com/MoneyTransferAPI/repository"
	"github.com/MoneyTransferAPI/usecase"
	"github.com/labstack/echo"
)

// const (
// 	projectName = "money-transfer-api"
// )

// func main() {
// 	// Drop table and re-create.
// 	force, _ := strconv.ParseBool(beego.AppConfig.String("dbforce"))
// 	debug, _ := strconv.ParseBool(beego.AppConfig.String("dbdebug"))

// 	// Print log.
// 	dbDriver := beego.AppConfig.String("dbdriver")
// 	dbString := beego.AppConfig.String("dbstring")
// 	repository.StartDB("default", dbDriver, dbString, force, debug)

//		beego.Run()
//	}
const (
	dbDsn string = "postgres://postgres:cicilaja@localhost:5432/moneytransfer?sslmode=disable"
)

func main() {
	e := echo.New()
	server := newServer()

	e.GET("/check", server.Check)
	e.POST("/validate-account", server.ValidateAccount)
	e.POST("/disbursement", server.Disbursement)
	e.POST("/disbursement/callback", server.DisbursementCallback)

	e.Logger.Fatal(e.Start(":8080"))
}

func newServer() *handler.Server {
	// dbDsn := os.Getenv("DATABASE_URL")
	repo := repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	repo.Db.AutoMigrate(&entity.AccountInfo{})
	repo.Db.AutoMigrate(&entity.DisbursementRecord{})

	usecase := usecase.NewUseCase(repo)
	return handler.NewServer(usecase)
}
