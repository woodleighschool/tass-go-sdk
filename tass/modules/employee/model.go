package tassemployee

import (
	"time"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// Employees

type EmployeeResponse struct {
	AddressLine1          *string           `json:"add1_text,omitempty"`
	AddressLine2          *string           `json:"add2_text,omitempty"`
	AlternateID           *string           `json:"alt_id,omitempty"`
	BirthDate             *time.Time        `json:"birth_date,omitempty"` // Must be a date
	CessationType         *string           `json:"cessation_type,omitempty"`
	City                  *string           `json:"city_text,omitempty"`
	CompanyCode           string            `json:"cmpy_code"`
	CountryCode           *string           `json:"country_code,omitempty"`
	CountryText           *string           `json:"country_text,omitempty"`
	Deceased              bool              `json:"deceased_flg"`
	DriversLicense        *string           `json:"driv_lic_text,omitempty"`
	Email                 *string           `json:"e_mail,omitempty"`
	EmployeeCode          string            `json:"emp_code"`
	FirstName             *string           `json:"first_name,omitempty"`
	Gender                *string           `json:"gender,omitempty"`
	IndiginousStatus      *IndiginousStatus `json:"indig_status,omitempty"`
	Initials              *string           `json:"initials,omitempty"`
	Role                  *string           `json:"main_activity,omitempty"`
	MaritalStatus         *string           `json:"marital_stat_flag,omitempty"`
	MobilePhone           *string           `json:"mob_phone,omitempty"`
	NextOfKinAddressLine1 *string           `json:"nok_add1_text,omitempty"`
	NextOfKinAddressLine2 *string           `json:"nok_add2_text,omitempty"`
	NextOfKinCity         *string           `json:"nok_city_text,omitempty"`
	NextOfKinCountry      *string           `json:"nok_country_text,omitempty"`
	NextOfKinName         *string           `json:"nok_name_text,omitempty"`
	NextOfKinHomePhone    *string           `json:"nok_phone_h_text,omitempty"`
	NextOfKinWorkPhone    *string           `json:"nok_phone_w_text,omitempty"`
	NextOfKinPostCode     *string           `json:"nok_post_code,omitempty"`
	NextOfKinRelationship *string           `json:"nok_relat_text,omitempty"`
	NextOfKinState        *string           `json:"nok_state_text,omitempty"`
	OtherNames            *string           `json:"other_name,omitempty"`
	HomePhone             *string           `json:"phone_h_text,omitempty"`
	WorkPhone             *string           `json:"phone_w_text,omitempty"`
	PositionText          *string           `json:"position_text,omitempty"`
	PositionTitle         *string           `json:"position_title,omitempty"`
	PostCode              *string           `json:"post_code,omitempty"`
	PreferredName         *string           `json:"preferred_name,omitempty"`
	PreviousPayrollID     *string           `json:"previous_payroll_id,omitempty"`
	SchoolEmail           *string           `json:"school_email,omitempty"`
	SMSFlag               bool              `json:"sms_flg"`
	StartDate             *time.Time        `json:"start_date,omitempty"` // Must be a date
	State                 *string           `json:"state_text,omitempty"`
	Status                *string           `json:"status_text,omitempty"`
	Surname               *string           `json:"surname,omitempty"`
	Suffix                *string           `json:"suffix,omitempty"`
	SupervisorCode        *string           `json:"supervisor_code,omitempty"`
	Supervisor2Code       *string           `json:"supervisor2_code,omitempty"`
	TeacherCode           *string           `json:"tch_code,omitempty"`
	TerminationDate       *time.Time        `json:"term_date,omitempty"` // Must be a date
	Title                 *string           `json:"title,omitempty"`
	UpdatedOn             *time.Time        `json:"update_on,omitempty"` // Must be a date
	VendorCode            *string           `json:"vend_code,omitempty"`
}

type UpdateEmployeeRequest struct {
	BirthDate     string `json:"birth_date"`     // Must be in format (yyyy-mm-dd) and < today
	EmployeeCode  string `json:"emp_code"`       // Max length 7
	FirstName     string `json:"first_name"`     // Max length 50
	Gender        string `json:"gender"`         // Max length 3, Must conform to GetAllGenderOptions response
	Initials      string `json:"initials"`       // Max length 5
	PostCode      string `json:"post_code"`      // Max length 10
	PreferredName string `json:"preferred_name"` // Max length 50
	StartDate     string `json:"start_date"`     // Must be a date
	State         string `json:"state_text"`     // Max length 3
	Status        string `json:"status_text"`    // Max length 1, Must conform to GetAllEmployeeStatusOptions response
	Surname       string `json:"surname"`        // Max length 50
	Title         string `json:"title"`          // Max length 15, Must conform to GetAllTitleOptions response

	AddressLine1          *string    `json:"add1_text,omitempty"`      // Max length 60
	AddressLine2          *string    `json:"add2_text,omitempty"`      // Max length 60
	AlternateID           *string    `json:"alt_id,omitempty"`         // Max length 40, Must be unique
	CessationType         *string    `json:"cessation_type,omitempty"` // Max length 1, Must conform to CessationType
	City                  *string    `json:"city_text,omitempty"`      // Max length 46
	CountryCode           *string    `json:"country_code,omitempty"`   // Max length 2
	CountryText           *string    `json:"country_text,omitempty"`   // Max length 20
	Deceased              bool       `json:"deceased_flg"`
	DriversLicense        *string    `json:"driv_lic_text,omitempty"`       // Max length 10
	Email                 *string    `json:"e_mail,omitempty"`              // Max length 60
	IndiginousStatus      *string    `json:"indig_status,omitempty"`        // Max length 1, Must conform to IndiginousStatus
	Role                  *string    `json:"main_activity,omitempty"`       // Max length 4, Must conform to GetAllMainActivityOptions response
	MaritalStatus         *string    `json:"marital_stat_flag,omitempty"`   // Max length 1, Must conform to GetAllMaritalStatusOptions response
	MobilePhone           *string    `json:"mob_phone,omitempty"`           // Max length 30
	NextOfKinAddressLine1 *string    `json:"nok_add1_text,omitempty"`       // Max length 60
	NextOfKinAddressLine2 *string    `json:"nok_add2_text,omitempty"`       // Max length 60
	NextOfKinCity         *string    `json:"nok_city_text,omitempty"`       // Max length 46
	NextOfKinCountry      *string    `json:"nok_country_text,omitempty"`    // Max length 20
	NextOfKinName         *string    `json:"nok_name_text,omitempty"`       // Max length 30
	NextOfKinHomePhone    *string    `json:"nok_phone_h_text,omitempty"`    // Max length 30
	NextOfKinWorkPhone    *string    `json:"nok_phone_w_text,omitempty"`    // Max length 30
	NextOfKinPostCode     *string    `json:"nok_post_code,omitempty"`       // Max length 10
	NextOfKinRelationship *string    `json:"nok_relat_text,omitempty"`      // Max length 20
	NextOfKinState        *string    `json:"nok_state_text,omitempty"`      // Max length 3
	OtherNames            *string    `json:"other_name,omitempty"`          // Max length 50
	HomePhone             *string    `json:"phone_h_text,omitempty"`        // Max length 30
	WorkPhone             *string    `json:"phone_w_text,omitempty"`        // Max length 30
	PositionText          *string    `json:"position_text,omitempty"`       // Max length 20
	PositionTitle         *string    `json:"position_title,omitempty"`      // Max length 100
	PreviousPayrollID     *string    `json:"previous_payroll_id,omitempty"` // Max length 200
	SchoolEmail           *string    `json:"school_email,omitempty"`        // Max length 60
	SMSFlag               bool       `json:"sms_flg"`
	Suffix                *string    `json:"suffix,omitempty"`           // Max length 30
	SupervisorCode        *string    `json:"supervisor_code,omitempty"`  // Max length 7, must be an existing employee and cannot be this employee
	Supervisor2Code       *string    `json:"supervisor2_code,omitempty"` // Max length 7, must be an existing employee and cannot be this employee
	TerminationDate       *time.Time `json:"term_date,omitempty"`        // Must be a date (yyyy-mm-dd) and be after start_date
	VendorCode            *string    `json:"vend_code,omitempty"`        // Max length 8, must conform to GetAllVendorOptions response
}

type AddEmployeeRequest struct {
	BirthDate     string `json:"birth_date"`     // Must be in format (yyyy-mm-dd) and < today
	EmployeeCode  string `json:"emp_code"`       // Max length 7, can be blank for auto-numbering
	FirstName     string `json:"first_name"`     // Max length 50
	Gender        string `json:"gender"`         // Max length 3, Must conform to GetAllGenderOptions response
	Initials      string `json:"initials"`       // Max length 5
	PostCode      string `json:"post_code"`      // Max length 10
	PreferredName string `json:"preferred_name"` // Max length 50
	StartDate     string `json:"start_date"`     // Must be a date
	State         string `json:"state_text"`     // Max length 3
	Status        string `json:"status_text"`    // Max length 1, Must conform to GetAllEmployeeStatusOptions response
	Surname       string `json:"surname"`        // Max length 50
	Title         string `json:"title"`          // Max length 15, Must conform to GetAllTitleOptions response

	AddressLine1          *string    `json:"add1_text,omitempty"`      // Max length 60
	AddressLine2          *string    `json:"add2_text,omitempty"`      // Max length 60
	AlternateID           *string    `json:"alt_id,omitempty"`         // Max length 40, Must be unique
	CessationType         *string    `json:"cessation_type,omitempty"` // Max length 1, Must conform to CessationType
	City                  *string    `json:"city_text,omitempty"`      // Max length 46
	CountryCode           *string    `json:"country_code,omitempty"`   // Max length 2
	CountryText           *string    `json:"country_text,omitempty"`   // Max length 20
	Deceased              bool       `json:"deceased_flg"`
	DriversLicense        *string    `json:"driv_lic_text,omitempty"`       // Max length 10
	Email                 *string    `json:"e_mail,omitempty"`              // Max length 60
	IndiginousStatus      *string    `json:"indig_status,omitempty"`        // Max length 1, Must conform to IndiginousStatus
	Role                  *string    `json:"main_activity,omitempty"`       // Max length 4, Must conform to GetAllMainActivityOptions response
	MaritalStatus         *string    `json:"marital_stat_flag,omitempty"`   // Max length 1, Must conform to GetAllMaritalStatusOptions response
	MobilePhone           *string    `json:"mob_phone,omitempty"`           // Max length 30
	NextOfKinAddressLine1 *string    `json:"nok_add1_text,omitempty"`       // Max length 60
	NextOfKinAddressLine2 *string    `json:"nok_add2_text,omitempty"`       // Max length 60
	NextOfKinCity         *string    `json:"nok_city_text,omitempty"`       // Max length 46
	NextOfKinCountry      *string    `json:"nok_country_text,omitempty"`    // Max length 20
	NextOfKinName         *string    `json:"nok_name_text,omitempty"`       // Max length 30
	NextOfKinHomePhone    *string    `json:"nok_phone_h_text,omitempty"`    // Max length 30
	NextOfKinWorkPhone    *string    `json:"nok_phone_w_text,omitempty"`    // Max length 30
	NextOfKinPostCode     *string    `json:"nok_post_code,omitempty"`       // Max length 10
	NextOfKinRelationship *string    `json:"nok_relat_text,omitempty"`      // Max length 20
	NextOfKinState        *string    `json:"nok_state_text,omitempty"`      // Max length 3
	OtherNames            *string    `json:"other_name,omitempty"`          // Max length 50
	HomePhone             *string    `json:"phone_h_text,omitempty"`        // Max length 30
	WorkPhone             *string    `json:"phone_w_text,omitempty"`        // Max length 30
	PositionText          *string    `json:"position_text,omitempty"`       // Max length 20
	PositionTitle         *string    `json:"position_title,omitempty"`      // Max length 100
	PreviousPayrollID     *string    `json:"previous_payroll_id,omitempty"` // Max length 200
	SchoolEmail           *string    `json:"school_email,omitempty"`        // Max length 60
	SMSFlag               bool       `json:"sms_flg"`
	Suffix                *string    `json:"suffix,omitempty"`           // Max length 30
	SupervisorCode        *string    `json:"supervisor_code,omitempty"`  // Max length 7, must be an existing employee and cannot be this employee
	Supervisor2Code       *string    `json:"supervisor2_code,omitempty"` // Max length 7, must be an existing employee and cannot be this employee
	TerminationDate       *time.Time `json:"term_date,omitempty"`        // Must be a date (yyyy-mm-dd) and be after start_date
	VendorCode            *string    `json:"vend_code,omitempty"`        // Max length 8, must conform to GetAllVendorOptions response
}

type EmployeeStandardNoteResponse struct {
	CompanyCode   string     `json:"cmpy_code"`
	EmployeeCode  string     `json:"emp_code"`
	Category      *string    `json:"note_cat,omitempty"`
	Date          *time.Time `json:"note_date,omitempty"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text          *string    `json:"note_text,omitempty"`
	ID            string     `json:"note_uid"` // TODO: UUID? Must be a UUID
	HasAttachment bool       `json:"has_attachment"`
}

type AddEmployeeStandardNoteRequest struct {
	Category string `json:"note_cat"`  // Max length 3, Must conform to GetAllEmployeeNoteCategoryOptions
	Date     string `json:"note_date"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text     string `json:"note_text"` // Max length 4000
}

type UpdateEmployeeStandardNoteRequest struct {
	Category string `json:"note_cat"`  // Max length 3, Must conform to GetAllEmployeeNoteCategoryOptions
	Date     string `json:"note_date"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text     string `json:"note_text"` // Max length 4000
}

type EmployeeConfidentialNoteResponse struct {
	CompanyCode   string     `json:"cmpy_code"`
	EmployeeCode  string     `json:"emp_code"`
	Category      *string    `json:"note_cat,omitempty"`
	Date          *time.Time `json:"note_date,omitempty"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text          *string    `json:"note_text,omitempty"`
	ID            string     `json:"note_uid"` // TODO: UUID? Must be a UUID
	HasAttachment bool       `json:"has_attachment"`
}

type AddEmployeeConfidentialNoteRequest struct {
	Category string `json:"note_cat"`  // Max length 3, Must conform to GetAllEmployeeNoteCategoryOptions
	Date     string `json:"note_date"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text     string `json:"note_text"` // Max length 4000
}

type UpdateEmployeeConfidentialNoteRequest struct {
	Category string `json:"note_cat"`  // Max length 3, Must conform to GetAllEmployeeNoteCategoryOptions
	Date     string `json:"note_date"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text     string `json:"note_text"` // Max length 4000
}

type EmployeePhotoChangesResponse struct {
	CompanyCode string                `json:"cmpy_code"`
	ChangeKey   string                `json:"change_key"`
	Changes     []EmployeePhotoChange `json:"changes"`
}

type EmployeePhotoChange struct {
	EmployeeCode string    `json:"emp_code"`
	UpdatedOn    time.Time `json:"photo_update_on"`
}

type EmployeeQualificationResponse struct {
	CompanyCode     string     `json:"cmpy_code"`
	EmployeeCode    string     `json:"emp_code"`
	InstitutionCode *string    `json:"inst_code,omitempty"` // GetAllQualificationInstitutionOptions
	Category        *string    `json:"qual_cat,omitempty"`  // GetAllQualificationCategoryOptions
	Text            *string    `json:"qual_text,omitempty"`
	ID              string     `json:"qual_uid"` // TODO: UUID? Must be a UUID
	ReminderFlag    bool       `json:"reminder_flg"`
	ValidDate       *time.Time `json:"valid_date,omitempty"`
	HasAttachment   bool       `json:"has_attachment"`
}

type AddEmployeeQualificationRequest struct {
	InstitutionCode string `json:"inst_code"` // Max length 4, Must conform to GetAllQualificationInstitutionOptions
	Category        string `json:"qual_cat"`  // Max length 4, Must conform to GetAllQualificationCategoryOptions
	Text            string `json:"qual_text"` // Max length 200
	ReminderFlag    bool   `json:"reminder_flg"`

	ValidDate *string `json:"valid_date,omitempty"` // Must be a date (yyyy-mm-dd)
}

type UpdateEmployeeQualificationRequest struct {
	InstitutionCode string `json:"inst_code"` // Max length 4, Must conform to GetAllQualificationInstitutionOptions
	Category        string `json:"qual_cat"`  // Max length 4, Must conform to GetAllQualificationCategoryOptions
	Text            string `json:"qual_text"` // Max length 200
	ReminderFlag    bool   `json:"reminder_flg"`

	ValidDate *string `json:"valid_date,omitempty"` // Must be a date (yyyy-mm-dd)
}

type EmployeeUDAreaResponse struct {
	CompanyCode  string                          `json:"cmpy_code"`
	AreaCode     string                          `json:"area_code"`
	EmployeeCode string                          `json:"emp_code"`
	UpdatedOn    *time.Time                      `json:"update_on,omitempty"`
	UDFields     tasscommon.UDAreaFieldsResponse `json:"ud_fields"`
}

type AddEmployeeUDAreaRequest struct {
	UDFields tasscommon.UDAreaFieldsRequest `json:"ud_fields"`
}

type UpdateEmployeeUDAreaRequest struct {
	UDFields tasscommon.UDAreaFieldsRequest `json:"ud_fields"`
}

// Personal Development

type EmployeePDActivityResponse struct {
	CompanyCode  string             `json:"cmpy_code"`
	CostAmount   *float64           `json:"cost_amt,omitempty"`
	Duration     *float64           `json:"duration,omitempty"`
	EmployeeCode string             `json:"emp_code"`
	FinishDate   *time.Time         `json:"finish_date,omitempty"`
	OtherAmount  *float64           `json:"other_amt,omitempty"`
	ProviderCode *string            `json:"pd_prov_code,omitempty"`
	Status       *string            `json:"pd_stat_code,omitempty"`
	Type         *string            `json:"pd_type_code,omitempty"`
	ID           int                `json:"pdact_num"`
	Details      *string            `json:"pdact_text,omitempty"`
	StartDate    *time.Time         `json:"start_date,omitempty"`
	UpdatedOn    *time.Time         `json:"updated_on,omitempty"`
	UDFields     PDUDFieldsResponse `json:"ud_fields"`
}

type AddEmployeePDActivityRequest struct {
	Duration   float64 `json:"duration"`
	FinishDate string  `json:"finish_date"`  // Must be a date
	Status     string  `json:"pd_stat_code"` // Max length 3
	Details    string  `json:"pdact_text"`   // Max length 200
	StartDate  string  `json:"start_date"`   // Must be a date

	CostAmount   *float64                  `json:"cost_amt,omitempty"`
	OtherAmount  *float64                  `json:"other_amt,omitempty"`
	ProviderCode *string                   `json:"pd_prov_code,omitempty"` // Max length 3, Must conform to GetAllPDProviderOptions
	Type         *string                   `json:"pd_type_code,omitempty"` // Max length 3, Must conform to GetAllPDTypesOptions
	UDFields     PDActivityUDFieldsRequest `json:"ud_fields"`
}

type UpdateEmployeePDActivityRequest struct {
	Duration   float64 `json:"duration"`
	FinishDate string  `json:"finish_date"`  // Must be a date
	Status     string  `json:"pd_stat_code"` // Max length 3
	Details    string  `json:"pdact_text"`   // Max length 200
	StartDate  string  `json:"start_date"`   // Must be a date

	CostAmount   *float64                  `json:"cost_amt,omitempty"`
	OtherAmount  *float64                  `json:"other_amt,omitempty"`
	ProviderCode *string                   `json:"pd_prov_code,omitempty"` // Max length 3, Must conform to GetAllPDProviderOptions
	Type         *string                   `json:"pd_type_code,omitempty"` // Max length 3, Must conform to GetAllPDTypesOptions
	UDFields     PDActivityUDFieldsRequest `json:"ud_fields"`
}

type PDUDFieldsResponse struct {
	Flags       PDActivityUDFlagResponse       `json:"ud_flags"`
	Codes       PDActivityUDCodesResponse      `json:"ud_codes"`
	Text        PDActivityUDTextResponse       `json:"ud_text"`
	Attachments PDActivityUDAttachmentResponse `json:"ud_attachments"`
}

type PDActivityUDFlagResponse struct {
	UD1Flag *string `json:"ud1_flg,omitempty"`
	UD2Flag *string `json:"ud2_flg,omitempty"`
	UD3Flag *string `json:"ud3_flg,omitempty"`
	UD4Flag *string `json:"ud4_flg,omitempty"`
	UD5Flag *string `json:"ud5_flg,omitempty"`
}

type PDActivityUDCodesResponse struct {
	UD6Code  *string `json:"ud6_code,omitempty"`
	UD7Code  *string `json:"ud7_code,omitempty"`
	UD8Code  *string `json:"ud8_code,omitempty"`
	UD9Code  *string `json:"ud9_code,omitempty"`
	UD10Code *string `json:"ud10_code,omitempty"`
}

type PDActivityUDTextResponse struct {
	UD11Text *string `json:"ud11_text,omitempty"`
	UD12Text *string `json:"ud12_text,omitempty"`
	UD13Text *string `json:"ud13_text,omitempty"`
	UD14Text *string `json:"ud14_text,omitempty"`
	UD15Text *string `json:"ud15_text,omitempty"`
}

type PDActivityUDAttachmentResponse struct {
	UD16Attachment tasscommon.AttachmentDetails `json:"ud16_attachment_details"`
	UD17Attachment tasscommon.AttachmentDetails `json:"ud17_attachment_details"`
	UD18Attachment tasscommon.AttachmentDetails `json:"ud18_attachment_details"`
	UD19Attachment tasscommon.AttachmentDetails `json:"ud19_attachment_details"`
	UD20Attachment tasscommon.AttachmentDetails `json:"ud20_attachment_details"`
}

type PDActivityUDFieldsRequest struct {
	Flags PDActivityUDFlagUpdateRequest `json:"ud_flags"`
	Codes PDActivityUDCodeUpdateRequest `json:"ud_codes"`
	Text  PDActivityUDTextUpdateRequest `json:"ud_text"`
}

type PDActivityUDFlagUpdateRequest struct {
	UDFlag1 *string `json:"ud1_flg,omitempty"` // Max length 1
	UDFlag2 *string `json:"ud2_flg,omitempty"` // Max length 1
	UDFlag3 *string `json:"ud3_flg,omitempty"` // Max length 1
	UDFlag4 *string `json:"ud4_flg,omitempty"` // Max length 1
	UDFlag5 *string `json:"ud5_flg,omitempty"` // Max length 1
}

type PDActivityUDCodeUpdateRequest struct {
	UDCode6  *string `json:"ud6_code,omitempty"`  // Max length 3
	UDCode7  *string `json:"ud7_code,omitempty"`  // Max length 3
	UDCode8  *string `json:"ud8_code,omitempty"`  // Max length 3
	UDCode9  *string `json:"ud9_code,omitempty"`  // Max length 3
	UDCode10 *string `json:"ud10_code,omitempty"` // Max length 3
}

type PDActivityUDTextUpdateRequest struct {
	UD11Text *string `json:"ud11_text,omitempty"` // Max length 100
	UD12Text *string `json:"ud12_text,omitempty"` // Max length 100
	UD13Text *string `json:"ud13_text,omitempty"` // Max length 100
	UD14Text *string `json:"ud14_text,omitempty"` // Max length 100
	UD15Text *string `json:"ud15_text,omitempty"` // Max length 100
}

// Payroll

type EmployeeLeaveBalanceResponse struct {
	AccrualCode         string     `json:"acr_code"`
	CompanyCode         string     `json:"cmpy_code"`
	EmployeeCode        string     `json:"emp_code"`
	CurrentEntitlement  *float64   `json:"ent_qty,omitempty"`
	EntitlementCalcDate *time.Time `json:"lst_up_date,omitempty"`
	NonAccrualDays      *int       `json:"non_acr_day_qty,omitempty"`
	NTSCode             *string    `json:"nts_code,omitempty"` // TODO: Review
	RateAmount          *float64   `json:"rate_amt,omitempty"`
	CommencementDate    *string    `json:"str_ent_date,omitempty"`
}

// Enums

type CessationType string

const (
	ContractFinishedCessation CessationType = "C"
	DeceasedCessation         CessationType = "D"
	DismissedCessation        CessationType = "F"
	IllnessCessation          CessationType = "I"
	RedundancyCessation       CessationType = "R"
	TransferCessation         CessationType = "T"
	VoluntaryCessation        CessationType = "V"
)

type IndiginousStatus string

const (
	AboriginalIS                     IndiginousStatus = "1"
	TorresStraitIslanderIS           IndiginousStatus = "2"
	AboriginalTorresStraitIslanderIS IndiginousStatus = "3"
	NeitherIS                        IndiginousStatus = "4"
	UnknownIS                        IndiginousStatus = "9"
)
