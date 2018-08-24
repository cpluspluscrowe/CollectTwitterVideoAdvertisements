package main

import (
	"net/http"
	"os"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
	"fmt"
)

func main() {
	client := getTwitterClient()
	var userId int64 = 53514896
	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{UserID: userId})
	if err != nil {
		panic(err)
	}
	for _, tweet := range tweets {
		fmt.Println(tweet.Entities.Media)
	}

}
func getClient() *http.Client {
	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")
	accessToken := os.Getenv("accessToken")
	accessSecret := os.Getenv("accessSecret")
	fmt.Println(consumerKey,consumerSecret,accessToken,accessSecret)
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return httpClient
}

func getTwitterClient() *twitter.Client {
	httpClient := getClient()
	client := twitter.NewClient(httpClient)
	return client
}
