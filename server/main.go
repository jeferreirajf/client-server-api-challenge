package main

import (
	"github.com/jeferreirajf/client-server-api-challenge/server/repository"
	"github.com/jeferreirajf/client-server-api-challenge/server/server"
	route "github.com/jeferreirajf/client-server-api-challenge/server/server/routes"
	services "github.com/jeferreirajf/client-server-api-challenge/server/services/request-usd-quotation-service/implementation"
	usecase "github.com/jeferreirajf/client-server-api-challenge/server/usecase/request-quotation"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	sqliteDsn := "dev.db"
	db, err := gorm.Open(sqlite.Open(sqliteDsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	aService := services.NewRequestUsdQuotationService()
	aRepository := repository.NewQuotationRepository(*db)
	anUsecase := usecase.NewRequestQuotationUsecase(aRepository, aService)

	aRoute := route.NewRequestUsdQuotationRoute(*anUsecase)

	aServer := server.NewServer(8080)

	aServer.AddRoute("/cotacao", aRoute.ServeHTTP)

	aServer.Start()
}
