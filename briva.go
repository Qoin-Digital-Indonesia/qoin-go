package qoin

import (
	"encoding/json"
	"fmt"
	"strings"
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

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[BRI VA Create Order] JSON encode body got error:", err)
	}

	bodyLastIndex := strings.LastIndex(string(jsonBody), "}")
	trimmedBody := string(jsonBody)[:bodyLastIndex]
	payload := trimmedBody + ",\"SecretKey\":\"" + secretKey + "\"}"
	signature := generateSignature(payload)
	fmt.Println(payload)
	fmt.Println(signature)

	return endpoint
}
