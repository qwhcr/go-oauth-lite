package firebasewrapper

import (
	"log"

	configUtil "go-oauth-lite/util/config"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App = nil
var firestoreClient *firestore.Client = nil
var authClient *auth.Client = nil

func setupApp(ctx context.Context) {
	opt := option.WithCredentialsFile(configUtil.GetConfig().FirebaseSecretURL)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	firebaseApp = app
}

func GetFirebaseAuthClient() *auth.Client {
	if authClient != nil {
		return authClient
	}

	ctx := context.Background()

	if firebaseApp == nil {
		setupApp(ctx)
	}

	var err error
	authClient, err = firebaseApp.Auth(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	return authClient
}

func GetFirestoreClient() *firestore.Client {
	if firestoreClient != nil {
		return firestoreClient
	}

	ctx := context.Background()

	if firebaseApp == nil {
		setupApp(ctx)
	}

	var err error
	firestoreClient, err = firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return firestoreClient
}
