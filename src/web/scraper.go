package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type TargetPage interface {
	GetPage(key string)
}

type yahooNewsComment struct {
	url string
}
func (site *yahooNewsComment) GetPage(key string) {
	doc, _ := goquery.NewDocument(site.url + key)
	doc.Find("span.cmtBody").Each(func(_ int, s *goquery.Selection) {
		//link, _ := s.Attr("href")
		text := s.Text()
		fmt.Println(text)
	})
}
type ruijiManga struct {
	url string
}
func (site *ruijiManga) GetPage(key string) {
	doc, _ := goquery.NewDocument(site.url + key)
	doc.Find("div.sm_one > h2 > a").Each(func(_ int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		fmt.Println(text)
	})
}

func NewYahoo() TargetPage {
	return &yahooNewsComment{url: "http://plugin.news.yahoo.co.jp/v1/comment/full/?"}
}
func NewRuiji() TargetPage {
	return &ruijiManga{url: "http://ruijianime.com/comic/ruiji/ruiji.php?"}
}
