package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	INR = "INR"
)

func IsSupportedCurency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, INR:
		return true
	}
	return false
}
