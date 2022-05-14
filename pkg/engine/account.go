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
	"errors"
	"fmt"
	"net/http"

	"github.com/bhojpur/bank/pkg/types"
)

// AccountService handles communication with Bhojpur Bank API
type AccountService struct {
	client *Client
}

//TODO: CreateNewIdentity

// Get account info
func (s *AccountService) Get(id string) (*types.Account, *Response, error) {

	path := fmt.Sprintf("/v1/accounts/%s", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var account types.Account
	resp, err := s.client.Do(req, &account)
	if err != nil {
		return nil, resp, err
	}

	return &account, resp, err
}

// List accounts
func (s *AccountService) List() ([]types.Account, *Response, error) {

	path := "/v1/accounts?paginate=true"

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor    `json:"cursor"`
		Data   []types.Account `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

//Get Account Balance
func (s *AccountService) GetBalance(id string) (*types.Balance, *Response, error) {

	path := fmt.Sprintf("/v1/accounts/%s/balance", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var balance types.Balance
	resp, err := s.client.Do(req, &balance)
	if err != nil {
		return nil, resp, err
	}

	return &balance, resp, err
}

//Get Account Statement
func (s *AccountService) GetStatement(id string) ([]types.Statement, *Response, error) {

	path := fmt.Sprintf("/v1/accounts/%s/statement", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor      `json:"cursor"`
		Data   []types.Statement `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

// Get Statement Entry
func (s *AccountService) GetStatementEntry(id string) (*types.Statement, *Response, error) {

	path := fmt.Sprintf("/v1/statement/entries/%s", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var statement types.Statement
	resp, err := s.client.Do(req, &statement)
	if err != nil {
		return nil, resp, err
	}

	return &statement, resp, err
}

// Get Account Fees of FeeType
func (s *AccountService) GetFees(accountID string, feeType string) (*types.Fee, *Response, error) {
	if feeType == "" {
		return nil, nil, errors.New("missing feeType value")
	}

	path := fmt.Sprintf("/v1/accounts/%s/fees/%s", accountID, feeType)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var fee types.Fee
	resp, err := s.client.Do(req, &fee)
	if err != nil {
		return nil, resp, err
	}

	return &fee, resp, err
}

// List Account Fees
func (s *AccountService) ListFees(accountID string) ([]types.Fee, *Response, error) {
	path := fmt.Sprintf("/v1/accounts/%s/fees", accountID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor `json:"cursor"`
		Data   []types.Fee  `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}
