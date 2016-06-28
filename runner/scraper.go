package main

import (
	"fmt"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Representation of a dish on a menu
// In the future we can also include information about the query from the user, including geolocation
type Dish struct {
	Name     string
	Cuisine  string
	Category string
}

// Initialize GORM. connect to the postgres db, and initialize the schema
func InitORM(auth string, shouldLog bool) *gorm.DB {
	db, err := gorm.Open("postgres", auth)
	checkError(err)
	db.LogMode(shouldLog)

	db.AutoMigrate(&Dish{})

	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
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

// Returns a list of names and urls
func scrapeCuisines(base string) ([]string, []string) {
	doc, err := goquery.NewDocument(base)
	checkError(err)
	urls, names := []string{}, []string{}

	doc.Find("#all_cuisines").Children().Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		link, _ := a.Attr("href")

		urls = append(urls, base+link)
		names = append(names, a.Text())
	})

	return names, urls
}

func scrapeRestaurants(page string, base string) []string {
	doc, err := goquery.NewDocument(page)
	checkError(err)

	citys := make([]string, 0)
	doc.Find("#restaurant_list").Children().Children().Each(func(i int, s *goquery.Selection) {
		name := s.Find("a")
		if u, ok := name.Attr("href"); ok {
			citys = append(citys, base+u)
		}
	})

	return citys
}

func scrapeMenus(page string) []Dish {
	doc, err := goquery.NewDocument(page)
	checkError(err)
	dishes := []Dish{}

	doc.Find("#menu").Find(".category").Each(func(i int, s *goquery.Selection) {
		cat := s.Find(".category_head").Find("h3").Text()

		s.Find(".name").Each(func(i int, q *goquery.Selection) {
			d := Dish{}
			d.Category = cat
			d.Name = q.Text()
			dishes = append(dishes, d)
		})
	})

	return dishes
}

func runScrape(base string) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=allmenudump sslmode=disable password=postgres")
	checkError(err)

	db.LogMode(false)
	db.AutoMigrate(&Dish{})

	cities := scrapeCities(base)

	for _, city := range cities {
		cuisineNames, cuisineUrls := scrapeCuisines(city)

		var wg sync.WaitGroup
		wg.Add(len(cuisineUrls))

		for i, cuisine := range cuisineUrls {
			go func(cuisineName string) {
				defer wg.Done()

				for _, rest := range scrapeRestaurants(cuisine, base) {
					dishes := scrapeMenus(rest)

					for _, dish := range dishes {
						dish.Cuisine = cuisineName
						// fmt.Println(dish)
						db.Save(&dish)
					}

					fmt.Printf("Added %d from %s", len(dishes), rest)
				}

			}(cuisineNames[i])
		}

		wg.Wait()
		fmt.Println("Finished city: ", city)
	}
}

func main() {
	runScrape("http://www.allmenus.com/")

	// restaurants := restaurants("http://www.allmenus.com/wi/milwaukee/-/american/", base)
	// dishes := menu("http://www.allmenus.com/wi/greenfield/101286-cousins-subs/menu/")
}
