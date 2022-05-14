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

//Account represents a Bhojpur Bank Account
type Account struct {
	AccountCode        string `json:"account_code"`
	BranchCode         string `json:"branch_code"`
	ID                 string `json:"id"`
	OwnerDocument      string `json:"owner_document"`
	OwnerID            string `json:"owner_id"`
	OwnerName          string `json:"owner_name"`
	RestrictedFeatures bool   `json:"restricted_features"`
	SortCode           string `json:"sortCode"`
	Currency           string `json:"currency"`
	IBAN               string `json:"iban"`
	BIC                string `json:"bic"`
	CreatedAt          string `json:"created_at,omitempty"`
}

//Balance represents a Bhojpur Bank Account Balance
type Balance struct {
	Cleared          float64 `json:"clearedBalance"`
	Effective        float64 `json:"effectiveBalance"`
	PendingTxns      float64 `json:"pendingTransactions"`
	Available        float64 `json:"availableToSpend"`
	Overdraft        float64 `json:"acceptedOverdraft"`
	Currency         string  `json:"currency"`
	Amount           float64 `json:"amount"`
	BlockedBalance   float64 `json:"blocked_balance"`
	ScheduledBalance float64 `json:"scheduled_balance"`
}

type Statement struct {
	ID                      string  `json:"id"`
	Type                    string  `json:"type"`
	Currency                string  `json:"currency"`
	Amount                  float64 `json:"amount"`
	BalanceAfter            int     `json:"balance_after,omitempty"`
	BalanceBefore           int     `json:"balance_before,omitempty"`
	CreatedAt               string  `json:"created_at,omitempty"`
	UpdatedAt               string  `json:"updated_at,omitempty"`
	Status                  string  `json:"status,omitempty"`
	Operation               string  `json:"operation,omitempty"`
	OperationID             string  `json:"operation_id,omitempty"`
	Description             string  `json:"description,omitempty"`
	OperationAmount         float64 `json:"operation_amount,omitempty"`
	FeeAmount               float64 `json:"fee_amount,omitempty"`
	RefundReasonCode        string  `json:"refund_reason_code,omitempty"`
	RefundReasonDescription string  `json:"refund_reason_description,omitempty"`
	OriginalOperationID     string  `json:"original_operation_id,omitempty"`
	RefundedAt              string  `json:"refunded_at,omitempty"`
	Barcode                 string  `json:"barcode,omitempty"`

	CardNetworkCode string `json:"card_network_code,omitempty"`
	CardNetworkName string `json:"card_network_name,omitempty"`
	CardType        string `json:"card_type,omitempty"`
	IsPrepayment    bool   `json:"is_prepayment,omitempty"`

	Details struct {
		BankName         string `json:"bank_name,omitempty"`
		RecipientCpfCnpj string `json:"recipient_cpf_cnpj,omitempty"`
		RecipientName    string `json:"recipient_name,omitempty"`
		WritableLine     string `json:"writable_line,omitempty"`
		ExpirationDate   string `json:"expiration_date,omitempty"`
	} `json:"details,omitempty"`

	CounterParty CounterParty `json:"counter_party,omitempty"`

	DelayedToNextBusinessDay bool `json:"delayed_to_next_business_day,omitempty"`
}

type CounterParty struct {
	Account struct {
		Institution     string `json:"institution,omitempty"`
		InstitutionName string `json:"institution_name,omitempty"`
		AccountCode     string `json:"account_code,omitempty"`
		BranchCode      string `json:"branch_code,omitempty"`
		AccountType     string `json:"account_type,omitempty"`
	} `json:"account"`
	Entity Entity `json:"entity"`
}

type Fee struct {
	Currency                    string  `json:"currency"`
	Amount                      float64 `json:"amount"`
	FeeType                     string  `json:"fee_type"`
	BillingExemptionParticipant bool    `json:"billing_exemption_participant"`
	OriginalFee                 int     `json:"original_fee"`
	MaxFreeTransfers            int     `json:"max_free_transfers"`
	RemainingFreeTransfers      int     `json:"remaining_free_transfers"`
}

func ListFeeTypes() []string {
	return []string{
		"internal_transfer",
		"external_transfer",
		"barcode_payment",
		"outbound_bhojpur_prepaid_card_withdrawal",
		"barcode_payment_invoice",
	}
}
