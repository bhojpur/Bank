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

const BhojpurISPBCode = "19730825"

//ListEntries list the UPI keys of an account
func (s *UpiService) ListEntries(accountID string) ([]types.UpiEntry, *Response, error) {
	path := fmt.Sprintf("/v1/upi/%s/entries", accountID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor     `json:"cursor"`
		Data   []types.UpiEntry `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

type CreateUpiEntryOutput struct {
	ID             string `json:"id"`
	VerificationID string `json:"verification_id"`
}

// CreateEntry creates a new Key Entry
func (s *UpiService) CreateEntry(input types.CreateUpiEntryInput, idempotencyKey string) (CreateUpiEntryOutput, *Response, error) {
	var output CreateUpiEntryOutput

	if input.AccountID == "" {
		return output, nil, errors.New("accountID cannot be empty")
	}

	path := fmt.Sprintf("/v1/upi/%s/entries", input.AccountID)

	if input.ParticipantISPB == "" {
		input.ParticipantISPB = BhojpurISPBCode
	}

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return output, nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return output, nil, err
	}

	if input.VerificationID != "" {
		req.Header.Add("x-bhojpur-verification-id", input.VerificationID)
	}
	if input.VerificationCode != "" {
		req.Header.Add("x-bhojpur-verification-code", input.VerificationCode)
	}

	resp, err := s.client.Do(req, &output)
	if err != nil {
		return output, resp, err
	}

	return output, resp, err
}
