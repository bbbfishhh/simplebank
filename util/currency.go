package util

// constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	CNY = "CNY"
)

// IsSupportCurrency returns true if the currency is supported
func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, CNY:
		return true
	}
	return false
}
