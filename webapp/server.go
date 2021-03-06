package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	//"google.golang.org/appengine" //Needed for Google Cloud Platform
)

type FlexStatements struct {
	FlexStatement []struct {
		AccountInformation []struct {
			Name       string `xml:"name,attr"`
			Street     string `xml:"street,attr"`
			Street2    string `xml:"street2,attr"`
			PostalCode string `xml:"postalCode,attr"`
			Email      string `xml:"primaryEmail,attr"`
		} `xml:"AccountInformation"`
		OpenPositions []struct {
			Isin        string `xml:"isin,attr"`
			Currency    string `xml:"currency,attr"`
			FxRate      string `xml:"fxRateToBase,attr"`
			Symbol      string `xml:"symbol,attr"`
			Position    string `xml:"position,attr"`
			MarkPrice   string `xml:"markPrice,attr"`
			Description string `xml:"description,attr"`
		} `xml:"OpenPositions>OpenPosition"`
		Order []struct {
			IBOrderID            string `xml:"ibOrderID,attr"`
			BuySell              string `xml:"buySell,attr"`
			Isin                 string `xml:"isin,attr"`
			Currency             string `xml:"currency,attr"`
			FxRate               string `xml:"fxRateToBase,attr"`
			Symbol               string `xml:"symbol,attr"`
			Quantity             string `xml:"quantity,attr"`
			TradePrice           string `xml:"tradePrice,attr"`
			TradeMoney           string `xml:"tradeMoney,attr"`
			Taxes                string `xml:"taxes,attr"`
			IBComission          string `xml:"ibCommission,attr"`
			IBCommissionCurrency string `xml:"ibCommissionCurrency,attr"`
			Cost                 string `xml:"cost,attr"`
			Description          string `xml:"description,attr"`
			TradeDate            string `xml:"tradeDate,attr"`
			TradeTime            string `xml:"tradeTime,attr"`
		} `xml:"Trades>Order"`
	} `xml:"FlexStatements>FlexStatement"`
}

type OPMap struct {
	Index       int
	Isin        string
	Currency    string
	FxRate      string
	Symbol      string
	Position    string
	MarkPrice   string
	Description string
}

type OrderMap struct {
	IBOrderID            string
	BuySell              string
	Isin                 string
	Currency             string
	FxRate               string
	Symbol               string
	Quantity             string
	TradePrice           string
	TradeMoney           string
	Taxes                string
	IBComission          string
	IBCommissionCurrency string
	Cost                 string
	Description          string
	TradeDate            string
	TradeTime            string
}

type OpenPositionsPage struct {
	Title         string
	OpenPositions map[int]OPMap
}

type OrdersPage struct {
	Title  string
	Orders map[int]OrderMap
}

type D6Map struct {
	Deposit          bool
	Year             string
	D6ReplacementNum string
	ClientName       string
	NIF              string
	WayType          string
	WayName          string
	WayNum           string
	PostalCode       string
	City             string
	State            string
}

type D6Page struct {
	Title         string
	D6            D6Map
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
		op_map[index] = OPMap{index + 1, i.Isin, i.Currency, i.FxRate, i.Symbol, i.Position, i.MarkPrice, i.Description}
	}

	//Template construction
	p := OpenPositionsPage{Title: "Open Positions", OpenPositions: op_map}
	t, _ := template.ParseFiles("./templates/openpositiontemplate.html")
	t.Execute(w, p)
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
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
	order_map := make(map[int]OrderMap)

	for index, i := range fx.FlexStatement[0].Order {
		order_map[index] = OrderMap{i.IBOrderID, i.BuySell, i.Isin, i.Currency, i.FxRate, i.Symbol, i.Quantity, i.TradePrice, i.TradeMoney, i.Taxes, i.IBComission, i.IBCommissionCurrency, i.Cost, i.Description, i.TradeDate, i.TradeTime}
	}

	//Template construction
	p := OrdersPage{Title: "Orders", Orders: order_map}
	t, _ := template.ParseFiles("./templates/orderstemplate.html")
	t.Execute(w, p)
}

func D6Handler(w http.ResponseWriter, r *http.Request) {

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
		op_map[index] = OPMap{index + 1, i.Isin, i.Currency, i.FxRate, i.Symbol, i.Position, i.MarkPrice, i.Description}
	}
	fmt.Printf("\n%s", fx)

	/**** TESTING ****/
	client_info := D6Map{true, "2018", "", "Adrián Ortiz Pimentel", "78982122F", "Urbanización", "El Rocío 1", "5", "29670", "San Pedro de Alcántara", "Málaga"}

	//Template construction
	p := D6Page{Title: "Modelo D-6", D6: client_info, OpenPositions: op_map}
	t, _ := template.ParseFiles("./templates/d6template.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ib/open", OpenPositionsHandler)
	http.HandleFunc("/ib/orders", OrdersHandler)
	http.HandleFunc("/ib/d6", D6Handler)

	http.ListenAndServe(":8000", nil)

	//Needed for Google Cloud Platform
	//appengine.Main()
}
