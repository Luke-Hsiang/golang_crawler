package main

import (
	"encoding/json"
	"fmt"
	"golang_crawler/apps/news"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var task = func() {
	fmt.Printf("colly start at %v\n", time.Now())
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	c.OnResponse(func(r *colly.Response) {
		dsn := os.Getenv("DATABASE_URL")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Printf("error open db: %v\n", err)
			return
		}

		var jsonMap map[string]interface{}
		json.Unmarshal(r.Body, &jsonMap)
		articles := jsonMap["articles"].([]interface{})
		for _, a := range articles {
			article := a.(map[string]interface{})
			author := article["author"].(string)
			title := article["title"].(string)
			desc := article["description"].(string)

			var reporter news.Reporter
			var newNews news.News

			db.FirstOrCreate(&reporter, news.Reporter{Name: author})
			db.Where(news.News{Title: title}).Attrs(news.News{Reporter: reporter, Content: desc}).FirstOrCreate(&newNews)
		}
	})
	c.OnError(func(rp *colly.Response, e error) {
		fmt.Printf("OnError, rp.Body: %v, e.Error: %v\n", string(rp.Body), e.Error())
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})
	yesterday := time.Now().AddDate(0, 0, -1)
	from := yesterday.Format("2006-01-02")
	apiKey := os.Getenv("NEWS_API_KEY")
	domains := "udn.com,tw.news.yahoo.com,ltn.com.tw,www.ettoday.net,news.google.com,www.appledaily.com.tw,www.chinatimes.com"
	lang := "zh"
	newsAPI := fmt.Sprintf("https://newsapi.org/v2/everything?from=%v&sortBy=publishedAt&apiKey=%v&domains=%v&language=%v", from, apiKey, domains, lang)
	c.Visit(newsAPI)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("error open db: %v\n", err)
		return
	}

	var s *gocron.Scheduler = gocron.NewScheduler(time.UTC)
	job, _ := s.Every(1).Hour().Do(task)
	s.StartAsync()
	fmt.Printf("job scheduled time: %v\n", job.ScheduledTime())

	r := gin.Default()
	v1 := r.Group("/v1")
	news.Init(db, v1)
	r.Run()
}
