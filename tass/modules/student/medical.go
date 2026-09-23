package tassstudent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// op: GetAsthmaManagement, path: /{cmpy_code}/students/{stud_code}/medical/asthmamanagement
func (c *Client) GetAsthmaManagement(ctx context.Context, studentCode string) (AsthmaManagementResponse, error) {
	var result AsthmaManagementResponse
	url := fmt.Sprintf("/students/%s/medical/asthmamanagement", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return AsthmaManagementResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return AsthmaManagementResponse{}, err
	}
	return result, nil
}

// op: UpdateAsthmaManagement, path: /{cmpy_code}/students/{stud_code}/medical/asthmamanagement
func (c *Client) UpdateAsthmaManagement(ctx context.Context, studentCode string, payload UpdateAsthmaManagementRequest) error {
	url := fmt.Sprintf("/students/%s/medical/asthmamanagement", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchAsthmaManagement, path: /{cmpy_code}/students/{stud_code}/medical/asthmamanagement
func (c *Client) PatchAsthmaManagement(ctx context.Context, studentCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/asthmamanagement", studentCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicalConditions, path: /{cmpy_code}/students/{stud_code}/medical/conditions
func (c *Client) GetAllStudentMedicalConditions(ctx context.Context, studentCode string) ([]StudentMedicalConditionResponse, error) {
	var result []StudentMedicalConditionResponse
	url := fmt.Sprintf("/students/%s/medical/conditions", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicalCondition, path: /{cmpy_code}/students/{stud_code}/medical/conditions
func (c *Client) AddStudentMedicalCondition(ctx context.Context, studentCode string, payload AddStudentMedicalConditionRequest) (StudentMedicalConditionResponse, error) {
	var result StudentMedicalConditionResponse
	url := fmt.Sprintf("/students/%s/medical/conditions", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentMedicalConditionResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentMedicalConditionResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalConditionResponse{}, err
	}
	return result, nil
}

// op: GetStudentMedicalCondition, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}
func (c *Client) GetStudentMedicalCondition(ctx context.Context, studentCode string, medicalConditionCode string) (StudentMedicalConditionResponse, error) {
	var result StudentMedicalConditionResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s", studentCode, medicalConditionCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMedicalConditionResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalConditionResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentMedicalCondition, path:/{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}
func (c *Client) UpdateStudentMedicalCondition(ctx context.Context, studentCode string, medicalConditionCode string, payload UpdateStudentMedicalConditionRequest) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s", studentCode, medicalConditionCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentMedicalCondition, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}
func (c *Client) PatchStudentMedicalCondition(ctx context.Context, studentCode string, medicalConditionCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s", studentCode, medicalConditionCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteStudentMedicalCondition, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}
func (c *Client) DeleteStudentMedicalCondition(ctx context.Context, studentCode string, medicalConditionCode string) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s", studentCode, medicalConditionCode)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicalConditionNotes, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes
func (c *Client) GetAllStudentMedicalConditionNotes(ctx context.Context, studentCode string, medicalConditionCode string) ([]StudentMedicalConditionNoteResponse, error) {
	var result []StudentMedicalConditionNoteResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/notes", studentCode, medicalConditionCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicalConditionNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes
func (c *Client) AddStudentMedicalConditionNote(ctx context.Context, studentCode string, medicalConditionCode string, payload AddStudentMedicalConditionNoteRequest) (StudentMedicalConditionNoteResponse, error) {
	var result StudentMedicalConditionNoteResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/notes", studentCode, medicalConditionCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentMedicalConditionNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentMedicalConditionNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalConditionNoteResponse{}, err
	}
	return result, nil
}

// op: GetStudentMedicalConditionNoteByCode, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes/{note_uid}
func (c *Client) GetStudentMedicalConditionNote(ctx context.Context, studentCode, medicalConditionCode string, noteID string) (StudentMedicalConditionNoteResponse, error) {
	var result StudentMedicalConditionNoteResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/notes/%s", studentCode, medicalConditionCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMedicalConditionNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalConditionNoteResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentMedicalConditionNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes/{note_uid}
func (c *Client) UpdateStudentMedicalConditionNote(ctx context.Context, studentCode string, medicalConditionCode string, noteID string, payload UpdateStudentMedicalConditionNoteRequest) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/notes/%s", studentCode, medicalConditionCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentMedicalConditionNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes/{note_uid}
func (c *Client) PatchStudentMedicalConditionNote(ctx context.Context, studentCode string, medicalConditionCode string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/notes/%s", studentCode, medicalConditionCode, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteStudentMedicalConditionNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes/{note_uid}
func (c *Client) DeleteStudentMedicalConditionNote(ctx context.Context, studentCode string, medicalConditionCode string, noteID string) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/notes/%s", studentCode, medicalConditionCode, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicalConditionAttachments, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/attachments
func (c *Client) GetAllStudentMedicalConditionAttachments(ctx context.Context, studentCode string, medicalConditionCode string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/attachments", studentCode, medicalConditionCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicalConditionAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/attachments
func (c *Client) AddStudentMedicalConditionAttachment(ctx context.Context, studentCode string, medicalConditionCode string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/attachments", studentCode, medicalConditionCode)
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

// op: DownloadStudentMedicalConditionAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/attachments/{attach_id}
func (c *Client) GetStudentMedicalConditionAttachment(ctx context.Context, studentCode string, medicalConditionCode string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/attachments/%s", studentCode, medicalConditionCode, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteStudentMedicalConditionAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/attachments/{attach_id}
func (c *Client) DeleteStudentMedicalConditionAttachment(ctx context.Context, studentCode string, medicalConditionCode string, attachmentID string) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/attachments/%s", studentCode, medicalConditionCode, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentIllnesses, path: /{cmpy_code}/students/{stud_code}/medical/illnesses
func (c *Client) GetAllStudentIllnesses(ctx context.Context, studentCode string) ([]StudentIllnessResponse, error) {
	var result []StudentIllnessResponse
	url := fmt.Sprintf("/students/%s/medical/illnesses", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentIllness, path: /{cmpy_code}/students/{stud_code}/medical/illnesses
func (c *Client) AddStudentIllness(ctx context.Context, studentCode string, payload AddStudentIllnessRequest) (StudentIllnessResponse, error) {
	var result StudentIllnessResponse
	url := fmt.Sprintf("/students/%s/medical/illnesses", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentIllnessResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentIllnessResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentIllnessResponse{}, err
	}
	return result, nil
}

// op: GetStudentIllnessByID, path: /{cmpy_code}/students/{stud_code}/medical/illnesses/{illness_uid}
func (c *Client) GetStudentIllness(ctx context.Context, studentCode string, illnessID string) (StudentIllnessResponse, error) {
	var result StudentIllnessResponse
	url := fmt.Sprintf("/students/%s/medical/illnesses/%s", studentCode, illnessID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentIllnessResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentIllnessResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentIllness, path: /{cmpy_code}/students/{stud_code}/medical/illnesses/{illness_uid}
func (c *Client) UpdateStudentIllness(ctx context.Context, studentCode string, illnessID string, payload UpdateStudentIllnessRequest) error {
	url := fmt.Sprintf("/students/%s/medical/illnesses/%s", studentCode, illnessID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentIllness, path: /{cmpy_code}/students/{stud_code}/medical/illnesses/{illness_uid}
func (c *Client) PatchStudentIllness(ctx context.Context, studentCode string, illnessID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/illnesses/%s", studentCode, illnessID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteStudentIllness, path: /{cmpy_code}/students/{stud_code}/medical/illnesses/{illness_uid}
func (c *Client) DeleteStudentIllness(ctx context.Context, studentCode string, illnessID string) error {
	url := fmt.Sprintf("/students/%s/medical/illnesses/%s", studentCode, illnessID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentImmunisations, path: /{cmpy_code}/students/{stud_code}/medical/immunisations
func (c *Client) GetAllStudentImmunisations(ctx context.Context, studentCode string) ([]StudentImmunisationResponse, error) {
	var result []StudentImmunisationResponse
	url := fmt.Sprintf("/students/%s/medical/immunisations", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentImmunisation, path: /{cmpy_code}/students/{stud_code}/medical/immunisations
func (c *Client) AddStudentImmunisation(ctx context.Context, studentCode string, payload AddStudentImmunisationRequest) (StudentImmunisationResponse, error) {
	var result StudentImmunisationResponse
	url := fmt.Sprintf("/students/%s/medical/immunisations", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentImmunisationResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentImmunisationResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentImmunisationResponse{}, err
	}
	return result, nil
}

// op: GetStudentImmunisation, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/{imm_code}
func (c *Client) GetStudentImmunisation(ctx context.Context, studentCode string, immunisationCode string) ([]StudentImmunisationResponse, error) {
	var result []StudentImmunisationResponse
	url := fmt.Sprintf("/students/%s/medical/immunisations/%s", studentCode, immunisationCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteStudentImmunisation, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/{imm_code}
func (c *Client) DeleteStudentImmunisation(ctx context.Context, studentCode string, immunisationCode string) error {
	url := fmt.Sprintf("/students/%s/medical/immunisations/%s", studentCode, immunisationCode)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: UpdateStudentImmunisation, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/{imm_code}
func (c *Client) UpdateStudentImmunisation(ctx context.Context, studentCode string, immunisationCode string, payload UpdateStudentImmunisationRequest) error {
	url := fmt.Sprintf("/students/%s/medical/immunisations/%s", studentCode, immunisationCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetStudentImmunisationRegister, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register
func (c *Client) GetStudentImmunisationRegister(ctx context.Context, studentCode string) (StudentImmunisationRegisterResponse, error) {
	var result StudentImmunisationRegisterResponse
	url := fmt.Sprintf("/students/%s/medical/immunisations/register", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentImmunisationRegisterResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentImmunisationRegisterResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentImmunisationRegister, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register
func (c *Client) UpdateStudentImmunisationRegister(ctx context.Context, studentCode string, payload UpdateStudentImmunisationRegisterRequest) error {
	url := fmt.Sprintf("/students/%s/medical/immunisations/register", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentImmunisationRegisterAttachments, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register/attachments
func (c *Client) GetAllStudentImmunisationRegisterAttachments(ctx context.Context, studentCode string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/students/%s/medical/immunisations/register/attachments", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentImmunisationRegisterAttachment, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register/attachments
func (c *Client) AddStudentImmunisationRegisterAttachment(ctx context.Context, studentCode string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/students/%s/medical/immunisations/register/attachments", studentCode)
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

// op: DownloadStudentImmunisationRegisterAttachment, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register/attachments/{attach_id}
func (c *Client) GetStudentImmunisationRegisterAttachment(ctx context.Context, studentCode string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/medical/immmunisations/register/attachments/%s", studentCode, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteStudentImmunisationRegisterAttachment, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register/attachments/{attach_id}
func (c *Client) DeleteStudentImmunisationRegisterAttachment(ctx context.Context, studentCode string, attachmentID string) error {
	url := fmt.Sprintf("/students/%s/medical/immunisations/register/attachments/%s", studentCode, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicalMedications, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications
func (c *Client) GetAllStudentMedicalMedications(ctx context.Context, studentCode string, medicalConditionCode string) ([]StudentMedicationResponse, error) {
	var result []StudentMedicationResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications", studentCode, medicalConditionCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicalMedication, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications
func (c *Client) AddStudentMedicalMedication(ctx context.Context, studentCode string, medicalConditionCode string, payload AddStudentMedicationRequest) (StudentMedicationResponse, error) {
	var result StudentMedicationResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications", studentCode, medicalConditionCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentMedicationResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentMedicationResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicationResponse{}, err
	}
	return result, nil
}

// op: GetStudentMedicalMedicationByCode, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}
func (c *Client) GetStudentMedicalMedication(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string) (StudentMedicationResponse, error) {
	var result StudentMedicationResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s", studentCode, medicalConditionCode, medicationID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMedicationResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicationResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentMedicalMedication, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}
func (c *Client) UpdateStudentMedicalMedication(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, payload UpdateStudentMedicationRequest) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s", studentCode, medicalConditionCode, medicationID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentMedicalMedication, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}
func (c *Client) PatchStudentMedicalMedication(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s", studentCode, medicalConditionCode, medicationID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteStudentMedicalMedication, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}
func (c *Client) DeleteStudentMedicalMedication(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s", studentCode, medicalConditionCode, medicationID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicationNotes, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes
func (c *Client) GetAllStudentMedicationNotes(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string) ([]StudentMedicationNoteResponse, error) {
	var result []StudentMedicationNoteResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/notes", studentCode, medicalConditionCode, medicationID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicationNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes
func (c *Client) AddStudentMedicationNote(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, payload AddStudentMedicationNoteRequest) (StudentMedicationNoteResponse, error) {
	var result StudentMedicationNoteResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/notes", studentCode, medicalConditionCode, medicationID)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentMedicationNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentMedicationNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicationNoteResponse{}, err
	}
	return result, nil
}

// op: GetStudentMedicationNoteByCode, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes/{note_uid}
func (c *Client) GetStudentMedicationNote(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, noteID string) (StudentMedicationNoteResponse, error) {
	var result StudentMedicationNoteResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/notes/%s", studentCode, medicalConditionCode, medicationID, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMedicationNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicationNoteResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentMedicationNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes/{note_uid}
func (c *Client) UpdateStudentMedicationNote(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, noteID string, payload UpdateStudentMedicationNoteRequest) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/notes/%s", studentCode, medicalConditionCode, medicationID, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentMedicationNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes/{note_uid}
func (c *Client) PatchStudentMedicationNote(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/notes/%s", studentCode, medicalConditionCode, medicationID, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteStudentMedicationNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes/{note_uid}
func (c *Client) DeleteStudentMedicationNote(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, noteID string) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/notes/%s", studentCode, medicalConditionCode, medicationID, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicationAttachments, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/attachments
func (c *Client) GetAllStudentMedicationAttachments(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/attachments", studentCode, medicalConditionCode, medicationID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicationAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/attachments
func (c *Client) AddStudentMedicationAttachment(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/attachments", studentCode, medicalConditionCode, medicationID)
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

// op: DownloadStudentMedicationAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/attachments/{attach_id}
func (c *Client) GetStudentMedicationAttachment(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/attachments/%s", studentCode, medicalConditionCode, medicationID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteStudentMedicationAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentMedicationAttachment(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, attachmentID string) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/attachments/%s", studentCode, medicalConditionCode, medicationID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicationSchedules, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules
func (c *Client) GetAllStudentMedicationSchedules(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string) ([]StudentMedicationScheduleResponse, error) {
	var result []StudentMedicationScheduleResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/schedules", studentCode, medicalConditionCode, medicationID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicationSchedule, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules
func (c *Client) AddStudentMedicationSchedule(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, payload AddStudentMedicationScheduleRequest) (StudentMedicationScheduleResponse, error) {
	var result StudentMedicationScheduleResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/schedules", studentCode, medicalConditionCode, medicationID)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentMedicationScheduleResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentMedicationScheduleResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicationScheduleResponse{}, err
	}
	return result, nil
}

// op: GetStudentMedicationScheduleByCode, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules/{sched_uid}
func (c *Client) GetStudentMedicationSchedule(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, scheduleID string) (StudentMedicationScheduleResponse, error) {
	var result StudentMedicationScheduleResponse
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/schedules/%s", studentCode, medicalConditionCode, medicationID, scheduleID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMedicationScheduleResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicationScheduleResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentMedicationSchedule, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules/{sched_uid}
func (c *Client) UpdateStudentMedicationSchedule(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, scheduleID string, payload UpdateStudentMedicationScheduleRequest) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/schedules/%s", studentCode, medicalConditionCode, medicationID, scheduleID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentMedicationSchedule, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules/{sched_uid}
func (c *Client) PatchStudentMedicationSchedule(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, scheduleID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/schedules/%s", studentCode, medicalConditionCode, medicationID, scheduleID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteStudentMedicationSchedule, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules/{sched_uid}
func (c *Client) DeleteStudentMedicationSchedule(ctx context.Context, studentCode string, medicalConditionCode string, medicationID string, scheduleID string) error {
	url := fmt.Sprintf("/students/%s/medical/conditions/%s/medications/%s/schedules/%s", studentCode, medicalConditionCode, medicationID, scheduleID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentStandardMedicalNotes, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard
func (c *Client) GetAllStudentStandardMedicalNotes(ctx context.Context, studentCode string) ([]StudentMedicalStandardNoteResponse, error) {
	var result []StudentMedicalStandardNoteResponse
	url := fmt.Sprintf("/students/%s/medical/notes/standard", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicalStandardNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard
func (c *Client) AddStudentMedicalStandardNote(ctx context.Context, studentCode string, payload AddStudentMedicalStandardNoteRequest) (StudentMedicalStandardNoteResponse, error) {
	var result StudentMedicalStandardNoteResponse
	url := fmt.Sprintf("/students/%s/medical/notes/standard", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentMedicalStandardNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentMedicalStandardNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalStandardNoteResponse{}, err
	}
	return result, nil
}

// op: GetStudentMedicalStandardNoteByID, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}
func (c *Client) GetStudentMedicalStandardNote(ctx context.Context, studentCode string, noteID string) (StudentMedicalStandardNoteResponse, error) {
	var result StudentMedicalStandardNoteResponse
	url := fmt.Sprintf("/students/%s/medical/notes/standard/%s", studentCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMedicalStandardNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalStandardNoteResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentMedicalStandardNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}
func (c *Client) UpdateStudentMedicalStandardNote(ctx context.Context, studentCode string, noteID string, payload UpdateStudentMedicalStandardNoteRequest) error {
	url := fmt.Sprintf("/students/%s/medical/notes/standard/%s", studentCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentMedicalStandardNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}
func (c *Client) PatchStudentMedicalStandardNote(ctx context.Context, studentCode string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/notes/standard/%s", studentCode, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteStudentMedicalStandardNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}
func (c *Client) DeleteStudentMedicalStandardNote(ctx context.Context, studentCode string, noteID string) error {
	url := fmt.Sprintf("/students/%s/medical/notes/standard/%s", studentCode, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicalStandardNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}/attachments
func (c *Client) GetAllStudentMedicalStandardNotesAttachment(ctx context.Context, studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/students/%s/medical/notes/standard/%s/attachments", studentCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicalStandardNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}/attachments
func (c *Client) AddStudentMedicalStandardNotesAttachment(ctx context.Context, studentCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/students/%s/medical/notes/standard/%s/attachments", studentCode, noteID)
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

// op: DownloadStudentMedicalStandardNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) GetStudentMedicalStandardNotesAttachment(ctx context.Context, studentCode string, noteID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/medical/notes/standard/%s/attachments/%s", studentCode, noteID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteStudentMedicalStandardNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentMedicalStandardNotesAttachment(ctx context.Context, studentCode string, noteID string, attachmentID string) error {
	url := fmt.Sprintf("/students/%s/medical/notes/standard/%s/attachments/%s", studentCode, noteID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentConfidentialMedicalNotes, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential
func (c *Client) GetAllStudentConfidentialMedicalNotes(ctx context.Context, studentCode string) ([]StudentMedicalConfidentialNoteResponse, error) {
	var result []StudentMedicalConfidentialNoteResponse
	url := fmt.Sprintf("/students/%s/medical/notes/confidential", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicalConfidentialNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential
func (c *Client) AddStudentMedicalConfidentialNote(ctx context.Context, studentCode string, payload AddStudentMedicalConfidentialNoteRequest) (StudentMedicalConfidentialNoteResponse, error) {
	var result StudentMedicalConfidentialNoteResponse
	url := fmt.Sprintf("/students/%s/medical/notes/confidential", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentMedicalConfidentialNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentMedicalConfidentialNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalConfidentialNoteResponse{}, err
	}
	return result, nil
}

// op: GetStudentMedicalConfidentialNoteByID, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}
func (c *Client) GetStudentMedicalConfidentialNote(ctx context.Context, studentCode string, noteID string) (StudentMedicalConfidentialNoteResponse, error) {
	var result StudentMedicalConfidentialNoteResponse
	url := fmt.Sprintf("/students/%s/medical/notes/confidential/%s", studentCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMedicalConfidentialNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalConfidentialNoteResponse{}, err
	}
	return result, nil
}

// op: UpdateStudentMedicalConfidentialNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}
func (c *Client) UpdateStudentMedicalConfidentialNote(ctx context.Context, studentCode string, noteID string, payload UpdateStudentMedicalConfidentialNoteRequest) error {
	url := fmt.Sprintf("/students/%s/medical/notes/confidential/%s", studentCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentMedicalConfidentialNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}
func (c *Client) PatchStudentMedicalConfidentialNote(ctx context.Context, studentCode string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/notes/confidential/%s", studentCode, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteStudentMedicalConfidentialNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}
func (c *Client) DeleteStudentMedicalConfidentialNote(ctx context.Context, studentCode string, noteID string) error {
	url := fmt.Sprintf("/students/%s/medical/notes/confidential/%s", studentCode, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicalConfidentialNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}/attachments
func (c *Client) GetAllStudentMedicalConfidentialNotesAttachment(ctx context.Context, studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/students/%s/medical/notes/confidential/%s/attachments", studentCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentMedicalConfidentialNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}/attachments
func (c *Client) AddStudentMedicalConfidentialNotesAttachment(ctx context.Context, studentCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/students/%s/medical/notes/confidential/%s/attachments", studentCode, noteID)
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

// op: DownloadStudentMedicalConfidentialNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) GetStudentMedicalConfidentialNotesAttachment(ctx context.Context, studentCode string, noteID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/medical/notes/confidential/%s/attachments/%s", studentCode, noteID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteStudentMedicalConfidentialNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentMedicalConfidentialNotesAttachment(ctx context.Context, studentCode string, noteID string, attachmentID string) error {
	url := fmt.Sprintf("/students/%s/medical/notes/confidential/%s/attachments/%s", studentCode, noteID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentPractitioners, path: /{cmpy_code}/students/{stud_code}/medical/practitioners
func (c *Client) GetAllStudentPractitioners(ctx context.Context, studentCode string) ([]StudentPractitionerResponse, error) {
	var result []StudentPractitionerResponse
	url := fmt.Sprintf("/students/%s/medical/practitioners", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddStudentPractitioner, path: /{cmpy_code}/students/{stud_code}/medical/practitioners
func (c *Client) AddStudentPractitioner(ctx context.Context, studentCode string, payload AddStudentPractitionerRequest) (StudentPractitionerResponse, error) {
	var result StudentPractitionerResponse
	url := fmt.Sprintf("/students/%s/medical/practitioners", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentPractitionerResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentPractitionerResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentPractitionerResponse{}, err
	}
	return result, nil
}

// op: GetStudentPractitionerByPracNum, path: /{cmpy_code}/students/{stud_code}/medical/practitioners/{prac_num}
func (c *Client) GetStudentPractitioner(ctx context.Context, studentCode string, pracNum int) (StudentPractitionerResponse, error) {
	var result StudentPractitionerResponse
	url := fmt.Sprintf("/students/%s/medical/practitioners/%d", studentCode, pracNum)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentPractitionerResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentPractitionerResponse{}, err
	}
	return result, nil
}

// op: DeleteStudentPractitioner, path: /{cmpy_code}/students/{stud_code}/medical/practitioners/{prac_num}
func (c *Client) DeleteStudentPractitioner(ctx context.Context, studentCode string, pracNum int) error {
	url := fmt.Sprintf("/students/%s/medical/practitioners/%d", studentCode, pracNum)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: UpdateStudentPractitioner, path: /{cmpy_code}/students/{stud_code}/medical/practitioners/{prac_num}
func (c *Client) UpdateStudentPractitioner(ctx context.Context, studentCode string, pracNum int, payload UpdateStudentPractitionerRequest) error {
	url := fmt.Sprintf("/students/%s/medical/practitioners/%d", studentCode, pracNum)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentPractitioner, path: /{cmpy_code}/students/{stud_code}/medical/practitioners/{prac_num}
func (c *Client) PatchStudentPractitioner(ctx context.Context, studentCode string, pracNum int, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/practitioners/%d", studentCode, pracNum)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllStudentMedicalSupplementaries, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries
func (c *Client) GetAllStudentMedicalSupplementaries(ctx context.Context, studentCode string) ([]StudentMedicalSupplementaryResponse, error) {
	var result []StudentMedicalSupplementaryResponse
	url := fmt.Sprintf("/students/%s/medical/supplementaries", studentCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: CreateStudentMedicalSupplementary, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries
func (c *Client) CreateStudentMedicalSupplementary(ctx context.Context, studentCode string, payload AddStudentMedicalSupplementaryRequest) (StudentMedicalSupplementaryResponse, error) {
	var result StudentMedicalSupplementaryResponse
	url := fmt.Sprintf("/students/%s/medical/supplementaries", studentCode)
	if err := tasscommon.Validate(payload); err != nil {
		return StudentMedicalSupplementaryResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return StudentMedicalSupplementaryResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalSupplementaryResponse{}, err
	}
	return result, nil
}

// op: GetStudentMedicalSupplementaryByCode, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries/{msupp_code}
func (c *Client) GetStudentMedicalSupplementary(ctx context.Context, studentCode string, supplementaryCode string) (StudentMedicalSupplementaryResponse, error) {
	var result StudentMedicalSupplementaryResponse
	url := fmt.Sprintf("/students/%s/medical/supplementaries/%s", studentCode, supplementaryCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return StudentMedicalSupplementaryResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentMedicalSupplementaryResponse{}, err
	}
	return result, nil
}

// op: DeleteStudentMedicalSupplementary, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries/{msupp_code}
func (c *Client) DeleteStudentMedicalSupplementary(ctx context.Context, studentCode string, supplementaryCode string) error {
	url := fmt.Sprintf("/students/%s/medical/supplementaries/%s", studentCode, supplementaryCode)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: UpdateStudentMedicalSupplementary, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries/{msupp_code}
func (c *Client) UpdateStudentMedicalSupplementary(ctx context.Context, studentCode string, supplementaryCode string, payload UpdateStudentMedicalSupplementaryRequest) error {
	url := fmt.Sprintf("/students/%s/medical/supplementaries/%s", studentCode, supplementaryCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchStudentMedicalSupplementary, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries/{msupp_code}
func (c *Client) PatchStudentMedicalSupplementary(ctx context.Context, studentCode string, supplementaryCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/students/%s/medical/supplementaries/%s", studentCode, supplementaryCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}
