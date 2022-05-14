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
	"net/http"

	"github.com/bhojpur/bank/pkg/types"
)

// UpiService handles communication with Bhojpur Bank API
type UpiService struct {
	client *Client
}

// GetOutboundUpi is a service used to retrieve information details from a UPI.
func (s *UpiService) GetOutboundUpi(id string) (*types.UPIOutBoundOutput, *Response, error) {
	path := fmt.Sprintf("/v1/upi/outbound_upi_payments/%s", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var upi types.UPIOutBoundOutput
	resp, err := s.client.Do(req, &upi)
	if err != nil {
		return nil, resp, err
	}

	return &upi, resp, err
}

// GetQRCodeData is a service used to retrieve information details from a UPI QRCode.
func (s *UpiService) GetQRCodeData(input types.GetQRCodeInput) (*types.QRCode, *Response, error) {
	const path = "/v1/upi/outbound_upi_payments/brcodes"

	req, err := s.client.NewAPIRequest(http.MethodGet, path, input)
	if err != nil {
		return nil, nil, err
	}

	var qrcode types.QRCode
	resp, err := s.client.Do(req, &qrcode)
	if err != nil {
		return nil, resp, err
	}

	return &qrcode, resp, err
}

//ListQRCodeDynamic list the dynamic qrcodes of an account
func (s *UpiService) ListDynamicQRCodes(accountID string) ([]types.QRCodeDynamic, *Response, error) {
	path := fmt.Sprintf("/v1/upi_payment_invoices/?account_id=%s", accountID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.AddAccountIdHeader(req, accountID)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor          `json:"cursor"`
		Data   []types.QRCodeDynamic `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

// CreateDynamicQRCode make a bar code payment invoice
func (s *UpiService) CreateDynamicQRCode(input types.CreateDynamicQRCodeInput, idempotencyKey string) (*types.UPIInvoiceOutput, *Response, error) {
	const path = "/v1/upi_payment_invoices"

	if err := input.Validate(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return nil, nil, err
	}

	var upiInvoiceOutput types.UPIInvoiceOutput
	resp, err := s.client.Do(req, &upiInvoiceOutput)
	if err != nil {
		return nil, resp, err
	}

	return &upiInvoiceOutput, resp, err
}

// CreatePedingPayment is a service used to create a pending payment.
func (s *UpiService) CreatePedingPayment(input types.CreatePedingPaymentInput, idempotencyKey string) (*types.PendingPaymentOutput, *Response, error) {
	const path = "/v1/upi/outbound_upi_payments"

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return nil, nil, err
	}

	var pendingPaymentOutput types.PendingPaymentOutput
	resp, err := s.client.Do(req, &pendingPaymentOutput)
	if err != nil {
		return nil, resp, err
	}

	return &pendingPaymentOutput, resp, err
}

// ConfirmPedingPayment is a service used to confirm a pending payment.
func (s *UpiService) ConfirmPedingPayment(input types.ConfirmPendingPaymentInput, idempotencyKey, upiID string) (*Response, error) {
	path := fmt.Sprintf("/v1/upi/outbound_upi_payments/%s/actions/confirm", upiID)

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
