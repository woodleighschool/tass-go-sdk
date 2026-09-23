package tassstudent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetAllAbsenceReasonOptions, path=/{cmpy_code}/options/student/attendance/absencereasons
func (c *Client) GetAllAbsenceReasonOptions(ctx context.Context) ([]AbsenceReasonOptionsResponse, error) {
	var result []AbsenceReasonOptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencereasons", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllAbsenceTypeOptions, path=/{cmpy_code}/options/student/attendance/absencetypes
func (c *Client) GetAllAbsenceTypeOptions(ctx context.Context) ([]tasscommon.OptionsResponseActive, error) {
	var result []tasscommon.OptionsResponseActive
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencetypes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllAbsenceTypeOptions, path=/{cmpy_code}/options/student/attendance/campuses
func (c *Client) GetAllCampusOptionsForAttendance(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/campuses", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllAbsenceTypeOptions, path=/{cmpy_code}/options/student/attendance/houses
func (c *Client) GetAllHouseOptionsForAttendance(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/houses", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllAbsenceTypeOptions, path=/{cmpy_code}/options/student/attendance/yeargroups
func (c *Client) GetAllYearGroupOptionsForAttendance(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/yeargroups", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllCommunicationRuleGenderOptions, path=/{cmpy_code}/options/students/communicationrules/genders
func (c *Client) GetAllCommunicationRuleGenderOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/communicationrules/genders", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllCommunicationRuleTypeOptions, path=/{cmpy_code}/options/students/communicationrules/types
func (c *Client) GetAllCommunicationRuleTypeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/communicationrules/types", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllMedicalConditionTypesOptions, path=/{cmpy_code}/options/students/medical/conditions/types
func (c *Client) GetAllMedicalConditionTypesOptions(ctx context.Context) ([]tasscommon.OptionsResponseActive, error) {
	var result []tasscommon.OptionsResponseActive
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/medical/conditions/types", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllMedicalConditionsOptions, path=/{cmpy_code}/options/students/medical/illnesses/conditions
func (c *Client) GetAllMedicalConditionsOptions(ctx context.Context) ([]tasscommon.OptionsResponseActive, error) {
	var result []tasscommon.OptionsResponseActive
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencereasons", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllMedicalTreatmentsOptions, path=/{cmpy_code}/options/students/medical/illnesses/treatments
func (c *Client) GetAllMedicalTreatmentOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencereasons", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllImmunisationTypesOptions, path=/{cmpy_code}/options/students/medical/immunisations/types
func (c *Client) GetAllImmunisationTypesOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencereasons", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllImmunisationStatusesOptions, path=/{cmpy_code}/options/students/medical/immunisations/register/statuses
func (c *Client) GetAllImmunisationStatusesOptions(ctx context.Context) ([]tasscommon.OptionsResponseActive, error) {
	var result []tasscommon.OptionsResponseActive
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencereasons", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentMedicalNoteCategoryOptions, path=/{cmpy_code}/options/students/medical/notes/categories
func (c *Client) GetAllStudentMedicalNoteCategoryOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencereasons", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllPractitionerTypesOptions, path=/{cmpy_code}/options/students/medical/practitioners/types
func (c *Client) GetAllPractitionerTypesOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencereasons", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllSupplementaryTypesOptions, path=/{cmpy_code}/options/students/medical/supplementaries/types
func (c *Client) GetAllSupplementaryTypesOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/attendance/absencereasons", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllReligionOptions, path=/{cmpy_code}/options/students/religions
func (c *Client) GetAllReligionOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/religions", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllResidencyStatusOptions, path=/{cmpy_code}/options/students/residencystatuses
func (c *Client) GetAllResidencyStatusOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/residencystatuses", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllCampusOptions, path=/{cmpy_code}/options/students/campuses
func (c *Client) GetAllCampusOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/campuses", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllFeederSchoolOptions, path=/{cmpy_code}/options/students/feederschools
func (c *Client) GetAllFeederSchoolOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/feederschools", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllHouseOptions, path=/{cmpy_code}/options/students/houses
func (c *Client) GetAllHouseOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/houses", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllYearGroupOptions, path=/{cmpy_code}/options/students/yeargroups
func (c *Client) GetAllYearGroupOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/yeargroups", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllNextYearIndicatorOptions, path=/{cmpy_code}/options/students/nextyearindicators
func (c *Client) GetAllNextYearIndicatorOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/nextyearindicators", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllComparativeReportingTypeOptions, path=/{cmpy_code}/options/students/comparativereportingtypes
func (c *Client) GetAllComparativeReportingTypeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/comparativereportingtypes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllPCTutorGroupOptions, path=/{cmpy_code}/options/students/pctutorgroups
func (c *Client) GetAllPCTutorGroupOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/pctutorgroups", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentNoteCategoryOptions, path=/{cmpy_code}/options/students/notes/categories
func (c *Client) GetAllStudentNoteCategoryOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/notes/categories", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentUDAreaOptions, path=/{cmpy_code}/options/students/udareas
func (c *Client) GetAllStudentUDAreaOptions(ctx context.Context) ([]tasscommon.UDAreaOptionsResponse, error) {
	var result []tasscommon.UDAreaOptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/udareas", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetSingleStudentUDAreaOptions, path=/{cmpy_code}/options/students/udareas/{area_code}
func (c *Client) GetSingleStudentUDAreaOptions(ctx context.Context, areaCode string) (tasscommon.UDAreaOptionsResponse, error) {
	var result tasscommon.UDAreaOptionsResponse
	url := fmt.Sprintf("/options/students/udareas/%s", areaCode)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return tasscommon.UDAreaOptionsResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.UDAreaOptionsResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentUDFieldOptions, path=/{cmpy_code}/options/students/udfields
func (c *Client) GetAllStudentUDFieldOptions(ctx context.Context) (StudentUDFieldOptionResponse, error) {
	var result StudentUDFieldOptionResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/udfields", nil, nil, http.StatusOK)
	if err != nil {
		return StudentUDFieldOptionResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return StudentUDFieldOptionResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentCountryOptions, path=/{cmpy_code}/options/students/mceecdya/countries
func (c *Client) GetAllStudentCountryOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/mceecdya/countries", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentLanguageOptions, path=/{cmpy_code}/options/students/mceecdya/languages
func (c *Client) GetAllStudentLanguageOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/mceecdya/languages", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllOccupationalGroupOptions, path=/{cmpy_code}/options/students/mceecdya/occupationalgroups
func (c *Client) GetAllOccupationalGroupOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/mceecdya/occupationalgroups", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllStudentIndigenousTypeOptions, path=/{cmpy_code}/options/students/mceecdya/indigenoustypes
func (c *Client) GetAllStudentIndigenousTypeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/mceecdya/indigenoustypes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllSchoolEducationTypeOptions, path=/{cmpy_code}/options/students/mceecdya/schooleducation
func (c *Client) GetAllSchoolEducationTypeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/mceecdya/schooleducation", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllNonSchoolEducationTypeOptions, path=/{cmpy_code}/options/students/mceecdya/nonschooleducation
func (c *Client) GetAllNonSchoolEducationTypeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/students/mceecdya/nonschooleducation", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
