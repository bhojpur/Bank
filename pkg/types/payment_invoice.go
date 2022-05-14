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

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	invoiceAmountMin = 2000
	invoiceAmountMax = 1000000

	invoiceTypeDeposit        = "deposit"
	invoiceTypeProposal       = "proposal"
	invoiceTypeBillOfExchange = "bill_of_exchange"
)

var digitsRegex = regexp.MustCompile("[0-9]+")

func onlyDigits(key string) string {
	return strings.Join(digitsRegex.FindAllString(key, -1), "")
}

type PaymentInvoiceInput struct {
	AccountID      string                   `json:"account_id"`
	Amount         int                      `json:"amount"`
	ExpirationDate string                   `json:"expiration_date"`
	LimitDate      string                   `json:"limit_date,omitempty"`
	InvoiceType    string                   `json:"invoice_type"`
	Payer          PaymentInvoicePayerInput `json:"payer,omitempty" `
}

type PaymentInvoicePayerInput struct {
	Document  string `json:"document"`
	LegalName string `json:"legal_name"`
	TradeName string `json:"trade_name,omitempty"`
}

func (p *PaymentInvoiceInput) Validate() error {
	if strings.TrimSpace(p.AccountID) == "" {
		return errors.New("account_id can't be empty")
	}

	if p.Amount < invoiceAmountMin || p.Amount > invoiceAmountMax {
		return fmt.Errorf("amount can't be < %v or > %v", invoiceAmountMin, invoiceAmountMax)
	}

	_, err := time.Parse("2006-01-02", p.ExpirationDate)
	if err != nil || time.Now().Format("2006-01-02") > p.ExpirationDate {
		return errors.New("invalid expiration_date")
	}

	switch p.InvoiceType {
	case invoiceTypeDeposit, invoiceTypeProposal:
		p.LimitDate = p.ExpirationDate
	case invoiceTypeBillOfExchange:
		if strings.TrimSpace(p.LimitDate) == "" {
			p.LimitDate = p.ExpirationDate
		} else {
			_, err := time.Parse("2006-01-02", p.LimitDate)
			if err != nil || p.LimitDate < p.ExpirationDate {
				return errors.New("invalid limit_date")
			}
		}
	default:
		return errors.New("invalid invoice_type")
	}

	if p.InvoiceType != invoiceTypeDeposit {
		if strings.TrimSpace(p.Payer.LegalName) == "" {
			return errors.New("payer legal_name can't be empty")
		}

		p.Payer.Document = onlyDigits(p.Payer.Document)
		if strings.TrimSpace(p.Payer.Document) == "" {
			return errors.New("payer document can't be empty")
		}
	}

	return nil
}

type PaymentInvoice struct {
	ID             string                    `json:"id"`
	AccountID      string                    `json:"account_id"`
	CreatedBy      string                    `json:"created_by"`
	CreatedAt      string                    `json:"created_at"`
	RegisteredAt   string                    `json:"registered_at"`
	SettledAt      string                    `json:"settled_at"`
	Amount         int                       `json:"amount"`
	Barcode        string                    `json:"barcode"`
	WritableLine   string                    `json:"writable_line"`
	ExpirationDate string                    `json:"expiration_date"`
	InvoiceType    string                    `json:"invoice_type"`
	IssuanceDate   string                    `json:"issuance_date"`
	LimitDate      string                    `json:"limit_date"`
	Status         string                    `json:"status"`
	OurNumber      string                    `json:"our_number"`
	Beneficiary    PaymentInvoiceBeneficiary `json:"beneficiary"`
	Payer          PaymentInvoicePayer       `json:"payer"`
}

type PaymentInvoiceBeneficiary struct {
	AccountCode  string `json:"account_code"`
	BranchCode   string `json:"branch_code"`
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
	LegalName    string `json:"legal_name"`
	TradeName    string `json:"trade_name"`
}

type PaymentInvoicePayer struct {
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
	LegalName    string `json:"legal_name"`
	TradeName    string `json:"trade_name"`
}
