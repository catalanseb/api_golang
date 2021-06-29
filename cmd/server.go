package main

import (
	flag "api-go/cmd/config"
	httprouter "api-go/cmd/handler"
	"api-go/internal/database"
	"api-go/internal/repositories"
	users "api-go/internal/services"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {

	configFlag := flag.NewFlagConfig()

	// connect to database
	db, err := database.InitDatabaseConnection()
	if err != nil {
		log.Fatalln(err)
	}
	// the schemes are created
	db.Exec(database.Schema)

	defer db.Close()

	e := echo.New()
	e.Debug = true

	// repositories
	usersRepository := repositories.NewUserRepository(db)
	// services
	usersServices := users.NewServiceUsers(usersRepository)
	// routes and handlers
	httprouter.NewUsersHandler(e, usersServices)

	go func() {
		e.Logger.Fatal(e.Start(":" + configFlag.HTTPPort))
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Printf("Recieved terminate, graceful shutdown %s \n", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	e.Shutdown(tc)
}
