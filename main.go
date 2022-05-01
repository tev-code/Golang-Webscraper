package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	//Create a file called data
	fName:= "data.csv"
	file, err := os.Create(fName)

    //Catch any errors along the way
	if err != nil {
      log.Fatalf("Could not create file, err :%q" ,err)
	  return
	}

	//Once done working with the file GO will close it
	defer file.Close()

	//What ever data we collecting will write into a CSV file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		//Enter URL of website you wish to scrape
		colly.AllowedDomains(""),
	)

		c.OnHTML(".standing-table__row", func(e *colly.HTMLElement) {

			//Function to write data toe the CSV file
			writer.Write([]string{
				e.ChildText("a"),
				e.ChildText("span"),
			})

		})

		log.Printf("Scraping Complete\n")
		log.Println(c)

}