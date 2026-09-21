package tassemployee

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// op: GetAllEmployeePDActivities, path: /{cmpy_code}/employees/{emp_code}/pdactivities
func (c *Client) GetAllEmployeePDActivities(ctx context.Context, employeeCode string) ([]EmployeePDActivityResponse, error) {
	var result []EmployeePDActivityResponse
	url := fmt.Sprintf("/employees/%s/pdactivities", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddEmployeePDActivities, path: /{cmpy_code}/employees/{emp_code}/pdactivities
func (c *Client) AddEmployeePDActivities(ctx context.Context, employeeCode string, payload AddEmployeePDActivityRequest) (EmployeePDActivityResponse, error) {
	var result EmployeePDActivityResponse
	url := fmt.Sprintf("/employees/%s/pdactivities", employeeCode)
	if err := payload.validate(); err != nil {
		return EmployeePDActivityResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return EmployeePDActivityResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeePDActivityResponse{}, err
	}
	return EmployeePDActivityResponse{}, nil
}

// op: GetEmployeePDActivityByID, path: /{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}
func (c *Client) GetEmployeePDActivity(employeeCode string, pdActivityCode string) (EmployeePDActivityResponse, error) {
	// TODO: Implementation
	return EmployeePDActivityResponse{}, nil
}

// op: UpdateEmployeePDActivity, path: /{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}
func (c *Client) UpdateEmployeePDActivity(employeeCode string, pdActivityCode string, payload UpdateEmployeePDActivityRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchEmployeePDActivity, path: /{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}
func (c *Client) PatchEmployeePDActivity(employeeCode string, pdActivityCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteEmployeePDActivity, path: /{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}
func (c *Client) DeleteEmployeePDActivity(employeeCode string, pdActivityCode string) error {
	// TODO: Implementation
	return nil
}

// op: DownloadEmployeePDActivityAttachment, path: /{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}/attachments/{field_number}/{attach_id}
func (c *Client) GetEmployeePDActivityAttachment(employeeCode string, pdActivityCode string, udFieldID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteEmployeePDActivityAttachment, path: /{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}/attachments/{field_number}/{attach_id}
func (c *Client) DeleteEmployeePDActivityAttachment(employeeCode string, pdActivityCode string, udFieldID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: AddEmployeePDActivityAttachment, path: /{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}/attachments/{field_number}
func (c *Client) AddEmployeePDActivityAttachment(employeeCode string, pdActivityCode string, udFieldID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}
