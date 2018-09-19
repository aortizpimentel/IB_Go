package main


import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"os"
)

/*
FlexQueryResponse (1)
	FlexStatements (1)
		FlexStatement (1)
			OpenPositions (1)
				OpenPosition (n)
			Trades (1)
				AssetSummary (1)
				SymbolSummary (n)
				Order (n)
				Trade (n)
			TransactionTaxes (1)
				TransactionTax (n)
			CashTransactions (1)
				CashTransaction (n)
			Transfers (¿?)
*/

type FlexStatements struct {
	FlexStatement []struct {
		OpenPositions []struct {
			Symbols string `xml:"symbol,attr"`
		} `xml:"OpenPositions>OpenPosition"`
	}`xml:"FlexStatements>FlexStatement"`
	
}

func main() {

	// Open xmlFile
	xmlFile, err := os.Open("./xml/All.xml")

	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened All.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var fx FlexStatements
	xml.Unmarshal(byteValue, &fx)
	fmt.Println(fx)

	for _, i := range fx.FlexStatement[0].OpenPositions {
		fmt.Printf("\n%s",i)
	}
}