package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/deev/data/backend/internal/handler"
	"github.com/deev/data/backend/internal/repository"
	"github.com/deev/data/backend/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()

	if err := godotenv.Load(); err != nil {
    	log.Println("Файл .env не найден, используются переменные окружения")
    }

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("db_host"),
		os.Getenv("db_port"),
		os.Getenv("db_user"),
		os.Getenv("db_pass"),
		os.Getenv("db_name"),
	)

	dsn, err := sqlx.ConnectContext(ctx,"postgres",connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepo(dsn)
	srv := service.NewService(repo)
	handle := handler.NewHandler(srv)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://instagram-orpin-one.vercel.app"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Group(func(public chi.Router) {
		public.Post("/api/login",handle.Login)
	})
	r.Group(func(private chi.Router) {
		private.Post("/api/email",handle.SendMsg)
	})

	server := &http.Server{
		Addr: ":8080",
		MaxHeaderBytes: 1 << 20,
		Handler: r,
	}

	go func () {
		slog.Info("server is running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	} ()

	quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit 

    if err := server.Shutdown(ctx); err != nil {
        slog.Error("server shutdown error")
    }
    
    slog.Info("server stopped gracefully")
	
}
