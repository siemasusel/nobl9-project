package randomorg

import (
	"context"
	"encoding/json"
	"net/http"
	"stddevapi"

	"github.com/pkg/errors"
)

const maxLength = 10_000

func (c *Client) GetRandomIntegers(ctx context.Context, length int) ([]int, error) {
	if length > maxLength {
		return nil, stddevapi.NewValidationError(nil, "length cannot be greater than %d", maxLength)
	}

	params := map[string]any{
		"apiKey": c.apiKey,
		"n":      length,
		"min":    apiMin,
		"max":    apiMax,
	}

	req, err := c.buildRequestBody(ctx, "generateIntegers", params)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody response
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	if err := checkResponsBody(resp, respBody); err != nil {
		return nil, err
	}

	return respBody.Result.Random.Data, nil
}

func checkResponsBody(resp *http.Response, respBody response) error {
	if resp.StatusCode != 200 {
		return errors.Errorf("randomorg API respond with non-200 status code: %d", resp.StatusCode)
	}

	if respBody.Error != nil {
		return ClientError{Code: respBody.Error.Code, Message: respBody.Error.Message}
	}

	if respBody.Result == nil {
		return errors.New("missing result in api response")
	}

	if len(respBody.Result.Random.Data) == 0 {
		return errors.New("recieved empty data array")
	}

	return nil
}

type request struct {
	JSONRPC string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Params  map[string]any `json:"params"`
	ID      string         `json:"id"`
}

type response struct {
	Result *generateIntegersResult `json:"result"`
	Error  *errorResponse          `json:"error"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type generateIntegersResult struct {
	Random struct {
		Data []int `json:"data"`
	} `json:"random"`
}
