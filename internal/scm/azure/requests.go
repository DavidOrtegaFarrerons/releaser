package azure

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"release-handler/config"
	"time"
)

func (c *Client) ReleasePullRequests() (Response, error) {
	repositoryId := viper.GetString(config.AzureRepositoryId)
	now := time.Now()
	lastFifteenDays := now.AddDate(0, 0, -15)

	endpoint := fmt.Sprintf("/git/repositories/%s/pullrequests?searchCriteria.targetRefName=refs/heads/release&searchCriteria.maxTime=%s&searchCriteria.status=all&api-version=7.1",
		repositoryId,
		lastFifteenDays.Format(time.RFC3339),
	)

	data, err := c.DoRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		fmt.Println("There has been an error while trying to get the ReleasePullRequests: "+err.Error(), err)
		fmt.Println("Error:", err)
	}

	var response Response

	err = json.Unmarshal(data, &response)

	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func (c *Client) SetAutocompletionInPullRequest(pullRequestId int) (Response, error) {
	repositoryId := viper.GetString(config.AzureRepositoryId)

	endpoint := fmt.Sprintf("/git/repositories/%s/pullrequests/%d?api-version=7.1", repositoryId, pullRequestId)

	// Payload structure that will be sent in the PATCH request
	//TODO Refactor to struct
	body := map[string]interface{}{
		"completionOptions": map[string]interface{}{
			"autoCompleteIgnoreConfigIds": []int{},
			"bypassPolicy":                false,
			"deleteSourceBranch":          false,
			"mergeCommitMessage":          "Merged PR",
			"mergeStrategy":               1,
			"transitionWorkItems":         true,
		},
		"autoCompleteSetBy": map[string]interface{}{
			"id": viper.GetString(config.AzureUserId),
		},
	}

	// Call the DoRequest method to send the PATCH request
	responseBody, err := c.DoRequest("PATCH", endpoint, body)
	if err != nil {
		return Response{}, fmt.Errorf("request failed: %w", err)
	}

	var response Response
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return Response{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}
