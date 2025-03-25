package azure

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"release-handler/config"
	"time"
)

const username = "azure" //Azure ignores username as it only uses the PAT

func ReleaseMergeRequests() Response {

	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + viper.GetString(config.AzureApiKey)))

	now := time.Now()
	lastFifteenDays := now.AddDate(0, 0, -15)

	url := "https://dev.azure.com/" + viper.GetString(config.AzureOrganization) + "/" + viper.GetString(config.AzureProject) + "/_apis/git/repositories/" + viper.GetString(config.AzureRepositoryId) + "/pullrequests?searchCriteria.targetRefName=refs/heads/release&searchCriteria.maxTime=" + lastFifteenDays.Format(time.RFC3339) + "&searchCriteria.status=all&api-version=7.1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error decoding json:", err)
	}

	return response
}
