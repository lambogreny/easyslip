package easyslip

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// VerifySlipByPayload verifies a bank slip using the QR code payload
func (c *Client) VerifySlipByPayload(ctx context.Context, payload string) (*VerificationResponse, error) {
	// Construct URL with query parameter
	endpoint, err := url.Parse(fmt.Sprintf("%s/verify", c.baseURL))
	if err != nil {
		return nil, fmt.Errorf("parsing base URL: %w", err)
	}

	// Add query parameter
	query := endpoint.Query()
	query.Set("payload", payload)
	endpoint.RawQuery = query.Encode()

	// Create request
	req, err := c.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

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
