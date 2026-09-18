package tassstudent

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetAsthmaManagement, path: /{cmpy_code}/students/{stud_code}/medical/asthmamanagement
func (c *Client) GetAsthmaManagement(studentCode string) (AsthmaManagementResponse, error) {
	// TODO: Implementation
	return AsthmaManagementResponse{}, nil
}

// op: UpdateAsthmaManagement, path: /{cmpy_code}/students/{stud_code}/medical/asthmamanagement
func (c *Client) UpdateAsthmaManagement(studentCode string, payload UpdateAsthmaManagementRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchAsthmaManagement, path: /{cmpy_code}/students/{stud_code}/medical/asthmamanagement
func (c *Client) PatchAsthmaManagement(studentCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicalConditions, path: /{cmpy_code}/students/{stud_code}/medical/conditions
func (c *Client) GetAllStudentMedicalConditions(studentCode string) ([]StudentMedicalConditionResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicalCondition, path: /{cmpy_code}/students/{stud_code}/medical/conditions
func (c *Client) AddStudentMedicalCondition(studentCode string, payload AddStudentMedicalConditionRequest) (StudentMedicalConditionResponse, error) {
	// TODO: Implementation
	return StudentMedicalConditionResponse{}, nil
}

// op: GetStudentMedicalCondition, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}
func (c *Client) GetStudentMedicalCondition(studentCode string, medicalConditionCode string) (StudentMedicalConditionResponse, error) {
	// TODO: Implementation
	return StudentMedicalConditionResponse{}, nil
}

// op: UpdateStudentMedicalCondition, path:/{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}
func (c *Client) UpdateStudentMedicalCondition(studentCode string, medicalConditionCode string, payload UpdateStudentMedicalConditionRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMedicalCondition, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}
func (c *Client) PatchStudentMedicalCondition(studentCode string, medicalConditionCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentMedicalCondition, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}
func (c *Client) DeleteStudentMedicalCondition(studentCode string, medicalConditionCode string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicalConditionNotes, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes
func (c *Client) GetAllStudentMedicalConditionNotes(studentCode string, medicalConditionCode string) ([]StudentMedicalConditionNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicalConditionNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes
func (c *Client) AddStudentMedicalConditionNote(studentCode string, medicalConditionCode string, payload AddStudentMedicalConditionNoteRequest) (StudentMedicalConditionNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalConditionNoteResponse{}, nil
}

// op: GetStudentMedicalConditionNoteByCode, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes/{note_uid}
func (c *Client) GetStudentMedicalConditionNoteByCode(studentCode, medicalConditionCode string, noteID string) (StudentMedicalConditionNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalConditionNoteResponse{}, nil
}

// op: UpdateStudentMedicalConditionNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes/{note_uid}
func (c *Client) UpdateStudentMedicalConditionNote(studentCode string, medicalConditionCode string, noteID string, payload UpdateStudentMedicalConditionNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMedicalConditionNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes/{note_uid}
func (c *Client) PatchStudentMedicalConditionNote(studentCode string, medicalConditionCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentMedicalConditionNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/notes/{note_uid}
func (c *Client) DeleteStudentMedicalConditionNote(studentCode string, medicalConditionCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicalConditionAttachments, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/attachments
func (c *Client) GetAllStudentMedicalConditionAttachments(studentCode string, medicalConditionCode string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicalConditionAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/attachments
func (c *Client) AddStudentMedicalConditionAttachment(studentCode string, medicalConditionCode string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadStudentMedicalConditionAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/attachments/{attach_id}
func (c *Client) GetStudentMedicalConditionAttachment(studentCode string, medicalConditionCode string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentMedicalConditionAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/attachments/{attach_id}
func (c *Client) DeleteStudentMedicalConditionAttachment(studentCode string, medicalConditionCode string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentIllnesses, path: /{cmpy_code}/students/{stud_code}/medical/illnesses
func (c *Client) GetAllStudentIllnesses(studentCode string) ([]StudentIllnessResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentIllness, path: /{cmpy_code}/students/{stud_code}/medical/illnesses
func (c *Client) AddStudentIllness(studentCode string, payload AddStudentIllnessRequest) (StudentIllnessResponse, error) {
	// TODO: Implementation
	return StudentIllnessResponse{}, nil
}

// op: GetStudentIllnessByID, path: /{cmpy_code}/students/{stud_code}/medical/illnesses/{illness_uid}
func (c *Client) GetStudentIllnessByID(studentCode string, illnessID string) (StudentIllnessResponse, error) {
	// TODO: Implementation
	return StudentIllnessResponse{}, nil
}

// op: UpdateStudentIllness, path: /{cmpy_code}/students/{stud_code}/medical/illnesses/{illness_uid}
func (c *Client) UpdateStudentIllness(studentCode string, illnessID string, payload UpdateStudentIllnessRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentIllness, path: /{cmpy_code}/students/{stud_code}/medical/illnesses/{illness_uid}
func (c *Client) PatchStudentIllness(studentCode string, illnessID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentIllness, path: /{cmpy_code}/students/{stud_code}/medical/illnesses/{illness_uid}
func (c *Client) DeleteStudentIllness(studentCode string, illnessID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentImmunisations, path: /{cmpy_code}/students/{stud_code}/medical/immunisations
func (c *Client) GetAllStudentImmunisations(studentCode string) ([]StudentImmunisationResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentImmunisation, path: /{cmpy_code}/students/{stud_code}/medical/immunisations
func (c *Client) AddStudentImmunisation(studentCode string, payload AddStudentImmunisationRequest) (StudentImmunisationResponse, error) {
	// TODO: Implementation
	return StudentImmunisationResponse{}, nil
}

// op: GetStudentImmunisation, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/{imm_code}
func (c *Client) GetStudentImmunisation(studentCode string, immunisationCode string) ([]StudentImmunisationResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentImmunisation, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/{imm_code}
func (c *Client) DeleteStudentImmunisation(studentCode string, immunisationCode string) error {
	// TODO: Implementation
	return nil
}

// op: UpdateStudentImmunisation, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/{imm_code}
func (c *Client) UpdateStudentImmunisation(studentCode string, immunisationCode string, payload UpdateStudentImmunisationRequest) error {
	// TODO: Implementation
	return nil
}

// op: GetStudentImmunisationRegister, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register
func (c *Client) GetStudentImmunisationRegister(studentCode string) (StudentImmunisationRegisterResponse, error) {
	// TODO: Implementation
	return StudentImmunisationRegisterResponse{}, nil
}

// op: UpdateStudentImmunisationRegister, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register
func (c *Client) UpdateStudentImmunisationRegister(studentCode string, payload UpdateStudentImmunisationRegisterRequest) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentImmunisationRegisterAttachments, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register/attachments
func (c *Client) GetAllStudentImmunisationRegisterAttachments(studentCode string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentImmunisationRegisterAttachment, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register/attachments
func (c *Client) AddStudentImmunisationRegisterAttachment(studentCode string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadStudentImmunisationRegisterAttachment, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register/attachments/{attach_id}
func (c *Client) GetStudentImmunisationRegisterAttachment(studentCode string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentImmunisationRegisterAttachment, path: /{cmpy_code}/students/{stud_code}/medical/immunisations/register/attachments/{attach_id}
func (c *Client) DeleteStudentImmunisationRegisterAttachment(studentCode string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicalMedications, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications
func (c *Client) GetAllStudentMedicalMedications(studentCode string, medicalConditionCode string) ([]StudentMedicationResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicalMedication, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications
func (c *Client) AddStudentMedicalMedication(studentCode string, medicalConditionCode string, payload AddStudentMedicationRequest) (StudentMedicationResponse, error) {
	// TODO: Implementation
	return StudentMedicationResponse{}, nil
}

// op: GetStudentMedicalMedicationByCode, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}
func (c *Client) GetStudentMedicalMedicationByCode(studentCode string, medicalConditionCode string, medicationID string) (StudentMedicationResponse, error) {
	// TODO: Implementation
	return StudentMedicationResponse{}, nil
}

// op: UpdateStudentMedicalMedication, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}
func (c *Client) UpdateStudentMedicalMedication(studentCode string, medicalConditionCode string, medicationID string, payload UpdateStudentMedicationRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMedicalMedication, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}
func (c *Client) PatchStudentMedicalMedication(studentCode string, medicalConditionCode string, medicationID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentMedicalMedication, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}
func (c *Client) DeleteStudentMedicalMedication(studentCode string, medicalConditionCode string, medicationID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicationNotes, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes
func (c *Client) GetAllStudentMedicationNotes(studentCode string, medicalConditionCode string, medicationID string) ([]StudentMedicationNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicationNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes
func (c *Client) AddStudentMedicationNote(studentCode string, medicalConditionCode string, medicationID string, payload AddStudentMedicationNoteRequest) (StudentMedicationNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicationNoteResponse{}, nil
}

// op: GetStudentMedicationNoteByCode, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes/{note_uid}
func (c *Client) GetStudentMedicationNoteByCode(studentCode string, medicalConditionCode string, medicationID string, noteID string) (StudentMedicationNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicationNoteResponse{}, nil
}

// op: UpdateStudentMedicationNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes/{note_uid}
func (c *Client) UpdateStudentMedicationNote(studentCode string, medicalConditionCode string, medicationID string, noteID string, payload UpdateStudentMedicationNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMedicationNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes/{note_uid}
func (c *Client) PatchStudentMedicationNote(studentCode string, medicalConditionCode string, medicationID string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentMedicationNote, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/notes/{note_uid}
func (c *Client) DeleteStudentMedicationNote(studentCode string, medicalConditionCode string, medicationID string, noteID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicationAttachments, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/attachments
func (c *Client) GetAllStudentMedicationAttachments(studentCode string, medicalConditionCode string, medicationID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicationAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/attachments
func (c *Client) AddStudentMedicationAttachment(studentCode string, medicalConditionCode string, medicationID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadStudentMedicationAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/attachments/{attach_id}
func (c *Client) GetStudentMedicationAttachment(studentCode string, medicalConditionCode string, medicationID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentMedicationAttachment, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentMedicationAttachment(studentCode string, medicalConditionCode string, medicationID string, attachmedID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicationSchedules, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules
func (c *Client) GetAllStudentMedicationSchedules(studentCode string, medicalConditionCode string, medicationID string) ([]StudentMedicationScheduleResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicationSchedule, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules
func (c *Client) AddStudentMedicationSchedule(studentCode string, medicalConditionCode string, medicationID string, payload AddStudentMedicationScheduleRequest) (StudentMedicationScheduleResponse, error) {
	// TODO: Implementation
	return StudentMedicationScheduleResponse{}, nil
}

// op: GetStudentMedicationScheduleByCode, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules/{sched_uid}
func (c *Client) GetStudentMedicationScheduleByCode(studentCode string, medicalConditionCode string, medicationID string, scheduleID string) (StudentMedicationScheduleResponse, error) {
	// TODO: Implementation
	return StudentMedicationScheduleResponse{}, nil
}

// op: UpdateStudentMedicationSchedule, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules/{sched_uid}
func (c *Client) UpdateStudentMedicationSchedule(studentCode string, medicalConditionCode string, medicationID string, scheduleID string, payload UpdateStudentMedicationScheduleRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMedicationSchedule, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules/{sched_uid}
func (c *Client) PatchStudentMedicationSchedule(studentCode string, medicalConditionCode string, medicationID string, scheduleID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentMedicationSchedule, path: /{cmpy_code}/students/{stud_code}/medical/conditions/{mcond_code}/medications/{medication_uid}/schedules/{sched_uid}
func (c *Client) DeleteStudentMedicationSchedule(studentCode string, medicalConditionCode string, medicationID string, scheduleID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentStandardMedicalNotes, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard
func (c *Client) GetAllStudentStandardMedicalNotes(studentCode string) ([]StudentMedicalStandardNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicalStandardNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard
func (c *Client) AddStudentMedicalStandardNote(studentCode string, payload AddStudentMedicalStandardNoteRequest) (StudentMedicalStandardNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalStandardNoteResponse{}, nil
}

// op: GetStudentMedicalStandardNoteByID, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}
func (c *Client) GetStudentMedicalStandardNoteByID(studentCode string, noteID string) (StudentMedicalStandardNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalStandardNoteResponse{}, nil
}

// op: UpdateStudentMedicalStandardNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}
func (c *Client) UpdateStudentMedicalStandardNote(studentCode string, noteID string, payload UpdateStudentMedicalStandardNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMedicalStandardNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}
func (c *Client) PatchStudentMedicalStandardNote(studentCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentMedicalStandardNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}
func (c *Client) DeleteStudentMedicalStandardNote(studentCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicalStandardNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}/attachments
func (c *Client) GetAllStudentMedicalStandardNotesAttachment(studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicalStandardNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}/attachments
func (c *Client) AddStudentMedicalStandardNotesAttachment(studentCode string, noteID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadStudentMedicalStandardNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) GetStudentMedicalStandardNotesAttachment(studentCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentMedicalStandardNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/standard/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentMedicalStandardNotesAttachment(studentCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentConfidentialMedicalNotes, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential
func (c *Client) GetAllStudentConfidentialMedicalNotes(studentCode string) ([]StudentMedicalConfidentialNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicalConfidentialNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential
func (c *Client) AddStudentMedicalConfidentialNote(studentCode string, payload AddStudentMedicalConfidentialNoteRequest) (StudentMedicalConfidentialNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalConfidentialNoteResponse{}, nil
}

// op: GetStudentMedicalConfidentialNoteByID, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}
func (c *Client) GetStudentMedicalConfidentialNoteByID(studentCode string, noteID string) (StudentMedicalConfidentialNoteResponse, error) {
	// TODO: Implementation
	return StudentMedicalConfidentialNoteResponse{}, nil
}

// op: UpdateStudentMedicalConfidentialNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}
func (c *Client) UpdateStudentMedicalConfidentialNote(studentCode string, noteID string, payload UpdateStudentMedicalConfidentialNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMedicalConfidentialNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}
func (c *Client) PatchStudentMedicalConfidentialNote(studentCode string, noteID string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteStudentMedicalConfidentialNote, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}
func (c *Client) DeleteStudentMedicalConfidentialNote(studentCode string, noteID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicalConfidentialNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}/attachments
func (c *Client) GetAllStudentMedicalConfidentialNotesAttachment(studentCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentMedicalConfidentialNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}/attachments
func (c *Client) AddStudentMedicalConfidentialNotesAttachment(studentCode string, noteID string) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadStudentMedicalConfidentialNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) GetStudentMedicalConfidentialNotesAttachment(studentCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteStudentMedicalConfidentialNotesAttachment, path: /{cmpy_code}/students/{stud_code}/medical/notes/confidential/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteStudentMedicalConfidentialNotesAttachment(studentCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentPractitioners, path: /{cmpy_code}/students/{stud_code}/medical/practitioners
func (c *Client) GetAllStudentPractitioners(studentCode string) ([]StudentPractitionerResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddStudentPractitioner, path: /{cmpy_code}/students/{stud_code}/medical/practitioners
func (c *Client) AddStudentPractitioner(studentCode string, payload AddStudentPractitionerRequest) (StudentPractitionerResponse, error) {
	// TODO: Implementation
	return StudentPractitionerResponse{}, nil
}

// op: GetStudentPractitionerByPracNum, path: /{cmpy_code}/students/{stud_code}/medical/practitioners/{prac_num}
func (c *Client) GetStudentPractitionerByPracNum(studentCode string, pracNum string) (StudentPractitionerResponse, error) {
	// TODO: Implementation
	return StudentPractitionerResponse{}, nil
}

// op: DeleteStudentPractitioner, path: /{cmpy_code}/students/{stud_code}/medical/practitioners/{prac_num}
func (c *Client) DeleteStudentPractitioner(studentCode string, pracNum string) error {
	// TODO: Implementation
	return nil
}

// op: UpdateStudentPractitioner, path: /{cmpy_code}/students/{stud_code}/medical/practitioners/{prac_num}
func (c *Client) UpdateStudentPractitioner(studentCode string, pracNum string, payload UpdateStudentPractitionerRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentPractitioner, path: /{cmpy_code}/students/{stud_code}/medical/practitioners/{prac_num}
func (c *Client) PatchStudentPractitioner(studentCode string, pracNum string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAllStudentMedicalSupplementaries, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries
func (c *Client) GetAllStudentMedicalSupplementaries(studentCode string) ([]StudentMedicalSupplementaryResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: CreateStudentMedicalSupplementary, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries
func (c *Client) CreateStudentMedicalSupplementary(stuentCode string, payload AddStudentMedicalSupplementaryRequest) (StudentMedicalSupplementaryResponse, error) {
	// TODO: Implementation
	return StudentMedicalSupplementaryResponse{}, nil
}

// op: GetStudentMedicalSupplementaryByCode, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries/{msupp_code}
func (c *Client) GetStudentMedicalSupplementaryByCode(studentCode string, supplementaryCode string) (StudentMedicalSupplementaryResponse, error) {
	// TODO: Implementation
	return StudentMedicalSupplementaryResponse{}, nil
}

// op: DeleteStudentMedicalSupplementary, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries/{msupp_code}
func (c *Client) DeleteStudentMedicalSupplementary(studentCode string, supplementaryCode string) error {
	// TODO: Implementation
	return nil
}

// op: UpdateStudentMedicalSupplementary, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries/{msupp_code}
func (c *Client) UpdateStudentMedicalSupplementary(studentCode string, supplementaryCode string, payload UpdateStudentMedicalSupplementaryRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchStudentMedicalSupplementary, path: /{cmpy_code}/students/{stud_code}/medical/supplementaries/{msupp_code}
func (c *Client) PatchStudentMedicalSupplementary(studentCode string, supplementaryCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}
