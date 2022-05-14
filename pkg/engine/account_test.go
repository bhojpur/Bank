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

func TestAccountGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/accounts/8cbeb3d2-750f-4b14-81a1-143ad715c273", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		response := `
		{
				 "account_code": "403881",
  			 "branch_code": "1",
  			 "id": "8cbeb3d2-750f-4b14-81a1-143ad715c273",
  			 "owner_document": "31455351881",
  			 "owner_id": "user:08807157-f8e1-439e-a2ec-154ecb4bee13",
  			 "owner_name": "Nome da Usuária",
  			 "created_at": "2019-07-31T19:13:56Z"
		}`

		fmt.Fprint(w, response)
	})

	acct, _, err := client.Account.Get("8cbeb3d2-750f-4b14-81a1-143ad715c273")
	if err != nil {
		t.Errorf("account.Get returned error: %v", err)
	}

	expected := &types.Account{AccountCode: "403881", BranchCode: "1", OwnerDocument: "31455351881", OwnerID: "user:08807157-f8e1-439e-a2ec-154ecb4bee13",
		ID: "8cbeb3d2-750f-4b14-81a1-143ad715c273", OwnerName: "Nome da Usuária", RestrictedFeatures: false, CreatedAt: "2019-07-31T19:13:56Z"}
	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("account.Get returned %+v, expected %+v", acct, expected)
	}
}
