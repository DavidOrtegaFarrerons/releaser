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
