package main

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
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/uuid"

	engine "github.com/bhojpur/bank/pkg/engine"
	"github.com/bhojpur/bank/pkg/types"
)

func main() {
	clientID := os.Getenv("BHOJPUR_BANK_CLIENT_ID")
	privKeyPath := os.Getenv("BHOJPUR_BANK_PRIVATE_KEY")
	consentURL := os.Getenv("BHOJPUR_BANK_CONSENT_REDIRECT_URL")

	pemPrivKey := readFileContent(privKeyPath)

	client, err := engine.NewClient(
		engine.WithClientID(clientID),
		engine.WithPEMPrivateKey(pemPrivKey),
		engine.SetConsentURL(consentURL),
		engine.UseSandbox(),
		//	engine.EnableDebug(),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Authenticate(); err != nil {
		log.Fatal(err)
	}

	consentLink, err := client.ConsentLink("")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nconsent_link: %s\n", consentLink)

	//list all game providers
	gameProviders, _, err := client.Topups.ListGameProviders()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(gameProviders)

	if len(gameProviders.Providers) > 1 {
		// list all product values from psn store
		products, _, err := client.Topups.GetValuesFromGameProvider(gameProviders.Providers[0].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(products)
	}

	// returns institutions
	allinstitutions, _, err := client.Institution.List(engine.AllInstitutions)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(allinstitutions), allinstitutions[0])

	// returns institutions participating in the SPI. Useful for UPI operations
	SPIinstitutions, _, err := client.Institution.List(engine.SPIParticipants)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(SPIinstitutions), SPIinstitutions[0])

	// returns institutions participating in the STR. Useful for TED operations
	STRinstitutions, _, err := client.Institution.List(engine.STRParticipants)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(STRinstitutions), STRinstitutions[0])

	// return institution by code or ISPB code
	institution, _, err := client.Institution.Get(SPIinstitutions[0].ISPBCode)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(institution)

	accounts, _, err := client.Account.List()
	if err != nil {
		log.Fatal(err)
	}
	for i := range accounts {
		log.Printf("acc[%d]: %v\n\n", i, accounts[i])
		acc, _, err := client.Account.Get(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Detailed account: %+v", acc)

		balance, _, err := client.Account.GetBalance(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Balance: %+v", balance)

		statement, _, err := client.Account.GetStatement(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Statement: %+v", statement)

		fees, _, err := client.Account.ListFees(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("AllFees: %+v", fees)

		fee, _, err := client.Account.GetFees(accounts[i].ID, "internal_transfer")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Fee: %+v", fee)

		//Internal DryRun Transfer
		transfInput := types.TransferInput{
			AccountID: accounts[i].ID,
			Amount:    100,
			Target: types.Target{
				Account: types.TransferAccount{
					AccountCode: "334201",
				},
			},
		}

		idempotencyKey := uuid.New().String()
		transfer, _, err := client.Transfer.DryRunTransfer(transfInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Transfer(dry-run): %+v", transfer)

		//Internal Transfer
		transfer, _, err = client.Transfer.Transfer(transfInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Transfer: %+v", transfer)

		//External DryRun Transfer
		transfExtInput := types.TransferInput{
			AccountID: accounts[i].ID,
			Amount:    100,
			Target: types.Target{
				Account: types.TransferAccount{
					AccountCode:     "1234",
					BranchCode:      "7032",
					InstitutionCode: "001",
				},
				Entity: types.Entity{
					Name:         "James Bond",
					Document:     "00700700700",
					DocumentType: "cpf",
				},
			},
		}
		transfer, _, err = client.Transfer.DryRunTransfer(transfExtInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("External Transfer(dry-run): %+v", transfer)

		//External  Transfer
		transfer, _, err = client.Transfer.Transfer(transfExtInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("External Transfer: %+v", transfer)

		//List Internal Transfers
		internalTransfers, _, err := client.Transfer.ListInternal(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Internal Transfers: %+v\n", internalTransfers)
		for i, t := range internalTransfers {
			//Get an internal transfer
			transfer, _, err := client.Transfer.GetInternal(t.ID)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Internal Transfer[%d]: %+v\n", i, transfer)
		}

		//List External Transfers
		externalTransfers, _, err := client.Transfer.ListExternal(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("External Transfers: %+v\n", externalTransfers)
		for i, t := range externalTransfers {
			//Get an external transfer
			transfer, _, err := client.Transfer.GetExternal(t.ID)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("External Transfer[%d]: %+v\n", i, transfer)
		}

		//Schedule and Cancel an internal Transfer
		ScheduleAndCancelTransfer(accounts[i].ID, client)

		//Payment Invoice
		paymentInvoiceInput := types.PaymentInvoiceInput{
			AccountID:      accounts[i].ID,
			Amount:         5000,
			ExpirationDate: time.Now().Format("2006-01-02"),
			InvoiceType:    "deposit",
		}

		// Make Payment Invoice
		paymentInvoice, _, err := client.PaymentInvoice.PaymentInvoice(paymentInvoiceInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Payment Invoice: %+v", paymentInvoice)

		//List Payment Invoices
		paymentInvoices, _, err := client.PaymentInvoice.List(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Payment Invoices: %+v\n", paymentInvoices)
		for i, t := range paymentInvoices {
			//Get a payment invoice
			invoice, _, err := client.PaymentInvoice.Get(t.ID)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Payment Invoice[%d]: %+v\n", i, invoice)
		}
	}

	//List UPI Keys
	accountID := "968cc34d-d827-448b-ac1b-e6e29836a160"
	//idempotencyKey := uuid.New().String()
	upiKeys, _, err := client.Upi.ListDynamicQRCodes(accountID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Keys UPI: %+v\n", upiKeys)

	//Get Outbound UPI
	UpiID := "b5c2354c-91a0-4837-bb15-7f88fcd9d4c5"
	outBoundUpi, _, err := client.Upi.GetOutboundUpi(UpiID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Outbound UPI: %+v\n", outBoundUpi)

	//Get QRCode Data
	getQRCodeInput := types.GetQRCodeInput{BRCode: "00020186580014in.gov.bcb.upi0136123e4567-e12b-12d1-a456-4266554400005204000053039865802BR5913Fulano de Tal6008BRASILIA62070503***63041D3D"}
	qrCode, _, err := client.Upi.GetQRCodeData(getQRCodeInput)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("QRCode Data: %+v\n", qrCode)
}

func readFileContent(path string) []byte {
	content, _ := ioutil.ReadFile(path)
	return content
}

func ScheduleAndCancelTransfer(accID string, client *engine.Client) {
	transfInput := types.TransferInput{
		AccountID:   accID,
		Amount:      100,
		ScheduledTo: "2020-03-25",
		Target: types.Target{
			Account: types.TransferAccount{
				AccountCode: "334201",
			},
		},
	}

	idempotencyKey := uuid.New().String()
	transfer, _, err := client.Transfer.Transfer(transfInput, idempotencyKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transfer: %+v", transfer)

	//Check transfer status
	intTransfer, _, err := client.Transfer.GetInternal(transfer.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Scheduled Transfer Status: %s\n", intTransfer.Status)

	//Cancel transfer
	resp, err := client.Transfer.CancelInternal(transfer.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response from Cancel Transfer: %+v\n", resp.Response)

	//Check if transfer was canceled
	canceledTransfer, resp, err := client.Transfer.GetInternal(transfer.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transfer Status: %s\n", canceledTransfer.Status)
}
