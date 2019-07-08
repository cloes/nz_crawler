package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Shareholder struct {
	Name    string
	Address string
}

type Allocation struct {
	Percentage   float64
	Shareholders []Shareholder
}

type Director struct {
	FullLegalName      string
	ResidentialAddress string
	AppointmentDate    string
}

type PageData struct {
	CompanyNumber           string
	CompanyName             string
	NZBN                    string
	IncorporationDate       string
	CompanyStatus           string
	EntityType              string
	ConstitutionFiled       string
	RegisteredOfficeAddress string
	AddressforService       string
	AddressforShareRegister string
	Directors               []Director
	ShareholderAllocations  []Allocation
}

/*
func handelfunc(e *colly.HTMLElement){
	//fmt.Println("found!")

	fmt.Print(e.ChildText("label"))
	e.DOM.Find("label").Remove()
	fmt.Println(strings.TrimSpace(e.DOM.Text()))

}
*/

var Data = new(PageData)

func handelCompanySummaryfunc(e *colly.HTMLElement) {
	e.DOM.Find("div.readonly.companySummary > div.row:nth-child(1) > label.SCR011_04_003").Remove()
	CompanyNumber := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(1)").Text()
	CompanyNumber = strings.Trim(CompanyNumber, "\n")
	Data.CompanyNumber = CompanyNumber
	println(Data.CompanyNumber)

	e.DOM.Find("div.readonly.companySummary > div.row:nth-child(2) > label.SCR011_04_003").Remove()
	NZBN := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(2)").Text()
	NZBN = strings.Trim(NZBN, "\n")
	Data.NZBN = NZBN
	println(Data.NZBN)

	e.DOM.Find("div.readonly.companySummary > div.row:nth-child(3) > label.SCR011_04_002").Remove()
	IncorporationDate := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(3)").Text()
	IncorporationDate = strings.Trim(IncorporationDate, "\n")
	Data.IncorporationDate = IncorporationDate
	println(Data.IncorporationDate)

	e.DOM.Find("div.readonly.companySummary > div.row:nth-child(4) > label.SCR011_04_022").Remove()
	CompanyStatus := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(4)").Text()
	CompanyStatus = strings.Trim(CompanyStatus, "\n")
	CompanyStatus = strings.TrimSpace(CompanyStatus)
	Data.CompanyStatus = CompanyStatus
	println(Data.CompanyStatus)
}

func handelCompanyNamefunc(e *colly.HTMLElement) {
	e.DOM.Find("span.entityIdentifier").Remove()
	Data.CompanyName = strings.TrimSpace(e.DOM.Text())
}

//func handelNZBNfunc(e *colly.HTMLElement) {
//	e.DOM.Find("label.SCR011_04_003").Remove()
//	value := strings.Trim(e.DOM.Text(), "\n")
//	Data.NZBN = value
//	fmt.Println(Data.NZBN)
//}

func handelOfficeAddressfunc(e *colly.HTMLElement) {
	registeredOfficeAddress := e.ChildText("div:nth-child(3) > div.addressLine")
	Data.RegisteredOfficeAddress = strings.Replace(registeredOfficeAddress, "\n", "", -1)

	addressforService := e.ChildText("div:nth-child(5) > div.addressLine")
	Data.AddressforService = strings.Replace(addressforService, "\n", "", -1)

	addressforShareRegister := e.ChildText("div:nth-child(8) > div.addressLine")
	Data.AddressforShareRegister = strings.Replace(addressforShareRegister, "\n", "", -1)
}

func handelDirectorfunc(e *colly.HTMLElement) {
	e.ForEach("table", func(i int, element *colly.HTMLElement) {
		Director := new(Director)

		element.DOM.Find("div.row:nth-child(1) > label").Remove()
		FullLegalName := element.DOM.Find("div.row:nth-child(1)").Text()
		FullLegalName = strings.Replace(FullLegalName, "\n", "", -1)
		FullLegalName = strings.Trim(FullLegalName, " ")
		Director.FullLegalName = FullLegalName

		element.DOM.Find("div.row:nth-child(2) > label").Remove()
		ResidentialAddress := element.DOM.Find("div.row:nth-child(2)").Text()
		ResidentialAddress = strings.Replace(ResidentialAddress, "\n", "", -1)
		Director.ResidentialAddress = ResidentialAddress

		element.DOM.Find("div.row:nth-child(3) > label").Remove()
		AppointmentDate := element.DOM.Find("div.row:nth-child(3)").Text()
		AppointmentDate = strings.Trim(AppointmentDate, "\n")
		Director.AppointmentDate = AppointmentDate

		Data.Directors = append(Data.Directors, *Director)
	})

	//for _,v := range Data.Directors {
	//	fmt.Println(v.FullLegalName)
	//	fmt.Println(v.ResidentialAddress)
	//	fmt.Println(v.AppointmentDate)
	//}
}

func handelShareholderfunc(e *colly.HTMLElement) {
	e.ForEach("div.allocationDetail", func(i int, element *colly.HTMLElement) {
		Allocation := new(Allocation)

		SharePercentage := element.ChildText("span.shareLabel")
		SharePercentage = strings.ReplaceAll(SharePercentage, "(", "")
		SharePercentage = strings.ReplaceAll(SharePercentage, ")", "")
		SharePercentage = strings.ReplaceAll(SharePercentage, "%", "")
		Allocation.Percentage, _ = strconv.ParseFloat(SharePercentage, 64)

		Shareholder := new(Shareholder)
		element.ForEach("div.labelValue.col2", func(j int, DivElement *colly.HTMLElement) {
			if j%2 == 0 {
				ShareholderName := strings.TrimSpace(DivElement.Text)
				Shareholder.Name = ShareholderName
			} else {
				ShareholderAddress := strings.TrimSpace(DivElement.Text)
				ShareholderAddress = strings.ReplaceAll(ShareholderAddress, "\n", "")
				Shareholder.Address = ShareholderAddress
				Allocation.Shareholders = append(Allocation.Shareholders, *Shareholder)
			}
		})
		Data.ShareholderAllocations = append(Data.ShareholderAllocations, *Allocation)
	})

	//for _, Allocation := range Data.ShareholderAllocations {
	//	fmt.Println(Allocation.Percentage)
	//	for _, holder := range Allocation.Shareholders {
	//		fmt.Println(holder.Name2)
	//		fmt.Println("***************")
	//		fmt.Println(holder.Address)
	//		fmt.Println("*******end********")
	//	}
	//}

}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("app.companiesoffice.govt.nz"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./cz_cache"),
		colly.UserAgent("None"),
	)
	// 设置超时时间为20秒
	c.SetRequestTimeout(30 * time.Second)

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

	c.OnHTML("div.readonly.companySummary", handelCompanySummaryfunc)
	c.OnHTML("div.panelContainer > div.leftPanel > div.row:first-child", handelCompanyNamefunc)
	c.OnHTML("div #addressPanel", handelOfficeAddressfunc)
	c.OnHTML("div #directorsPanel", handelDirectorfunc)
	c.OnHTML("div #shareholdersPanel", handelShareholderfunc)

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
		for CompanyNumber :=0; CompanyNumber < 10000000; CompanyNumber++  {
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
		//https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1830488/detail
		//fmt.Println(CompanyNumber)
		c.Visit("https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/" + strconv.Itoa(companyNumber) + "/detail")
	}

	insert(Data)
	end := time.Since(start)
	fmt.Println(end)

	//c.Visit("https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1908322") //OK

	//enc := json.NewEncoder(os.Stdout)
	//enc.SetIndent("", "  ")

	// Dump json to the standard output
	//enc.Encode(courses)
}
