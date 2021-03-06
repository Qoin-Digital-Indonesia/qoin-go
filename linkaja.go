package qoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// LinkAjaCreateOrder returns the response from LinkAja Create Order API
func LinkAjaCreateOrder(body map[string]interface{}) map[string]interface{} {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/espay-order"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/wallet/order"
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[LinkAja Create Order] JSON encode body got error:", err)
	}

	payload, err := json.Marshal(map[string]interface{}{
		"ReferenceNo": body["ReferenceNo"],
		"ReqTime":     body["ReqTime"],
		"SecretKey":   secretKey,
	})
	if err != nil {
		fmt.Println("[LinkAja Create Order] JSON encode payload got error:", err)
	}

	signature := generateSignature(string(payload))

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("[LinkAja Create Order] Create HTTP request got error:", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[LinkAja Create Order] Send HTTP request got error:", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("[LinkAja Create Order] JSON decode response got error:", err)
	}

	return result
}

// LinkAjaGetStatus returns the response from LinkAja Get Status API
func LinkAjaGetStatus(body map[string]string) map[string]interface{} {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/espay-check-status"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/wallet/status"
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[LinkAja Get Status] JSON encode body got error:", err)
	}

	payload, err := json.Marshal(map[string]string{
		"ReferenceNo": body["ReferenceNo"],
		"ReqTime":     body["RequestTime"],
		"SecretKey":   secretKey,
	})
	if err != nil {
		fmt.Println("[LinkAja Get Status] JSON encode payload got error:", err)
	}

	signature := generateSignature(string(payload))

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("[LinkAja Get Status] Create HTTP request got error:", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[LinkAja Get Status] Send HTTP request got error:", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("[LinkAja Get Status] JSON decode response got error:", err)
	}

	return result
}
