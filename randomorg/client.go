package randomorg

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	basicAPIURL = "https://api.random.org/json-rpc/2/invoke"
	apiMin      = 0
	apiMax      = 1000000
)

type Client struct {
	apiKey string
	client *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) buildRequestBody(ctx context.Context, method string, params map[string]any) (*http.Request, error) {
	reqUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "could not generate request ID")
	}

	reqBody := request{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      reqUUID.String(),
	}

	bodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, basicAPIURL, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}
