package structs

// Credentials which stores google ids.
type TransactionDetails struct {
	Account       string
	Address       string
	Category      string
	Amount        float64
	Fee           float64
	Confirmations int64
	TimeReceived  int64
}
