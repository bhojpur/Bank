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
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

func (c *Client) Authenticate() error {
	if c.validToken() {
		return nil
	}

	claims := c.authClaims()
	tokenString, err := c.generateToken(claims)
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Set("client_assertion", tokenString)
	data.Set("client_id", c.ClientID)
	data.Set("grant_type", "client_credentials")

	u, err := c.AccountURL.Parse("/auth/realms/bhojpur_bank/protocol/openid-connect/token")
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("user-agent", c.UserAgent)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	var token oauth2.Token
	_, err = c.Do(req, &token)
	if err != nil {
		return err
	}
	ctx := context.Background()
	config := &oauth2.Config{}
	ts := config.TokenSource(ctx, &token)

	c.m.Lock()
	defer c.m.Unlock()
	c.client = oauth2.NewClient(ctx, ts)
	c.token = token

	return nil
}

func (c *Client) authClaims() jwt.MapClaims {
	now := time.Now()
	u, _ := c.AccountURL.Parse("/auth/realms/bhojpur_bank")
	claims := jwt.MapClaims{
		"aud":       u.String(),
		"client_id": c.ClientID,
		"exp":       now.Add(time.Hour * time.Duration(2)).Unix(),
		"iat":       now.Unix(),
		"jti":       uuid.New().String(),
		"iss":       c.ClientID,
		"nbf":       now.Unix(),
		"realm":     "bhojpur_bank",
		"sub":       c.ClientID,
	}
	return claims
}

func (c *Client) validToken() bool {
	if !c.token.Valid() {
		return false
	}

	src := strings.Split(c.token.AccessToken, ".")
	if len(src) != 3 {
		return false
	}

	if l := len(src[1]) % 4; l > 0 {
		src[1] += strings.Repeat("=", 4-l)
	}

	decoded, err := base64.URLEncoding.DecodeString(src[1])
	if err != nil {
		c.log.Error(fmt.Errorf("decoding base64 error %s", err))
		return false
	}

	var output tokenData
	err = json.Unmarshal(decoded, &output)
	if err != nil {
		c.log.Error(fmt.Errorf("decoding json error %s", err))
		return false
	}

	tm := time.Unix(int64(output.Exp), 0)
	remainder := tm.Sub(time.Now())
	if remainder < 30 {
		return false
	}

	return true
}

type tokenData struct {
	Exp int `json:"exp"`
}
