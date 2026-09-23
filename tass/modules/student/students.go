package tassstudent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetAllStudents, path=/{cmpy_code}/students
func (c *Client) GetAllStudents(ctx context.Context) ([]StudentResponse, error) {
	var result []StudentResponse
	body, err := c.t.request(ctx, http.MethodGet, "/students", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetStudentByID, path=/{cmpy_code}/students/{stud_code}
func (c *Client) GetStudentByID(ctx context.Context, studentCode string) (StudentResponse, error) {
	var result StudentResponse
	url := fmt.Sprintf("/students/%s", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateStudentByID, path=/{cmpy_code}/students/{stud_code}
func (c *Client) UpdateStudentByID(ctx context.Context, studentCode string, payload UpdateStudentRequest) error {
	url := fmt.Sprintf("/students/%s", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchStudentByID, path=/{cmpy_code}/students/{stud_code}
func (c *Client) PatchStudentByID(ctx context.Context, studentCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s", studentCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllStudentStandardNotes, path=/{cmpy_code}/students/{stud_code}/notes/standard
func (c *Client) GetAllStudentStandardNotes(ctx context.Context, studentCode string) ([]StudentStandardNoteResponse, error) {
	var result []StudentStandardNoteResponse
	url := fmt.Sprintf("/students/%s/notes/standard", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddStudentStandardNote, path=/{cmpy_code}/students/{stud_code}/notes/standard
func (c *Client) AddStudentStandardNote(ctx context.Context, studentCode string, payload AddStudentStandardNoteRequest) (StudentStandardNoteResponse, error) {
	var result StudentStandardNoteResponse
	url := fmt.Sprintf("/students/%s/notes/standard", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentStandardNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentStandardNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentStandardNoteResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetStudentStandardNoteByID, path=/{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}
func (c *Client) GetStudentStandardNoteByID(ctx context.Context, studentCode string, noteID string) (StudentStandardNoteResponse, error) {
	var result StudentStandardNoteResponse
	url := fmt.Sprintf("/students/%s/notes/%s", studentCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentStandardNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentStandardNoteResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateStudentStandardNote, path=/{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}
func (c *Client) UpdateStudentStandardNote(ctx context.Context, studentCode string, noteID string, payload UpdateStudentStandardNoteRequest) error {
	url := fmt.Sprintf("/students/%s/notes/standard/%s", studentCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchStudentStandardNote, path=/{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}
func (c *Client) PatchStudentStandardNote(ctx context.Context, studentCode string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/notes/standard/%s", studentCode, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteStudentStandardNote, path=/{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}
func (c *Client) DeleteStudentStandardNote(ctx context.Context, studentCode string, noteID string) error {
	url := fmt.Sprintf("/students/%s/notes/standard/%s", studentCode, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllStudentStandardNoteAttachments, path=/{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}/attachments
func (c *Client) GetAllStudentStandardNoteAttachments(ctx context.Context, studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/students/%s/notes/standard/%s/attachments", studentCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddStudentStandardNoteAttachment, path=/{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}/attachments
func (c *Client) AddStudentStandardNoteAttachment(ctx context.Context, studentCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/students/%s/notes/standard/%s/attachments", studentCode, noteID)
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

// CODEGEN(none): op=GetStudentStandardNoteAttachment, path=/{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) GetStudentStandardNoteAttachment(ctx context.Context, studentCode string, noteID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/notes/standard/%s/attachments/%s", studentCode, noteID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DeleteStudentStandardNoteAttachment, path=/{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentStandardNoteAttachment(ctx context.Context, studentCode string, noteID string, attachmentID string) error {
	url := fmt.Sprintf("/students/%s/notes/standard/%s/attachments/%s", studentCode, noteID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllStudentConfidentialNotes, path=/{cmpy_code}/students/{stud_code}/notes/confidential
func (c *Client) GetAllStudentConfidentialNotes(ctx context.Context, studentCode string) ([]StudentConfidentialNoteResponse, error) {
	var result []StudentConfidentialNoteResponse
	url := fmt.Sprintf("/students/%s/notes/confidential", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddStudentConfidentialNote, path=/{cmpy_code}/students/{stud_code}/notes/confidential
func (c *Client) AddStudentConfidentialNote(ctx context.Context, studentCode string, payload AddStudentConfidentialNoteRequest) (StudentConfidentialNoteResponse, error) {
	var result StudentConfidentialNoteResponse
	url := fmt.Sprintf("/students/%s/notes/confidential", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentConfidentialNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentConfidentialNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentConfidentialNoteResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetStudentConfidentialNoteByID, path=/{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}
func (c *Client) GetStudentConfidentialNoteByID(ctx context.Context, studentCode string, noteID string) (StudentConfidentialNoteResponse, error) {
	var result StudentConfidentialNoteResponse
	url := fmt.Sprintf("/students/%s/notes/confidential/%s", studentCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentConfidentialNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentConfidentialNoteResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateStudentConfidentialNote, path=/{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}
func (c *Client) UpdateStudentConfidentialNote(ctx context.Context, studentCode string, noteID string, payload UpdateStudentConfidentialNoteRequest) error {
	url := fmt.Sprintf("/students/%s/notes/confidential/%s", studentCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchStudentConfidentialNote, path=/{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}
func (c *Client) PatchStudentConfidentialNote(ctx context.Context, studentCode string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/notes/confidential/%s", studentCode, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteStudentConfidentialNote, path=/{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}
func (c *Client) DeleteStudentConfidentialNote(ctx context.Context, studentCode string, noteID string) error {
	url := fmt.Sprintf("/students/%s/notes/confidential/%s", studentCode, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllStudentConfidentialNoteAttachments, path=/{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}/attachments
func (c *Client) GetAllStudentConfidentialNoteAttachments(ctx context.Context, studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/students/%s/notes/confidential/%s/attachments", studentCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddStudentConfidentialNoteAttachment, path=/{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}/attachments
func (c *Client) AddStudentConfidentialNoteAttachment(ctx context.Context, studentCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/students/%s/notes/confidential/%s/attachments", studentCode, noteID)
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

// CODEGEN(none): op=GetStudentConfidentialNoteAttachment, path=/{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) GetStudentConfidentialNoteAttachment(ctx context.Context, studentCode string, noteID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/notes/confidential/%s/attachments/%s", studentCode, noteID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DeleteStudentConfidentialNoteAttachment, path=/{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentConfidentialNoteAttachment(ctx context.Context, studentCode string, noteID string, attachmentID string) error {
	url := fmt.Sprintf("/students/%s/notes/confidential/%s/attachments/%s", studentCode, noteID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetStudentPhotoChanges, path=/{cmpy_code}/students/photos/changes
func (c *Client) GetStudentPhotoChanges(ctx context.Context, changeKey string) (StudentPhotoChangesResponse, error) {
	var result StudentPhotoChangesResponse
	query, err := url.ParseQuery(fmt.Sprintf("change_key=%s", changeKey))
	if err != nil {
		return StudentPhotoChangesResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodGet, "/students/photos/changes", query, nil, http.StatusOK)
	if err != nil {
		return StudentPhotoChangesResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentPhotoChangesResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetStudentPhoto, path=/{cmpy_code}/students/{stud_code}/photo
func (c *Client) GetStudentPhoto(ctx context.Context, studentCode string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/photo", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddStudentPhoto, path=/{cmpy_code}/students/{stud_code}/photo
func (c *Client) AddStudentPhoto(ctx context.Context, studentCode string, payload tasscommon.FileRequest) error {
	url := fmt.Sprintf("/students/%s/photo", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.upload(ctx, url, payload, http.StatusCreated)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllStudentUDAreas, path=/{cmpy_code}/students/{stud_code}/udareas
func (c *Client) GetAllStudentUDAreas(ctx context.Context, studentCode string) ([]StudentUDAreaResponse, error) {
	var result []StudentUDAreaResponse
	url := fmt.Sprintf("/students/%s/udareas", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetStudentUDAreaByID, path=/{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) GetStudentUDAreaByID(ctx context.Context, studentCode string, areaCode string) (StudentUDAreaResponse, error) {
	var result StudentUDAreaResponse
	url := fmt.Sprintf("/students/%s/udareas/%s", studentCode, areaCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentUDAreaResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentUDAreaResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=AddStudentUDArea, path=/{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) AddStudentUDArea(ctx context.Context, studentCode string, areaCode string, payload AddStudentUDAreaRequest) (StudentUDAreaResponse, error) {
	var result StudentUDAreaResponse
	url := fmt.Sprintf("/students/%s/udareas/%s", studentCode, areaCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentUDAreaResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentUDAreaResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentUDAreaResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateStudentUDArea, path=/{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) UpdateStudentUDArea(ctx context.Context, studentCode string, areaCode string, payload UpdateStudentUDAreaRequest) error {
	url := fmt.Sprintf("/students/%s/udareas/%s", studentCode, areaCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchStudentUDArea, path=/{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) PatchStudentUDArea(ctx context.Context, studentCode string, areaCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/udareas/%s", studentCode, areaCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteStudentUDArea, path=/{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) DeleteStudentUDArea(ctx context.Context, studentCode string, areaCode string) error {
	url := fmt.Sprintf("/students/%s/udareas/%s", studentCode, areaCode)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DownloadStudentUDAreaAttachment, path=/{cmpy_code}/students/{stud_code}/udareas/{area_code}/attachments/{field_number}/{attach_id}
func (c *Client) GetStudentUDAreaAttachment(ctx context.Context, studentCode string, areaCode string, fieldID int, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/udareas/%s/attachments/%d/%s", studentCode, areaCode, fieldID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DeleteStudentUDAreaAttachment, path=/{cmpy_code}/students/{stud_code}/udareas/{area_code}/attachments/{field_number}/{attach_id}
func (c *Client) DeleteStudentUDAreaAttachment(ctx context.Context, studentCode string, areaCode string, fieldID int, attachmentID string) error {
	url := fmt.Sprintf("/students/%s/udareas/%s/attachments/%d/%s", studentCode, areaCode, fieldID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=AddStudentUDAreaAttachment, path=/{cmpy_code}/students/{stud_code}/udareas/{area_code}/attachments/{field_number}
func (c *Client) AddStudentUDAreaAttachment(ctx context.Context, studentCode string, areaCode string, fieldID int, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/students/%s/udareas/%s/attachments/%d", studentCode, areaCode, fieldID)
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

// CODEGEN(none): op=GetStudentUDFields, path=/{cmpy_code}/students/{stud_code}/udfields
func (c *Client) GetStudentUDFields(ctx context.Context, studentCode string) ([]StudentUDFieldsResponse, error) {
	var result []StudentUDFieldsResponse
	url := fmt.Sprintf("/students/%s/udfields", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateStudentUDFields, path=/{cmpy_code}/students/{stud_code}/udfields
func (c *Client) UpdateStudentUDFields(ctx context.Context, studentCode string, payload UpdateStudentUDFieldsRequest) error {
	url := fmt.Sprintf("/students/%s/udfields", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchStudentUDFields, path=/{cmpy_code}/students/{stud_code}/udfields
func (c *Client) PatchStudentUDFields(ctx context.Context, studentCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/udfields", studentCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetStudentMCEECDYA, path=/{cmpy_code}/students/{stud_code}/mceecdya
func (c *Client) GetStudentMCEECDYA(ctx context.Context, studentCode string) (StudentMCEECDYAResponse, error) {
	var result StudentMCEECDYAResponse
	url := fmt.Sprintf("/students/%s/mceecdya", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMCEECDYAResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMCEECDYAResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateStudentMCEECDYA, path=/{cmpy_code}/students/{stud_code}/mceecdya
func (c *Client) UpdateStudentMCEECDYA(ctx context.Context, studentCode string, payload UpdateStudentMCEECDYARequest) error {
	url := fmt.Sprintf("/students/%s/mceecdya", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchStudentMCEECDYA, path=/{cmpy_code}/students/{stud_code}/mceecdya
func (c *Client) PatchStudentMCEECDYA(ctx context.Context, studentCode string, payload tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/mceecdya", studentCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}
