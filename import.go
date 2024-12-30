package com_n26

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// test
func ImportFile(path string) []*Transaction {
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
		bankTrans = append(bankTrans, newTransaction(row))
	}
	return bankTrans
}

func newTransaction(transaction []string) *Transaction {
	amount, _ := strconv.ParseFloat(transaction[7], 64)
	originalAmount, _ := strconv.ParseFloat(transaction[8], 64)
	exchangeRate, _ := strconv.ParseFloat(transaction[10], 64)
	return &Transaction{
		BookingDate:      transaction[0],
		ValueDate:        transaction[1],
		PartnerName:      transaction[2],
		PartnerIBAN:      transaction[3],
		Type:             transaction[4],
		PaymentReference: transaction[5],
		AccountName:      transaction[6],
		AmountEUR:        amount,
		OriginalAmount:   originalAmount,
		OriginalCurrency: transaction[9],
		ExchangeRate:     exchangeRate,
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

type Transaction_N26 struct {
	BookingDate      time.Time `json:"booking_date"`
	ValueDate        time.Time `json:"value_date"`
	PartnerName      string    `json:"partner_name"`
	PartnerIBAN      string    `json:"partner_iban"`
	Type             string    `json:"type"`
	PaymentReference string    `json:"payment_reference"`
	AccountName      string    `json:"account_name"`
	AmountEUR        float32   `json:"amount_eur"`
	OriginalAmount   float32   `json:"original_amount"`
	OriginalCurrency string    `json:"original_currency"`
	ExchangeRate     float32   `json:"exchange_rate"`
}
