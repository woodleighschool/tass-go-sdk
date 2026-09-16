package tassstudent

func (c *Client) GetAllActivities(year string, semester string) ([]ActivityResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllActivityStudents(activityID string) ([]ActivityStudentResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentActivityResult(activityID string, studentCode string) (ActivityStudentResultsResponse, error) {
	// TODO: Implementation
	return ActivityStudentResultsResponse{}, nil
}

func (c *Client) UpdateStudentActivityResults(activityID string, studentCode string, payload UpdateActivityStudentResultsRequest) error {
	// TODO: Implementation
	return nil
}
