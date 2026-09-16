package tassstudent

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

func (c *Client) GetAsthmaManagement(studentCode string) (AsthmaManagementResponse, error) {
	// TODO: Implementation
	return AsthmaManagementResponse{}, nil
}

func (c *Client) UpdateAsthmaManagement(studentCode string, payload UpdateAsthmaManagementRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchAsthmaManagement(studentCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicalConditions(studentCode string) ([]StudentMedicalConditionResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicalCondition(studentCode string, payload AddStudentMedicalConditionRequest) (StudentMedicalConditionResponse, error) {
	// TODO: Implementation
	return StudentMedicalConditionResponse{}, nil
}

func (c *Client) GetStudentMedicalCondition(studentCode string, medicalConditionCode string) (StudentMedicalConditionResponse, error) {
	// TODO: Implementation
	return StudentMedicalConditionResponse{}, nil
}

func (c *Client) UpdateStudentMedicalCondition(studentCode string, medicalConditionCode string, payload UpdateStudentMedicalConditionRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMedicalCondition(studentCode string, medicalConditionCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentMedicalCondition(studentCode string, medicalConditionCode string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicalConditionNotes(studentCode string, medicalConditionCode string) ([]StudentMedicalConditionNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicalConditionNote(studentCode string, medicalConditionCode string, payload AddStudentMedicalConditionNoteRequest) (StudentMedicalConditionNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalConditionNoteResponse{}, nil
}

func (c *Client) GetStudentMedicalConditionNoteByCode(studentCode, medicalConditionCode string, noteID string) (StudentMedicalConditionNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalConditionNoteResponse{}, nil
}

func (c *Client) UpdateStudentMedicalConditionNote(studentCode string, medicalConditionCode string, noteID string, payload UpdateStudentMedicalConditionNoteRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMedicalConditionNote(studentCode string, medicalConditionCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentMedicalConditionNote(studentCode string, medicalConditionCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicalConditionAttachments(studentCode string, medicalConditionCode string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicalConditionAttachment(studentCode string, medicalConditionCode string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

func (c *Client) GetStudentMedicalConditionAttachment(studentCode string, medicalConditionCode string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentMedicalConditionAttachment(studentCode string, medicalConditionCode string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllMedicalConditionTypesOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return []tasscommon.OptionsResponseActive{}, nil
}

func (c *Client) GetAllStudentIllnesses(studentCode string) ([]StudentIllnessResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentIllness(studentCode string, payload AddStudentIllnessRequest) (StudentIllnessResponse, error) {
	// TODO: Implementation
	return StudentIllnessResponse{}, nil
}

func (c *Client) GetStudentIllnessByID(studentCode string, illnessID string) (StudentIllnessResponse, error) {
	// TODO: Implementation
	return StudentIllnessResponse{}, nil
}

func (c *Client) UpdateStudentIllness(studentCode string, illnessID string, payload UpdateStudentIllnessRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentIllness(studentCode string, illnessID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentIllness(studentCode string, illnessID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllMedicalConditionsOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllMedicalTreatmentOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllImmunisationTypesOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllStudentImmunisations(studentCode string) ([]StudentImmunisationResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentImmunisation(studentCode string, payload AddStudentImmunisationRequest) (StudentImmunisationResponse, error) {
	// TODO: Implementation
	return StudentImmunisationResponse{}, nil
}

func (c *Client) GetStudentImmunisation(studentCode string, immunisationCode string) ([]StudentImmunisationResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentImmunisation(studentCode string, immunisationCode string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) UpdateStudentImmunisation(studentCode string, immunisationCode string, payload UpdateStudentImmunisationRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllImmunisationStatusesOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentImmunisationRegister(studentCode string) (StudentImmunisationRegisterResponse, error) {
	// TODO: Implementation
	return StudentImmunisationRegisterResponse{}, nil
}

func (c *Client) UpdateStudentImmunisationRegister(studentCode string, payload UpdateStudentImmunisationRegisterRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentImmunisationRegisterAttachments(studentCode string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentImmunisationRegisterAttachment(studentCode string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

func (c *Client) GetStudentImmunisationRegisterAttachment(studentCode string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentImmunisationRegisterAttachment(studentCode string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicalMedications(studentCode string, medicalConditionCode string) ([]StudentMedicationResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicalMedication(studentCode string, medicalConditionCode string, payload AddStudentMedicationRequest) (StudentMedicationResponse, error) {
	// TODO: Implementation
	return StudentMedicationResponse{}, nil
}

func (c *Client) GetStudentMedicalMedicationByCode(studentCode string, medicalConditionCode string, medicationID string) (StudentMedicationResponse, error) {
	// TODO: Implementation
	return StudentMedicationResponse{}, nil
}

func (c *Client) UpdateStudentMedicalMedication(studentCode string, medicalConditionCode string, medicationID string, payload UpdateStudentMedicationRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMedicalMedication(studentCode string, medicalConditionCode string, medicationID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentMedicalMedication(studentCode string, medicalConditionCode string, medicationID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicationNotes(studentCode string, medicalConditionCode string, medicationID string) ([]StudentMedicationNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicationNote(studentCode string, medicalConditionCode string, medicationID string, payload AddStudentMedicationNoteRequest) (StudentMedicationNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicationNoteResponse{}, nil
}

func (c *Client) GetStudentMedicationNoteByCode(studentCode string, medicalConditionCode string, medicationID string, noteID string) (StudentMedicationNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicationNoteResponse{}, nil
}

func (c *Client) UpdateStudentMedicationNote(studentCode string, medicalConditionCode string, medicationID string, noteID string, payload UpdateStudentMedicationNoteRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMedicationNote(studentCode string, medicalConditionCode string, medicationID string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentMedicationNote(studentCode string, medicalConditionCode string, medicationID string, noteID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicationAttachments(studentCode string, medicalConditionCode string, medicationID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicationAttachment(studentCode string, medicalConditionCode string, medicationID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

func (c *Client) GetStudentMedicationAttachment(studentCode string, medicalConditionCode string, medicationID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentMedicationAttachment(studentCode string, medicalConditionCode string, medicationID string, attachmedID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicationSchedules(studentCode string, medicalConditionCode string, medicationID string) ([]StudentMedicationScheduleResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicationSchedule(studentCode string, medicalConditionCode string, medicationID string, payload AddStudentMedicationScheduleRequest) (StudentMedicationScheduleResponse, error) {
	// TODO: Implementation
	return StudentMedicationScheduleResponse{}, nil
}

func (c *Client) GetStudentMedicationScheduleByCode(studentCode string, medicalConditionCode string, medicationID string, scheduleID string) (StudentMedicationScheduleResponse, error) {
	// TODO: Implementation
	return StudentMedicationScheduleResponse{}, nil
}

func (c *Client) UpdateStudentMedicationSchedule(studentCode string, medicalConditionCode string, medicationID string, scheduleID string, payload UpdateStudentMedicationScheduleRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMedicationSchedule(studentCode string, medicalConditionCode string, medicationID string, scheduleID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentMedicationSchedule(studentCode string, medicalConditionCode string, medicationID string, scheduleID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentStandardMedicalNotes(studentCode string) ([]StudentMedicalStandardNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicalStandardNote(studentCode string, payload AddStudentMedicalStandardNoteRequest) (StudentMedicalStandardNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalStandardNoteResponse{}, nil
}

func (c *Client) GetStudentMedicalStandardNoteByID(studentCode string, noteID string) (StudentMedicalStandardNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalStandardNoteResponse{}, nil
}

func (c *Client) UpdateStudentMedicalStandardNote(studentCode string, noteID string, payload UpdateStudentMedicalStandardNoteRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMedicalStandardNote(studentCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentMedicalStandardNote(studentCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicalStandardNotesAttachment(studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicalStandardNotesAttachment(studentCode string, noteID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

func (c *Client) GetStudentMedicalStandardNotesAttachment(studentCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentMedicalStandardNotesAttachment(studentCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentConfidentialMedicalNotes(studentCode string) ([]StudentMedicalConfidentialNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicalConfidentialNote(studentCode string, payload AddStudentMedicalConfidentialNoteRequest) (StudentMedicalConfidentialNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalConfidentialNoteResponse{}, nil
}

func (c *Client) GetStudentMedicalConfidentialNoteByID(studentCode string, noteID string) (StudentMedicalConfidentialNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalConfidentialNoteResponse{}, nil
}

func (c *Client) UpdateStudentMedicalConfidentialNote(studentCode string, noteID string, payload UpdateStudentMedicalConfidentialNoteRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMedicalConfidentialNote(studentCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) DeleteStudentMedicalConfidentialNote(studentCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicalConfidentialNotesAttachment(studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentMedicalConfidentialNotesAttachment(studentCode string, noteID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

func (c *Client) GetStudentMedicalConfidentialNotesAttachment(studentCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) DeleteStudentMedicalConfidentialNotesAttachment(studentCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicalNoteCategoryOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllPractitionerTypesOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllStudentPractitioners(studentCode string) ([]StudentPractitionerResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) AddStudentPractitioner(studentCode string, payload AddStudentPractitionerRequest) (StudentPractitionerResponse, error) {
	// TODO: Implementation
	return StudentPractitionerResponse{}, nil
}

func (c *Client) GetStudentPractitionerByPracNum(studentCode string, pracNum string) (StudentPractitionerResponse, error) {
	// TODO: Implementation
	return StudentPractitionerResponse{}, nil
}

func (c *Client) DeleteStudentPractitioner(studentCode string, pracNum string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) UpdateStudentPractitioner(studentCode string, pracNum string, payload UpdateStudentPractitionerRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentPractitioner(studentCode string, pracNum string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllStudentMedicalSupplementaries(studentCode string) ([]StudentMedicalSupplementaryResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) CreateStudentMedicalSupplementary(stuentCode string, payload AddStudentMedicalSupplementaryRequest) (StudentMedicalSupplementaryResponse, error) {
	// TODO: Implementation
	return StudentMedicalSupplementaryResponse{}, nil
}

func (c *Client) GetStudentMedicalSupplementaryByCode(studentCode string, supplementaryCode string) (StudentMedicalSupplementaryResponse, error) {
	// TODO: Implementation
	return StudentMedicalSupplementaryResponse{}, nil
}

func (c *Client) DeleteStudentMedicalSupplementary(studentCode string, supplementaryCode string) error {
	// TODO: Implementation
	return nil
}

func (c *Client) UpdateStudentMedicalSupplementary(studentCode string, supplementaryCode string, payload UpdateStudentMedicalSupplementaryRequest) error {
	// TODO: Implementation
	return nil
}

func (c *Client) PatchStudentMedicalSupplementary(studentCode string, supplementaryCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

func (c *Client) GetAllSupplementaryTypesOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}
