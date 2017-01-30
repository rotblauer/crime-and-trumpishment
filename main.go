package main

import (
	"flag"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/neurosnap/sentences.v1/english"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Specification struct {
	Consumerkey    string
	Consumersecret string
	Accesstoken    string
	Accesssecret   string
}

func getBookmark(where string) int {
	b, err := ioutil.ReadFile(where)
	if err != nil {
		panic(err)
	}
	i, _ := strconv.Atoi(string(b))
	return i
}

func setBookmark(where string, i int) error {
	b := []byte(strconv.Itoa(i))
	err := ioutil.WriteFile(where, b, 0644)
	if err != nil {
		panic(err)
	}
	return err
}

func trumpize(s string) string {

	var trumpableWords = map[string]string{
		" his ": " Trump's ",
		" him ": " Trump ",
		" he ":  " Trump ",
		"His ":  "Trump's ",
		"Him ":  "Trump ",
		"He ":   "Trump ",
	}

	for word, trumpizedWord := range trumpableWords {
		s = strings.Replace(s, word, trumpizedWord, -1)
	}

	return s
}

func formatSentence(s string) string {
	var t = strings.TrimSpace(s)
	if len(t) > 140 {
		t = t[0:139]
		var ta = strings.Split(t, " ")
		ta = ta[0 : len(ta)-2]
		t = strings.Join(ta, " ")
		t = t + "!"
	}
	return t
}

// point it at a .txt of a book and run to p
func getOriginalSentence(i int, f string) (string, error) {
	// reads through original text and cleverly parses sentences feeding them to setContent(i, sentece)

	// can we really read all of crime&punishment into memory? we'll find out...
	content, err := ioutil.ReadFile(f)
	if err != nil {
		//Do something
	}

	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	sentences := tokenizer.Tokenize(string(content))

	return sentences[i].Text, err
}

func main() {

	var appenvname string
	var bookpath string
	var bookmarkpath string
	flag.StringVar(&appenvname, "env-prefix", "trumper", "prefix_ for exported twitter api app config vars, like TRUMPER_CONSUMERKEY and TRUMPER_CONSUMERSECRET")
	flag.StringVar(&bookpath, "book", "crimeandpunishment.txt", "absolute (preferably) path to the book you want trweempyize")
	flag.StringVar(&bookmarkpath, "bookmark", "bookmark.txt", "absolute (preferable) path to a file to hold your bookmark index")
	flag.Parse()

	// export TRUMPER_CONSUMERKEY=asdfasdfasdfasdfawerwer890232
	// export TRUMPER_CONSUMERSECRET=asdfasfq2345qwefasdfasdfqw345q3tae
	// export TRUMPER_ACCESSTOKEN=aslfjq345ouqp8ufp89u3495r8awiejpqo348urtpaoisjdfasjf
	// export TRUMPER_ACCESSSECRET=a;alskjdf;aljsf;laksjfpaoiweurpqo384urpa8sdfasdfasdf
	// * note that you can set the TRUMPER_ prefix as a flag
	var s Specification
	err := envconfig.Process(appenvname, &s)
	if err != nil {
		fmt.Println(err.Error())
	}

	// env vars set twitter authors (get it?)
	config := oauth1.NewConfig(s.Consumerkey, s.Consumersecret)
	token := oauth1.NewToken(s.Accesstoken, s.Accesssecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Get bookmark
	bm := getBookmark(bookmarkpath)

	currentSentence, e := getOriginalSentence(bm, bookpath)
	if e != nil {
		fmt.Println("shit e", e)
	}

	trumpizedSentence := trumpize(currentSentence)
	formattedTrumpizedSentence := formatSentence(trumpizedSentence)

	// Send a Tweet
	_, _, err = client.Statuses.Update(formattedTrumpizedSentence, nil) // tweet, response, err
	if err != nil {
		fmt.Println("shit err tweetin trumpy", err)
	} else {
		setBookmark(bookmarkpath, bm+1)
		fmt.Println(bm, time.Now(), formattedTrumpizedSentence)
	}

	// TODO: de Sade
}
