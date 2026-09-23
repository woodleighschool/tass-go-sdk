package tassstudent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CODEGEN(none): op=GetStudentCommunicationRules, path=/{cmpy_code}/students/{stud_code}/communicationrules
func (c *Client) GetStudentCommunicationRules(ctx context.Context, studentCode string) (StudentCommunicationRulesResponse, error) {
	var result StudentCommunicationRulesResponse
	url := fmt.Sprintf("/students/%s/communicationrules", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentCommunicationRulesResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentCommunicationRulesResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetStudentCommunicationRulesByCommTypeCode, path=/{cmpy_code}/students/{stud_code}/communicationrules/{commtype_code}
func (c *Client) GetStudentCommunicationRulesByCommType(ctx context.Context, studentCode string, communicationType string) (StudentCommunicationRulesResponse, error) {
	var result StudentCommunicationRulesResponse
	url := fmt.Sprintf("/students/%s/communicationrules/%s", studentCode, communicationType)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentCommunicationRulesResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentCommunicationRulesResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentCommunicationRules, path=/{cmpy_code}/students/communicationrules
func (c *Client) GetAllStudentCommunicationRules(ctx context.Context) ([]StudentCommunicationRulesResponse, error) {
	var result []StudentCommunicationRulesResponse
	body, err := c.t.request(ctx, http.MethodGet, "/students/communicationrules", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentCommunicationRulesByCommType, path=/{cmpy_code}/students/communicationrules/{commtype_code}
func (c *Client) GetAllStudentCommunicationRulesByCommType(ctx context.Context, communicationType string) ([]StudentCommunicationRulesResponse, error) {
	var result []StudentCommunicationRulesResponse
	url := fmt.Sprintf("/students/communicationrules/%s", communicationType)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
