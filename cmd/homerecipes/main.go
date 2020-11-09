package main

import (
	"context"
	"fmt"
	"github.com/mikeletux/home-recipes/cmd/sample-data"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"github.com/mikeletux/home-recipes/pkg/localstorage"
	"github.com/mikeletux/home-recipes/pkg/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	envPort := os.Getenv("RECIPES_PORT")
	envWithData := os.Getenv("RECIPES_SAMPLE_DATA")
	if len(envPort) == 0 {
		envPort = "8080"
	}
	var recipes map[string]*recipe.Recipe
	if envWithData == "yes" {
		log.Print("Loading sample data into server")
		recipes = data.SampleRecipes
	}
	storage := localstorage.NewLocalStorage(recipes)
	s := server.New(storage)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", envPort),
		Handler: s.Router(),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()
	log.Printf("Server started on port %s", envPort)

	<-done
	log.Print("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		//Extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Print("Server exited propertly")

}
