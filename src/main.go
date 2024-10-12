package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"portfolio-tracker/controller"
	"portfolio-tracker/repository"
	"portfolio-tracker/service"
	"syscall"
	"time"
)
import "github.com/gin-gonic/gin"

func main() {
	log.Println("Starting server...")

	db, err := initDb()
	if err != nil {
		log.Fatalf("Unable to initialize database: %v\n", err)
	}

	router := gin.Default()
	transactionRepository := repository.NewTransactionRepository(db.DB)
	transactionService := service.NewTransactionSerivce(transactionRepository)

	controller.NewController(&controller.Config{
		R:                  router,
		TransactionService: transactionService,
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if err := db.close(); err != nil {
		log.Fatalf("A problem occured gracefully shutting down the database connection: %v\n", err)
	}
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}

	//catching ctx.Done()
	<-ctx.Done()
	log.Println("Timeout of 5 seconds")
	log.Println("Server exiting")
}
