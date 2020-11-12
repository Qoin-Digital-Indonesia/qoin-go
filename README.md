# Qoin-Go

## Installation
`go get github.com/Qoin-Digital-Indonesia/qoin-go`

## Usage
### Transaction Description Example
```
import "encoding/json"

type DescriptionDetail struct {
	Item   uint8
	Desc   string
	Amount float32
}

type Description struct {
	Objects []DescriptionDetail
}

var description Description
description.Objects = make([]DescriptionDetail, 0)
description.Objects = append(description.Objects, DescriptionDetail{1, "T-Shirt", 15000})
description.Objects = append(description.Objects, DescriptionDetail{2, "Admin", 5000})
description.Objects = append(description.Objects, DescriptionDetail{3, "Shipping", 1000})
descriptionJSON, err := json.Marshal(description.Objects)
```

### 1. BRI VA
#### a. Create Order
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.BriVaCreateOrder(map[string]interface{}{
    "MerchantNumber":  "<your merchant code>",
    "ReferenceNumber": "<reference number>",
    "Amount":          21000,
    "Currency":        "IDR",
    "Description":     string(descriptionJSON),
    "UserName":        "Giovanni Reinard",
    "UserContact":     "628123456789;giovanni@qoin.id", // format: phone_number;email_address
    "RequestTime":     time.Now().Format("2006-01-02 15:04:05"),
})
```
#### b. Get Status
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.BriVaGetStatus(map[string]string{
    "OrderNumber": "<order number>",
    "ReqTime":     time.Now().Format("2006-01-02 15:04:05"),
})
```

### 2. Credit Card
#### a. Create Order
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.CreditCardCreateOrder(map[string]interface{}{
    "reference_no":   "<reference number>",
    "account_number": "4000000000000002",
    "exp_month":      "12",
    "exp_year":       "2020",
    "card_cvn":       "123",
    "amount":         25000,
    "request_time":   time.Now().Format("2006-01-02 15:04:05"),
    "merchant_code":  "<your merchant code>",
})
```
#### b. Charge
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
response := qoin.CreditCardCharge(map[string]string{}{
    "order_no": "<order number>",
})
```
#### c. Get Status
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
qoin.SetReferenceNumber("<reference number>")
response := qoin.CreditCardGetStatus(map[string]string{
    "MerchantCode":  "<your merchant code>",
    "OrderNo":       "<order number>",
    "ReqTime":       time.Now().Format("2006-01-02 15:04:05"),
})
```

### 3. OVO
#### a. Create Order
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.OvoCreateOrder(map[string]interface{}{
    "Amount":         10,
    "Currency":       "IDR",
    "Description":    string(descriptionJSON),
    "ReqTime":        time.Now().Format("2006-01-02 15:04:05"),
    "MerchantCode":   "<your merchant code>",
    "ReferenceNo":    "<reference number>",
    "CustomerName":   "Giovanni Reinard",
    "CustomerPhone":  "08123456789",
    "CustomerEmail":  "giovanni@qoin.id",
    "WalletType":     "OVO",
})
```
#### b. Get Status
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.OvoGetStatus(map[string]string{
    "RequestTime":  time.Now().Format("2006-01-02 15:04:05"),
    "MerchantCode": "<your merchant code>",
    "ReferenceNo":  "<reference number>",
})
```

### 4. LinkAja
#### a. Create Order
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.LinkAjaCreateOrder(map[string]interface{}{
    "Amount":         10,
    "Currency":       "IDR",
    "Description":    string(descriptionJSON),
    "ReqTime":        time.Now().Format("2006-01-02 15:04:05"),
    "MerchantCode":   "<your merchant code>",
    "ReferenceNo":    "<reference number>",
    "CustomerName":   "Giovanni Reinard",
    "CustomerPhone":  "08123456789",
    "CustomerEmail":  "giovanni@qoin.id",
    "WalletType":     "LINKAJA",
})
```
#### b. Get Status
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.LinkAjaGetStatus(map[string]string{
    "RequestTime":  time.Now().Format("2006-01-02 15:04:05"),
    "MerchantCode": "<your merchant code>",
    "ReferenceNo":  "<reference number>",
})
```

### 5. Snap
#### - Create Order
```
import "github.com/Qoin-Digital-Indonesia/qoin-go"

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.SnapCreateOrder(map[string]interface{}{
    "merchantCode": "<your merchant code>",
    "linkPayment": "12345",
    "referenceNo": "<reference number>",
    "expiredDate": "",
    "requestTime": time.Now().Format("2006-01-02 15:04:05"),
    "currency": "IDR",
    "paymentMethod": "",
    "paymentChannel": "",
    "customerName": "Giovanni Reinard",
    "customerPhone": "628123456789",
    "customerEmail": "giovanni@qoin.id",
    "product": description.Objects,
    "totalPrice": 21000,
})
```