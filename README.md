# EasySlip | The Ultimate Golang Library for Payment Slip Verification

**EasySlip** is your go-to **Golang library** for secure and efficient **payment slip verification**, supporting:
- **File uploads**
- **Base64 data**
- **String payloads**
- **TrueMoney Wallet integration**

Whether you're a startup or an enterprise, EasySlip streamlines payment validation with **accuracy**, **reliability**, and **scalability**.

> 🌏 **ภาษาไทย**: EasySlip เป็นไลบรารี Golang ที่ช่วยยืนยันสลิปการชำระเงิน รองรับไฟล์ Base64 String Payload และ TrueMoney Wallet สำหรับธุรกิจทุกขนาด

---

## 🚀 Why Choose EasySlip?

- **Comprehensive Verification**: Supports multiple methods to suit diverse use cases.
- **Seamless Integration**: Plug-and-play with any Golang application.
- **Robust Security**: Built with enterprise-grade encryption and API security.
- **Scalable**: Handles high transaction volumes with ease.
- **Error Transparency**: Provides clear and actionable error messages.

> **Keywords for SEO**: Golang library, payment slip verification, TrueMoney Wallet, Base64 validation, secure APIs.

---

## 🔧 Installation

Add EasySlip to your Golang project:

```bash
go get github.com/lambogreny/easyslip
```

---

## 📂 Project Structure

| File                           | Description                                                   |
|--------------------------------|---------------------------------------------------------------|
| `pkg/easyslip/models.go`       | Defines data models for requests and responses.               |
| `pkg/easyslip/client.go`       | Manages API client creation and HTTP requests.                |
| `pkg/easyslip/verify.go`       | Handles general slip verification through file uploads.       |
| `pkg/easyslip/verify_by_base64.go` | Processes Base64-encoded slip verification.                |
| `pkg/easyslip/verify_by_payload.go`| Verifies slips using string payloads.                    |
| `pkg/easyslip/verify_truewallet.go`| Manages TrueMoney Wallet slip verification.               |
| `pkg/easyslip/errors.go`       | Provides error-handling utilities.                           |

---

## 🧑‍💻 Usage Examples

### 1️⃣ Initialize the Client

```go
client := easyslip.NewClient("https://api.easyslip.com", "your_api_key")
```

---

### 2️⃣ Verify a Payment Slip Using a File

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

---

### 3️⃣ Verify a Slip Using Base64 Data

```go
base64Data := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD..."
resp, err := client.VerifyByBase64(context.Background(), base64Data)
if err != nil {
    panic(err)
}

fmt.Printf("Verification result: %+v
", resp)
```

---

### 4️⃣ Verify a Slip Using String Payload

```go
payload := "your_encoded_payload_string"

resp, err := client.VerifySlipByPayload(context.Background(), payload)
if err != nil {
    panic(err)
}

fmt.Printf("Payload verification result: %+v
", resp)
```

---

### 5️⃣ Verify a TrueMoney Wallet Slip

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

## ⚙️ Error Handling

```go
if err != nil {
    fmt.Printf("Error: %s
", err)
}
```

---

## 🛠️ Debugging Tips

- **Check API Keys**: Ensure your API key is valid.
- **Inspect HTTP Requests**: Use tools like Postman, curl, or Wireshark.
- **Enable Logging**: Add debug flags in your application.

---

## 🤝 Contributions

We welcome developers to contribute to EasySlip! Here's how:

1. Fork the repository.
2. Create a branch for your changes.
3. Commit your updates.
4. Open a pull request.

> **ภาษาไทย**: ร่วมพัฒนา EasySlip ได้ง่ายๆ!

---

## 📜 License

This project is licensed under the **MIT License**.

---

## 📞 Support

For assistance, contact:
- **Email**: support@easyslip.com
- **Documentation**: [EasySlip Docs](https://document.easyslip.com)

> 🔥 Maximize your payment verification efficiency today with EasySlip!
