package main

import (
	"context"
	"log"

	"github.com/sivchari/gotwtr"
)

func main() {
        bearerToken := "" // Input your Bearer Twitter Token
        client := gotwtr.New(bearerToken)

        searchTweet := "ぽよ〜"
        tsr, err := client.SearchRecentTweets(context.Background(), searchTweet, &gotwtr.SearchTweetsOption{
            TweetFields: []gotwtr.TweetField{
                gotwtr.TweetFieldAuthorID,
                gotwtr.TweetFieldAttachments,
            },
            MaxResults: 10,
        })
        if err != nil {
            log.Fatal(err)
        }

        for _, t := range tsr.Tweets {
            log.Println("----------------------------------------")
            log.Println(t.Text)
        }
}
