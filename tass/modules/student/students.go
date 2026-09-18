package tassstudent

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetAllStudents, path: /{cmpy_code}/students
func (c *Client) GetAllStudents() ([]StudentResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetStudentByID, path: /{cmpy_code}/students/{stud_code}
func (c *Client) GetStudentByID(studentCode string) (StudentResponse, error) {
	// TODO: Implementation
	return StudentResponse{}, nil
}

// op: UpdateStudentByID, path: /{cmpy_code}/students/{stud_code}
func (c *Client) UpdateStudentByID(studentCode string, payload UpdateStudentRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentByID, path: /{cmpy_code}/students/{stud_code}
func (c *Client) PatchStudentByID(studentCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentStandardNotes, path: /{cmpy_code}/students/{stud_code}/notes/standard
func (c *Client) GetAllStudentStandardNotes(studentCode string) ([]StudentStandardNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentStandardNote, path: /{cmpy_code}/students/{stud_code}/notes/standard
func (c *Client) AddStudentStandardNote(studentCode string, payload AddStudentStandardNoteRequest) (StudentStandardNoteResponse, error) {
	// TODO: Implementation
	return StudentStandardNoteResponse{}, nil
}

// op: GetStudentStandardNoteByID, path: /{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}
func (c *Client) GetStudentStandardNoteByID(studentCode string, noteID string) (StudentStandardNoteResponse, error) {
	// TODO: Implementation
	return StudentStandardNoteResponse{}, nil
}

// op: UpdateStudentStandardNote, path: /{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}
func (c *Client) UpdateStudentStandardNote(studentCode string, noteID string, payload UpdateStudentStandardNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentStandardNote, path: /{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}
func (c *Client) PatchStudentStandardNote(studentCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentStandardNote, path: /{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}
func (c *Client) DeleteStudentStandardNote(studentCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentStandardNoteAttachments, path: /{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}/attachments
func (c *Client) GetAllStudentStandardNoteAttachments(studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentStandardNoteAttachment, path: /{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}/attachments
func (c *Client) AddStudentStandardNoteAttachment(studentCode string, noteID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: GetStudentStandardNoteAttachment, path: /{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) GetStudentStandardNoteAttachment(studentCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentStandardNoteAttachment, path: /{cmpy_code}/students/{stud_code}/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentStandardNoteAttachment(studentCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentConfidentialNotes, path: /{cmpy_code}/students/{stud_code}/notes/confidential
func (c *Client) GetAllStudentConfidentialNotes(studentCode string) ([]StudentConfidentialNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentConfidentialNote, path: /{cmpy_code}/students/{stud_code}/notes/confidential
func (c *Client) AddStudentConfidentialNote(studentCode string, payload AddStudentConfidentialNoteRequest) (StudentConfidentialNoteResponse, error) {
	// TODO: Implementation
	return StudentConfidentialNoteResponse{}, nil
}

// op: GetStudentConfidentialNoteByID, path: /{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}
func (c *Client) GetStudentConfidentialNoteByID(studentCode string, noteID string) (StudentConfidentialNoteResponse, error) {
	// TODO: Implementation
	return StudentConfidentialNoteResponse{}, nil
}

// op: UpdateStudentConfidentialNote, path: /{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}
func (c *Client) UpdateStudentConfidentialNote(studentCode string, noteID string, payload UpdateStudentConfidentialNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentConfidentialNote, path: /{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}
func (c *Client) PatchStudentConfidentialNote(studentCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentConfidentialNote, path: /{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}
func (c *Client) DeleteStudentConfidentialNote(studentCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentConfidentialNoteAttachments, path: /{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}/attachments
func (c *Client) GetAllStudentConfidentialNoteAttachments(studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentConfidentialNoteAttachment, path: /{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}/attachments
func (c *Client) AddStudentConfidentialNoteAttachment(studentCode string, noteID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: GetStudentConfidentialNoteAttachment, path: /{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) GetStudentConfidentialNoteAttachment(studentCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentConfidentialNoteAttachment, path: /{cmpy_code}/students/{stud_code}/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentConfidentialNoteAttachment(studentCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetStudentPhotoChanges, path: /{cmpy_code}/students/photos/changes
func (c *Client) GetStudentPhotoChanges(changeKey string) (StudentPhotoChangesResponse, error) {
	// TODO: Implementation
	return StudentPhotoChangesResponse{}, nil
}

// op: GetStudentPhoto, path: /{cmpy_code}/students/{stud_code}/photo
func (c *Client) GetStudentPhoto(studentCode string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentPhoto, path: /{cmpy_code}/students/{stud_code}/photo
func (c *Client) AddStudentPhoto(studentCode string, file any) error {
	// TODO: How to supply photo
	// TODO: Implementation
	return nil
}

// op: GetAllStudentUDAreas, path: /{cmpy_code}/students/{stud_code}/udareas
func (c *Client) GetAllStudentUDAreas(studentCode string) ([]StudentUDAreaResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetStudentUDAreaByID, path: /{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) GetStudentUDAreaByID(studentCode string, areaCode string) (StudentUDAreaResponse, error) {
	// TODO: Implementation
	return StudentUDAreaResponse{}, nil
}

// op: AddStudentUDArea, path: /{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) AddStudentUDArea(studentCode string, areaCode string, payload AddStudentUDAreaRequest) (StudentUDAreaResponse, error) {
	// TODO: Implementation
	return StudentUDAreaResponse{}, nil
}

// op: UpdateStudentUDArea, path: /{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) UpdateStudentUDArea(studentCode string, areaCode string, payload UpdateStudentUDAreaRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentUDArea, path: /{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) PatchStudentUDArea(studentCode string, areaCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentUDArea, path: /{cmpy_code}/students/{stud_code}/udareas/{area_code}
func (c *Client) DeleteStudentUDArea(studentCode string, areaCode string) error {
	// TODO: Implementation
	return nil
}

// op: GetStudentUDAreaAttachment, path: /{cmpy_code}/students/{stud_code}/udareas/{area_code}/attachments/{field_number}/{attach_id}
func (c *Client) GetStudentUDAreaAttachment(studentCode string, areaCode string, fieldID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentUDAreaAttachment, path: /{cmpy_code}/students/{stud_code}/udareas/{area_code}/attachments/{field_number}/{attach_id}
func (c *Client) DeleteStudentUDAreaAttachment(studentCode string, areaCode string, fieldID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: AddStudentUDAreaAttachment, path: /{cmpy_code}/students/{stud_code}/udareas/{area_code}/attachments/{field_number}
func (c *Client) AddStudentUDAreaAttachment(studentCode string, areaCode string, fieldID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: GetStudentUDFields, path: /{cmpy_code}/students/{stud_code}/udfields
func (c *Client) GetStudentUDFields(studentCode string) (StudentUDFieldsResponse, error) {
	// TODO: Implementation
	return StudentUDFieldsResponse{}, nil
}

// op: UpdateStudentUDFields, path: /{cmpy_code}/students/{stud_code}/udfields
func (c *Client) UpdateStudentUDFields(studentCode string, payload UpdateStudentUDFieldsRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentUDFields, path: /{cmpy_code}/students/{stud_code}/udfields
func (c *Client) PatchStudentUDFields(studentCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentUDFieldOptions, path: /{cmpy_code}/options/students/udfields
func (c *Client) GetAllStudentUDFieldOptions() (StudentUDFieldOptionResponse, error) {
	// TODO: Implementation
	return StudentUDFieldOptionResponse{}, nil
}

// MCEECDYA

// op: GetStudentMCEECDYA, path: /{cmpy_code}/students/{stud_code}/mceecdya
func (c *Client) GetStudentMCEECDYA(studentCode string) (StudentMCEECDYAResponse, error) {
	// TODO: Implementation
	return StudentMCEECDYAResponse{}, nil
}

// op: UpdateStudentMCEECDYA, path: /{cmpy_code}/students/{stud_code}/mceecdya
func (c *Client) UpdateStudentMCEECDYA(studentCode string, request UpdateStudentMCEECDYARequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMCEECDYA, path: /{cmpy_code}/students/{stud_code}/mceecdya
func (c *Client) PatchStudentMCEECDYA(studentCode string, request tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}
