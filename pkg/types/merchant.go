package types

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

// Merchant represents details of a merchant
type Merchant struct {
	ID              string `json:"merchantId"`
	Name            string `json:"name"`
	Website         string `json:"website"`
	PhoneNumber     string `json:"phoneNumber"`
	TwitterUsername string `json:"twitterUsername"`
}

// MerchantLocation represents details of a merchant location
type MerchantLocation struct {
	ID                       string `json:"merchantLocationId"`
	MerchantID               string `json:"merchantId"`
	MerchantName             string `json:"merchantName"`
	LocationName             string `json:"locationName"`
	Address                  string `json:"address"`
	PhoneNumber              string `json:"phoneNumber"`
	GooglePlaceID            string `json:"googlePlaceId"`
	CardMerchantCategoryCode int32  `json:"cardMerchantCategoryCode"`
}
