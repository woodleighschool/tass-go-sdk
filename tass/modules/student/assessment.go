package tassstudent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetAllActivities, path=/{cmpy_code}/assessment/activities/{year}/{semester}
func (c *Client) GetAllActivities(ctx context.Context, year string, semester string) ([]ActivityResponse, error) {
	var result []ActivityResponse
	url := fmt.Sprintf("/assessment/activities/%s/%s", year, semester)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllActivityStudents, path=/{cmpy_code}/assessment/activities/{activity_id}/students
func (c *Client) GetAllActivityStudents(ctx context.Context, activityID int) ([]ActivityStudentResponse, error) {
	var result []ActivityStudentResponse
	url := fmt.Sprintf("/assessment/activities/%d/students", activityID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetStudentActivityResult, path=/{cmpy_code}/assessment/activities/{activity_id}/students/{stud_code}/results
func (c *Client) GetStudentActivityResult(ctx context.Context, activityID int, studentCode string) (ActivityStudentResultsResponse, error) {
	var result ActivityStudentResultsResponse
	url := fmt.Sprintf("/assessment/activities/%d/students/%s/results", activityID, studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return ActivityStudentResultsResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return ActivityStudentResultsResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateStudentActivityResults, path=/{cmpy_code}/assessment/activities/{activity_id}/students/{stud_code}/results
func (c *Client) UpdateStudentActivityResults(ctx context.Context, activityID int, studentCode string, payload UpdateActivityStudentResultsRequest) error {
	url := fmt.Sprintf("/assessment/activities/%d/students/%s/results", activityID, studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}
