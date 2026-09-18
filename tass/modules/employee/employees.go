package tassemployee

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetEmployeeByCode, path: /{cmpy_code}/employees/{emp_code}
func (c *Client) GetEmployee(employeeCode string) (EmployeeResponse, error) {
	// TODO: Implementation
	return EmployeeResponse{}, nil
}

// op: UpdateEmployee, path: /{cmpy_code}/employees/{emp_code}
func (c *Client) UpdateEmployee(employeeCode string, payload UpdateEmployeeRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchEmployee, path: /{cmpy_code}/employees/{emp_code}
func (c *Client) PatchEmployee(employeeCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAllEmployees, path: /{cmpy_code}/employees
func (c *Client) GetAllEmployees() ([]EmployeeResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddEmployee, path: /{cmpy_code}/employees
func (c *Client) AddEmployee(payload AddEmployeeRequest) (EmployeeResponse, error) {
	// TODO: Implementation
	return EmployeeResponse{}, nil
}

// op: GetAllEmployeeStandardNotes, path: /{cmpy_code}/employees/{emp_code}/notes/standard
func (c *Client) GetAllEmployeeStandardNotes(employeeCode string) ([]EmployeeStandardNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddEmployeeStandardNote, path: /{cmpy_code}/employees/{emp_code}/notes/standard
func (c *Client) AddEmployeeStandardNote(employeeCode string, payload AddEmployeeStandardNoteRequest) (EmployeeStandardNoteResponse, error) {
	// TODO: Implementation
	return EmployeeStandardNoteResponse{}, nil
}

// op: GetEmployeeStandardNoteByID, path: /{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}
func (c *Client) GetEmployeeStandardNote(employeeCode string, noteID string) (EmployeeStandardNoteResponse, error) {
	// TODO: Implementation
	return EmployeeStandardNoteResponse{}, nil
}

// op: UpdateEmployeeStandardNote, path: /{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}
func (c *Client) UpdateEmployeeStandardNote(employeeCode string, noteID string, payload UpdateEmployeeStandardNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchEmployeeStandardNote, path: /{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}
func (c *Client) PatchEmployeeStandardNote(employeeCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteEmployeeStandardNote, path: /{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}
func (c *Client) DeleteEmployeeStandardNote(employeeCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllEmployeeStandardNotesAttachments, path: /{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}/attachments
func (c *Client) GetAllEmployeeStandardNotesAttachments(employeeCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddEmployeeStandardNotesAttachments, path: /{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}/attachments
func (c *Client) AddEmployeeStandardNotesAttachments(employeeCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadEmployeeStandardNotesAttachment, path: /{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) GetEmployeeStandardNotesAttachment(employeeCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteEmployeeStandardNotesAttachment, path: /{cmpy_code}/employees/{emp_code}/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteEmployeeStandardNotesAttachment(employeeCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllEmployeeConfidentialNotes, path: /{cmpy_code}/employees/{emp_code}/notes/confidential
func (c *Client) GetAllEmployeeConfidentialNotes(employeeCode string) ([]EmployeeConfidentialNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddEmployeeConfidentialNote, path: /{cmpy_code}/employees/{emp_code}/notes/confidential
func (c *Client) AddEmployeeConfidentialNote(employeeCode string, payload AddEmployeeConfidentialNoteRequest) (EmployeeConfidentialNoteResponse, error) {
	// TODO: Implementation
	return EmployeeConfidentialNoteResponse{}, nil
}

// op: GetEmployeeConfidentialNoteByID, path: /{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}
func (c *Client) GetEmployeeConfidentialNote(employeeCode string, noteID string) (EmployeeConfidentialNoteResponse, error) {
	// TODO: Implementation
	return EmployeeConfidentialNoteResponse{}, nil
}

// op: UpdateEmployeeConfidentialNote, path: /{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}
func (c *Client) UpdateEmployeeConfidentialNote(employeeCode string, noteID string, payload UpdateEmployeeConfidentialNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchEmployeeConfidentialNote, path: /{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}
func (c *Client) PatchEmployeeConfidentialNote(employeeCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteEmployeeConfidentialNote, path: /{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}
func (c *Client) DeleteEmployeeConfidentialNote(employeeCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllEmployeeConfidentialNotesAttachments, path: /{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}/attachments
func (c *Client) GetAllEmployeeConfidentialNotesAttachments(employeeCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddEmployeeConfidentialNotesAttachments, path: /{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}/attachments
func (c *Client) AddEmployeeConfidentialNotesAttachments(employeeCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadEmployeeConfidentialNotesAttachment, path: /{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) GetEmployeeConfidentialNotesAttachment(employeeCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteEmployeeConfidentialNotesAttachment, path: /{cmpy_code}/employees/{emp_code}/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteEmployeeConfidentialNotesAttachment(employeeCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetEmployeePhotoChanges, path: /{cmpy_code}/employees/photo/changes
func (c *Client) GetEmployeePhotoChanges(changeKey string) (EmployeePhotoChangesResponse, error) {
	// TODO: Implementation
	return EmployeePhotoChangesResponse{}, nil
}

// op: GetEmployeePhoto, path: /{cmpy_code}/employees/{emp_code}/photo
func (c *Client) GetEmployeePhoto(employeeCode string) (any, error) {
	// TODO: Implementation
	// What does this return?!
	return nil, nil
}

// op: AddEmployeePhoto, path: /{cmpy_code}/employees/{emp_code}/photo
func (c *Client) AddEmployeePhoto(employeeCode string, payload tasscommon.FileRequest) error {
	// TODO: Implementation
	return nil
}

// op: GetAllEmployeeQualifications, path: /{cmpy_code}/employees/{emp_code}/qualifications
func (c *Client) GetAllEmployeeQualifications(employeeCode string) ([]EmployeeQualificationResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddEmployeeQualification, path: /{cmpy_code}/employees/{emp_code}/qualifications
func (c *Client) AddEmployeeQualification(employeeCode string, payload AddEmployeeQualificationRequest) (EmployeeQualificationResponse, error) {
	// TODO: Implementation
	return EmployeeQualificationResponse{}, nil
}

// op: GetEmployeeQualificationByID, path: /{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}
func (c *Client) GetEmployeeQualification(employeeCode string, qualificationID string) (EmployeeQualificationResponse, error) {
	// TODO: Implementation
	return EmployeeQualificationResponse{}, nil
}

// op: UpdateEmployeeQualification, path: /{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}
func (c *Client) UpdateEmployeeQualification(employeeCode string, qualificationID string, payload UpdateEmployeeQualificationRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchEmployeeQualification, path: /{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}
func (c *Client) PatchEmployeeQualification(employeeCode string, qualificationID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteEmployeeQualification, path: /{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}
func (c *Client) DeleteEmployeeQualification(employeeCode string, qualificationID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllEmployeeQualificationsAttachments, path: /{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}/attachments
func (c *Client) GetAllEmployeeQualificationsAttachments(employeeCode string, qualificationID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddEmployeeQualificationsAttachment, path: /{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}/attachments
func (c *Client) AddEmployeeQualificationsAttachment(employeeCode string, qualificationID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadEmployeeQualificationsAttachment, path: /{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}/attachments/{attach_id}
func (c *Client) GetEmployeeQualificationsAttachment(employeeCode string, qualificationID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteEmployeeQualificationsAttachment, path: /{cmpy_code}/employees/{emp_code}/qualifications/{qual_uid}/attachments/{attach_id}
func (c *Client) DeleteEmployeeQualificationsAttachment(employeeCode string, qualificationID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllEmployeeUDAreas, path: /{cmpy_code}/employees/{emp_code}/udareas
func (c *Client) GetAllEmployeeUDAreas(employeeCode string) ([]EmployeeUDAreaResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetEmployeeUDArea, path: /{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) GetEmployeeUDArea(employeeCode string, areaCode string) (EmployeeUDAreaResponse, error) {
	// TODO: Implementation
	return EmployeeUDAreaResponse{}, nil
}

// op: AddEmployeeUDArea, path: /{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) AddEmployeeUDArea(employeeCode string, areaCode string, payload AddEmployeeUDAreaRequest) (EmployeeUDAreaResponse, error) {
	// TODO: Implementation
	return EmployeeUDAreaResponse{}, nil
}

// op: UpdateEmployeeUDArea, path: /{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) UpdateEmployeeUDArea(employeeCode string, areaCode string, payload UpdateEmployeeUDAreaRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchEmployeeUDArea, path: /{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) PatchEmployeeUDArea(employeeCode string, areaCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteEmployeeUDArea, path: /{cmpy_code}/employees/{emp_code}/udareas/{area_code}
func (c *Client) DeleteEmployeeUDArea(employeeCode string, areaCode string) error {
	// TODO: Implementation
	return nil
}

// op: DownloadEmployeeUDAreaAttachment, path: /{cmpy_code}/employees/{emp_code}/udareas/{area_code}/attachments/{field_number}/{attach_id}
func (c *Client) GetEmployeeUDAreaAttachment(employeeCode string, areaCode string, udFieldID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DownloadEmployeeUDAreaAttachment, path: /{cmpy_code}/employees/{emp_code}/udareas/{area_code}/attachments/{field_number}/{attach_id}
func (c *Client) DeleteEmployeeUDAreaAttachment(employeeCode string, areaCode string, udFieldID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: AddEmployeeUDAreaAttachment, path: /{cmpy_code}/employees/{emp_code}/udareas/{area_code}/attachments/{field_number}
func (c *Client) AddEmployeeUDAreaAttachment(employeeCode string, areaCode string, udFieldID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}
