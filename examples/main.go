package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/lambogreny/easyslip/pkg/easyslip"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// Create client without setting timeout in HTTP client
	client := easyslip.NewClient(
		"YOUR_TOKEN",
		easyslip.WithUserAgent("your-app/1.0.0"),
	)

	// ตัวอย่างการตรวจสอบสลิปโดยการส่งภาพ
	verifySlipExample(client)

	// ตัวอย่างการตรวจสอบสลิป TrueMoney Wallet
	verifyTrueWalletSlipExample(client)

	// ตัวอย่างการตรวจสอบสลิปโดยการส่ง Payload
	verifyByPayloadExample(client)

	// ตัวอย่างการตรวจสอบสลิปโดยการส่งภาพ Base64
	verifyByBase64Example(client)
}

func verifySlipExample(client *easyslip.Client) {
	// Open slip image
	file, err := os.Open("slip.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Verify slip
	resp, err := client.VerifySlip(ctx, file, "slip.jpg")
	if err != nil {
		var apiErr *easyslip.Error
		if errors.As(err, &apiErr) {
			log.Printf("API error: %s (status: %d)", apiErr.Message, apiErr.Status)
			return
		}
		log.Fatal(err)
	}

	// Print verification result
	fmt.Println("=== Verify Slip ===")
	fmt.Printf("Transaction Reference: %s\n", resp.Data.TransRef)
	fmt.Printf("Amount: %.2f\n", resp.Data.Amount.Amount)
	fmt.Printf("Date: %s\n", resp.Data.Date.Format(time.RFC3339))
	fmt.Printf("Sender: %s\n", resp.Data.Sender.Account.Name.Thai)
	fmt.Printf("Receiver: %s\n", resp.Data.Receiver.Account.Name.Thai)
}

func verifyTrueWalletSlipExample(client *easyslip.Client) {
	// Open TrueMoney Wallet slip image
	file, err := os.Open("truemoney_slip.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Verify TrueMoney Wallet slip
	resp, err := client.VerifyTrueWalletSlip(ctx, file, "truemoney_slip.jpg")
	if err != nil {
		var apiErr *easyslip.Error
		if errors.As(err, &apiErr) {
			log.Printf("API error: %s (status: %d)", apiErr.Message, apiErr.Status)
			return
		}
		log.Fatal(err)
	}

	// Print verification result
	fmt.Println("=== Verify TrueMoney Wallet Slip ===")
	fmt.Printf("Transaction ID: %s\n", resp.Data.TransactionID)
	fmt.Printf("Amount: %.2f\n", resp.Data.Amount)
	fmt.Printf("Date: %s\n", resp.Data.Date.Format(time.RFC3339))
	fmt.Printf("Sender Name: %s\n", resp.Data.Sender.Name)
	fmt.Printf("Receiver Name: %s\n", resp.Data.Receiver.Name)
	fmt.Printf("Receiver Phone: %s\n", resp.Data.Receiver.Phone)
}

func verifyByPayloadExample(client *easyslip.Client) {
	// ตัวอย่าง Payload ที่ได้จากการสแกน QR Code
	payload := "YOUR_PAYLOAD"

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Verify slip by payload
	resp, err := client.VerifySlipByPayload(ctx, payload)
	if err != nil {
		var apiErr *easyslip.Error
		if errors.As(err, &apiErr) {
			log.Printf("API error: %s (status: %d)", apiErr.Message, apiErr.Status)
			return
		}
		log.Fatal(err)
	}

	// Print verification result
	fmt.Println("=== Verify Slip By Payload ===")
	fmt.Printf("Transaction Reference: %s\n", resp.Data.TransRef)
	fmt.Printf("Amount: %.2f\n", resp.Data.Amount.Amount)
	fmt.Printf("Date: %s\n", resp.Data.Date.Format(time.RFC3339))
	fmt.Printf("Sender: %s\n", resp.Data.Sender.Account.Name.Thai)
	fmt.Printf("Receiver: %s\n", resp.Data.Receiver.Account.Name.Thai)
}

func verifyByBase64Example(client *easyslip.Client) {
	// Read image file
	imageBytes, err := ioutil.ReadFile("slip.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// Encode image to base64
	base64Image := base64.StdEncoding.EncodeToString(imageBytes)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Verify slip by base64 image
	resp, err := client.VerifySlipByBase64(ctx, base64Image)
	if err != nil {
		var apiErr *easyslip.Error
		if errors.As(err, &apiErr) {
			log.Printf("API error: %s (status: %d)", apiErr.Message, apiErr.Status)
			return
		}
		log.Fatal(err)
	}

	// Print verification result
	fmt.Println("=== Verify Slip By Base64 Image ===")
	fmt.Printf("Transaction Reference: %s\n", resp.Data.TransRef)
	fmt.Printf("Amount: %.2f\n", resp.Data.Amount.Amount)
	fmt.Printf("Date: %s\n", resp.Data.Date.Format(time.RFC3339))
	fmt.Printf("Sender: %s\n", resp.Data.Sender.Account.Name.Thai)
	fmt.Printf("Receiver: %s\n", resp.Data.Receiver.Account.Name.Thai)
}
