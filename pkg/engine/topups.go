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

// TopupsService handles communication with Bhojpur Bank API
type TopupsService struct {
	client *Client
}

// ListGameProviders list all game providers
func (s *TopupsService) ListGameProviders() (*types.Providers, *Response, error) {
	const path = "/v1/topups/games/providers"

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var providers types.Providers
	resp, err := s.client.Do(req, &providers)
	if err != nil {
		return nil, resp, err
	}

	return &providers, resp, err
}

// GetValuesFromGameProvider list all values from a game provider
func (s *TopupsService) GetValuesFromGameProvider(id int) (*types.Products, *Response, error) {
	path := fmt.Sprintf("/v1/topups/games/values/%v", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var products types.Products
	resp, err := s.client.Do(req, &products)
	if err != nil {
		return nil, resp, err
	}

	return &products, resp, err
}
