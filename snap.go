package qoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SnapCreateOrder returns the response from Snap Create Order API
func SnapCreateOrder(body map[string]interface{}) map[string]interface{} {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/snap/order"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/snap/order"
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[Snap Create Order] JSON encode body got error:", err)
	}

	payload, err := json.Marshal(map[string]interface{}{
		"ReferenceNo": body["referenceNo"],
		"ReqTime":     body["requestTime"],
		"SecretKey":   secretKey,
	})
	if err != nil {
		fmt.Println("[Snap Create Order] JSON encode payload got error:", err)
	}

	signature := generateSignature(string(payload))

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("[Snap Create Order] Create HTTP request got error:", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[Snap Create Order] Send HTTP request got error:", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("[Snap Create Order] JSON decode response got error:", err)
	}

	return result
}
