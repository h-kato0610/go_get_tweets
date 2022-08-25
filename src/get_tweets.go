package main

import (
	"context"
	"log"

	"github.com/sivchari/gotwtr"
)

func main() {
	bearerToken := "" // Input your Bearer Twitter Token
	client := gotwtr.New(bearerToken)

	searchWord := "ぽよ〜"
	tsr := getSearchPubTweets(client, searchWord)

	for _, t := range tsr.Tweets {
		log.Println("----------------------------------------")
		log.Println(t.Text)
	}
}

func getSearchPubTweets(client *gotwtr.Client, searchWord string) (tsr *gotwtr.SearchTweetsResponse) {
	maxResult := 10
	tsr, err := client.SearchRecentTweets(context.Background(), searchWord, &gotwtr.SearchTweetsOption{
		TweetFields: []gotwtr.TweetField{
			gotwtr.TweetFieldAuthorID,
			gotwtr.TweetFieldAttachments,
		},
		MaxResults: maxResult,
	})
	if err != nil {
		log.Fatal(err)
	}

	return tsr
}
