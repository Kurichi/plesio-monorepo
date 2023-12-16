package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/Kurichi/plesio-monorepo/services/user/pkg/config"
	"google.golang.org/api/option"
)

func InitializeFirebaseApp() *firebase.App {
	ctx := context.Background()
	cfg := config.New()
	opt := option.WithCredentialsFile(cfg.FirebaseAuthCredentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}
	return app
}

func InitializeFBAuthClient(app *firebase.App) *auth.Client {
	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing firebase auth client: %v", err)
	}
	return client
}
