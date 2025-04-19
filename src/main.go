package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reuros-api/handlers"
	"reuros-api/repositories"
	"reuros-api/services"
	"time"

	_ "reuros-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiKey := os.Getenv("EXCHANGE_RATE_API_KEY")
	if apiKey == "" {
		log.Fatal("EXCHANGE_RATE_API_KEY is not set")
	}

	currencyService := services.NewCurrencyService(apiKey)

	db, err := services.NewDatabaseConnection()
	if err != nil {
		log.Fatal("Database connection error %v", err)
	}

	userRepo := repositories.NewUserRepository(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/currency/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCurrencyRates(w, r, currencyService)
	})
	mux.HandleFunc("/users", handlers.CreateUserHandler(userRepo))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: corsMiddleware(mux),
	}

	go func() {
		log.Printf("Starting server on :%s", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop // Wait for a signal to shut down

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
