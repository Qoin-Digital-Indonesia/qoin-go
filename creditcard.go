package qoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreditCardCreateOrder returns the response from Credit Card Create Order API
func CreditCardCreateOrder(body map[string]interface{}) map[string]interface{} {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/credit-card/create-order"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/credit-card/create-order"
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[Credit Card Create Order] JSON encode body got error:", err)
	}

	payload, err := json.Marshal(map[string]interface{}{
		"ReferenceNo": body["reference_no"],
		"ReqTime":     body["request_time"],
		"SecretKey":   secretKey,
	})
	if err != nil {
		fmt.Println("[Credit Card Create Order] JSON encode payload got error:", err)
	}

	signature := generateSignature(string(payload))

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("[Credit Card Create Order] Create HTTP request got error:", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[Credit Card Create Order] Send HTTP request got error:", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("[Credit Card Create Order] JSON decode response got error:", err)
	}

	return result
}

// CreditCardCharge returns the response from Credit Card Charge API
func CreditCardCharge(body map[string]string) map[string]interface{} {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/credit-card/charge"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/credit-card/charge"
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[Credit Card Charge] JSON encode body got error:", err)
	}

	payload, err := json.Marshal(map[string]string{
		"order_no": body["order_no"],
	})
	if err != nil {
		fmt.Println("[Credit Card Charge] JSON encode payload got error:", err)
	}

	signature := generateSignature(string(payload))

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("[Credit Card Charge] Create HTTP request got error:", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[Credit Card Charge] Send HTTP request got error:", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("[Credit Card Charge] JSON decode response got error:", err)
	}

	return result
}

// CreditCardGetStatus returns the response from Credit Card Get Status API
func CreditCardGetStatus(body map[string]string) map[string]interface{} {
	var endpoint string

	switch isProduction {
	case true:
		endpoint = "https://dev-apipg.qoin.id/credit-card/status"
	default:
		endpoint = "https://dev-sandbox-apipg.qoin.id/sandbox/credit-card/status"
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[Credit Card Get Status] JSON encode body got error:", err)
	}

	payload, err := json.Marshal(map[string]string{
		"RefNo":     referenceNumber,
		"ReqTime":   body["ReqTime"],
		"SecretKey": secretKey,
	})
	if err != nil {
		fmt.Println("[Credit Card Get Status] JSON encode payload got error:", err)
	}

	signature := generateSignature(string(payload))

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("[Credit Card Get Status] Create HTTP request got error:", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[Credit Card Get Status] Send HTTP request got error:", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("[Credit Card Get Status] JSON decode response got error:", err)
	}

	return result
}
