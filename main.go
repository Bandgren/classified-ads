package main

import (
	"context"
	"github.com/bandgren/classified-ads/database"
	"github.com/bandgren/classified-ads/http"
	"github.com/bandgren/classified-ads/servicemanager"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}
}

func main() {
	_, cancel := start()
	defer shutdown(cancel)
	servicemanager.WaitShutdown()
}

func start() (ctx context.Context, cancel context.CancelFunc) {
	// This is the main context for the service. When it is canceled it means the service is going down.
	// All the tasks must be canceled
	ctx, cancel = context.WithCancel(context.Background())
	if err := database.Start; err != nil {
		log.Printf("couldnt start database error [%s]\n", err)
	}
	http.Start()
	return
}

func shutdown(cancel context.CancelFunc) {
	cancel()
	ctx, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelTimeout()
	doneHTTP := http.Shutdown(ctx)
	err := servicemanager.WaitUntilIsDoneOrCanceled(ctx, doneHTTP)
	if err != nil {
		log.Printf("service stopped by timeout %s\n", err)
	}
	time.Sleep(time.Millisecond * 200)
	log.Println("bye bye")
}
