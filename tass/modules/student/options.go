package tassstudent

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// Absences

// op: GetAllAbsenceReasonOptions, path: /{cmpy_code}/options/student/attendance/absencereasons
func (c *Client) GetAllAbsenceReasonOptions() ([]AbsenceReasonOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAbsenceTypeOptions, path: /{cmpy_code}/options/student/attendance/absencetypes
func (c *Client) GetAllAbsenceTypeOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAbsenceTypeOptions, path: /{cmpy_code}/options/student/attendance/campuses
func (c *Client) GetAllCampusOptionsForAttendance() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAbsenceTypeOptions, path: /{cmpy_code}/options/student/attendance/houses
func (c *Client) GetAllHouseOptionsForAttendance() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAbsenceTypeOptions, path: /{cmpy_code}/options/student/attendance/yeargroups
func (c *Client) GetAllYearGroupOptionsForAttendance() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// Communication Rules

// op: GetAllCommunicationRuleGenderOptions, path: /{cmpy_code}/options/students/communicationrules/genders
func (c *Client) GetAllCommunicationRuleGenderOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllCommunicationRuleTypeOptions, path: /{cmpy_code}/options/students/communicationrules/types
func (c *Client) GetAllCommunicationRuleTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// Medical

func (c *Client) GetAllMedicalConditionTypesOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return []tasscommon.OptionsResponseActive{}, nil
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

func (c *Client) GetAllImmunisationStatusesOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllStudentMedicalNoteCategoryOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllPractitionerTypesOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllSupplementaryTypesOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// Students

// op: GetAllReligionOptions, path: /{cmpy_code}/options/students/religions
func (c *Client) GetAllReligionOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllResidencyStatusOptions, path: /{cmpy_code}/options/students/residencystatuses
func (c *Client) GetAllResidencyStatusOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllCampusOptions, path: /{cmpy_code}/options/students/campuses
func (c *Client) GetAllCampusOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllFeederSchoolOptions, path: /{cmpy_code}/options/students/feederschools
func (c *Client) GetAllFeederSchoolOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllHouseOptions, path: /{cmpy_code}/options/students/houses
func (c *Client) GetAllHouseOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllYearGroupOptions, path: /{cmpy_code}/options/students/yeargroups
func (c *Client) GetAllYearGroupOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllNextYearIndicatorOptions, path: /{cmpy_code}/options/students/nextyearindicators
func (c *Client) GetAllNextYearIndicatorOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllComparativeReportingTypeOptions, path: /{cmpy_code}/options/students/camparativereportingtypes
func (c *Client) GetAllComparativeReportingTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllPCTutorGroupOptions, path: /{cmpy_code}/options/students/pctutorgroups
func (c *Client) GetAllPCTutorGroupOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllStudentNoteCategoryOptions, path: /{cmpy_code}/options/students/notes/categories
func (c *Client) GetAllStudentNoteCategoryOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllStudentUDAreaOptions, path: /{cmpy_code}/options/students/udareas
func (c *Client) GetAllStudentUDAreaOptions() ([]tasscommon.UDAreaOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetSingleStudentUDAreaOptions, path: /{cmpy_code}/options/students/udareas/{area_code}
func (c *Client) GetSingleStudentUDAreaOptions(areaCode string) (tasscommon.UDAreaOptionsResponse, error) {
	// TODO: Implementation
	return tasscommon.UDAreaOptionsResponse{}, nil
}

// op: GetAllStudentCountryOptions, path: /{cmpy_code}/options/students/mceecdya/countries
func (c *Client) GetAllStudentCountryOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllStudentLanguageOptions, path: /{cmpy_code}/options/students/mceecdya/languages
func (c *Client) GetAllStudentLanguageOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllOccupationalGroupOptions, path: /{cmpy_code}/options/students/mceecdya/occupationalgroups
func (c *Client) GetAllOccupationalGroupOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllStudentIndigenousTypeOptions, path: /{cmpy_code}/options/students/mceecdya/indigenoustypes
func (c *Client) GetAllStudentIndigenousTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllSchoolEducationTypeOptions, path: /{cmpy_code}/options/students/mceecdya/schooleducation
func (c *Client) GetAllSchoolEducationTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllNonSchoolEducationTypeOptions, path: /{cmpy_code}/options/students/mceecdya/nonschooleducation
func (c *Client) GetAllNonSchoolEducationTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}
