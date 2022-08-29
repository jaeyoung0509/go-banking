package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	KOR = "KOR"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, KOR:
		return true
	}
	return false
}
