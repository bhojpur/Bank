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

// Receipt is a receipt for a transaction
type Receipt struct {
	ID                 string        `json:"receiptId"`
	EventID            string        `json:"eventId"`
	MetadataSource     string        `json:"metadataSource"`
	ReceiptIdentifier  string        `json:"receiptIdentifier"`
	MerchantIdentifier string        `json:"merchantIdentifier"`
	MerchantAddress    string        `json:"merchantAddress"`
	TotalAmount        float64       `json:"totalAmount"`
	TotalTax           float64       `json:"totalTax"`
	TaxReference       string        `json:"taxNumber"`
	AuthCode           string        `json:"authCode"`
	CardLast4          string        `json:"cardLast4"`
	ProviderName       string        `json:"providerName"`
	Items              []ReceiptItem `json:"items"`
	Notes              []ReceiptNote `json:"notes"`
}

// ReceiptItem is a single item on a Receipt
type ReceiptItem struct {
	ID          string  `json:"receiptItemId"`
	Description string  `json:"description"`
	Quantity    int32   `json:"quantity"`
	Amount      float64 `json:"amount"`
	Tax         float64 `json:"tax"`
	URL         string  `json:"url"`
}

// ReceiptNote is a single item on a Receipt
type ReceiptNote struct {
	ID          string `json:"noteId"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
