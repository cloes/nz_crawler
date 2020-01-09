package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type Shareholder struct {
	Name    string
	Address string
}

type HistoricShareholder struct {
	FullLegalName string
	VacationDate  string
}

type Allocation struct {
	Percentage   float64
	Shareholders []Shareholder //有可能两个shareholder共同占有一定比例的股份
}

type Director struct {
	FullLegalName      string
	ResidentialAddress string
	AppointmentDate    string
}

type PreviousName struct {
	Name string
	From string
	To   string
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
	PreviousNames           []PreviousName
}

var wg sync.WaitGroup

//var Data *PageData

/*
func handelCompanySummaryfunc(e *colly.HTMLElement) {
	e.DOM.Find("div.readonly.companySummary > div.row:nth-child(1) > label.SCR011_04_003").Remove()
	CompanyNumber := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(1)").Text()
	CompanyNumber = strings.Trim(CompanyNumber, "\n")
	Data.CompanyNumber = CompanyNumber
	fmt.Printf("CompanyNumber:%v\n", Data.CompanyNumber)

	e.DOM.Find("div.readonly.companySummary > div.row:nth-child(2) > label.SCR011_04_003").Remove()
	NZBN := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(2)").Text()
	NZBN = strings.Trim(NZBN, "\n")
	Data.NZBN = NZBN
	fmt.Printf("NZBN:%v\n", Data.NZBN)

	e.DOM.Find("div.readonly.companySummary > div.row:nth-child(3) > label.SCR011_04_002").Remove()
	IncorporationDate := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(3)").Text()
	IncorporationDate = strings.Trim(IncorporationDate, "\n")
	Data.IncorporationDate = IncorporationDate
	fmt.Printf("IncorporationDate:%v\n", Data.IncorporationDate)

	e.DOM.Find("div.readonly.companySummary > div.row:nth-child(4) > label.SCR011_04_022").Remove()
	CompanyStatus := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(4)").Text()
	CompanyStatus = strings.TrimSpace(CompanyStatus)
	Data.CompanyStatus = CompanyStatus
	fmt.Printf("CompanyStatus:%v\n", Data.CompanyStatus)

	tmpElement := e.DOM.Find("div.readonly.companySummary > div.row").Has("label[for='entityType']")
	EntityType := tmpElement.Contents().Eq(2).Text()
	EntityType = strings.TrimSpace(EntityType)
	Data.EntityType = EntityType
	fmt.Printf("EntityType:%v\n", Data.EntityType)

	tmpElement = e.DOM.Find("div.readonly.companySummary > div.row").Has("label[for='constitutionFiled']")
	var ConstitutionFiled string
	if tmpElement.Find("a").Text() == "" {
		ConstitutionFiled = strings.TrimSpace(tmpElement.Contents().Eq(2).Text())
	} else {
		ConstitutionFiled = tmpElement.Find("a").Text()
	}
	Data.ConstitutionFiled = ConstitutionFiled
	fmt.Printf("ConstitutionFiled:%v\n", Data.ConstitutionFiled)

}

func handelCompanyNamefunc(e *colly.HTMLElement) {
	e.DOM.Find("span.entityIdentifier").Remove()
	Data.CompanyName = strings.TrimSpace(e.DOM.Find("div.row:first-child").Text())
	fmt.Printf("CompanyName:%v\n", Data.CompanyName)

	e.ForEach("div.previousNames > label", func(i int, element *colly.HTMLElement) {
		r, _ := regexp.Compile(`(.*)\(from (.*) to (.*)\)`)
		PreviousNameData := r.FindStringSubmatch(strings.TrimSpace(element.Text))

		PreviousName := new(PreviousName)
		PreviousName.Name = strings.TrimSpace(PreviousNameData[1])
		PreviousName.From = strings.TrimSpace(PreviousNameData[2])
		PreviousName.To = strings.TrimSpace(PreviousNameData[3])

		Data.PreviousNames = append(Data.PreviousNames, *PreviousName)
	})
}

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
*/

/*
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

*/

func main() {
	start := time.Now()
	for companyNumber := 1830488; companyNumber < 1830488+1; companyNumber++ {
		//c.Visit("https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1908322")
		//https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1908322/detail
		//https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/1830488/detail

		//Data = new(PageData)
		//c.Visit("https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/" + strconv.Itoa(companyNumber) + "/detail")
		wg.Add(1)
		url := "https://app.companiesoffice.govt.nz/companies/app/ui/pages/companies/" + strconv.Itoa(companyNumber) + "/detail"
		go work(url)
	}

	wg.Wait()
	end := time.Since(start)
	fmt.Println(end)
}

func work(url string) {

	defer wg.Done()

	Data := new(PageData)

	handelShareholderfunc := func(e *colly.HTMLElement) {
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
	}

	handleHistoricShareholdersfunc := func(e *colly.HTMLElement) {
		e.ForEach("div.shareholder", func(i int, element *colly.HTMLElement) {
			HistoricShareholder := new(HistoricShareholder)
			RowContent := element.ChildText("div.row")
			RowContents := strings.Split(RowContent, ":")
			HistoricShareholder.FullLegalName = strings.TrimSpace(RowContents[1])
			HistoricShareholder.VacationDate = strings.TrimSpace(RowContents[2])
			fmt.Printf("%+v\n", HistoricShareholder)
		})
	}

	handelCompanySummaryfunc := func(e *colly.HTMLElement) {
		e.DOM.Find("div.readonly.companySummary > div.row:nth-child(1) > label.SCR011_04_003").Remove()
		CompanyNumber := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(1)").Text()
		CompanyNumber = strings.Trim(CompanyNumber, "\n")
		Data.CompanyNumber = CompanyNumber
		fmt.Printf("CompanyNumber:%v\n", Data.CompanyNumber)

		e.DOM.Find("div.readonly.companySummary > div.row:nth-child(2) > label.SCR011_04_003").Remove()
		NZBN := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(2)").Text()
		NZBN = strings.Trim(NZBN, "\n")
		Data.NZBN = NZBN
		fmt.Printf("NZBN:%v\n", Data.NZBN)

		e.DOM.Find("div.readonly.companySummary > div.row:nth-child(3) > label.SCR011_04_002").Remove()
		IncorporationDate := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(3)").Text()
		IncorporationDate = strings.Trim(IncorporationDate, "\n")
		Data.IncorporationDate = IncorporationDate
		fmt.Printf("IncorporationDate:%v\n", Data.IncorporationDate)

		e.DOM.Find("div.readonly.companySummary > div.row:nth-child(4) > label.SCR011_04_022").Remove()
		CompanyStatus := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(4)").Text()
		CompanyStatus = strings.TrimSpace(CompanyStatus)
		Data.CompanyStatus = CompanyStatus
		fmt.Printf("CompanyStatus:%v\n", Data.CompanyStatus)

		tmpElement := e.DOM.Find("div.readonly.companySummary > div.row").Has("label[for='entityType']")
		EntityType := tmpElement.Contents().Eq(2).Text()
		EntityType = strings.TrimSpace(EntityType)
		Data.EntityType = EntityType
		fmt.Printf("EntityType:%v\n", Data.EntityType)

		tmpElement = e.DOM.Find("div.readonly.companySummary > div.row").Has("label[for='constitutionFiled']")
		var ConstitutionFiled string
		if tmpElement.Find("a").Text() == "" {
			ConstitutionFiled = strings.TrimSpace(tmpElement.Contents().Eq(2).Text())
		} else {
			ConstitutionFiled = tmpElement.Find("a").Text()
		}
		Data.ConstitutionFiled = ConstitutionFiled
		fmt.Printf("ConstitutionFiled:%v\n", Data.ConstitutionFiled)

		/*
			e.DOM.Find("div.readonly.companySummary > div.row:nth-child(8) > label.SCR011_04_028").Remove()
			EntityType := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(8)").Text()
			EntityType = strings.TrimSpace(EntityType)
			Data.EntityType = EntityType
			fmt.Printf("EntityType:%v\n",Data.EntityType)

			e.DOM.Find("div.readonly.companySummary > div.row:nth-child(9) > label.SCR011_04_029").Remove()
			ConstitutionFiled := e.DOM.Find("div.readonly.companySummary > div.row:nth-child(9)").Text()
			ConstitutionFiled = strings.TrimSpace(ConstitutionFiled)
			Data.ConstitutionFiled = ConstitutionFiled
			fmt.Printf("ConstitutionFiled:%v\n",Data.ConstitutionFiled)
		*/
	}

	handelCompanyNamefunc := func(e *colly.HTMLElement) {
		e.DOM.Find("span.entityIdentifier").Remove()
		Data.CompanyName = strings.TrimSpace(e.DOM.Find("div.row:first-child").Text())
		fmt.Printf("CompanyName:%v\n", Data.CompanyName)

		e.ForEach("div.previousNames > label", func(i int, element *colly.HTMLElement) {
			r, _ := regexp.Compile(`(.*)\(from (.*) to (.*)\)`)
			PreviousNameData := r.FindStringSubmatch(strings.TrimSpace(element.Text))

			PreviousName := new(PreviousName)
			PreviousName.Name = strings.TrimSpace(PreviousNameData[1])
			PreviousName.From = strings.TrimSpace(PreviousNameData[2])
			PreviousName.To = strings.TrimSpace(PreviousNameData[3])

			Data.PreviousNames = append(Data.PreviousNames, *PreviousName)
		})
	}

	handelOfficeAddressfunc := func(e *colly.HTMLElement) {
		registeredOfficeAddress := e.ChildText("div:nth-child(3) > div.addressLine")
		Data.RegisteredOfficeAddress = strings.Replace(registeredOfficeAddress, "\n", "", -1)

		addressforService := e.ChildText("div:nth-child(5) > div.addressLine")
		Data.AddressforService = strings.Replace(addressforService, "\n", "", -1)

		addressforShareRegister := e.ChildText("div:nth-child(8) > div.addressLine")
		Data.AddressforShareRegister = strings.Replace(addressforShareRegister, "\n", "", -1)
	}

	handelDirectorfunc := func(e *colly.HTMLElement) {
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
	}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("app.companiesoffice.govt.nz"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./cz_cache"),
		colly.UserAgent("None"),
	)
	// 设置超时时间为60秒
	c.SetRequestTimeout(60 * time.Second)
	c.MaxDepth = 1

	c.OnHTML("div.readonly.companySummary", handelCompanySummaryfunc)
	c.OnHTML("div.panelContainer > div.leftPanel", handelCompanyNamefunc)
	c.OnHTML("div #addressPanel", handelOfficeAddressfunc)
	c.OnHTML("div #directorsPanel", handelDirectorfunc)
	c.OnHTML("div #shareholdersPanel", handelShareholderfunc)
	c.OnHTML("div .historic.wideLabel", handleHistoricShareholdersfunc)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnScraped(func(response *colly.Response) {
		if response.StatusCode == 200 {
			insert(Data)
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
		Data = new(PageData)
	})

	c.Visit(url)
}
