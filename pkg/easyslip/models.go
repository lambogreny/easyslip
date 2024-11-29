package easyslip

import "time"

// Bank represents bank information
type Bank struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Short string `json:"short,omitempty"`
}

// AccountName represents account holder name in different languages
type AccountName struct {
	Thai    string `json:"th,omitempty"`
	English string `json:"en,omitempty"`
}

// BankAccount represents bank account information
type BankAccount struct {
	Type    string `json:"type"`
	Account string `json:"account"`
}

// ProxyAccount represents proxy payment information
type ProxyAccount struct {
	Type    string `json:"type"`
	Account string `json:"account"`
}

// Account represents complete account information
type Account struct {
	Name  AccountName   `json:"name"`
	Bank  *BankAccount  `json:"bank,omitempty"`
	Proxy *ProxyAccount `json:"proxy,omitempty"`
}

// Amount represents transaction amount information
type Amount struct {
	Amount float64 `json:"amount"`
	Local  struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"local"`
}

// PartyInfo represents sender or receiver information
type PartyInfo struct {
	Bank       Bank    `json:"bank"`
	Account    Account `json:"account"`
	MerchantID string  `json:"merchantId,omitempty"`
}

// VerificationResponse represents the API response for slip verification
type VerificationResponse struct {
	Status int `json:"status"`
	Data   struct {
		Payload     string    `json:"payload"`
		TransRef    string    `json:"transRef"`
		Date        time.Time `json:"date"`
		CountryCode string    `json:"countryCode"`
		Amount      Amount    `json:"amount"`
		Fee         float64   `json:"fee,omitempty"`
		Ref1        string    `json:"ref1,omitempty"`
		Ref2        string    `json:"ref2,omitempty"`
		Ref3        string    `json:"ref3,omitempty"`
		Sender      PartyInfo `json:"sender"`
		Receiver    PartyInfo `json:"receiver"`
	} `json:"data"`
}

// TrueWalletVerificationResponse represents the API response for TrueMoney Wallet slip verification
type TrueWalletVerificationResponse struct {
	Status int `json:"status"`
	Data   struct {
		TransactionID string    `json:"transactionId"`
		Date          time.Time `json:"date"`
		Amount        float64   `json:"amount"`
		Sender        struct {
			Name string `json:"name"`
		} `json:"sender"`
		Receiver struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
		} `json:"receiver"`
	} `json:"data"`
}
