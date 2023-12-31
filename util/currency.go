package util

// Constant For All supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	IDR = "IDR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, IDR:
		return true
	}
	return false
}
