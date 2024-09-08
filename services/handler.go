package services

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sparkeexd/hoyoapi/internal/errors"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
)

// Default HTTP client timeout duration.
const clientDefaultTimeout = 10 * time.Second

// Base handler for dealing with HTTP request processes.
// This ranges from sending HTTP requests to HoYoLab endpoints, parsing responses, and setting cookies.
type Handler struct {
	client http.Client
}

// Constructor.
func NewHandler() Handler {
	client := http.DefaultClient
	client.Timeout = clientDefaultTimeout
	client.Transport = &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 2,
	}

	handler := Handler{client: *client}
	return handler
}

// Sends a HTTP request.
// Returns a generic map from the unmarshalled response.
// Specific retcode errors are handled by their respective clients.
func (handler Handler) Send(request Request, res any) error {
	// Build HTTP request.
	httpRequest, err := handler.createHttpRequest(request)
	if err != nil {
		return err
	}

	// Add cookies.
	for _, token := range request.cookie.Tokens() {
		httpRequest.AddCookie(&token)
	}

	// Send HTTP request.
	response, err := handler.client.Do(httpRequest)
	if err != nil {
		return errors.NewError(
			errors.HTTP_REQUEST_SEND_ERROR,
			fmt.Sprintf("URL: %s\nError: %s", request.endpoint, err.Error()),
		)
	}

	defer response.Body.Close()

	// Parse response body into readable format.
	body, err := handler.parseResponse(response)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.NewError(
			response.StatusCode,
			fmt.Sprintf("URL: %s\nError: %+v", request.endpoint, string(body)),
		)
	}

	utilities.UnmarshalJSON(body, &res)

	return err
}

// Create HTTP request structure.
func (handler Handler) createHttpRequest(request Request) (*http.Request, error) {
	var body io.Reader

	// JSON marshal request body.
	if request.body != nil {
		jsonData, err := json.Marshal(request.body)
		if err != nil {
			return nil,
				errors.NewError(
					errors.JSON_SERIALIZATION_ERROR,
					fmt.Sprintf("Error: %s", err.Error()),
				)
		}

		body = bytes.NewBuffer(jsonData)
	}

	// Create HTTP request.
	httpRequest, err := http.NewRequest(request.method, request.endpoint, body)
	if err != nil {
		return nil,
			errors.NewError(
				errors.HTTP_REQUEST_CREATE_ERROR,
				fmt.Sprintf("URL: %s\nError: %s", request.endpoint, err.Error()),
			)
	}

	// Add query parameters to HTTP request.
	query := httpRequest.URL.Query()
	for k, v := range request.params {
		query.Add(k, v)
	}

	httpRequest.URL.RawQuery = query.Encode()

	// Add headers to HTTP request.
	for k, v := range request.headers {
		httpRequest.Header.Set(k, v)
	}

	return httpRequest, nil
}

// Parse response body by decompressing content according to its encoding.
// HoYoLab endpoints uses the standard gzip, deflate and brotli encodings.
func (handler Handler) parseResponse(response *http.Response) ([]byte, error) {
	var err error
	var reader io.ReadCloser

	switch encoding := response.Header.Get("Content-Encoding"); encoding {
	case encodingGzip:
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil,
				errors.NewError(
					errors.HTTP_RESPONSE_READ_ERROR,
					fmt.Sprintf("Error: %s", err.Error()),
				)
		}

	case encodingDeflate:
		reader = flate.NewReader(response.Body)

	default:
		reader = response.Body
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil,
			errors.NewError(
				errors.HTTP_RESPONSE_READ_ERROR,
				fmt.Sprintf("Error: %s", err.Error()),
			)
	}

	return body, nil
}
