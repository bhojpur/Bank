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
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (c *Client) ConsentLink(sessionID string) (string, error) {
	claims := c.consentClaims(sessionID)
	tokenString, err := c.generateToken(claims)
	if err != nil {
		return "", err
	}

	pathURL := fmt.Sprintf("/consentimento?client_id=%s&jwt=%s", c.ClientID, tokenString)
	u, err := c.SiteURL.Parse(pathURL)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func (c *Client) consentClaims(sessionID string) jwt.MapClaims {
	now := time.Now()
	if sessionID == "" {
		sessionID = uuid.New().String()
	}
	claims := jwt.MapClaims{
		"aud":              "accounts-hubid@bank.bhojpur.net",
		"client_id":        c.ClientID,
		"exp":              now.Add(time.Hour * time.Duration(2)).Unix(),
		"iat":              now.Unix(),
		"iss":              c.ClientID,
		"jti":              sessionID,
		"nbf":              now.Unix(),
		"redirect_uri":     c.ConsentRedirectURL,
		"session_metadata": map[string]string{"client_session": sessionID},
		"type":             "consent",
	}
	return claims
}
