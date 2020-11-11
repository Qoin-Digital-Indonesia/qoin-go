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

var body = map[string]interface{}{
    "MerchantNumber":  "<your merchant code>",
    "ReferenceNumber": "<reference number>",
    "Amount":          21000,
    "Currency":        "IDR",
    "Description":     string(descriptionJSON),
    "UserName":        "Giovanni Reinard",
    "UserContact":     "628123456789;giovanni@qoin.id", // format: phone_number;email_address
    "RequestTime":     time.Now().Format("2006-01-02 15:04:05"),
}

qoin.SetEnvironment("sandbox") // sandbox || production
qoin.SetPrivateKey(`<your private key>`) // must use back quote (`) symbol
qoin.SetSecretKey("<your secret key>")
response := qoin.BriVaCreateOrder(body)
```