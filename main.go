package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
	"log"
)

const appEnvName = "crimeandtrumpishment"

type Specification struct {
	Consumerkey    string
	Consumersecret string
	Accesstoken    string
	Accesssecret   string
}

func getBookmark() []byte {
	// read val from "bookmarker_index"
}

func setBookmark(b []byte) error {
	// update bookmark after success twttering, where byte is last key
}

func getContent(b []byte) string {
	// return v from k+v where k is b as return from getBookmark
}

func setContent(sentence string) error {
	// should auto++ the key

}

// point it at a .txt of a book and run to p
func parseOriginal(f string) error {
	// reads through original text and cleverly parses sentences feeding them to setContent(i, sentece)

}

func main() {

	var s Specification
	err := envconfig.Process(appEnvName, &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	config := oauth1.NewConfig(s.Consumerkey, s.Consumersecret)
	token := oauth1.NewToken(s.Accesstoken, s.Accesssecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// break every paragraph into a bucket.
	// break every sentence into a k-v
	// for both buckets and k/v's, id's should auto++.

	// create bucket 'bookmarker', where key:'iam' -> {paragraph: pp, sentence: ss} as jsonable struct

	// do reading.
	// parse sentences. sentences are separated by a period that is not prefixed by a capital letter
	// ["K." is not a sentence.]
	// ie "They went to the church in K. and prayed."
	// replace all pronouns with Trump, where pronouns are Capitalized words.
	// replace all "K." with "T."
	// replace all "his" with "Trump's"
	// replace all "him" with "Trump"

	// TODO: de Sade

	// Send a Tweet
	client.Statuses.Update("just setting up my twttr", nil) // tweet, response, err

}
