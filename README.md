# EasySlip

EasySlip is a powerful and flexible Golang library designed to streamline the process of verifying payment slips through various methods, including file uploads, Base64 data, JSON payloads, and TrueMoney Wallet integration. This library makes it easier for businesses to integrate payment verification into their systems with high accuracy and reliability.

---

## Features

- **Comprehensive Slip Verification**: Supports file, Base64, payload, and TrueMoney Wallet slip verification.
- **High Compatibility**: Easily integrates with any Golang application.
- **Robust Error Handling**: Provides detailed error messages to help diagnose and resolve issues.
- **Secure and Scalable**: Designed for enterprise use with secure API interactions.

---

## Installation

To start using EasySlip, simply add it to your project:

```bash
go get github.com/lambogreny/easyslip
```

---

## Project Structure

- `pkg/easyslip/models.go`: Defines data structures for requests and responses.
- `pkg/easyslip/client.go`: Manages API client creation and HTTP requests.
- `pkg/easyslip/verify.go`: General slip verification through file uploads.
- `pkg/easyslip/verify_truewallet.go`: Verification for TrueMoney Wallet slips.
- `pkg/easyslip/verify_by_base64.go`: Verification through Base64-encoded data.
- `pkg/easyslip/verify_by_payload.go`: Verification using JSON payloads.
- `pkg/easyslip/errors.go`: Handles errors returned by the API.

---

## Usage Examples

### 1. Initializing the Client

```go
client := easyslip.NewClient("https://api.easyslip.com", "your_api_key")
```

### 2. Verifying a Slip Using a File

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

### 3. Verifying a Slip Using Base64 Data

```go
base64Data := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD..."
resp, err := client.VerifyByBase64(context.Background(), base64Data)
if err != nil {
    panic(err)
}

fmt.Printf("Verification result: %+v
", resp)
```

### 4. Verifying a Slip Using JSON Payload

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

fmt.Printf("Payload verification result: %+v
", resp)
```

### 5. Verifying a TrueMoney Wallet Slip

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

## Error Handling

EasySlip provides detailed error responses for every API call:

```go
if err != nil {
    fmt.Printf("Error: %s
", err)
}
```

---

## Debugging Tips

- **Check API Keys**: Ensure your API key is valid and properly configured.
- **Inspect HTTP Requests**: Use tools like Postman, curl, or Wireshark to analyze API calls.
- **Enable Logging**: Add debug flags to your application to track detailed logs.

---

## Contributing

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Commit and push your changes.
4. Open a pull request.

---

## License

This project is licensed under the MIT License.

---

## Support

For support, contact:
- Email: support@easyslip.com
- Documentation: [EasySlip Docs](https://document.easyslip.com)
