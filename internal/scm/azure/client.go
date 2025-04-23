package azure

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"release-handler/config"
)

const username = "azure" //Azure ignores username as it only uses the PAT

type Client struct {
	BaseURL    string
	AuthHeader string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL:    fmt.Sprintf("https://dev.azure.com/%s/%s/_apis", viper.GetString(config.AzureOrganization), viper.GetString(config.AzureProject)),
		AuthHeader: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, viper.GetString(config.AzureApiKey)))),
		HTTPClient: &http.Client{},
	}
}

func (c *Client) DoRequest(method string, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBytes)
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+c.AuthHeader)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}
