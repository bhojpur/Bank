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

type InstitutionService struct {
	client *Client
}

type InstitutionContext string

const (
	AllInstitutions InstitutionContext = "all"
	SPIParticipants InstitutionContext = "spi"
	STRParticipants InstitutionContext = "str"
)

// Get institution info
func (s InstitutionService) Get(context string) (*types.Institution, *Response, error) {

	path := fmt.Sprintf("/v1/institutions/%s", context)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var institution types.Institution
	resp, err := s.client.Do(req, &institution)
	if err != nil {
		return nil, resp, err
	}

	return &institution, resp, err
}

// List institutions
func (s InstitutionService) List(context InstitutionContext) ([]types.Institution, *Response, error) {

	path := fmt.Sprintf("/v1/institutions?context=%s", context)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var institution []types.Institution
	resp, err := s.client.Do(req, &institution)
	if err != nil {
		return nil, resp, err
	}

	return institution, resp, err
}
