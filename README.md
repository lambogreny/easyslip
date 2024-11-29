# EasySlip

EasySlip is a Golang library designed to simplify the process of verifying bank slips through an HTTP API. It supports various methods to verify slips efficiently and accurately, making it a valuable tool for businesses requiring automated payment validation.

---

## Features

- **Slip Verification**: Send images of payment slips for validation through API calls.
- **Flexible Inputs**: Supports multiple verification methods (base64, payload, and file uploads).
- **TrueMoney Wallet Support**: Verify TrueMoney Wallet slips with ease.
- **Error Handling**: Provides clear error messages for failed verifications.
- **Seamless Integration**: Easily integrate with any Golang application.

---

## Installation

To use EasySlip, add it to your Go module:

```bash
go get github.com/lambogreny/easyslip
```

---

## Usage

### Initializing the Client with an API Key

To use EasySlip, you need to initialize the client with your API key:

```go
client := easyslip.NewClient("https://api.easyslip.com", "your_api_key")
```

### Verification Examples

#### Verify Slip Using File

```go
file, err := os.Open("slip.jpg")
if err != nil {
    panic(err)
}
defer file.Close()

resp, err := client.VerifySlip(context.Background(), file, "slip.jpg")
if err != nil {
    panic(err)
}

fmt.Printf("Verification result: %+v
", resp)
```

#### Verify Slip Using Base64

```go
base64Data := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD..."
resp, err := client.VerifyByBase64(context.Background(), base64Data)
if err != nil {
    panic(err)
}

fmt.Printf("Verification result: %+v
", resp)
```

#### Verify Slip Using Payload

```go
payload := easyslip.Payload{
    TransactionID: "1234567890",
    Amount:        150.00,
    Timestamp:     "2024-11-29T10:00:00Z",
}
resp, err := client.VerifyByPayload(context.Background(), payload)
if err != nil {
    panic(err)
}

fmt.Printf("Verification result: %+v
", resp)
```

#### Verify TrueMoney Wallet Slip

```go
file, err := os.Open("truemoney_slip.jpg")
if err != nil {
    panic(err)
}
defer file.Close()

resp, err := client.VerifyTrueWalletSlip(context.Background(), file, "truemoney_slip.jpg")
if err != nil {
    panic(err)
}

fmt.Printf("TrueMoney Wallet verification result: %+v
", resp)
```

---

## API Endpoints

The library communicates with the EasySlip API. Below is an example of a request:

**Endpoint**: `POST /verify`

**Request**:
- **Form Data**: File upload with field `file`
- **Base64**: JSON payload with a base64-encoded string
- **Payload**: JSON payload containing transaction details

**Response**:
```json
{
   "status": 200,
   "message": "Verification successful",
   "data": {
      "verified": true
   }
}
```

---

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

---

## License

This project is licensed under the MIT License.

---

## Contact

For more information or support, reach out to:
- Email: support@easyslip.com
- Website: [EasySlip](https://easyslip.com)
