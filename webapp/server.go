package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type FlexStatements struct {
	FlexStatement []struct {
		OpenPositions []struct {
			isin string `xml:"isin,attr"`
			//currency   string `xml:"currency,attr"`
			//fxRate     string `xml:"fxRateToBase,attr"`
			symbol string `xml:"symbol,attr"`
			//reportDate string `xml:"reportDate,attr"`
			//position   string `xml:"position,attr"`
			//markPrice  string `xml:"markPrice,attr"`
		} `xml:"OpenPositions>OpenPosition"`
	} `xml:"FlexStatements>FlexStatement"`
}

type OPMap struct {
	isin string
	//currency   string
	//fxRate     string
	symbol string
	//reportDate string
	//position   string
	//markPrice  string
}

type OpenPositionsPage struct {
	Title         string
	OpenPositions map[int]OPMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>IB GO</h1>")
}

func OpenPositionsHandler(w http.ResponseWriter, r *http.Request) {
	var fx FlexStatements

	// Open xmlFile
	xmlFile, err := os.Open("./xml/All.xml")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &fx)
	op_map := make(map[int]OPMap)

	for index, i := range fx.FlexStatement[0].OpenPositions {

		fmt.Printf("\n %d: %s | %s", index, i.isin, i.symbol)
		op_map[index] = OPMap{i.isin /*, i.currency, i.fxRate*/, i.symbol /*, i.reportDate, i.position, i.markPrice*/}
	}
	fmt.Printf("\n%s", op_map)

	//Template construction
	p := OpenPositionsPage{Title: "Open Positions", OpenPositions: op_map}
	t, _ := template.ParseFiles("./templates/openpositiontemplate.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ib/", OpenPositionsHandler)
	http.ListenAndServe(":8000", nil)
}
