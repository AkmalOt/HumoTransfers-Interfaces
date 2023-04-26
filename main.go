package main

import (
	"awesomeProject2/Handler"
	"awesomeProject2/db"
	"awesomeProject2/repository"
	"awesomeProject2/repository/Country"
	"awesomeProject2/repository/Language"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	route := gin.Default()
	DB := db.SetupGorm()

	repoLayer := &repository.Repository{
		CountryRepo:  Country.CountryRepo{Connection: DB},
		LanguageRepo: Language.LanguageRepo{Connection: DB},
	}
	newHandler := Handler.NewHandler(route, repoLayer)
	//newHandler := routs.NewHandler(route, handler)
	newHandler.Init()

	err := route.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
