package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

type Credentials struct {
	ConsumerKey       string
    ConsumerSecret    string
    AccessToken       string
    AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error) {
    // Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
    config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
    // Pass in your Access Token and your Access Token Secret
    token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

    httpClient := config.Client(oauth1.NoContext, token)
    client := twitter.NewClient(httpClient)

    // Verify Credentials
    verifyParams := &twitter.AccountVerifyParams{
        SkipStatus:   twitter.Bool(true),
        IncludeEmail: twitter.Bool(true),
    }

    // we can retrieve the user and verify if the credentials
    // we have used successfully allow us to log in!
    user, _, err := client.Accounts.VerifyCredentials(verifyParams)
    if err != nil {
        return nil, err
    }

    log.Printf("User's ACCOUNT: %s\n", user.Name)
    return client, nil
}

func FromDotEnv() (*twitter.Client, error) {
	// Loading .env variables
	err := godotenv.Load(".env")
	if err != nil {
		panic("Cannot load dotenv file")
	}

	c := &Credentials{
        AccessToken:       os.Getenv("TW_ACCESS_KEY"),
        AccessTokenSecret: os.Getenv("TW_ACCESS_SECRET"),
        ConsumerKey:       os.Getenv("TW_CONSUMER_KEY"),
        ConsumerSecret:    os.Getenv("TW_CONSUMER_SECRET"),
    }

	return getClient(c)
}
