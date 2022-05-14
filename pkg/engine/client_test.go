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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/bhojpur/bank/pkg/types"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = NewClient()
	url, _ := url.Parse(server.URL)
	client.AccountURL = url
	client.ApiBaseURL = url
	client.SiteURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func testClientServices(t *testing.T, c *Client) {
	services := []string{
		"Account",
		"Institution",
		"PaymentLink",
		"PaymentInvoice",
		"Upi",
		"Topups",
		"Transfer",
	}

	cp := reflect.ValueOf(c)
	cv := reflect.Indirect(cp)

	for _, s := range services {
		if cv.FieldByName(s).IsNil() {
			t.Errorf("c.%s shouldn't be nil", s)
		}
	}
}

func testClientDefaultURLs(t *testing.T, c *Client) {
	if c.ApiBaseURL == nil || c.ApiBaseURL.String() != prodAPIBaseURL {
		t.Errorf("NewClient ApiBaseURL = %v, expected %v", c.ApiBaseURL, prodAPIBaseURL)
	}

	if c.AccountURL == nil || c.AccountURL.String() != prodAccountURL {
		t.Errorf("NewClient AccountURL = %v, expected %v", c.AccountURL, prodAccountURL)
	}

	if c.SiteURL == nil || c.SiteURL.String() != prodSiteURL {
		t.Errorf("NewClient SiteURL = %v, expected %v", c.AccountURL, prodAccountURL)
	}
}

func testClientDefaults(t *testing.T, c *Client) {
	testClientDefaultURLs(t, c)
	testClientServices(t, c)
}

func TestNewClient(t *testing.T) {
	c, _ := NewClient()
	testClientDefaults(t, c)
}

func TestClientWithSandboxURLs(t *testing.T) {
	c, _ := NewClient(UseSandbox())
	if c.ApiBaseURL == nil || c.ApiBaseURL.String() != sandboxAPIBaseURL {
		t.Errorf("NewClient ApiBaseURL = %v, expected %v", c.ApiBaseURL, sandboxAPIBaseURL)
	}

	if c.AccountURL == nil || c.AccountURL.String() != sandboxAccountURL {
		t.Errorf("NewClient AccountURL = %v, expected %v", c.AccountURL, sandboxAccountURL)
	}

	if c.SiteURL == nil || c.SiteURL.String() != sandboxSiteURL {
		t.Errorf("NewClient SiteURL = %v, expected %v", c.AccountURL, sandboxAccountURL)
	}
}

func TestNewAPIRequest(t *testing.T) {
	c, _ := NewClient()

	inURL, outURL := "/test", prodAPIBaseURL+"/test"
	inBody, outBody := &types.PaymentLinkInput{AccountID: "abc123"},
		`{"account_id":"abc123","items":null,"customer":{"name":""},"closed":false,"payments":null}`+"\n"
	req, _ := c.NewAPIRequest(http.MethodPost, inURL, inBody)

	if req.URL.String() != outURL {
		t.Errorf("NewAPIRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewAPIRequest(%v)Body = %v, expected %v", inBody, string(body), outBody)
	}

	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewAPIRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}
}
