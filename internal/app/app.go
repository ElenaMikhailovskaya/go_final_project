package app

import (
	_ "errors"
	"fmt"
	"github.com/ElenaMikhailovskaya/go_final_project/internal/interfaces/database"
	"github.com/ElenaMikhailovskaya/go_final_project/internal/service"
	"github.com/ElenaMikhailovskaya/go_final_project/internal/transport/http_transport"
)

func Start() {
	db, errDB := database.New()
	if errDB != nil {
		fmt.Println("Database error", errDB)
	}
	db.CreateTable()
	defer db.Close()

	logic, err := service.New(service.WithDatabase(db))
	if err != nil {
		fmt.Println("Service error", err)
		panic(err)
	}

	transport, err := http_transport.New(logic)
	if err != nil {
		fmt.Println("Transport error", err)
		panic(err)
	}

	err = transport.Listen()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
