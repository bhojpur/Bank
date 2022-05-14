package types

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Transaction represents the details of a transaction.
type Transaction struct {
	ID        string  `json:"id"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
	Direction string  `json:"direction"`
	Created   string  `json:"created"`
	Narrative string  `json:"narrative"`
	Source    string  `json:"source"`
	Balance   float64 `json:"balance,omitempty"`
}

// DDTransaction represents the details of a direct debit transaction.
type DDTransaction struct {
	ID                 string  `json:"id"`
	Currency           string  `json:"currency"`
	Amount             float64 `json:"amount"`
	Direction          string  `json:"direction"`
	Created            string  `json:"created"`
	Narrative          string  `json:"narrative"`
	Source             string  `json:"source"`
	MandateID          string  `json:"mandateId"`
	Type               string  `json:"type"`
	MerchantID         string  `json:"merchantId"`
	MerchantLocationID string  `json:"merchantLocationId"`
	SpendingCategory   string  `json:"spendingCategory"`
}

// CardTransaction represents the details of a Payment Card transaction
type CardTransaction struct {
	Transaction
	Method            string  `json:"cardTransactionMethod"`
	Status            string  `json:"status"`
	SourceAmount      float64 `json:"sourceAmount"`
	SourceCurrency    string  `json:"sourceCurrency"`
	MerchantID        string  `json:"merchantId"`
	SpendingCategory  string  `json:"spendingCategory"`
	Country           string  `json:"country"`
	POSTimestamp      string  `json:"posTimestamp"`
	AuthorisationCode string  `json:"authorisationCode"`
	EventID           string  `json:"eventId"`
	Receipt           Receipt `json:"receipt"`
	CardLast4         string  `json:"cardLast4"`
}

// SpendingCategory is the category associated with a transaction
type SpendingCategory struct {
	SpendingCategory string `json:"spendingCategory"`
}
