package tassemployee

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetEmployeeByCode, path=/{cmpy_code}/employees/{emp_code}
func (c *Client) GetEmployee(ctx context.Context, employeeCode string) (EmployeeResponse, error) {
	var result EmployeeResponse
	url := fmt.Sprintf("/employees/%s", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return EmployeeResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateEmployee, path=/{cmpy_code}/employees/{emp_code}
func (c *Client) UpdateEmployee(ctx context.Context, employeeCode string, payload UpdateEmployeeRequest) error {
	url := fmt.Sprintf("/employees/%s", employeeCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchEmployee, path=/{cmpy_code}/employees/{emp_code}
func (c *Client) PatchEmployee(ctx context.Context, employeeCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/employees/%s", employeeCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllEmployees, path=/{cmpy_code}/employees
func (c *Client) GetAllEmployees(ctx context.Context) ([]EmployeeResponse, error) {
	var result []EmployeeResponse
	body, err := c.t.request(ctx, http.MethodGet, "/employees", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployee, path=/{cmpy_code}/employees
func (c *Client) AddEmployee(ctx context.Context, payload AddEmployeeRequest) (EmployeeResponse, error) {
	var result EmployeeResponse
	if err := tasscommon.Validate(payload); err != nil {
		return EmployeeResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, "/employees", nil, payload, http.StatusCreated)
	if err != nil {
		return EmployeeResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllEmployeeStandardNotes, path=/{cmpy_code}/employees/{emp_code}/notes/standard
func (c *Client) GetAllEmployeeStandardNotes(ctx context.Context, employeeCode string) ([]EmployeeStandardNoteResponse, error) {
	var result []EmployeeStandardNoteResponse
	url := fmt.Sprintf("/employees/%s/notes/standard", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeeStandardNote, path=/{cmpy_code}/employees/{emp_code}/notes/standard
func (c *Client) AddEmployeeStandardNote(ctx context.Context, employeeCode string, payload AddEmployeeStandardNoteRequest) (EmployeeStandardNoteResponse, error) {
	var result EmployeeStandardNoteResponse
	url := fmt.Sprintf("/employees/%s/notes/standard", employeeCode)
	if err := tasscommon.Validate(payload); err != nil {
		return EmployeeStandardNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return EmployeeStandardNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeStandardNoteResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetEmployeeStandardNoteByID, path=/{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}
func (c *Client) GetEmployeeStandardNote(ctx context.Context, employeeCode string, noteID string) (EmployeeStandardNoteResponse, error) {
	var result EmployeeStandardNoteResponse
	url := fmt.Sprintf("/employees/%s/notes/standard/%s", employeeCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return EmployeeStandardNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeStandardNoteResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateEmployeeStandardNote, path=/{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}
func (c *Client) UpdateEmployeeStandardNote(ctx context.Context, employeeCode string, noteID string, payload UpdateEmployeeStandardNoteRequest) error {
	url := fmt.Sprintf("/employees/%s/notes/standard/%s", employeeCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchEmployeeStandardNote, path=/{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}
func (c *Client) PatchEmployeeStandardNote(ctx context.Context, employeeCode string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/employees/%s/notes/standard/%s", employeeCode, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteEmployeeStandardNote, path=/{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}
func (c *Client) DeleteEmployeeStandardNote(ctx context.Context, employeeCode string, noteID string) error {
	url := fmt.Sprintf("/employees/%s/notes/standard/%s", employeeCode, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllEmployeeStandardNotesAttachments, path=/{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}/attachments
func (c *Client) GetAllEmployeeStandardNotesAttachments(ctx context.Context, employeeCode string, noteID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/employees/%s/notes/standard/%s/attachments", employeeCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeeStandardNotesAttachments, path=/{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}/attachments
func (c *Client) AddEmployeeStandardNotesAttachments(ctx context.Context, employeeCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/employees/%s/notes/standard/%s/attachments", employeeCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	body, err := c.t.upload(ctx, url, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=DownloadEmployeeStandardNotesAttachment, path=/{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) GetEmployeeStandardNotesAttachment(ctx context.Context, employeeCode string, noteID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/employees/%s/notes/standard/%s/attachments/%s", employeeCode, noteID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DeleteEmployeeStandardNotesAttachment, path=/{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteEmployeeStandardNotesAttachment(ctx context.Context, employeeCode string, noteID string, attachmentID string) error {
	url := fmt.Sprintf("/employees/%s/notes/standard/%s/attachments/%s", employeeCode, noteID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllEmployeeConfidentialNotes, path=/{cmpy_code}/employees/{emp_code}/notes/confidential
func (c *Client) GetAllEmployeeConfidentialNotes(ctx context.Context, employeeCode string) ([]EmployeeConfidentialNoteResponse, error) {
	var result []EmployeeConfidentialNoteResponse
	url := fmt.Sprintf("/employees/%s/notes/confidential", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeeConfidentialNote, path=/{cmpy_code}/employees/{emp_code}/notes/confidential
func (c *Client) AddEmployeeConfidentialNote(ctx context.Context, employeeCode string, payload AddEmployeeConfidentialNoteRequest) (EmployeeConfidentialNoteResponse, error) {
	var result EmployeeConfidentialNoteResponse
	url := fmt.Sprintf("/employees/%s/notes/confidential", employeeCode)
	if err := tasscommon.Validate(payload); err != nil {
		return EmployeeConfidentialNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return EmployeeConfidentialNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeConfidentialNoteResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetEmployeeConfidentialNoteByID, path=/{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}
func (c *Client) GetEmployeeConfidentialNote(ctx context.Context, employeeCode string, noteID string) (EmployeeConfidentialNoteResponse, error) {
	var result EmployeeConfidentialNoteResponse
	url := fmt.Sprintf("/employees/%s/notes/confidential/%s", employeeCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return EmployeeConfidentialNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeConfidentialNoteResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateEmployeeConfidentialNote, path=/{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}
func (c *Client) UpdateEmployeeConfidentialNote(ctx context.Context, employeeCode string, noteID string, payload UpdateEmployeeConfidentialNoteRequest) error {
	url := fmt.Sprintf("/employees/%s/notes/confidential/%s", employeeCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchEmployeeConfidentialNote, path=/{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}
func (c *Client) PatchEmployeeConfidentialNote(ctx context.Context, employeeCode string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/employees/%s/notes/confidential/%s", employeeCode, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteEmployeeConfidentialNote, path=/{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}
func (c *Client) DeleteEmployeeConfidentialNote(ctx context.Context, employeeCode string, noteID string) error {
	url := fmt.Sprintf("/employees/%s/notes/confidential/%s", employeeCode, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllEmployeeConfidentialNotesAttachments, path=/{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}/attachments
func (c *Client) GetAllEmployeeConfidentialNotesAttachments(ctx context.Context, employeeCode string, noteID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/employees/%s/notes/confidential/%s/attachments", employeeCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeeConfidentialNotesAttachments, path=/{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}/attachments
func (c *Client) AddEmployeeConfidentialNotesAttachments(ctx context.Context, employeeCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/employees/%s/notes/confidential/%s/attachments", employeeCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	body, err := c.t.upload(ctx, url, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=DownloadEmployeeConfidentialNotesAttachment, path=/{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) GetEmployeeConfidentialNotesAttachment(ctx context.Context, employeeCode string, noteID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/employees/%s/notes/confidential/%s/attachments/%s", employeeCode, noteID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DeleteEmployeeConfidentialNotesAttachment, path=/{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteEmployeeConfidentialNotesAttachment(ctx context.Context, employeeCode string, noteID string, attachmentID string) error {
	url := fmt.Sprintf("/employees/%s/notes/confidential/%s/attachments/%s", employeeCode, noteID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetEmployeePhotoChanges, path=/{cmpy_code}/employees/photo/changes
func (c *Client) GetEmployeePhotoChanges(ctx context.Context, changeKey string) (EmployeePhotoChangesResponse, error) {
	var result EmployeePhotoChangesResponse
	query, err := url.ParseQuery(fmt.Sprintf("change_key=%s", changeKey))
	if err != nil {
		return EmployeePhotoChangesResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodGet, "/employees/photo/changes", query, nil, http.StatusOK)
	if err != nil {
		return EmployeePhotoChangesResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeePhotoChangesResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetEmployeePhoto, path=/{cmpy_code}/employees/{emp_code}/photo
func (c *Client) GetEmployeePhoto(ctx context.Context, employeeCode string) ([]byte, error) {
	// TODO: Check this works?
	var result []byte
	url := fmt.Sprintf("/employees/%s/photo", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeePhoto, path=/{cmpy_code}/employees/{emp_code}/photo
func (c *Client) AddEmployeePhoto(ctx context.Context, employeeCode string, payload tasscommon.FileRequest) error {
	url := fmt.Sprintf("/employees/%s/photo", employeeCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.upload(ctx, url, payload, http.StatusCreated)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllEmployeeQualifications, path=/{cmpy_code}/employees/{emp_code}/qualifications
func (c *Client) GetAllEmployeeQualifications(ctx context.Context, employeeCode string) ([]EmployeeQualificationResponse, error) {
	var result []EmployeeQualificationResponse
	url := fmt.Sprintf("/employees/%s/qualifications", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeeQualification, path=/{cmpy_code}/employees/{emp_code}/qualifications
func (c *Client) AddEmployeeQualification(ctx context.Context, employeeCode string, payload AddEmployeeQualificationRequest) (EmployeeQualificationResponse, error) {
	var result EmployeeQualificationResponse
	url := fmt.Sprintf("/employees/%s/qualifications", employeeCode)
	if err := tasscommon.Validate(payload); err != nil {
		return EmployeeQualificationResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return EmployeeQualificationResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeQualificationResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetEmployeeQualificationByID, path=/{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}
func (c *Client) GetEmployeeQualification(ctx context.Context, employeeCode string, qualificationID string) (EmployeeQualificationResponse, error) {
	var result EmployeeQualificationResponse
	url := fmt.Sprintf("/employees/%s/qualifications/%s", employeeCode, qualificationID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return EmployeeQualificationResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeQualificationResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateEmployeeQualification, path=/{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}
func (c *Client) UpdateEmployeeQualification(ctx context.Context, employeeCode string, qualificationID string, payload UpdateEmployeeQualificationRequest) error {
	url := fmt.Sprintf("/employees/%s/qualifications/%s", employeeCode, qualificationID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchEmployeeQualification, path=/{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}
func (c *Client) PatchEmployeeQualification(ctx context.Context, employeeCode string, qualificationID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/employees/%s/qualifications/%s", employeeCode, qualificationID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteEmployeeQualification, path=/{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}
func (c *Client) DeleteEmployeeQualification(ctx context.Context, employeeCode string, qualificationID string) error {
	url := fmt.Sprintf("/employees/%s/qualifications/%s", employeeCode, qualificationID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllEmployeeQualificationsAttachments, path=/{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}/attachments
func (c *Client) GetAllEmployeeQualificationsAttachments(ctx context.Context, employeeCode string, qualificationID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/employees/%s/qualifications/%s/attachments", employeeCode, qualificationID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeeQualificationsAttachment, path=/{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}/attachments
func (c *Client) AddEmployeeQualificationsAttachment(ctx context.Context, employeeCode string, qualificationID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/employees/%s/qualifications/%s/attachments", employeeCode, qualificationID)
	if err := tasscommon.Validate(payload); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	body, err := c.t.upload(ctx, url, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=DownloadEmployeeQualificationsAttachment, path=/{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}/attachments/{attach_id}
func (c *Client) GetEmployeeQualificationsAttachment(ctx context.Context, employeeCode string, qualificationID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/employees/%s/qualifications/%s/attachments/%s", employeeCode, qualificationID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DeleteEmployeeQualificationsAttachment, path=/{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}/attachments/{attach_id}
func (c *Client) DeleteEmployeeQualificationsAttachment(ctx context.Context, employeeCode string, qualificationID string, attachmentID string) error {
	url := fmt.Sprintf("/employees/%s/qualifications/%s/attachments/%s", employeeCode, qualificationID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllEmployeeUDAreas, path=/{cmpy_code}/employees/{emp_code}/udareas
func (c *Client) GetAllEmployeeUDAreas(ctx context.Context, employeeCode string) ([]EmployeeUDAreaResponse, error) {
	var result []EmployeeUDAreaResponse
	url := fmt.Sprintf("/employees/%s/udareas", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetEmployeeUDArea, path=/{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) GetEmployeeUDArea(ctx context.Context, employeeCode string, areaCode string) (EmployeeUDAreaResponse, error) {
	var result EmployeeUDAreaResponse
	url := fmt.Sprintf("/employees/%s/udareas/%s", employeeCode, areaCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return EmployeeUDAreaResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeUDAreaResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=AddEmployeeUDArea, path=/{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) AddEmployeeUDArea(ctx context.Context, employeeCode string, areaCode string, payload AddEmployeeUDAreaRequest) (EmployeeUDAreaResponse, error) {
	var result EmployeeUDAreaResponse
	url := fmt.Sprintf("/employees/%s/udareas/%s", employeeCode, areaCode)
	if err := tasscommon.Validate(payload); err != nil {
		return EmployeeUDAreaResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return EmployeeUDAreaResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return EmployeeUDAreaResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateEmployeeUDArea, path=/{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) UpdateEmployeeUDArea(ctx context.Context, employeeCode string, areaCode string, payload UpdateEmployeeUDAreaRequest) error {
	url := fmt.Sprintf("/employees/%s/udareas/%s", employeeCode, areaCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchEmployeeUDArea, path=/{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) PatchEmployeeUDArea(ctx context.Context, employeeCode string, areaCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/employees/%s/udareas/%s", employeeCode, areaCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteEmployeeUDArea, path=/{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) DeleteEmployeeUDArea(ctx context.Context, employeeCode string, areaCode string) error {
	url := fmt.Sprintf("/employees/%s/udareas/%s", employeeCode, areaCode)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DownloadEmployeeUDAreaAttachment, path=/{cmpy_code}/employees/{emp_code}/udareas/{area_code}/attachments/{field_number}/{attach_id}
func (c *Client) GetEmployeeUDAreaAttachment(ctx context.Context, employeeCode string, areaCode string, udFieldID int, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/employees/%s/udareas/%s/attachments/%d/%s", employeeCode, areaCode, udFieldID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DownloadEmployeeUDAreaAttachment, path=/{cmpy_code}/employees/{emp_code}/udareas/{area_code}/attachments/{field_number}/{attach_id}
func (c *Client) DeleteEmployeeUDAreaAttachment(ctx context.Context, employeeCode string, areaCode string, udFieldID int, attachmentID string) error {
	url := fmt.Sprintf("/employees/%s/udareas/%s/attachments/%d/%s", employeeCode, areaCode, udFieldID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=AddEmployeeUDAreaAttachment, path=/{cmpy_code}/employees/{emp_code}/udareas/{area_code}/attachments/{field_number}
func (c *Client) AddEmployeeUDAreaAttachment(ctx context.Context, employeeCode string, areaCode string, udFieldID int, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/employees/%s/udareas/%s/attachments/%d", employeeCode, areaCode, udFieldID)
	if err := tasscommon.Validate(payload); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	body, err := c.t.upload(ctx, url, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}
