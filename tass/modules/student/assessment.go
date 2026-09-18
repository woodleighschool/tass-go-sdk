package tassstudent

// op: GetAllActivities, path: /{cmpy_code}/assessment/activities/{year}/{semester}
func (c *Client) GetAllActivities(year string, semester string) ([]ActivityResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllActivityStudents, path: /{cmpy_code}/assessment/activities/{activity_id}/students
func (c *Client) GetAllActivityStudents(activityID string) ([]ActivityStudentResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetStudentActivityResult, path: /{cmpy_code}/assessment/activities/{activity_id}/students/{stud_code}/results
func (c *Client) GetStudentActivityResult(activityID string, studentCode string) (ActivityStudentResultsResponse, error) {
	// TODO: Implementation
	return ActivityStudentResultsResponse{}, nil
}

// op: UpdateStudentActivityResults, path: /{cmpy_code}/assessment/activities/{activity_id}/students/{stud_code}/results
func (c *Client) UpdateStudentActivityResults(activityID string, studentCode string, payload UpdateActivityStudentResultsRequest) error {
	// TODO: Implementation
	return nil
}
