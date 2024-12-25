package com_n26

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func cuetest() {
	ctx := cuecontext.New()

	// Load the package "example" from the current directory.
	// We don't need to specify a config in this example.
	insts := load.Instances([]string{"."}, nil)

	// The current directory just has one file without any build tags,
	// and that file belongs to the example package.
	// So we get a single instance as a result.
	v := ctx.BuildInstance(insts[0])
	if err := v.Err(); err != nil {
		log.Fatal(err)
	}

	// Lookup the 'output' field and print it out
	output := v.LookupPath(cue.ParsePath("output"))
	fmt.Println(output)
}

func importFile(path string) []*Transaction {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	bankTrans := []*Transaction{}

	for _, row := range records {
		bankTrans = append(bankTrans, NewTransaction(row))
	}
	return bankTrans
}

func NewTransaction(transaction []string) *Transaction {
	amount, _ := strconv.ParseFloat(transaction[7], 64)
	return &Transaction{
		BookingDate:      transaction[0],
		ValueDate:        transaction[1],
		PartnerName:      transaction[2],
		PartnerIBAN:      transaction[3],
		Type:             transaction[4],
		PaymentReference: transaction[5],
		AccountName:      transaction[6],
		AmountEUR:        amount,
		OriginalAmount:   0,
		OriginalCurrency: transaction[9],
		ExchangeRate:     0,
	}
}

type Transaction struct {
	BookingDate      string  `json:"booking_date"`
	ValueDate        string  `json:"value_date"`
	PartnerName      string  `json:"partner_name"`
	PartnerIBAN      string  `json:"partner_iban"`
	Type             string  `json:"type"`
	PaymentReference string  `json:"payment_reference"`
	AccountName      string  `json:"account_name"`
	AmountEUR        float64 `json:"amount_eur"`
	OriginalAmount   float64 `json:"original_amount"`
	OriginalCurrency string  `json:"original_currency"`
	ExchangeRate     float64 `json:"exchange_rate"`
}
