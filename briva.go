package qoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// BriVaCreateOrder returns the response from BRI VA Create Order API
func BriVaCreateOrder(body map[string]interface{}) map[string]interface{} {
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

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("[BRI VA Create Order] Create HTTP request got error:", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[BRI VA Create Order] Send HTTP request got error:", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("[BRI VA Create Order] JSON decode response got error:", err)
	}

	return result
}

// BriVaGetStatus returns the response from BRI VA Get Status API
func BriVaGetStatus(body map[string]string) map[string]interface{} {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/bri/paymentstatus"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/bri/paymentstatus"
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[BRI VA Get Status] JSON encode body got error:", err)
	}

	bodyLastIndex := strings.LastIndex(string(jsonBody), "}")
	trimmedBody := string(jsonBody)[:bodyLastIndex]
	payload := trimmedBody + ",\"SecretKey\":\"" + secretKey + "\"}"
	signature := generateSignature(payload)

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("[BRI VA Get Status] Create HTTP request got error:", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[BRI VA Get Status] Send HTTP request got error:", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("[BRI VA Get Status] JSON decode response got error:", err)
	}

	return result
}
