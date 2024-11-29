package easyslip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// VerifySlipByBase64 verifies a bank slip image encoded in Base64
func (c *Client) VerifySlipByBase64(ctx context.Context, base64Image string) (*VerificationResponse, error) {
	// Create request body
	requestBody := map[string]string{
		"image": base64Image,
	}
	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("encoding request body: %w", err)
	}

	// Create request
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/verify", c.baseURL), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

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
	var result VerificationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &result, nil
}
