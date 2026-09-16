package tassstudent

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

func (c *Client) GetAllStudents() ([]StudentResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentByID(studentCode string) (StudentResponse, error) {
	// TODO: Implementation
	return StudentResponse{}, nil
}

func (c *Client) UpdateStudentByID(studentCode string, payload UpdateStudentRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentByID(studentCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllReligionOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllResidencyStatusOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllCampusOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllFeederSchoolOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllHouseOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllYearGroupOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllNextYearIndicatorOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllComparativeReportingTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllPCTutorGroupOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllStudentStandardNotes(studentCode string) ([]StudentStandardNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentStandardNote(studentCode string, payload AddStudentStandardNoteRequest) (StudentStandardNoteResponse, error) {
	// TODO: Implementation
	return StudentStandardNoteResponse{}, nil
}

func (c *Client) GetStudentStandardNoteByID(studentCode string, noteID string) (StudentStandardNoteResponse, error) {
	// TODO: Implementation
	return StudentStandardNoteResponse{}, nil
}

func (c *Client) UpdateStudentStandardNote(studentCode string, noteID string, payload UpdateStudentStandardNoteRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentStandardNote(studentCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentStandardNote(studentCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentStandardNoteAttachments(studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentStandardNoteAttachment(studentCode string, noteID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

func (c *Client) GetStudentStandardNoteAttachment(studentCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentStandardNoteAttachment(studentCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentConfidentialNotes(studentCode string) ([]StudentConfidentialNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentConfidentialNote(studentCode string, payload AddStudentConfidentialNoteRequest) (StudentConfidentialNoteResponse, error) {
	// TODO: Implementation
	return StudentConfidentialNoteResponse{}, nil
}

func (c *Client) GetStudentConfidentialNoteByID(studentCode string, noteID string) (StudentConfidentialNoteResponse, error) {
	// TODO: Implementation
	return StudentConfidentialNoteResponse{}, nil
}

func (c *Client) UpdateStudentConfidentialNote(studentCode string, noteID string, payload UpdateStudentConfidentialNoteRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentConfidentialNote(studentCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentConfidentialNote(studentCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentConfidentialNoteAttachments(studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentConfidentialNoteAttachment(studentCode string, noteID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

func (c *Client) GetStudentConfidentialNoteAttachment(studentCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentConfidentialNoteAttachment(studentCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentNoteCategoryOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentPhotoChanges(changeKey string) (StudentPhotoChangesResponse, error) {
	// TODO: Implementation
	return StudentPhotoChangesResponse{}, nil
}

func (c *Client) GetStudentPhoto(studentCode string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentPhoto(studentCode string, file any) error {
	// TODO: How to supply photo
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentUDAreaOptions() ([]UDAreaOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetSingleStudentUDAreaOptions(areaCode string) (UDAreaOptionsResponse, error) {
	// TODO: Implementation
	return UDAreaOptionsResponse{}, nil
}

func (c *Client) GetAllStudentUDAreas(studentCode string) ([]StudentUDAreaResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentUDAreaByID(studentCode string, areaCode string) (StudentUDAreaResponse, error) {
	// TODO: Implementation
	return StudentUDAreaResponse{}, nil
}

func (c *Client) AddStudentUDArea(studentCode string, areaCode string, payload AddStudentUDAreaRequest) (StudentUDAreaResponse, error) {
	// TODO: Implementation
	return StudentUDAreaResponse{}, nil
}

func (c *Client) UpdateStudentUDArea(studentCode string, areaCode string, payload UpdateStudentUDAreaRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentUDArea(studentCode string, areaCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentUDArea(studentCode string, areaCode string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetStudentUDAreaAttachment(studentCode string, areaCode string, fieldID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentUDAreaAttachment(studentCode string, areaCode string, fieldID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) AddStudentUDAreaAttachment(studentCode string, areaCode string, fieldID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

func (c *Client) GetStudentUDFields(studentCode string) (StudentUDFieldsResponse, error) {
	// TODO: Implementation
	return StudentUDFieldsResponse{}, nil
}

func (c *Client) UpdateStudentUDFields(studentCode string, payload UpdateStudentUDFieldsRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentUDFields(studentCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentUDFieldOptions() (StudentUDFieldOptionResponse, error) {
	// TODO: Implementation
	return StudentUDFieldOptionResponse{}, nil
}

// MCEECDYA

func (c *Client) GetStudentMCEECDYA(studentCode string) (StudentMCEECDYAResponse, error) {
	// TODO: Implementation
	return StudentMCEECDYAResponse{}, nil
}

func (c *Client) UpdateStudentMCEECDYA(studentCode string, request UpdateStudentMCEECDYARequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMCEECDYA(studentCode string, request tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentCountryOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllStudentLanguageOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllOccupationalGroupOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllStudentIndigenousTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllSchoolEducationTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllNonSchoolEducationTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}
