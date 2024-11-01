package main

import (
	"context"
	"log"

	consts "github.com/MGomed/chat_server/consts"
	app "github.com/MGomed/chat_server/internal/app"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), consts.ContextTimeout)
	defer cancel()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("couldn't create app: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("couldn't run app: %v", err)
	}
}
