package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Shop struct {
	Name     string
	manNum   int
	womenNum int
}

func GetAiseki(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalf("error!!")
	}

	doc.Find(".congestion_home .con_shop").Each(func(_ int, s *goquery.Selection) {
		name := s.Find(".shopname  p").Text()
		man := s.Find(".cong_man > span").Text()
		woman := s.Find(".cong_woman > span").Text()
		fmt.Println(name, man, woman)
	})
}

func GetJis(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalf("error!!")
	}

	script := doc.Find("script").Text()
	exp := regexp.MustCompile(`datas = (.*)}}}`)
	data := strings.Replace(exp.FindString(script), "datas = ", "", 1)

	fmt.Println(data)
}

func GetOriental(url string) {
	fmt.Println("Oriental start")
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalf("error!!")
	}

	doc.Find("#contents1000621 .shop_list .shop").Each(func(_ int, s *goquery.Selection) {
		name := s.Find(".shop_name span").Text()
		man := s.Find(".num .man").Text()
		woman := s.Find(".num .woman").Text()
		fmt.Println(name, man, woman)
	})
}

func main() {
	url := "https://aiseki-ya.com/"
	GetAiseki(url)

	//url2 := "https://jis.bar/"
	//GetJis(url2)

	// url3 := "https://oriental-lounge.com/"
	// GetOriental(url3)

}
