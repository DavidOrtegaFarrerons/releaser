package azure

func GetFinalReviewStatus(reviewers []Reviewer) string {
	hasRejected := false
	hasWaitingAuthor := false
	hasApproval := false

	for _, reviewer := range reviewers {
		switch reviewer.Vote {
		case -10:
			hasRejected = true
		case -5:
			hasWaitingAuthor = true
		case 10, 5:
			hasApproval = true
		}
	}

	if hasRejected {
		return "Rejected"
	}
	if hasWaitingAuthor {
		return "Waiting Author"
	}
	if hasApproval {
		return "Approved"
	}

	return "Not Approved"
}
