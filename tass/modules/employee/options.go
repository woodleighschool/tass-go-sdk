package tassemployee

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetAllCountryOptions, path=/{cmpy_code}/options/employees/countries
func (c *Client) GetAllCountryOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/countries", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllEmployeeStatusOptions, path=/{cmpy_code}/options/employees/statuses
func (c *Client) GetAllEmployeeStatusOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/statuses", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllGenderOptions, path=/{cmpy_code}/options/employees/genders
func (c *Client) GetAllGenderOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/genders", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllIndigneousTypeOptions, path=/{cmpy_code}/options/employees/indigenoustype
func (c *Client) GetAllIndigneousTypeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/indigenoustype", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllMainActivityOptions, path=/{cmpy_code}/options/employees/mainactivities
func (c *Client) GetAllMainActivityOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/mainactivities", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllMaritalStatusOptions, path=/{cmpy_code}/options/employees/maritalstatuses
func (c *Client) GetAllMaritalStatusOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/maritalstatuses", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllTerminationReasonOptions, path=/{cmpy_code}/options/employees/terminationreasons
func (c *Client) GetAllTerminationReasonOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/terminationreasons", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllTitleOptions, path=/{cmpy_code}/options/employees/titles
func (c *Client) GetAllTitleOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/titles", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllVendorOptions, path=/{cmpy_code}/options/employees/vendors
func (c *Client) GetAllVendorOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/vendors", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllEmployeeNoteCategoryOptions, path=/{cmpy_code}/options/employees/notes/categories
func (c *Client) GetAllEmployeeNoteCategoryOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/notes/categories", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllPDProviderOptions, path=/{cmpy_code}/options/employees/pdactivities/providers
func (c *Client) GetAllPDProviderOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/pdactivities/providers", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllPDStatusOptions, path=/{cmpy_code}/options/employees/pdactivities/statuses
func (c *Client) GetAllPDStatusOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/pdactivities/statuses", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllPDTypesOptions, path=/{cmpy_code}/options/employees/pdactivities/types
func (c *Client) GetAllPDTypesOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/pdactivities/types", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllPDUDFieldOptions, path=/{cmpy_code}/options/employees/pdactivities/udfields
func (c *Client) GetAllPDUDFieldOptions(ctx context.Context) ([]tasscommon.UDFieldOptionsResponse, error) {
	var result []tasscommon.UDFieldOptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/pdactivities/udfields", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllQualificationCategoryOptions, path=/{cmpy_code}/options/employees/qualifications/categories
func (c *Client) GetAllQualificationCategoryOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/qualifications/categories", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllQualificationInstitutionOptions, path=/{cmpy_code}/options/employees/qualifications/institutions
func (c *Client) GetAllQualificationInstitutionOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/qualifications/institutions", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllUDAreaOptions, path=/{cmpy_code}/options/employees/udareas
func (c *Client) GetAllUDAreaOptions(ctx context.Context) ([]tasscommon.UDAreaOptionsResponse, error) {
	var result []tasscommon.UDAreaOptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/employees/udareas", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetSingleUDAreaOptions, path=/{cmpy_code}/options/employees/udareas/{area_code}
func (c *Client) GetSingleUDAreaOptions(ctx context.Context, areaCode string) (tasscommon.UDAreaOptionsResponse, error) {
	var result tasscommon.UDAreaOptionsResponse
	url := fmt.Sprintf("/options/employees/udareas/%s", areaCode)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return tasscommon.UDAreaOptionsResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.UDAreaOptionsResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAllAccrualCodeOptions, path=/{cmpy_code}/options/payroll/leaveentitlements/accrualcodes
func (c *Client) GetAllAccrualCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/options/payroll/leaveentitlements/accrualcodes", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
