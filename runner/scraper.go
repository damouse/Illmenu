package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func scrapeCities(base string) []string {
	doc, err := goquery.NewDocument(base)
	checkError(err)

	citys := make([]string, 0)

	doc.Find("#nav_links").Children().Children().Each(func(i int, s *goquery.Selection) {
		if u, ok := s.Children().First().Attr("href"); ok {
			citys = append(citys, base+u)
		}
	})

	return citys
}

func scrapeCuisines(base string) []string {
	doc, err := goquery.NewDocument(base)
	checkError(err)

	citys := make([]string, 0)

	doc.Find("#all_cuisines").Children().Each(func(i int, s *goquery.Selection) {
		if u, ok := s.Children().First().Attr("href"); ok {
			citys = append(citys, base+u)
		}
	})

	return citys
}

func restaurants(page string, base string) []string {
	doc, err := goquery.NewDocument(page)
	checkError(err)

	citys := make([]string, 0)
	doc.Find("#restaurant_list").Children().Children().Each(func(i int, s *goquery.Selection) {
		name := s.Find("a")
		// fmt.Println(name.Text)

		if u, ok := name.Attr("href"); ok {
			citys = append(citys, base+u)
		}
	})

	return citys
}

func menu(page string) []string {
	doc, err := goquery.NewDocument(page)
	checkError(err)
	citys := make([]string, 0)

	doc.Find("#menu").Find(".category").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".category_head").Find("h3").Text()
		fmt.Println(name)

		s.Find(".name").Each(func(i int, q *goquery.Selection) {
			fmt.Println("\t", q.Text())
		})
	})

	return citys
}

// Representation of a dish on a menu
type dish struct {
	name         string
	images       []string
	lastSearched string
	createdAt

	// Metadata
	hits    int
	scrapes int
}

func main() {
	// base := "http://www.allmenus.com/"

	// cities := scrapeCities(base)
	// cuisines := scrapeCuisines("http://www.allmenus.com/wi/milwaukee/")
	// restaurants := restaurants("http://www.allmenus.com/wi/milwaukee/-/american/", base)

	dishes := menu("http://www.allmenus.com/wi/greenfield/101286-cousins-subs/menu/")
	fmt.Println(dishes)
}
