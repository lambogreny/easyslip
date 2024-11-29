package easyslip

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// VerifyTrueWalletSlip verifies a TrueMoney Wallet slip image
func (c *Client) VerifyTrueWalletSlip(ctx context.Context, image io.Reader, filename string) (*TrueWalletVerificationResponse, error) {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	// Use a goroutine to write data into the pipe
	go func() {
		defer pw.Close()
		defer writer.Close()

		part, err := writer.CreateFormFile("file", filename)
		if err != nil {
			pw.CloseWithError(fmt.Errorf("creating form file: %w", err))
			return
		}
		if _, err := io.Copy(part, image); err != nil {
			pw.CloseWithError(fmt.Errorf("copying image data: %w", err))
			return
		}
	}()

	// Create request
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/verify/truewallet", c.baseURL), pr)
	if err != nil {
		return nil, err
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		var errResp struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, fmt.Errorf("decoding error response: %w", err)
		}
		return nil, NewErrorFromMessage(errResp.Status, errResp.Message)
	}

	// Decode successful response
	var result TrueWalletVerificationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &result, nil
}
