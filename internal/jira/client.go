package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"release-handler/config"
)

func releaseVersionIssues(releaseName string) Response {
	jql, err := getJQL(releaseName)

	if err != nil {
		fmt.Errorf("JQL could not be parsed, are you using the correct syntax? Error: %v", err)
	}

	payload := map[string]interface{}{
		"jql": jql,
		"fields": []string{
			"assignee",
			"key",
			"status",
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	url := "https://" + viper.GetString(config.JiraDomain) + ".atlassian.net/rest/api/2/search/jql"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(viper.GetString(config.JiraEmail), viper.GetString(config.JiraApiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var result Response

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	return result
}
