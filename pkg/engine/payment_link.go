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
	"strings"

	"github.com/bhojpur/bank/pkg/types"
)

type PaymentLinkService struct {
	client *Client
}

func (s *PaymentLinkService) Get(accountID, orderID string) (types.PaymentLink, *Response, error) {
	accountID = strings.TrimSpace(accountID)
	orderID = strings.TrimSpace(orderID)

	if accountID == "" {
		return types.PaymentLink{}, nil, errors.New("account_id can't be empty")
	}

	if orderID == "" {
		return types.PaymentLink{}, nil, errors.New("order_id can't be empty")
	}

	path := fmt.Sprintf("/v1/payment_links/%s/orders/%s", accountID, orderID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return types.PaymentLink{}, nil, err
	}

	var paymentLink types.PaymentLink

	resp, err := s.client.Do(req, &paymentLink)
	if err != nil {
		return types.PaymentLink{}, resp, err
	}

	return paymentLink, resp, nil
}

func (s *PaymentLinkService) Create(input types.PaymentLinkInput) (types.PaymentLink, *Response, error) {
	path := "/v1/payment_links/orders"

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return types.PaymentLink{}, nil, err
	}

	var paymentLink types.PaymentLink

	resp, err := s.client.Do(req, &paymentLink)
	if err != nil {
		return types.PaymentLink{}, resp, err
	}

	return paymentLink, resp, nil
}

func (s *PaymentLinkService) Cancel(orderID string, input types.PaymentLinkCancelInput) (types.PaymentLink, *Response, error) {
	orderID = strings.TrimSpace(orderID)

	if orderID == "" {
		return types.PaymentLink{}, nil, errors.New("account_id can't be empty")
	}

	if err := input.Validate(); err != nil {
		return types.PaymentLink{}, nil, err
	}

	path := fmt.Sprintf("/v1/payment_links/orders/%s/closed", orderID)

	req, err := s.client.NewAPIRequest(http.MethodPatch, path, input)
	if err != nil {
		return types.PaymentLink{}, nil, err
	}

	var paymentLink types.PaymentLink

	resp, err := s.client.Do(req, &paymentLink)
	if err != nil {
		return types.PaymentLink{}, resp, err
	}

	return paymentLink, resp, nil
}
