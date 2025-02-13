package main

import (
	"context"
	"full_stack_application/controller"
	"full_stack_application/repository"
	"full_stack_application/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server ...")

	db, err := initDB()

	if err != nil {
		log.Fatalf("Unable to initialize database: %v\n", err)
	}

	router := gin.Default()

	transactionRepository := repository.NewTransactionRepository(db.DB)
	transactionService := service.NewTransactionService(transactionRepository)

	controller.NewController(&controller.Config{
		R:                  router,
		TransactionService: transactionService,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if err := db.close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down the database connection: %v\n", err)
	}

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shotdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
