package tasscommon

type OptionsResponse struct {
	Code        string `json:"code"`
	Description string `json:"desc"`
}

type OptionsResponseActive struct {
	Code        string `json:"code"`
	Description string `json:"desc"`
	Active      bool   `json:"is_active"`
}

type UDAreaOptionsResponse struct {
	CompanyCode     string       `json:"cmpy_code"`
	AreaCode        string       `json:"area_code"`
	AreaDescription string       `json:"area_desc"`
	UDFields        UDFieldTypes `json:"ud_fields"`
}

type UDFieldOptionsResponse struct {
	CompanyCode string `json:"cmpy_code"`
}
