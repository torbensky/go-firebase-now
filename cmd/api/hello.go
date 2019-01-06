package main

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var cred64 = os.Getenv("FIREBASE_ADMIN_CREDENTIAL")
var databaseURL = os.Getenv("FIREBASE_DB_URL")

func initFirebase() (*firebase.App, error) {
	dec, err := b64.StdEncoding.DecodeString(cred64)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(dec)
	return firebase.NewApp(context.Background(), nil, opt)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	_, err := initFirebase()
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	fmt.Fprintf(w, "Hello from Go on Now 2.0!")
}
