package crawldata

import (
	"CrawlS/dbconnect"
	"CrawlS/entity"
	"fmt"
	"log"
	"strconv"

	
	"github.com/gocolly/colly"
)

func Crawl() {
	db, err := dbconnect.DBConnect("crawldata") //mo connection
	if err != nil {                             //bat loi
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")

	db.DropTableIfExists(&entity.Questions{})
	db.CreateTable(&entity.Questions{})
	db.HasTable(&entity.Questions{})

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnHTML(".s-post-summary", func(e *colly.HTMLElement) {
		data := entity.Questions{}
		data.ID, _ = strconv.Atoi(e.Attr("data-post-id"))
		data.Author = e.ChildText(".s-user-card--link .flex--item")
		data.Votes = e.ChildText(".s-post-summary--stats-item:nth-child(1) .s-post-summary--stats-item-number")
		data.Answers = e.ChildText(".s-post-summary--stats-item:nth-child(2) .s-post-summary--stats-item-number")
		data.Views = e.ChildText(".s-post-summary--stats-item:nth-child(3) .s-post-summary--stats-item-number")
		data.Title = e.ChildText(".s-link")
		data.Link = "https://stackoverflow.com/" + e.Attr("href")
		// datas = append(datas, data)
		db.Create(&data)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	for i := 1; i < 6; i++ {
		fullURL := fmt.Sprintf("https://stackoverflow.com/questions/tagged/ibm-blockchain?tab=newest&page=%d&pagesize=50", i)
		c.Visit(fullURL)
	}
	// file, _ := json.MarshalIndent(datas, "", " ")

	// _ = ioutil.WriteFile("test.json", file, 0644)
}
