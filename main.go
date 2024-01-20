package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"twittergram/twittergram"
	"twittergram/twittergram/database"

	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)

type config struct {
	TelegramToken string `env:"TELEGRAM_TOKEN" validate:"required"`
}

func main() {
	// Get Bot from environment variables (.env)
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	// Create bot
	bot, err := telego.NewBot(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize signal handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{}, 1)

	// Get updates
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Handle updates
	bh, _ := telegohandler.NewBotHandler(bot, updates)
	handler := twittergram.NewHandler(bot, bh)
	handler.RegisterHandlers()

	// Open a new SQLite database file
	if err := database.Open(); err != nil {
		log.Fatal(err)
	}

	// Define the tables
	if err := database.CreateTables(); err != nil {
		log.Fatal("Error creating table:", err)
		return
	}

	go func() {
		// Wait for stop signal
		<-sigs
		fmt.Println("Stopping...")

		bot.StopLongPolling()
		fmt.Println("Long polling done")

		bh.Stop()
		fmt.Println("Bot handler done")

		// Close the database connection
		database.Close()

		done <- struct{}{}
	}()

	go bh.Start()

	<-done
	fmt.Println("Done")
}
