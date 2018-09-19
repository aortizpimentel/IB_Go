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
			Isin        string `xml:"isin,attr"`
			Currency    string `xml:"currency,attr"`
			FxRate      string `xml:"fxRateToBase,attr"`
			Symbol      string `xml:"symbol,attr"`
			Position    string `xml:"position,attr"`
			MarkPrice   string `xml:"markPrice,attr"`
			Description string `xml:"description,attr"`
		} `xml:"OpenPositions>OpenPosition"`
	} `xml:"FlexStatements>FlexStatement"`
}

type OPMap struct {
	Isin        string
	Currency    string
	FxRate      string
	Symbol      string
	Position    string
	MarkPrice   string
	Description string
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
		op_map[index] = OPMap{i.Isin, i.Currency, i.FxRate, i.Symbol, i.Position, i.MarkPrice, i.Description}
	}

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
