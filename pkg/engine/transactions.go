package engine

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
	"fmt"

	"github.com/bhojpur/bank/pkg/types"
)

// TransactionService handles communication with Bhojpur Bank API
type TransactionService struct {
	client *Client
}

// Transactions is a list of transaction summaries.
type transactions struct {
	Transactions []types.Transaction `json:"transactions"`
}

// HALTransactions is a HAL wrapper around the Transactions type.
type halTransactions struct {
	Embedded *transactions `json:"_embedded"`
}

// ddTransactions is a list of transaction summaries.
type ddTransactions struct {
	Transactions []types.DDTransaction `json:"transactions"`
}

// halDDTransactions is a HAL wrapper around the Transactions type.
type halDDTransactions struct {
	Embedded *ddTransactions `json:"_embedded"`
}

// cardTransactions is a list of Card transactions
type cardTransactions struct {
	Transactions []types.CardTransaction `json:"transactions"`
}

// HALCardTransactions is a HAL wrapper around the Transactions type.
type halCardTransactions struct {
	Embedded *cardTransactions `json:"_embedded"`
}

// Transactions returns a list of transaction summaries for the current user. It
// accepts optional time.Time values to request transactions within a given date
// range. If these values are not provided the API returns the last 100 transactions.
func (s *TransactionService) Transactions(dr *types.DateRange) ([]types.Transaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions")
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if dr != nil {
		q := req.URL.Query()
		q.Add("from", dr.From.Format("2006-01-02"))
		q.Add("to", dr.To.Format("2006-01-02"))
		req.URL.RawQuery = q.Encode()
	}

	hTxns := new(halTransactions)
	resp, err := s.client.Do(req, &hTxns)

	if hTxns == nil {
		return nil, resp, err
	}

	if hTxns.Embedded == nil {
		return nil, resp, err
	}

	return hTxns.Embedded.Transactions, resp, err
}

// Transaction returns an individual transaction for the current customer.
func (s *TransactionService) Transaction(uid string) (*types.Transaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/%s", uid)
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	txn := new(types.Transaction)
	resp, err := s.client.Do(req, txn)
	if err != nil {
		return nil, resp, err
	}
	return txn, resp, nil
}

// DDTransactions returns a list of direct debit transactions for the current user.
// It accepts optional time.Time values to request transactions within a given date
// range. If these values are not provided the API returns the last 100 transactions.
func (s *TransactionService) DDTransactions(dr *types.DateRange) ([]types.DDTransaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/direct-debit")
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if dr != nil {
		q := req.URL.Query()
		q.Add("from", dr.From.Format("2006-01-02"))
		q.Add("to", dr.To.Format("2006-01-02"))
		req.URL.RawQuery = q.Encode()
	}

	var halResp *halDDTransactions
	var txns *ddTransactions
	resp, err := s.client.Do(req, &halResp)
	if err != nil {
		return nil, resp, err
	}

	if halResp.Embedded != nil {
		txns = halResp.Embedded
	}

	return txns.Transactions, resp, nil
}

// DDTransaction returns an individual transaction for the current customer.
func (s *TransactionService) DDTransaction(uid string) (*types.DDTransaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/direct-debit/%s", uid)
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	ddTxn := new(types.DDTransaction)
	resp, err := s.client.Do(req, ddTxn)
	if err != nil {
		return nil, resp, err
	}

	return ddTxn, resp, nil
}

// SetDDSpendingCategory updates the spending category for a given direct debit.
func (s *TransactionService) SetDDSpendingCategory(uid, cat string) (*Response, error) {
	path := fmt.Sprintf("/v1/transactions/direct-debit/%s", uid)
	reqCat := types.SpendingCategory{SpendingCategory: cat}
	req, err := s.client.NewAPIRequest("PUT", path, reqCat)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}

// FPSTransactionsIn returns a list of inbound Fast Payments transaction summaries for
// the current user. It accepts optional time.Time values to request transactions within
// a given date range. If these values are not provided the API returns the last 100
// transactions.
func (s *TransactionService) FPSTransactionsIn(dr *types.DateRange) ([]types.Transaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/fps/in")
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if dr != nil {
		q := req.URL.Query()
		q.Add("from", dr.From.Format("2006-01-02"))
		q.Add("to", dr.To.Format("2006-01-02"))
		req.URL.RawQuery = q.Encode()
	}

	var halResp *halTransactions
	var txns *transactions
	resp, err := s.client.Do(req, &halResp)
	if err != nil {
		return nil, resp, err
	}

	if halResp.Embedded != nil {
		txns = halResp.Embedded
	}

	return txns.Transactions, resp, nil
}

// FPSTransactionIn returns an individual transaction for the current customer.
func (s *TransactionService) FPSTransactionIn(uid string) (*types.Transaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/fps/in/%s", uid)
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	txn := new(types.Transaction)
	resp, err := s.client.Do(req, txn)
	if err != nil {
		return nil, resp, err
	}

	return txn, resp, err
}

// FPSTransactionsOut returns a list of inbound Fast Payments transaction summaries for
// the current user. It accepts optional time.Time values to request transactions within
// a given date range. If these values are not provided the API returns the last 100
// transactions.
func (s *TransactionService) FPSTransactionsOut(dr *types.DateRange) ([]types.Transaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/fps/out")
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if dr != nil {
		q := req.URL.Query()
		q.Add("from", dr.From.Format("2006-01-02"))
		q.Add("to", dr.To.Format("2006-01-02"))
		req.URL.RawQuery = q.Encode()
	}

	var halResp *halTransactions
	var txns *transactions
	resp, err := s.client.Do(req, &halResp)
	if err != nil {
		return nil, resp, err
	}

	if halResp.Embedded != nil {
		txns = halResp.Embedded
	}

	return txns.Transactions, resp, nil
}

// FPSTransactionOut returns an individual transaction for the current customer.
func (s *TransactionService) FPSTransactionOut(uid string) (*types.Transaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/fps/out/%s", uid)
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	txn := new(types.Transaction)
	resp, err := s.client.Do(req, txn)
	if err != nil {
		return nil, resp, err
	}

	return txn, resp, err
}

// CardTransactions returns a list of transaction summaries for the current user. It accepts optional
// time.Time values to request transactions within a given date range. If these values are not provided
// the API returns the last 100 transactions.
func (s *TransactionService) CardTransactions(dr *types.DateRange) ([]types.CardTransaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/card")
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if dr != nil {
		q := req.URL.Query()
		q.Add("from", dr.From.Format("2006-01-02"))
		q.Add("to", dr.To.Format("2006-01-02"))
		req.URL.RawQuery = q.Encode()
	}

	hTxns := new(halCardTransactions)
	resp, err := s.client.Do(req, &hTxns)

	if hTxns == nil {
		return nil, resp, err
	}

	if hTxns.Embedded == nil {
		return nil, resp, err
	}

	return hTxns.Embedded.Transactions, resp, err
}

// CardTransaction returns an individual payment card transaction for the current customer.
func (s *TransactionService) CardTransaction(uid string) (*types.CardTransaction, *Response, error) {
	path := fmt.Sprintf("/v1/transactions/card/%s", uid)
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	txn := new(types.CardTransaction)
	resp, err := s.client.Do(req, txn)
	if err != nil {
		return nil, resp, err
	}

	return txn, resp, nil
}

// SetCardSpendingCategory updates the spending category for a given payment card transaction.
func (s *TransactionService) SetCardSpendingCategory(uid, cat string) (*Response, error) {
	path := fmt.Sprintf("/v1/transactions/card/%s", uid)
	reqCat := types.SpendingCategory{SpendingCategory: cat}
	req, err := s.client.NewAPIRequest("PUT", path, reqCat)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}
