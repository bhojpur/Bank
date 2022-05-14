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
	gopath "path"

	"github.com/bhojpur/bank/pkg/types"
)

// PaymentService handles communication with Bhojpur Bank API
type PaymentService struct {
	client *Client
}

// PaymentOrders is a list of PaymentOrders
type paymentOrders struct {
	PaymentOrders []types.PaymentOrder `json:"paymentOrders"`
}

// HALPaymentOrders is a HAL wrapper around the Transactions type.
type halPaymentOrders struct {
	Embedded *paymentOrders `json:"_embedded"`
}

// MakeLocalPayment creates a local payment.
func (s *PaymentService) MakeLocalPayment(p types.LocalPayment) (*Response, error) {
	req, err := s.client.NewAPIRequest("POST", "/v1/payments/local", p)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}

// CreateScheduledPayment creates a scheduled payment. It returns the ID for the scheduled payment.
func (s *PaymentService) CreateScheduledPayment(p types.ScheduledPayment) (string, *Response, error) {
	req, err := s.client.NewAPIRequest("POST", "/v1/payments/scheduled", p)
	if err != nil {
		return "", nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return "", resp, err
	}

	loc := resp.Header.Get("Location")
	uid := gopath.Base(loc)
	return uid, resp, err
}

// ScheduledPayments retrieves a list of all the payment orders on the customer account. These may be
// orders for previous immediate payments or scheduled payment orders for future or on-going payments.
func (s *PaymentService) ScheduledPayments() ([]types.PaymentOrder, *Response, error) {
	path := fmt.Sprintf("/v1/payments/scheduled")
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	hPO := new(halPaymentOrders)
	resp, err := s.client.Do(req, &hPO)

	if hPO == nil {
		return nil, resp, err
	}

	if hPO.Embedded == nil {
		return nil, resp, err
	}

	return hPO.Embedded.PaymentOrders, resp, err
}
