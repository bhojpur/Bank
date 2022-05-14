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
	"reflect"
	"testing"

	"github.com/bhojpur/bank/pkg/types"
)

func TestCreateEntry(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/upi/8cbeb3d2-750f-4b14-81a1-143ad715c273/entries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		response := `
		{
				 "verification_id": "a123456"
		}`

		fmt.Fprint(w, response)
	})

	input := types.CreateUpiEntryInput{
		AccountID: "8cbeb3d2-750f-4b14-81a1-143ad715c273",
		Key:       "c1@bhojpur.net",
		KeyType:   "email",
	}
	data, _, err := client.Upi.CreateEntry(input, "idempotencyKey123")
	if err != nil {
		t.Errorf("upi.CreateEntry returned error: %v", err)
	}

	expected := CreateUpiEntryOutput{VerificationID: "a123456"}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("upi.CreateEntry returned %+v, expected %+v", data, expected)
	}
}

func TestCreateEntryWithVerification(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/upi/8cbeb3d2-750f-4b14-81a1-143ad715c273/entries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		response := `
		{
				 "id": "abcd123"
		}`

		fmt.Fprint(w, response)
	})

	input := types.CreateUpiEntryInput{
		AccountID:      "8cbeb3d2-750f-4b14-81a1-143ad715c273",
		Key:            "c1@bhojpur.net",
		KeyType:        "email",
		VerificationID: "a123456",
	}
	data, _, err := client.Upi.CreateEntry(input, "idempotencyKey123")
	if err != nil {
		t.Errorf("upi.CreateEntry returned error: %v", err)
	}

	expected := CreateUpiEntryOutput{ID: "abcd123"}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("upi.CreateEntry returned %+v, expected %+v", data, expected)
	}
}
