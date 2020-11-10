package qoin

// BriVaCreateOrder returns BRI VA Create Order API endpoint
func BriVaCreateOrder() string {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/bri/order"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/bri/order"
	}

	return endpoint
}
