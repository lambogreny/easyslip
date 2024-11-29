# EasySlip

EasySlip is a Golang library designed to simplify the process of verifying bank slips through an HTTP API. It supports various methods to verify slips efficiently and accurately, making it a valuable tool for businesses requiring automated payment validation.

---

## Features

- **Slip Verification**: Send images of payment slips for validation through API calls.
- **Flexible Inputs**: Supports multiple verification methods (base64, payload, and file uploads).
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

### Basic Example

```go
package main

import (
    "context"
    "os"
    "github.com/lambogreny/easyslip/pkg/easyslip"
)

func main() {
    client := easyslip.NewClient("https://api.easyslip.com", nil)

    // Open a slip image
    file, err := os.Open("slip.jpg")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Verify the slip
    resp, err := client.VerifySlip(context.Background(), file, "slip.jpg")
    if err != nil {
        panic(err)
    }

    // Print the result
    fmt.Printf("Verification result: %+v
", resp)
}
```

### Other Methods

- **Verify by Base64**: Use `VerifyByBase64` for slips in base64 format.
- **Verify by Payload**: Use `VerifyByPayload` for JSON payload verification.

---

## API Endpoints

The library communicates with the EasySlip API. Below is an example of a request:

**Endpoint**: `POST /verify`

**Request**:
- **Form Data**: File upload with field `file`

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
