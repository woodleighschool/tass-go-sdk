package tassemployee

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetAllEmployeePDActivities, path=/{cmpy_code}/employees/{emp_code}/pdactivities
func (c *Client) GetAllEmployeePDActivities(ctx context.Context, employeeCode string) ([]EmployeePDActivityResponse, error) {
	var result []EmployeePDActivityResponse
	url := fmt.Sprintf("/employees/%s/pdactivities", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeePDActivities, path=/{cmpy_code}/employees/{emp_code}/pdactivities
func (c *Client) AddEmployeePDActivities(ctx context.Context, employeeCode string, payload AddEmployeePDActivityRequest) (EmployeePDActivityResponse, error) {
	var result EmployeePDActivityResponse
	url := fmt.Sprintf("/employees/%s/pdactivities", employeeCode)
	if err := tasscommon.Validate(payload); err != nil {
		return EmployeePDActivityResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return EmployeePDActivityResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeePDActivityResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetEmployeePDActivityByID, path=/{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}
func (c *Client) GetEmployeePDActivity(ctx context.Context, employeeCode string, pdActivityCode int) (EmployeePDActivityResponse, error) {
	var result EmployeePDActivityResponse
	url := fmt.Sprintf("/employees/%s/pdactivities/%d", employeeCode, pdActivityCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return EmployeePDActivityResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeePDActivityResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateEmployeePDActivity, path=/{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}
func (c *Client) UpdateEmployeePDActivity(ctx context.Context, employeeCode string, pdActivityCode int, payload UpdateEmployeePDActivityRequest) error {
	url := fmt.Sprintf("/employees/%s/pdactivities/%d", employeeCode, pdActivityCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchEmployeePDActivity, path=/{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}
func (c *Client) PatchEmployeePDActivity(ctx context.Context, employeeCode string, pdActivityCode int, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/employees/%s/pdactivities/%d", employeeCode, pdActivityCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteEmployeePDActivity, path=/{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}
func (c *Client) DeleteEmployeePDActivity(ctx context.Context, employeeCode string, pdActivityCode int) error {
	url := fmt.Sprintf("/employees/%s/pdactivities/%d", employeeCode, pdActivityCode)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DownloadEmployeePDActivityAttachment, path=/{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}/attachments/{field_number}/{attach_id}
func (c *Client) GetEmployeePDActivityAttachment(ctx context.Context, employeeCode string, pdActivityCode int, udFieldID int, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/employees/%s/pdactivities/%d/attachments/%d/%s", employeeCode, pdActivityCode, udFieldID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DeleteEmployeePDActivityAttachment, path=/{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}/attachments/{field_number}/{attach_id}
func (c *Client) DeleteEmployeePDActivityAttachment(ctx context.Context, employeeCode string, pdActivityCode int, udFieldID int, attachmentID string) error {
	url := fmt.Sprintf("/employees/%s/pdactivities/%d/attachments/%d/%s", employeeCode, pdActivityCode, udFieldID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=AddEmployeePDActivityAttachment, path=/{cmpy_code}/employees/{emp_code}/pdactivities/{pdact_num}/attachments/{field_number}
func (c *Client) AddEmployeePDActivityAttachment(ctx context.Context, employeeCode string, pdActivityCode int, udFieldID int, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/employees/%s/pdactivities/%d/attachments/%d", employeeCode, pdActivityCode, udFieldID)
	body, err := c.t.upload(ctx, url, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}
