package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	_"strings"
)

func main() {
	//intellijからの起動では追加の環境変数は取得できないっぽい・・・
	//for _, e := range os.Environ() {
	//	pair := strings.Split(e, "=")
	//	fmt.Println(pair[0])
	//}

	var word string = "golang"

	//flag使う場合、Parseしないと引数認識しない
	flag.Parse()
	if flag.NArg() > 0 {
		word = flag.Arg(0)
		fmt.Println("word = " + word)
	}

	anaconda.SetConsumerKey(os.Getenv("TWITTER_API_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_API_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))

	v := url.Values{}
	v.Set("lang", "ja")
	v.Set("count", "30")

	result, _ := api.GetSearch(word + " exclude:retweets", v)
	//result, _ := api.GetSearch("golang", nil)
	for _, tweet := range result.Statuses {
		fmt.Println(tweet.Text)
	}
}
