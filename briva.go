package qoin

import (
	"encoding/json"
	"fmt"
)

// BriVaCreateOrder returns BRI VA Create Order API endpoint
func BriVaCreateOrder(body map[string]interface{}) string {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/bri/order"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/bri/order"
	}

	body["SecretKey"] = secretKey
	payload, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[BRI VA Create Order] JSON encode body got error:", err)
	}
	fmt.Println(string(payload))
	// signature := generateSignature(string(payload))

	return endpoint
}
