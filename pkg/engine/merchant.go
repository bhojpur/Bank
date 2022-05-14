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

// MerchantService handles communication with Bhojpur Bank API
type MerchantService struct {
	client *Client
}

// Get returns the Merchant details for the current client.
func (s *MerchantService) Get(id string) (*types.Merchant, *Response, error) {
	path := fmt.Sprintf("/v1/merchants/%s", id)
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var merchant *types.Merchant
	resp, err := s.client.Do(req, &merchant)
	if err != nil {
		return merchant, resp, err
	}

	return merchant, resp, nil
}

// MerchantLocation returns an individual Merchant location based on the merchant ID and location ID.
func (s *MerchantService) MerchantLocation(mID, lID string) (*types.MerchantLocation, *Response, error) {
	path := fmt.Sprintf("/v1/merchants/%s/locations/%s", mID, lID)
	req, err := s.client.NewAPIRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	merLoc := new(types.MerchantLocation)
	resp, err := s.client.Do(req, merLoc)
	if err != nil {
		return nil, resp, err
	}

	return merLoc, resp, err
}