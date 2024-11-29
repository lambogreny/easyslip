package easyslip

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient defines the interface for an HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	defaultBaseURL   = "https://developer.easyslip.com/api/v1"
	defaultUserAgent = "easyslip-go/1.0.0"
)

// Client represents an API client for EasySlip services
type Client struct {
	baseURL    string
	httpClient HTTPClient
	token      string
	userAgent  string
}

// ClientOption is a function that modifies a Client
type ClientOption func(*Client)

// NewClient creates a new EasySlip API client
func NewClient(token string, opts ...ClientOption) *Client {
	c := &Client{
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		token:     token,
		userAgent: defaultUserAgent,
	}

	// Apply options
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient HTTPClient) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithUserAgent sets a custom user agent
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// newRequest creates a new HTTP request with the necessary headers
func (c *Client) newRequest(ctx context.Context, method, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, urlStr, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("User-Agent", c.userAgent)

	return req, nil
}
