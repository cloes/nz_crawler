package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Course stores information about a coursera course
type Course struct {
	Title       string
	Description string
	Creator     string
	Level       string
	URL         string
	Language    string
	Commitment  string
	HowToPass   string
	Rating      string
}

/*
func handelfunc(e *colly.HTMLElement){
	//fmt.Println("found!")

	fmt.Print(e.ChildText("label"))
	e.DOM.Find("label").Remove()
	fmt.Println(strings.TrimSpace(e.DOM.Text()))

}
*/

func handelCompanyNamefunc(e *colly.HTMLElement) {
	key := strings.TrimRight(e.ChildText("label.SCR011_04_003"), ":")

	e.DOM.Find("label.SCR011_04_003").Remove()
	value := strings.Trim(e.DOM.Text(), "\n")

	fmt.Println(key, value)
}

func handelNZBNfunc(e *colly.HTMLElement) {

	fmt.Println("found2!")
	fmt.Println(e.Text)
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("app.companiesoffice.govt.nz"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./coursera_cache"),
		colly.UserAgent("None"),
	)
	// 设置超时时间为3秒
	c.SetRequestTimeout(3 * time.Second)

	// Create another collector to scrape course details
	//detailCollector := c.Clone()

	//courses := make([]Course, 0, 200)

	// On every a element which has href attribute call callback
	/*
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			// If attribute class is this long string return from callback
			// As this a is irrelevant
			if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
				return
			}
			link := e.Attr("href")
			// If link start with browse or includes either signup or login return from callback
			if !strings.HasPrefix(link, "/browse") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
				return
			}
			// start scaping the page under the link found
			e.Request.Visit(link)
		})

	*/

	c.OnHTML("div.readonly.companySummary > div:first-child", handelCompanyNamefunc)
	c.OnHTML("div.readonly.companySummary > div:nth-child(2)", handelNZBNfunc)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// On every a HTML element which has name attribute call callback
	/*
		c.OnHTML(`a[name]`, func(e *colly.HTMLElement) {
			// Activate detailCollector if the link contains "coursera.org/learn"
			courseURL := e.Request.AbsoluteURL(e.Attr("href"))
			if strings.Index(courseURL, "coursera.org/learn") != -1 {
				detailCollector.Visit(courseURL)
			}
		})
	*/

	// Extract details of the course
	/*
		detailCollector.OnHTML(`div[id=rendered-content]`, func(e *colly.HTMLElement) {
			log.Println("Course found", e.Request.URL)
			title := e.ChildText(".course-title")
			if title == "" {
				log.Println("No title found", e.Request.URL)
			}
			course := Course{
				Title:       title,
				URL:         e.Request.URL.String(),
				Description: e.ChildText("div.content"),
				Creator:     e.ChildText("div.creator-names > span"),
			}
			// Iterate over rows of the table which contains different information
			// about the course
			e.ForEach("table.basic-info-table tr", func(_ int, el *colly.HTMLElement) {
				switch el.ChildText("td:first-child") {
				case "Language":
					course.Language = el.ChildText("td:nth-child(2)")
				case "Level":
					course.Level = el.ChildText("td:nth-child(2)")
				case "Commitment":
					course.Commitment = el.ChildText("td:nth-child(2)")
				case "How To Pass":
					course.HowToPass = el.ChildText("td:nth-child(2)")
				case "User Ratings":
					course.Rating = el.ChildText("td:nth-child(2) div:nth-of-type(2)")
				}
			})
			courses = append(courses, course)
		})
	*/

	/*
		for companyNumber :=0; companyNumber < 10000000; companyNumber++  {
			//https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1908322
			//http://app.companiesoffice.govt.nz/co/1908322

			// Start scraping on http://coursera.com/browse
			c.Visit("https://app.companiesoffice.govt.nz")
		}
	*/

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		//fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	start := time.Now()
	for companyNumber := 1908322; companyNumber < 1908322+1; companyNumber++ {
		//c.Visit("https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1908322")
		//https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1908322/detail
		//fmt.Println(companyNumber)
		c.Visit("https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/" + strconv.Itoa(companyNumber) + "/detail")
	}
	end := time.Since(start)
	fmt.Println(end)

	//c.Visit("https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1908322") //OK

	//enc := json.NewEncoder(os.Stdout)
	//enc.SetIndent("", "  ")

	// Dump json to the standard output
	//enc.Encode(courses)
}
