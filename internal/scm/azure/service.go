package azure

import "fmt"

func ReleasePullRequests() []PullRequest {
	response, err := NewClient().ReleasePullRequests()

	if err != nil {
		fmt.Println("There's been an error while trying to get information from Pull Requests")
		fmt.Println("Error:", err)
		return make([]PullRequest, 0)
	}

	return response.PullRequests
}
