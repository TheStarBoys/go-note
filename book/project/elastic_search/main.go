package main

import (
	elastic "gopkg.in/olivere/elastic.v2"
	"fmt"
)

type Tweet struct {
	User string
	Message string
}

func main() {
	url := "http://localhost:9200"
	client, err := elastic.NewClient(elastic.SetSniff(false,), elastic.SetURL(url))
	if err != nil {
		panic(fmt.Sprintf("connect es err: %s", err))
	}

	fmt.Println("conn elastic search success")

	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do()
	if err != nil {
		panic(err)
	}

	fmt.Println("insert success")
}
