package tassstudent

import (
	"time"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

type StudentResponse struct {
	AltID                *string    `json:"alt_id,omitempty"`
	Boarder              bool       `json:"boarder"`
	Campus               *string    `json:"campus,omitempty"`
	CEIDER               *string    `json:"ceider,omitempty"`
	CompanyCode          string     `json:"cmpy_code"`
	ComparativeReporting string     `json:"compare_flg"`
	DateOfArrival        *time.Time `json:"date_arrival,omitempty"`
	DistanceEducation    bool       `json:"distance_ed"`
	DateOfBirth          *time.Time `json:"dob,omitempty"`
	DateOfEntry          *time.Time `json:"doe,omitempty"`
	DateOfLeaving        *time.Time `json:"dol,omitempty"`
	Email                *string    `json:"e_mail,omitempty"`
	EntryYear            *string    `json:"entry_lev,omitempty"`
	FFPOS                bool       `json:"ffpos"`
	FirstName            *string    `json:"first_name,omitempty"`
	FormClass            *string    `json:"form_cls,omitempty"`
	FTE                  *string    `json:"fte"`
	Gender               string     `json:"gender"`
	House                *string    `json:"house,omitempty"`
	IDM                  *string    `json:"idm_id,omitempty"`
	MobilePhone          *string    `json:"mob_phone,omitempty"`
	MultiParenting       bool       `json:"multipar_flg"`
	NextYear             string     `json:"next_yr_ind"`
	OtherName            *string    `json:"other_name,omitempty"`
	ParentCode           *string    `json:"par_code,omitempty"`
	PCTutorGroup         *string    `json:"pctut_grp,omitempty"`
	PreferredName        *string    `json:"preferred_name,omitempty"`
	PreferredSurname     *string    `json:"preferred_surname,omitempty"`
	PreviousSchool       *string    `json:"prev_school,omitempty"`
	PrivacyFlag          bool       `json:"privacy_flg"`
	Religion             *string    `json:"religion,omitempty"`
	ResidencyStatus      bool       `json:"resident_sts"`
	SMSFlag              bool       `json:"sms_flg"`
	StudentCode          string     `json:"stud_code"`
	StudentGovernmentID  *string    `json:"stud_govt_id,omitempty"`
	StudentID            *string    `json:"stud_id,omitempty"`
	Surname              string     `json:"surname"`
	UpdatedOn            *time.Time `json:"update_on,omitempty"`
	USI                  *string    `json:"usi,omitempty"`
	VisaExpiry           *time.Time `json:"visa_expiry,omitempty"`
	VisaSubclass         *string    `json:"visa_subclass,omitempty"`
	WebAccess            bool       `json:"web_access_ind"`
	YearGroup            *string    `json:"year_grp,omitempty"`
}

type UpdateStudentRequest struct {
	ComparativeReporting string    `json:"compare_flg"` // Max length 1
	DateOfEntry          time.Time `json:"doe"`
	FirstName            string    `json:"first_name"` // Max length 50
	FTE                  string    `json:"fte"`
	Gender               string    `json:"gender"`            // Max length 3
	PreferredName        string    `json:"preferred_name"`    // Max length 20
	PreferredSurname     string    `json:"preferred_surname"` // Max length 50
	Surname              string    `json:"surname"`           // Max length 30
	YearGroup            string    `json:"year_grp"`

	AltID             *string    `json:"alt_id,omitempty"` // Max length 40
	Boarder           *bool      `json:"boarder,omitempty"`
	Campus            *string    `json:"campus,omitempty"` // Max length 3
	CEIDER            *string    `json:"ceider,omitempty"` // Max length 9
	DateOfArrival     *time.Time `json:"date_arrival,omitempty"`
	DistanceEducation *bool      `json:"distance_ed,omitempty"`
	DateOfBirth       *time.Time `json:"dob,omitempty"`
	DateOfLeaving     *time.Time `json:"dol,omitempty"`
	Email             *string    `json:"e_mail,omitempty"` // Max length 60
	EntryYear         *string    `json:"entry_lev,omitempty"`
	FFPOS             *bool      `json:"ffpos"`
	FormClass         *string    `json:"form_cls,omitempty"`    // Max length 2
	House             *string    `json:"house,omitempty"`       // Max length 2
	IDM               *string    `json:"idm_id,omitempty"`      // Max length 100
	MobilePhone       *string    `json:"mob_phone,omitempty"`   // Max length 30
	NextYear          *string    `json:"next_yr_ind"`           // Max length 1
	OtherName         *string    `json:"other_name,omitempty"`  // Max length 50
	PCTutorGroup      *string    `json:"pctut_grp,omitempty"`   // Max length 5
	PreviousSchool    *string    `json:"prev_school,omitempty"` // Max length 5
	PrivacyFlag       *bool      `json:"privacy_flg"`
	Religion          *string    `json:"religion,omitempty"` // Max length 2
	ResidencyStatus   *bool      `json:"resident_sts"`       // Max length 3
	SMSFlag           *bool      `json:"sms_flg"`
	StudentCode       *string    `json:"stud_code"`     // Max length 8
	USI               *string    `json:"usi,omitempty"` // Max length 10
	VisaExpiry        *time.Time `json:"visa_expiry,omitempty"`
	VisaSubclass      *string    `json:"visa_subclass,omitempty"` // Max length 6
	WebAccess         *bool      `json:"web_access_ind"`
}

type StudentStandardNoteResponse struct {
	CompanyCode   string    `json:"cmpy_code"`
	StudentCode   string    `json:"stud_code"`
	NoteCategory  *string   `json:"note_cat,omitempty"`
	NoteDate      time.Time `json:"note_date"`
	NoteText      *string   `json:"note_text,omitempty"`
	NoteID        string    `json:"note_uid"`
	HasAttachment bool      `json:"has_attachment"`
}

type AddStudentStandardNoteRequest struct {
	NoteCategory string `json:"note_cat"`  // Max length 3
	NoteDate     string `json:"note_date"` // Must be a date (yyyy-mm-dd)
	NoteText     string `json:"note_text"` // Max length 4000
}

type UpdateStudentStandardNoteRequest struct {
	NoteCategory string `json:"note_cat"`  // Max length 3
	NoteDate     string `json:"note_date"` // Must be a date (yyyy-mm-dd)
	NoteText     string `json:"note_text"` // Max length 4000
}

type StudentConfidentialNoteResponse struct {
	CompanyCode   string  `json:"cmpy_code"`
	StudentCode   string  `json:"stud_code"`
	NoteCategory  *string `json:"note_cat,omitempty"`
	NoteDate      string  `json:"note_date"`
	NoteText      *string `json:"note_text,omitempty"`
	NoteID        string  `json:"note_uid"`
	HasAttachment bool    `json:"has_attachment"`
}

type AddStudentConfidentialNoteRequest struct {
	NoteCategory string `json:"note_cat"`  // Max length 3
	NoteDate     string `json:"note_date"` // Must be a date (yyyy-mm-dd)
	NoteText     string `json:"note_text"` // Max length 4000
}

type UpdateStudentConfidentialNoteRequest struct {
	NoteCategory string `json:"note_cat"`  // Max length 3
	NoteDate     string `json:"note_date"` // Must be a date (yyyy-mm-dd)
	NoteText     string `json:"note_text"` // Max length 4000
}

type StudentPhotoChangesResponse struct {
	CompanyCode string               `json:"cmpy_code"`
	ChangeKey   string               `json:"change_key"`
	Changes     []StudentPhotoChange `json:"changes"`
}

type StudentPhotoChange struct {
	StudentCode string    `json:"stud_code"`
	UpdatedOn   time.Time `json:"photo_update_on"`
}

type UDAreaOptionsResponse struct {
	CompanyCode     string       `json:"cmpy_code"`
	AreaCode        string       `json:"area_code"`
	AreaDescription string       `json:"area_desc"`
	UDFields        UDFieldTypes `json:"ud_fields"`
}

type UDFieldTypes struct {
	UDFlags       []UDFieldDetails           `json:"ud_flags"`
	UDCodes       []UDCodeFieldDetails       `json:"ud_codes"`
	UDText        []UDFieldDetails           `json:"ud_text"`
	UDDates       []UDFieldDetails           `json:"ud_dates"`
	UDAttachments []UDAttachmentFieldDetails `json:"ud_attachments"`
}

type UDFieldDetails struct {
	Name        string  `json:"field_name"`
	Description *string `json:"field_desc,omitempty"`
	SortOrder   string  `json:"sort_order"` // Must conform to SortOrderValidation
}

type UDCodeFieldDetails struct {
	ReferenceValues []UDFieldReferenceValue `json:"reference_values"`

	UDFieldDetails
}

type UDFieldReferenceValue struct {
	Code        *string `json:"ud_code,omitempty"`
	Description string  `json:"ud_desc"`
	SortOrder   string  `json:"sort_order"` // Must conform to SortOrderValidation
}

type UDAttachmentFieldDetails struct {
	FieldNumber string `json:"field_number"` // Must conform to FieldNumberValidation

	UDFieldDetails
}

type StudentUDAreaResponse struct {
	CompanyCode string                      `json:"cmpy_code"`
	AreaCode    string                      `json:"area_code"`
	StudentCode string                      `json:"stud_code"`
	UpdatedOn   *time.Time                  `json:"update_on,omitempty"`
	UDFields    StudentUDAreaFieldsResponse `json:"ud_fields"`
}

type StudentUDAreaFieldsResponse struct {
	UDFlags       StudentUDAreaFlagResponse       `json:"ud_flags"`
	UDCodes       StudentUDAreaCodeResponse       `json:"ud_codes"`
	UDText        StudentUDAreaTextResponse       `json:"ud_text"`
	UDDates       StudentUDAreaDateResponse       `json:"ud_dates"`
	UDAttachments StudentUDAreaAttachmentResponse `json:"ud_attachments"`
}

type StudentUDAreaFlagResponse struct {
	UD1Flag  *string `json:"ud1_flg,omitempty"`  // Max length 1
	UD2Flag  *string `json:"ud2_flg,omitempty"`  // Max length 1
	UD3Flag  *string `json:"ud3_flg,omitempty"`  // Max length 1
	UD4Flag  *string `json:"ud4_flg,omitempty"`  // Max length 1
	UD5Flag  *string `json:"ud5_flg,omitempty"`  // Max length 1
	UD6Flag  *string `json:"ud6_flg,omitempty"`  // Max length 1
	UD7Flag  *string `json:"ud7_flg,omitempty"`  // Max length 1
	UD8Flag  *string `json:"ud8_flg,omitempty"`  // Max length 1
	UD9Flag  *string `json:"ud9_flg,omitempty"`  // Max length 1
	UD10Flag *string `json:"ud10_flg,omitempty"` // Max length 1
}

type StudentUDAreaCodeResponse struct {
	UD11Code *string `json:"ud11_code,omitempty"` // Max length 3
	UD12Code *string `json:"ud12_code,omitempty"` // Max length 3
	UD13Code *string `json:"ud13_code,omitempty"` // Max length 3
	UD14Code *string `json:"ud14_code,omitempty"` // Max length 3
	UD15Code *string `json:"ud15_code,omitempty"` // Max length 3
	UD16Code *string `json:"ud16_code,omitempty"` // Max length 3
	UD17Code *string `json:"ud17_code,omitempty"` // Max length 3
	UD18Code *string `json:"ud18_code,omitempty"` // Max length 3
	UD19Code *string `json:"ud19_code,omitempty"` // Max length 3
	UD20Code *string `json:"ud20_code,omitempty"` // Max length 3
}

type StudentUDAreaTextResponse struct {
	UD21Text *string `json:"ud21_text,omitempty"` // Max length 100
	UD22Text *string `json:"ud22_text,omitempty"` // Max length 100
	UD23Text *string `json:"ud23_text,omitempty"` // Max length 100
	UD24Text *string `json:"ud24_text,omitempty"` // Max length 100
	UD25Text *string `json:"ud25_text,omitempty"` // Max length 100
	UD26Text *string `json:"ud26_text,omitempty"` // Max length 100
	UD27Text *string `json:"ud27_text,omitempty"` // Max length 100
	UD28Text *string `json:"ud28_text,omitempty"` // Max length 100
	UD29Text *string `json:"ud29_text,omitempty"` // Max length 100
	UD30Text *string `json:"ud30_text,omitempty"` // Max length 100
}

type StudentUDAreaDateResponse struct {
	UD31Date *time.Time `json:"ud31_date,omitempty"`
	UD32Date *time.Time `json:"ud32_date,omitempty"`
	UD40Date *time.Time `json:"ud33_date,omitempty"`
	UD33Date *time.Time `json:"ud34_date,omitempty"`
	UD34Date *time.Time `json:"ud35_date,omitempty"`
	UD35Date *time.Time `json:"ud36_date,omitempty"`
	UD36Date *time.Time `json:"ud37_date,omitempty"`
	UD37Date *time.Time `json:"ud38_date,omitempty"`
	UD38Date *time.Time `json:"ud39_date,omitempty"`
	UD39Date *time.Time `json:"ud40_date,omitempty"`
}

type StudentUDAreaAttachmentResponse struct {
	UD41AttachmentDetails *tasscommon.AttachmentDetails `json:"ud41_attachment_details,omitempty"`
	UD42AttachmentDetails *tasscommon.AttachmentDetails `json:"ud42_attachment_details,omitempty"`
	UD43AttachmentDetails *tasscommon.AttachmentDetails `json:"ud43_attachment_details,omitempty"`
	UD44AttachmentDetails *tasscommon.AttachmentDetails `json:"ud44_attachment_details,omitempty"`
	UD45AttachmentDetails *tasscommon.AttachmentDetails `json:"ud45_attachment_details,omitempty"`
	UD46AttachmentDetails *tasscommon.AttachmentDetails `json:"ud46_attachment_details,omitempty"`
	UD47AttachmentDetails *tasscommon.AttachmentDetails `json:"ud47_attachment_details,omitempty"`
	UD48AttachmentDetails *tasscommon.AttachmentDetails `json:"ud48_attachment_details,omitempty"`
	UD49AttachmentDetails *tasscommon.AttachmentDetails `json:"ud49_attachment_details,omitempty"`
	UD50AttachmentDetails *tasscommon.AttachmentDetails `json:"ud50_attachment_details,omitempty"`
}

type AddStudentUDAreaRequest struct {
	UDFields StudentUDAreaFieldsRequest `json:"ud_fields"`
}

type StudentUDAreaFieldsRequest struct {
	UDFlags StudentUDAreaFlagRequest `json:"ud_flags"`
	UDCodes StudentUDAreaCodeRequest `json:"ud_codes"`
	UDText  StudentUDAreaTextRequest `json:"ud_text"`
	UDDates StudentUDAreaDateRequest `json:"ud_dates"`
}

type StudentUDAreaFlagRequest struct {
	UD1Flag  *string `json:"ud1_flg,omitempty"`  // Max length 1
	UD2Flag  *string `json:"ud2_flg,omitempty"`  // Max length 1
	UD3Flag  *string `json:"ud3_flg,omitempty"`  // Max length 1
	UD4Flag  *string `json:"ud4_flg,omitempty"`  // Max length 1
	UD5Flag  *string `json:"ud5_flg,omitempty"`  // Max length 1
	UD6Flag  *string `json:"ud6_flg,omitempty"`  // Max length 1
	UD7Flag  *string `json:"ud7_flg,omitempty"`  // Max length 1
	UD8Flag  *string `json:"ud8_flg,omitempty"`  // Max length 1
	UD9Flag  *string `json:"ud9_flg,omitempty"`  // Max length 1
	UD10Flag *string `json:"ud10_flg,omitempty"` // Max length 1
}

type StudentUDAreaCodeRequest struct {
	UD11Code *string `json:"ud11_code,omitempty"` // Max length 3
	UD12Code *string `json:"ud12_code,omitempty"` // Max length 3
	UD13Code *string `json:"ud13_code,omitempty"` // Max length 3
	UD14Code *string `json:"ud14_code,omitempty"` // Max length 3
	UD15Code *string `json:"ud15_code,omitempty"` // Max length 3
	UD16Code *string `json:"ud16_code,omitempty"` // Max length 3
	UD17Code *string `json:"ud17_code,omitempty"` // Max length 3
	UD18Code *string `json:"ud18_code,omitempty"` // Max length 3
	UD19Code *string `json:"ud19_code,omitempty"` // Max length 3
	UD20Code *string `json:"ud20_code,omitempty"` // Max length 3
}

type StudentUDAreaTextRequest struct {
	UD21Text *string `json:"ud21_text,omitempty"` // Max length 100
	UD22Text *string `json:"ud22_text,omitempty"` // Max length 100
	UD23Text *string `json:"ud23_text,omitempty"` // Max length 100
	UD24Text *string `json:"ud24_text,omitempty"` // Max length 100
	UD25Text *string `json:"ud25_text,omitempty"` // Max length 100
	UD26Text *string `json:"ud26_text,omitempty"` // Max length 100
	UD27Text *string `json:"ud27_text,omitempty"` // Max length 100
	UD28Text *string `json:"ud28_text,omitempty"` // Max length 100
	UD29Text *string `json:"ud29_text,omitempty"` // Max length 100
	UD30Text *string `json:"ud30_text,omitempty"` // Max length 100
}

type StudentUDAreaDateRequest struct {
	UD31Date *time.Time `json:"ud31_date,omitempty"`
	UD32Date *time.Time `json:"ud32_date,omitempty"`
	UD40Date *time.Time `json:"ud33_date,omitempty"`
	UD33Date *time.Time `json:"ud34_date,omitempty"`
	UD34Date *time.Time `json:"ud35_date,omitempty"`
	UD35Date *time.Time `json:"ud36_date,omitempty"`
	UD36Date *time.Time `json:"ud37_date,omitempty"`
	UD37Date *time.Time `json:"ud38_date,omitempty"`
	UD38Date *time.Time `json:"ud39_date,omitempty"`
	UD39Date *time.Time `json:"ud40_date,omitempty"`
}

type UpdateStudentUDAreaRequest struct {
	UDFields StudentUDAreaFieldsRequest `json:"ud_fields"`
}

type StudentUDFieldsResponse struct {
	CompanyCode string           `json:"cmpy_code"`
	StudentCode string           `json:"stud_code"`
	UDFields    UDFieldsResponse `json:"ud_fields"`
}

type UDFieldsResponse struct {
	UDFlags UDFlagResponse `json:"ud_flags"`
	UDCodes UDCodeResponse `json:"ud_codes"`
	UDText  UDTextResponse `json:"ud_text"`
}

type UDFlagResponse struct {
	UD1Flag  *string `json:"ud1_flg,omitempty"`  // Max length 1
	UD2Flag  *string `json:"ud2_flg,omitempty"`  // Max length 1
	UD3Flag  *string `json:"ud3_flg,omitempty"`  // Max length 1
	UD4Flag  *string `json:"ud4_flg,omitempty"`  // Max length 1
	UD5Flag  *string `json:"ud5_flg,omitempty"`  // Max length 1
	UD6Flag  *string `json:"ud6_flg,omitempty"`  // Max length 1
	UD7Flag  *string `json:"ud7_flg,omitempty"`  // Max length 1
	UD8Flag  *string `json:"ud8_flg,omitempty"`  // Max length 1
	UD9Flag  *string `json:"ud9_flg,omitempty"`  // Max length 1
	UD10Flag *string `json:"ud10_flg,omitempty"` // Max length 1
}

type UDCodeResponse struct {
	UD11Code *string `json:"ud11_code,omitempty"` // Max length 3
	UD12Code *string `json:"ud12_code,omitempty"` // Max length 3
	UD13Code *string `json:"ud13_code,omitempty"` // Max length 3
	UD14Code *string `json:"ud14_code,omitempty"` // Max length 3
	UD15Code *string `json:"ud15_code,omitempty"` // Max length 3
	UD16Code *string `json:"ud16_code,omitempty"` // Max length 3
	UD17Code *string `json:"ud17_code,omitempty"` // Max length 3
	UD18Code *string `json:"ud18_code,omitempty"` // Max length 3
	UD19Code *string `json:"ud19_code,omitempty"` // Max length 3
	UD20Code *string `json:"ud20_code,omitempty"` // Max length 3
}

type UDTextResponse struct {
	UD21Text *string `json:"ud21_text,omitempty"` // Max length 100
	UD22Text *string `json:"ud22_text,omitempty"` // Max length 100
	UD23Text *string `json:"ud23_text,omitempty"` // Max length 100
	UD24Text *string `json:"ud24_text,omitempty"` // Max length 100
	UD25Text *string `json:"ud25_text,omitempty"` // Max length 100
	UD26Text *string `json:"ud26_text,omitempty"` // Max length 100
	UD27Text *string `json:"ud27_text,omitempty"` // Max length 100
	UD28Text *string `json:"ud28_text,omitempty"` // Max length 100
	UD29Text *string `json:"ud29_text,omitempty"` // Max length 100
	UD30Text *string `json:"ud30_text,omitempty"` // Max length 100
}

type UpdateStudentUDFieldsRequest struct {
	UDFields UDFieldsRequest `json:"ud_fields"`
}

type UDFieldsRequest struct {
	UDFlags UDFlagRequest `json:"ud_flags"`
	UDCodes UDCodeRequest `json:"ud_codes"`
	UDText  UDTextRequest `json:"ud_text"`
}

type UDFlagRequest struct {
	UD1Flag  *string `json:"ud1_flg,omitempty"`  // Max length 1
	UD2Flag  *string `json:"ud2_flg,omitempty"`  // Max length 1
	UD3Flag  *string `json:"ud3_flg,omitempty"`  // Max length 1
	UD4Flag  *string `json:"ud4_flg,omitempty"`  // Max length 1
	UD5Flag  *string `json:"ud5_flg,omitempty"`  // Max length 1
	UD6Flag  *string `json:"ud6_flg,omitempty"`  // Max length 1
	UD7Flag  *string `json:"ud7_flg,omitempty"`  // Max length 1
	UD8Flag  *string `json:"ud8_flg,omitempty"`  // Max length 1
	UD9Flag  *string `json:"ud9_flg,omitempty"`  // Max length 1
	UD10Flag *string `json:"ud10_flg,omitempty"` // Max length 1
}

type UDCodeRequest struct {
	UD11Code *string `json:"ud11_code,omitempty"` // Max length 3
	UD12Code *string `json:"ud12_code,omitempty"` // Max length 3
	UD13Code *string `json:"ud13_code,omitempty"` // Max length 3
	UD14Code *string `json:"ud14_code,omitempty"` // Max length 3
	UD15Code *string `json:"ud15_code,omitempty"` // Max length 3
	UD16Code *string `json:"ud16_code,omitempty"` // Max length 3
	UD17Code *string `json:"ud17_code,omitempty"` // Max length 3
	UD18Code *string `json:"ud18_code,omitempty"` // Max length 3
	UD19Code *string `json:"ud19_code,omitempty"` // Max length 3
	UD20Code *string `json:"ud20_code,omitempty"` // Max length 3
}

type UDTextRequest struct {
	UD21Text *string `json:"ud21_text,omitempty"` // Max length 100
	UD22Text *string `json:"ud22_text,omitempty"` // Max length 100
	UD23Text *string `json:"ud23_text,omitempty"` // Max length 100
	UD24Text *string `json:"ud24_text,omitempty"` // Max length 100
	UD25Text *string `json:"ud25_text,omitempty"` // Max length 100
	UD26Text *string `json:"ud26_text,omitempty"` // Max length 100
	UD27Text *string `json:"ud27_text,omitempty"` // Max length 100
	UD28Text *string `json:"ud28_text,omitempty"` // Max length 100
	UD29Text *string `json:"ud29_text,omitempty"` // Max length 100
	UD30Text *string `json:"ud30_text,omitempty"` // Max length 100
}

type StudentUDFieldOptionResponse struct {
	CompanyCode string `json:"cmpy_code"`
}

type StudentMCEECDYAResponse struct {
	CompanyCode            string     `json:"cmpy_code"`
	StudentCode            string     `json:"stud_code"`
	ArrivalYear            *string    `json:"arrive_yr,omitempty"` // Must conform to ArrivalYearValidation
	Parent1LOTE            *string    `json:"mlote_code,omitempty"`
	Parent1NonSchoolLevel  *string    `json:"mnse_code,omitempty"`
	Parent1OccupationGroup *string    `json:"mocc_code,omitempty"`
	Parent1SchoolLevel     *string    `json:"mse_code,omitempty"`
	Parent2LOTE            *string    `json:"flote_code,omitempty"`
	Parent2NonSchoolLevel  *string    `json:"fnse_code,omitempty"`
	Parent2OccupationGroup *string    `json:"focc_code,omitempty"`
	Parent2SchoolLevel     *string    `json:"fse_code,omitempty"`
	IndiginousStatus       *string    `json:"s_indig_sts,omitempty"`
	CountryOfBirth         *string    `json:"scob_code,omitempty"`
	LOTE                   *string    `json:"slote_code,omitempty"`
	UpdatedOn              *time.Time `json:"update_on,omitempty"`
}

type UpdateStudentMCEECDYARequest struct {
	ArrivalYear            *string `json:"arrive_yr,omitempty"` // Must conform to ArrivalYearValidation
	Parent1LOTE            *string `json:"mlote_code,omitempty"`
	Parent1NonSchoolLevel  *string `json:"mnse_code,omitempty"`
	Parent1OccupationGroup *string `json:"mocc_code,omitempty"`
	Parent1SchoolLevel     *string `json:"mse_code,omitempty"`
	Parent2LOTE            *string `json:"flote_code,omitempty"`
	Parent2NonSchoolLevel  *string `json:"fnse_code,omitempty"`
	Parent2OccupationGroup *string `json:"focc_code,omitempty"`
	Parent2SchoolLevel     *string `json:"fse_code,omitempty"`
	IndiginousStatus       *string `json:"s_indig_sts,omitempty"`
	CountryOfBirth         *string `json:"scob_code,omitempty"`
	LOTE                   *string `json:"slote_code,omitempty"`
}

// MEDICAL

type AsthmaManagementResponse struct {
	CompanyCode    string                         `json:"cmpy_code"`
	StudentCode    string                         `json:"stud_code"`
	UsualSigns     AsthmaManagementUsualSigns     `json:"usual_signs"`
	WorseningSigns AsthmaManagementWorseningSigns `json:"worsening_signs"`
	Triggers       AsthmaManagementTriggers       `json:"triggers"`
}

type AsthmaManagementUsualSigns struct {
	Wheeze bool    `json:"wheez_flg"`
	Tight  bool    `json:"tight_flg"`
	Cough  bool    `json:"cough_flg"`
	Breath bool    `json:"breath_flg"`
	Speak  bool    `json:"speak_flg"`
	Other  *string `json:"comm_text,omitempty"`
}

type AsthmaManagementWorseningSigns struct {
	Wheeze bool    `json:"wheez_flg"`
	Tight  bool    `json:"tight_flg"`
	Cough  bool    `json:"cough_flg"`
	Breath bool    `json:"breath_flg"`
	Speak  bool    `json:"speak_flg"`
	Other  *string `json:"comm_text,omitempty"`
}

type AsthmaManagementTriggers struct {
	Exercise  bool    `json:"exercise_flg"`
	ColdVirus bool    `json:"cold_virus_flg"`
	Pollen    bool    `json:"pollen_flg"`
	Dust      bool    `json:"dust_flg"`
	Food      bool    `json:"food_flg"`
	Foods     *string `json:"food_text,omitempty"`
	Other     *string `json:"comm_text,omitempty"`
}

type UpdateAsthmaManagementRequest struct {
	// TODO: Can these sub-structs be pointers? They feel like at least 1 is required but not all
	UsualSigns     AsthmaManagementUsualSigns     `json:"usual_signs"`
	WorseningSigns AsthmaManagementWorseningSigns `json:"worsening_sign"`
	Triggers       AsthmaManagementTriggers       `json:"triggers"`
}

type StudentMedicalConditionResponse struct {
	CompanyCode          string     `json:"cmpy_code"`
	LastOccurance        *time.Time `json:"last_occ_date,omitempty"`
	MedicalConditionCode string     `json:"mcond_code"`
	Severe               bool       `json:"severe_ind"`
	StudentCode          string     `json:"stud_code"`
	TreatmentDetails     *string    `json:"treat_text,omitempty"`
	Active               bool       `json:"active_flg"`
	UDFields             UDFields   `json:"ud_fields"`
	HasAttachment        bool       `json:"has_attachment"`
	HasNote              bool       `json:"has_note"`
}

type UDFields struct {
	UDText UDText `json:"ud_text"`
}

type UDText struct {
	UD1Text  *string `json:"ud1_text,omitempty"`  // Max length 100
	UD2Text  *string `json:"ud2_text,omitempty"`  // Max length 100
	UD3Text  *string `json:"ud3_text,omitempty"`  // Max length 100
	UD4Text  *string `json:"ud4_text,omitempty"`  // Max length 100
	UD5Text  *string `json:"ud5_text,omitempty"`  // Max length 100
	UD6Text  *string `json:"ud6_text,omitempty"`  // Max length 100
	UD7Text  *string `json:"ud7_text,omitempty"`  // Max length 100
	UD8Text  *string `json:"ud8_text,omitempty"`  // Max length 100
	UD9Text  *string `json:"ud9_text,omitempty"`  // Max length 100
	UD10Text *string `json:"ud10_text,omitempty"` // Max length 100
}

type AddStudentMedicalConditionRequest struct {
	LastOccurance        *time.Time `json:"last_occ_date,omitempty"`
	MedicalConditionCode string     `json:"mcond_code"` // Max length 3
	Severe               bool       `json:"severe_ind"`
	TreatmentDetails     *string    `json:"treat_text,omitempty"` // Max length 4000
	Active               bool       `json:"active_flg"`
	UDFields             UDFields   `json:"ud_fields"`
}

type UpdateStudentMedicalConditionRequest struct {
	LastOccurance    *time.Time `json:"last_occ_date,omitempty"`
	Severe           bool       `json:"severe_ind"`
	TreatmentDetails *string    `json:"treat_text,omitempty"` // Max length 4000
	Active           bool       `json:"active_flg"`
	UDFields         UDFields   `json:"ud_fields"`
}

type StudentMedicalConditionNoteResponse struct {
	CompanyCode          string    `json:"cmpy_code"`
	StudentCode          string    `json:"stud_code"`
	MedicalConditionCode string    `json:"mcond_code"`
	Date                 time.Time `json:"note_date"`
	Text                 *string   `json:"note_text,omitempty"`
	ID                   string    `json:"note_uid"` // Must be a UUID
}

type AddStudentMedicalConditionNoteRequest struct {
	Date string  `json:"note_date"`           // Must be a date
	Text *string `json:"note_text,omitempty"` // Max length 4000
}

type UpdateStudentMedicalConditionNoteRequest struct {
	Date string  `json:"note_date"`           // Must be a date
	Text *string `json:"note_text,omitempty"` // Max length 4000
}

type StudentIllnessResponse struct {
	CompanyCode          string     `json:"cmpy_code"`
	StudentCode          string     `json:"stud_code"`
	IllnessUID           string     `json:"illness_uid"`
	IllnessDate          time.Time  `json:"ill_date"`
	IllnessTime          *time.Time `json:"ill_time,omitempty"`
	MedicalConditionCode *string    `json:"mcond_code,omitempty"`
	TreatmentCode        *string    `json:"treat_code,omitempty"`
	DischargeDate        *time.Time `json:"disch_date,omitempty"`
	DischargeTime        *time.Time `json:"disch_time,omitempty"`
	Hospitalised         bool       `json:"host_flg"`
	IllnessDescription   *string    `json:"ill_desc,omitempty"`
	IllnessNotes         *string    `json:"ill_note,omitempty"`
	HasMedications       bool       `json:"has_medications"`
}

type AddStudentIllnessRequest struct {
	IllnessDate          string `json:"ill_date"`   // Must be a date
	IllnessTime          string `json:"ill_time"`   // Must be a time
	MedicalConditionCode string `json:"mcond_code"` // Max length 3

	TreatmentCode      *string `json:"treat_code,omitempty"` // Max length 3
	DischargeDate      *string `json:"disch_date,omitempty"` // Must be a date
	DischargeTime      *string `json:"disch_time,omitempty"` // Must be a time
	Hospitalised       bool    `json:"host_flg"`
	IllnessDescription *string `json:"ill_desc,omitempty"` // Max length 60
	IllnessNotes       *string `json:"ill_note,omitempty"` // Max length 4000
}

type UpdateStudentIllnessRequest struct {
	IllnessDate          string  `json:"ill_date"`             // Must be a date
	IllnessTime          string  `json:"ill_time"`             // Must be a time
	MedicalConditionCode string  `json:"mcond_code"`           // Max length 3
	TreatmentCode        string  `json:"treat_code,omitempty"` // Max length 3
	DischargeDate        string  `json:"disch_date,omitempty"` // Must be a date
	DischargeTime        string  `json:"disch_time,omitempty"` // Must be a time
	Hospitalised         bool    `json:"host_flg"`
	IllnessDescription   *string `json:"ill_desc,omitempty"` // Max length 60
	IllnessNotes         *string `json:"ill_note,omitempty"` // Max length 4000
}

type StudentImmunisationResponse struct {
	CompanyCode      string  `json:"cmpy_code"`
	ImmunisationCode string  `json:"imm_code"`
	ImmunisationYear *string `json:"imm_year,omitempty"` // Must conform to ImmunisationYearValidation
	StudentCode      string  `json:"stud_code"`
}

type AddStudentImmunisationRequest struct {
	ImmunisationCode string  `json:"imm_code"`           // Max length 2
	ImmunisationYear *string `json:"imm_year,omitempty"` // Must conform to ImmunisationYearValidation
}

type UpdateStudentImmunisationRequest struct {
	ImmunisationYear *string `json:"imm_year,omitempty"` // Must conform to ImmunisationYearValidation
}

type StudentImmunisationRegisterResponse struct {
	CompanyCode   string     `json:"cmpy_code"`
	AIRStateDate  *time.Time `json:"air_state_date,omitempty"` // Must be a date
	NextDueDate   *time.Time `json:"next_due_date,omitempty"`  // Must be a date
	StudentCode   string     `json:"stud_code"`
	StatusCode    *string    `json:"status_code,omitempty"`
	HasAttachment bool       `json:"has_attachment"`
}

type UpdateStudentImmunisationRegisterRequest struct {
	NextDueDate *string `json:"next_due_date,omitempty"` // Must be a date
	StatusCode  *string `json:"status_code,omitempty"`
}

type MedicationAdminister string

const (
	SelfAdminister  MedicationAdminister = "S"
	NeedsAssistance MedicationAdminister = "A"
)

type StudentMedicationResponse struct {
	Active                bool                  `json:"active_flg"`
	Administer            *MedicationAdminister `json:"administer,omitempty"` // Uses MedicationAdminster enum
	CompanyCode           string                `json:"cmpy_code"`
	DoctorPhone           *string               `json:"doc_phone,omitempty"`
	EndDate               *time.Time            `json:"end_date,omitempty"`
	ExpiryDate            *time.Time            `json:"expiry_date,omitempty"`
	MedicalConditionCode  string                `json:"mcond_code"`
	FurtherDetails        *string               `json:"med_detl,omitempty"`
	MethodOfUse           *string               `json:"med_meth,omitempty"`
	Name                  *string               `json:"med_text,omitempty"`
	MedicationUID         string                `json:"medication_uid"`                   // Must be a UUID
	MinTimeBetweenDoses   *string               `json:"min_time_between_doses,omitempty"` // Must conform to MedicationMinTimeBetweenDosesValidation
	PrescribingDoctor     *string               `json:"script_doc,omitempty"`
	StartDate             *time.Time            `json:"start_date,omitempty"`
	StudentCode           string                `json:"stud_code"`
	StaffTrainingRequired bool                  `json:"training"`
	HasAttachment         bool                  `json:"has_attachment"`
	HasNote               bool                  `json:"has_note"`
	HasSchedule           bool                  `json:"has_schedule"`
}

type AddStudentMedicationRequest struct {
	Active                bool                 `json:"active_flg"`
	Administer            MedicationAdminister `json:"administer"`
	DoctorPhone           *string              `json:"doc_phone,omitempty"`              // Max length 25
	EndDate               *string              `json:"end_date,omitempty"`               // Must be a date
	ExpiryDate            *string              `json:"expiry_date,omitempty"`            // Must be a date
	FurtherDetails        *string              `json:"med_detl,omitempty"`               // Max length 200
	MethodOfUse           *string              `json:"med_meth,omitempty"`               // Max length 200
	Name                  *string              `json:"med_text,omitempty"`               // Max length 200
	MinTimeBetweenDoses   *string              `json:"min_time_between_doses,omitempty"` // Must conform to MedicationMinTimeBetweenDosesValidation
	PrescribingDoctor     *string              `json:"script_doc,omitempty"`             // Max length 30
	StartDate             *string              `json:"start_date,omitempty"`             // Must be a date
	StaffTrainingRequired bool                 `json:"training"`
}

type UpdateStudentMedicationRequest struct {
	Active            bool                 `json:"active_flg"`
	Administer        MedicationAdminister `json:"administer"`
	DoctorPhone       *string              `json:"doc_phone,omitempty"`   // Max length 25
	EndDate           *string              `json:"end_date,omitempty"`    // Must be a date
	ExpiryDate        *string              `json:"expiry_date,omitempty"` // Must be a date
	FurtherDetails    *string              `json:"med_detl,omitempty"`    // Max length 200
	MethodOfUse       *string              `json:"med_meth,omitempty"`    // Max length 200
	Name              *string              `json:"med_text,omitempty"`    // Max length 200
	MedicationUID     string               `json:"medication_uid"`        // Must be a UUID
	PrescribingDoctor *string              `json:"script_doc,omitempty"`  // Max length 30
	StartDate         *string              `json:"start_date,omitempty"`  // Must be a date
	StudentCode       string               `json:"stud_code"`
	Training          bool                 `json:"training"`
}

type StudentMedicationNoteResponse struct {
	CompanyCode          string    `json:"cmpy_code"`
	StudentCode          string    `json:"stud_code"`
	MedicalConditionCode string    `json:"mcond_code"`
	MedicationUID        string    `json:"medication_uid"` // Must be a UUID
	Date                 time.Time `json:"note_date"`
	Text                 *string   `json:"note_text,omitempty"`
	ID                   string    `json:"note_uid"` // Must be a UUID
}

type AddStudentMedicationNoteRequest struct {
	Date string  `json:"note_date"`           // Must be a date
	Text *string `json:"note_text,omitempty"` // Max length 4000
}

type UpdateStudentMedicationNoteRequest struct {
	Date string  `json:"note_date"`           // Must be a date
	Text *string `json:"note_text,omitempty"` // Max length 4000
}

type StudentMedicationScheduleResponse struct {
	CompanyCode          string       `json:"cmpy_code"`
	StudentCode          string       `json:"stud_code"`
	MedicalConditionCode string       `json:"mcond_code"`
	MedicationUID        string       `json:"medication_uid"` // Must be a UUID
	ID                   string       `json:"sched_uid"`      // Must be a UUID
	Dose                 *string      `json:"med_dose,omitempty"`
	DoseTime             string       `json:"med_time"`                  // Must be a time
	StartDate            *string      `json:"shed_start_date,omitempty"` // Must be a date
	EndDate              *string      `json:"shed_end_date,omitempty"`   // Must be a date
	Days                 DaysResponse `json:"days"`
}

type DaysResponse struct {
	Monday    bool `json:"mon_flg"`
	Tuesday   bool `json:"tue_flg"`
	Wednesday bool `json:"wed_flg"`
	Thursday  bool `json:"thu_flg"`
	Friday    bool `json:"fri_flg"`
	Saturday  bool `json:"sat_flg"`
	Sunday    bool `json:"sun_flg"`
}

type AddStudentMedicationScheduleRequest struct {
	Dose      string `json:"med_dose"`
	Time      string `json:"med_time"`        // Must be a time
	StartDate string `json:"shed_start_date"` // Must be a date

	EndDate *string      `json:"shed_end_date,omitempty"` // Must be a date
	Days    *DaysRequest `json:"days,omitempty"`
}

type DaysRequest struct {
	Monday    bool `json:"mon_flg"`
	Tuesday   bool `json:"tue_flg"`
	Wednesday bool `json:"wed_flg"`
	Thursday  bool `json:"thu_flg"`
	Friday    bool `json:"fri_flg"`
	Saturday  bool `json:"sat_flg"`
	Sunday    bool `json:"sun_flg"`
}

type UpdateStudentMedicationScheduleRequest struct {
	Dose      string `json:"med_dose"`
	Time      string `json:"med_time"`        // Must be a time
	StartDate string `json:"shed_start_date"` // Must be a date

	EndDate *string      `json:"shed_end_date,omitempty"` // Must be a date
	Days    *DaysRequest `json:"days,omitempty"`
}

type StudentMedicalStandardNoteResponse struct {
	CompanyCode   string    `json:"cmpy_code"`
	StudentCode   string    `json:"stud_code"`
	Category      *string   `json:"note_cat,omitempty"`
	Date          time.Time `json:"note_date"`
	Text          *string   `json:"note_text,omitempty"`
	ID            string    `json:"note_uid"` // Must be a UUID
	HasAttachment bool      `json:"has_attachment"`
}

type AddStudentMedicalStandardNoteRequest struct {
	Category string `json:"note_cat"`  // Max length 3
	Date     string `json:"note_date"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text     string `json:"note_text"` // Max length 4000
}

type UpdateStudentMedicalStandardNoteRequest struct {
	Category string `json:"note_cat"`  // Max length 3
	Date     string `json:"note_date"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text     string `json:"note_text"` // Max length 4000
}

type StudentMedicalConfidentialNoteResponse struct {
	CompanyCode   string    `json:"cmpy_code"`
	StudentCode   string    `json:"stud_code"`
	Category      *string   `json:"note_cat,omitempty"`
	Date          time.Time `json:"note_date"`
	Text          *string   `json:"note_text,omitempty"`
	ID            string    `json:"note_uid"` // Must be a UUID
	HasAttachment bool      `json:"has_attachment"`
}

type AddStudentMedicalConfidentialNoteRequest struct {
	Category string `json:"note_cat"`  // Max length 3
	Date     string `json:"note_date"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text     string `json:"note_text"` // Max length 4000
}

type UpdateStudentMedicalConfidentialNoteRequest struct {
	Category string `json:"note_cat"`  // Max length 3
	Date     string `json:"note_date"` // Must be a datetime (yyyy-mm-ddTHH:mm:ss.fff)
	Text     string `json:"note_text"` // Max length 4000
}

type StudentPractitionerResponse struct {
	CompanyCode          string  `json:"cmpy_code"`
	Name                 *string `json:"doct_name,omitempty"`
	Phone                *string `json:"doct_phone,omitempty"`
	PractitionerNumber   *string `json:"prac_num,omitempty"` // Must conform to PractitionerNumberValidation
	PractitionerTypeCode *string `json:"ptype_code,omitempty"`
	StudentCode          string  `json:"stud_code"`
}

type AddStudentPractitionerRequest struct {
	Name                 string `json:"doct_name"`  // Max length 30
	PractitionerTypeCode string `json:"ptype_code"` // Max length 3

	Phone *string `json:"doct_phone,omitempty"` // Max length 25
}

type UpdateStudentPractitionerRequest struct {
	Name                 *string `json:"doct_name,omitempty"`  // Max length 50
	Phone                *string `json:"doct_phone,omitempty"` // Max length 20
	PractitionerTypeCode *string `json:"ptype_code,omitempty"` // Max length 10
}

type StudentMedicalSupplementaryResponse struct {
	CompanyCode       string  `json:"cmpy_code"`
	StudentCode       string  `json:"stud_code"`
	Code              string  `json:"msupp_code"`
	AdditionalDetails *string `json:"comm_text,omitempty"`
}

type AddStudentMedicalSupplementaryRequest struct {
	Code              string  `json:"msupp_code"`          // Max length 3
	AdditionalDetails *string `json:"comm_text,omitempty"` // Max length 200
}

type UpdateStudentMedicalSupplementaryRequest struct {
	AdditionalDetails *string `json:"comm_text,omitempty"` // Max length 200
}

// Assessment

type ActivityResponse struct {
	CompanyCode        string                       `json:"cmpy_code"`
	ActivityID         string                       `json:"activity_id"` // Must conform to ActivityIDValidation
	Year               *string                      `json:"year,omitempty"`
	Period             *string                      `json:"period,omitempty"`
	SubjectCode        *string                      `json:"sub_code,omitempty"`
	YearGroup          *string                      `json:"year_grp,omitempty"` // Must conform to YearGroupValidation
	ActivityName       string                       `json:"activity_name"`
	TopicID            string                       `json:"topic_id"` // Must conform to TopicIDValidation
	TopicName          string                       `json:"topic_name"`
	AssessmentCriteria []AssessmentCriteriaResponse `json:"assessment_criteria"`
}

type AssessmentCriteriaResponse struct {
	ObjectCode        string                    `json:"obj_code"`
	ObjectDescription string                    `json:"obj_desc"`
	MaxValue          *string                   `json:"max_val,omitempty"` // Must conform to AssessmentMaxValueValidation
	AssessmentMethod  *AssessmentMethodResponse `json:"assessment_method,omitempty"`
}

type AssessmentMethodResponse struct {
	Code           string               `json:"ass_code"`
	Type           string               `json:"ass_type"`
	Range          string               `json:"ass_range"`
	Description    string               `json:"desc_text"`
	ValidationType string               `json:"val_type"`
	Validations    []ValidationResponse `json:"validations"`
}

type ValidationResponse struct {
	ValidResult string  `json:"valid_result"`
	MinValue    *string `json:"min_val,omitempty"` // Must conform to ValidationMinValueValidation
	MaxValue    *string `json:"max_val,omitempty"` // Must conform to ValidationMaxValueValidation
}

type ActivityStudentResponse struct {
	CompanyCode string  `json:"cmpy_code"`
	ActivityID  string  `json:"activity_id"` // Must conform to ActivityIDValidation
	StudentCode string  `json:"stud_code"`
	ClassCode   *string `json:"class_code,omitempty"`
	YearGroup   *string `json:"year_grp,omitempty"` // Must conform to YearGroupValidation
}

type ActivityStudentResultsResponse struct {
	CompanyCode      string            `json:"cmpy_code"`
	ActivityID       string            `json:"activity_id"` // Must conform to ActivityIDValidation
	StudentCode      string            `json:"stud_code"`
	ObjectiveResults []ObjectiveResult `json:"objective_results"`
}

type UpdateActivityStudentResultsRequest struct {
	ObjectiveResults []ObjectiveResult `json:"objective_results"`
}

type ObjectiveResult struct {
	ObjectiveCode string  `json:"obj_code"`
	StudentResult *string `json:"stud_result,omitempty"`
}

// Attendance

type StudentAttendanceResponse struct {
	CompanyCode              string     `json:"cmpy_code"`
	StudentCode              string     `json:"stud_code"`
	AbsentDate               time.Time  `json:"absent_date"`
	AbsentTime               *time.Time `json:"absent_time,omitempty"`
	AbsentType               string     `json:"absent_type"`
	ReasonCode               *string    `json:"reas_code,omitempty"`
	DoctorsCertificateFlag   bool       `json:"dcert_flg"`
	ParentAcknowledgedFlag   bool       `json:"par_flg"`
	ParentAcknowledgmentDate *time.Time `json:"par_date,omitempty"`
	CorrespondanceSent       bool       `json:"corr_flg"`
	CorrespondanceSentDate   *time.Time `json:"corr_date,omitempty"`
	PeriodCode               *string    `json:"prd_code,omitempty"`
	SourceReference          *string    `json:"ref_num,omitempty"`
	AbsentFromTime           *time.Time `json:"abs_from_time,omitempty"`
	AbsentToTime             *time.Time `json:"abs_to_time,omitempty"`
	ID                       string     `json:"key_num"`         // Must conform to AbsenceIDValidation
	TimetableID              *string    `json:"tt_id,omitempty"` // Must conform to TimetableIDValidation
	Comment                  *string    `json:"note_text,omitempty"`
	YearGroup                *string    `json:"year_grp,omitempty"` // Must conform to YearGroupValidation
	Boarder                  bool       `json:"boarder"`
	House                    *string    `json:"house,omitempty"`
	PCTutorGroup             *string    `json:"pctut_grp,omitempty"`
	Gender                   *string    `json:"gender,omitempty"`
	CampusCode               *string    `json:"campus_code,omitempty"`
	HasAttachment            bool       `json:"has_attachment"`
}

type AbsenceReasonOptionsResponse struct {
	Code             string `json:"code"`
	Description      string `json:"desc"`
	Active           bool   `json:"is_active"`
	AcceptableReason bool   `json:"acceptable_reason"`
}

// Communication Rules

type StudentCommunicationRulesResponse struct {
	CompanyCode        string                                           `json:"cmpy_code"`
	StudentCode        string                                           `json:"stud_code"`
	CommunicationRules []StudentCommunicationRulesParentDetailsResponse `json:"comm_rules"`
}

type StudentCommunicationRulesParentDetailsResponse struct {
	ParentCode      string                                           `json:"par_code"`
	ParentNames     []StudentCommunicationRulesParentNameResponse    `json:"parent_names"`
	ParentAddresses []StudentCommunicationRulesParentAddressResponse `json:"addresses"`
}

type StudentCommunicationRulesParentNameResponse struct {
	ParentType     string  `json:"parent_type"`
	Deceased       bool    `json:"deceased_flg"`
	FirstName      string  `json:"first_name"`
	Gender         *string `json:"gender,omitempty"`
	Initials       *string `json:"initials,omitempty"`
	PersonPosition string  `json:"person_posn"` // Must conform to PersonPositionValidation
	PreferredName  string  `json:"preferred_name"`
	Suffix         *string `json:"suffix,omitempty"`
	Surname        string  `json:"surname"`
	Title          *string `json:"title,omitempty"`
}

type StudentCommunicationRulesParentAddressResponse struct {
	AddressNumber  *string                                               `json:"add_num"` // Must conform to AddressNumberValidation
	AddressLine1   *string                                               `json:"addr1,omitempty"`
	AddressLine2   *string                                               `json:"addr2,omitempty"`
	AddressLine3   *string                                               `json:"addr3,omitempty"`
	Description    *string                                               `json:"addr_desc,omitempty"`
	Addresse       *string                                               `json:"addresse,omitempty"`
	BusinessPhone  *string                                               `json:"bus_phone,omitempty"`
	CallOrder      *string                                               `json:"call_order,omitempty"` // Must conform to CallOrderValidation
	Country        *string                                               `json:"country,omitempty"`
	Email1         *string                                               `json:"e_mail1,omitempty"`
	Email2         *string                                               `json:"e_mail2,omitempty"`
	Fax            *string                                               `json:"fax,omitempty"`
	HomePhone      *string                                               `json:"home_phone,omitempty"`
	MobilePhone1   *string                                               `json:"mobile1,omitempty"`
	MobilePhone2   *string                                               `json:"mobile2,omitempty"`
	PersonPosition *string                                               `json:"person_posn,omitempty"` // Must conform to PersonPositionValidation
	PostCode       *string                                               `json:"post_code,omitempty"`
	Relationship   *string                                               `json:"relationship,omitempty"`
	Salutation     *string                                               `json:"salutation,omitempty"`
	SMSFlag1       bool                                                  `json:"sms_flg1"`
	SMSFlag2       bool                                                  `json:"sms_flg2"`
	StateCode      *string                                               `json:"state_code,omitempty"`
	TownSuburb     *string                                               `json:"town_sub,omitempty"`
	CommTypes      []StudentCommunicationRulesCommunicationTypesResponse `json:"comm_types"`
}

type StudentCommunicationRulesCommunicationTypesResponse struct {
	Code string `json:"commtype_code"`
}
