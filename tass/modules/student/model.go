package tassstudent

import (
	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

type StudentResponse struct {
	AltID                *string              `json:"alt_id,omitempty"`
	Boarder              bool                 `json:"boarder"`
	Campus               *string              `json:"campus,omitempty"`
	CEIDER               *string              `json:"ceider,omitempty"`
	CompanyCode          string               `json:"cmpy_code"`
	ComparativeReporting string               `json:"compare_flg"`
	DateOfArrival        *tasscommon.Date     `json:"date_arrival,omitempty"`
	DistanceEducation    bool                 `json:"distance_ed"`
	DateOfBirth          *tasscommon.Date     `json:"dob,omitempty"`
	DateOfEntry          *tasscommon.Date     `json:"doe,omitempty"`
	DateOfLeaving        *tasscommon.Date     `json:"dol,omitempty"`
	Email                *string              `json:"e_mail,omitempty"`
	EntryYear            *int                 `json:"entry_lev,omitempty"`
	FFPOS                bool                 `json:"ffpos"`
	FirstName            *string              `json:"first_name,omitempty"`
	FormClass            *string              `json:"form_cls,omitempty"`
	FTE                  *float64             `json:"fte"`
	Gender               string               `json:"gender"`
	House                *string              `json:"house,omitempty"`
	IDM                  *string              `json:"idm_id,omitempty"`
	MobilePhone          *string              `json:"mob_phone,omitempty"`
	MultiParenting       bool                 `json:"multipar_flg"`
	NextYear             string               `json:"next_yr_ind"`
	OtherName            *string              `json:"other_name,omitempty"`
	ParentCode           *string              `json:"par_code,omitempty"`
	PCTutorGroup         *string              `json:"pctut_grp,omitempty"`
	PreferredName        *string              `json:"preferred_name,omitempty"`
	PreferredSurname     *string              `json:"preferred_surname,omitempty"`
	PreviousSchool       *string              `json:"prev_school,omitempty"`
	PrivacyFlag          bool                 `json:"privacy_flg"`
	Religion             *string              `json:"religion,omitempty"`
	ResidencyStatus      *string              `json:"resident_sts"`
	SMSFlag              bool                 `json:"sms_flg"`
	StudentCode          string               `json:"stud_code"`
	StudentGovernmentID  *string              `json:"stud_govt_id,omitempty"`
	StudentID            *string              `json:"stud_id,omitempty"`
	Surname              string               `json:"surname"`
	UpdatedOn            *tasscommon.DateTime `json:"update_on,omitempty"`
	USI                  *string              `json:"usi,omitempty"`
	VisaExpiry           *tasscommon.Date     `json:"visa_expiry,omitempty"`
	VisaSubclass         *string              `json:"visa_subclass,omitempty"`
	WebAccess            bool                 `json:"web_access_ind"`
	YearGroup            *int                 `json:"year_grp,omitempty"`
}

type UpdateStudentRequest struct {
	ComparativeReporting string          `json:"compare_flg" validate:"max=1"`
	DateOfEntry          tasscommon.Date `json:"doe"`
	FirstName            string          `json:"first_name" validate:"max=50"`
	FTE                  float64         `json:"fte"`
	Gender               string          `json:"gender" validate:"max=3"`
	PreferredName        string          `json:"preferred_name" validate:"max=20"`
	PreferredSurname     string          `json:"preferred_surname" validate:"max=50"`
	Surname              string          `json:"surname" validate:"max=30"`
	YearGroup            int             `json:"year_grp"`

	AltID             *string          `json:"alt_id,omitempty" validate:"max=40"`
	Boarder           *bool            `json:"boarder,omitempty"`
	Campus            *string          `json:"campus,omitempty" validate:"max=3"`
	CEIDER            *string          `json:"ceider,omitempty" validate:"max=9"`
	DateOfArrival     *tasscommon.Date `json:"date_arrival,omitempty"`
	DistanceEducation *bool            `json:"distance_ed,omitempty"`
	DateOfBirth       *tasscommon.Date `json:"dob,omitempty"`
	DateOfLeaving     *tasscommon.Date `json:"dol,omitempty"`
	Email             *string          `json:"e_mail,omitempty" validate:"max=60"`
	EntryYear         *int             `json:"entry_lev,omitempty"`
	FFPOS             *bool            `json:"ffpos"`
	FormClass         *string          `json:"form_cls,omitempty" validate:"max=2"`
	House             *string          `json:"house,omitempty" validate:"max=2"`
	IDM               *string          `json:"idm_id,omitempty" validate:"max=100"`
	MobilePhone       *string          `json:"mob_phone,omitempty" validate:"max=30"`
	NextYear          *string          `json:"next_yr_ind" validate:"max=1"`
	OtherName         *string          `json:"other_name,omitempty" validate:"max=50"`
	PCTutorGroup      *string          `json:"pctut_grp,omitempty" validate:"max=5"`
	PreviousSchool    *string          `json:"prev_school,omitempty" validate:"max=5"`
	PrivacyFlag       *bool            `json:"privacy_flg"`
	Religion          *string          `json:"religion,omitempty" validate:"max=2"`
	ResidencyStatus   *string          `json:"resident_sts" validate:"max=3"`
	SMSFlag           *bool            `json:"sms_flg"`
	StudentCode       *string          `json:"stud_code" validate:"max=8"`
	USI               *string          `json:"usi,omitempty" validate:"max=10"`
	VisaExpiry        *tasscommon.Date `json:"visa_expiry,omitempty"`
	VisaSubclass      *string          `json:"visa_subclass,omitempty" validate:"max=6"`
	WebAccess         *bool            `json:"web_access_ind"`
}

type StudentStandardNoteResponse struct {
	CompanyCode   string          `json:"cmpy_code"`
	StudentCode   string          `json:"stud_code"`
	NoteCategory  *string         `json:"note_cat,omitempty"`
	NoteDate      tasscommon.Date `json:"note_date"`
	NoteText      *string         `json:"note_text,omitempty"`
	NoteID        string          `json:"note_uid"`
	HasAttachment bool            `json:"has_attachment"`
}

type AddStudentStandardNoteRequest struct {
	NoteCategory string `json:"note_cat" validate:"max=3"`
	NoteDate     string `json:"note_date" validate:"datetime=2006-01-02"`
	NoteText     string `json:"note_text" validate:"max=4000"`
}

type UpdateStudentStandardNoteRequest struct {
	NoteCategory string `json:"note_cat" validate:"max=3"`
	NoteDate     string `json:"note_date" validate:"datetime=2006-01-02"`
	NoteText     string `json:"note_text" validate:"max=4000"`
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
	NoteCategory string `json:"note_cat" validate:"max=3"`
	NoteDate     string `json:"note_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	NoteText     string `json:"note_text" validate:"max=4000"`
}

type UpdateStudentConfidentialNoteRequest struct {
	NoteCategory string `json:"note_cat" validate:"max=3"`
	NoteDate     string `json:"note_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	NoteText     string `json:"note_text" validate:"max=4000"`
}

type StudentPhotoChangesResponse struct {
	CompanyCode string               `json:"cmpy_code"`
	ChangeKey   string               `json:"change_key"`
	Changes     []StudentPhotoChange `json:"changes"`
}

type StudentPhotoChange struct {
	StudentCode string              `json:"stud_code"`
	UpdatedOn   tasscommon.DateTime `json:"photo_update_on"`
}

type StudentUDAreaResponse struct {
	CompanyCode string                      `json:"cmpy_code"`
	AreaCode    string                      `json:"area_code"`
	StudentCode string                      `json:"stud_code"`
	UpdatedOn   *tasscommon.DateTime        `json:"update_on,omitempty"`
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
	UD1Flag  *string `json:"ud1_flg,omitempty" validate:"max=1"`
	UD2Flag  *string `json:"ud2_flg,omitempty" validate:"max=1"`
	UD3Flag  *string `json:"ud3_flg,omitempty" validate:"max=1"`
	UD4Flag  *string `json:"ud4_flg,omitempty" validate:"max=1"`
	UD5Flag  *string `json:"ud5_flg,omitempty" validate:"max=1"`
	UD6Flag  *string `json:"ud6_flg,omitempty" validate:"max=1"`
	UD7Flag  *string `json:"ud7_flg,omitempty" validate:"max=1"`
	UD8Flag  *string `json:"ud8_flg,omitempty" validate:"max=1"`
	UD9Flag  *string `json:"ud9_flg,omitempty" validate:"max=1"`
	UD10Flag *string `json:"ud10_flg,omitempty" validate:"max=1"`
}

type StudentUDAreaCodeResponse struct {
	UD11Code *string `json:"ud11_code,omitempty" validate:"max=3"`
	UD12Code *string `json:"ud12_code,omitempty" validate:"max=3"`
	UD13Code *string `json:"ud13_code,omitempty" validate:"max=3"`
	UD14Code *string `json:"ud14_code,omitempty" validate:"max=3"`
	UD15Code *string `json:"ud15_code,omitempty" validate:"max=3"`
	UD16Code *string `json:"ud16_code,omitempty" validate:"max=3"`
	UD17Code *string `json:"ud17_code,omitempty" validate:"max=3"`
	UD18Code *string `json:"ud18_code,omitempty" validate:"max=3"`
	UD19Code *string `json:"ud19_code,omitempty" validate:"max=3"`
	UD20Code *string `json:"ud20_code,omitempty" validate:"max=3"`
}

type StudentUDAreaTextResponse struct {
	UD21Text *string `json:"ud21_text,omitempty" validate:"max=100"`
	UD22Text *string `json:"ud22_text,omitempty" validate:"max=100"`
	UD23Text *string `json:"ud23_text,omitempty" validate:"max=100"`
	UD24Text *string `json:"ud24_text,omitempty" validate:"max=100"`
	UD25Text *string `json:"ud25_text,omitempty" validate:"max=100"`
	UD26Text *string `json:"ud26_text,omitempty" validate:"max=100"`
	UD27Text *string `json:"ud27_text,omitempty" validate:"max=100"`
	UD28Text *string `json:"ud28_text,omitempty" validate:"max=100"`
	UD29Text *string `json:"ud29_text,omitempty" validate:"max=100"`
	UD30Text *string `json:"ud30_text,omitempty" validate:"max=100"`
}

type StudentUDAreaDateResponse struct {
	UD31Date *tasscommon.Date `json:"ud31_date,omitempty"`
	UD32Date *tasscommon.Date `json:"ud32_date,omitempty"`
	UD40Date *tasscommon.Date `json:"ud33_date,omitempty"`
	UD33Date *tasscommon.Date `json:"ud34_date,omitempty"`
	UD34Date *tasscommon.Date `json:"ud35_date,omitempty"`
	UD35Date *tasscommon.Date `json:"ud36_date,omitempty"`
	UD36Date *tasscommon.Date `json:"ud37_date,omitempty"`
	UD37Date *tasscommon.Date `json:"ud38_date,omitempty"`
	UD38Date *tasscommon.Date `json:"ud39_date,omitempty"`
	UD39Date *tasscommon.Date `json:"ud40_date,omitempty"`
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
	UD1Flag  *string `json:"ud1_flg,omitempty" validate:"max=1"`
	UD2Flag  *string `json:"ud2_flg,omitempty" validate:"max=1"`
	UD3Flag  *string `json:"ud3_flg,omitempty" validate:"max=1"`
	UD4Flag  *string `json:"ud4_flg,omitempty" validate:"max=1"`
	UD5Flag  *string `json:"ud5_flg,omitempty" validate:"max=1"`
	UD6Flag  *string `json:"ud6_flg,omitempty" validate:"max=1"`
	UD7Flag  *string `json:"ud7_flg,omitempty" validate:"max=1"`
	UD8Flag  *string `json:"ud8_flg,omitempty" validate:"max=1"`
	UD9Flag  *string `json:"ud9_flg,omitempty" validate:"max=1"`
	UD10Flag *string `json:"ud10_flg,omitempty" validate:"max=1"`
}

type StudentUDAreaCodeRequest struct {
	UD11Code *string `json:"ud11_code,omitempty" validate:"max=3"`
	UD12Code *string `json:"ud12_code,omitempty" validate:"max=3"`
	UD13Code *string `json:"ud13_code,omitempty" validate:"max=3"`
	UD14Code *string `json:"ud14_code,omitempty" validate:"max=3"`
	UD15Code *string `json:"ud15_code,omitempty" validate:"max=3"`
	UD16Code *string `json:"ud16_code,omitempty" validate:"max=3"`
	UD17Code *string `json:"ud17_code,omitempty" validate:"max=3"`
	UD18Code *string `json:"ud18_code,omitempty" validate:"max=3"`
	UD19Code *string `json:"ud19_code,omitempty" validate:"max=3"`
	UD20Code *string `json:"ud20_code,omitempty" validate:"max=3"`
}

type StudentUDAreaTextRequest struct {
	UD21Text *string `json:"ud21_text,omitempty" validate:"max=100"`
	UD22Text *string `json:"ud22_text,omitempty" validate:"max=100"`
	UD23Text *string `json:"ud23_text,omitempty" validate:"max=100"`
	UD24Text *string `json:"ud24_text,omitempty" validate:"max=100"`
	UD25Text *string `json:"ud25_text,omitempty" validate:"max=100"`
	UD26Text *string `json:"ud26_text,omitempty" validate:"max=100"`
	UD27Text *string `json:"ud27_text,omitempty" validate:"max=100"`
	UD28Text *string `json:"ud28_text,omitempty" validate:"max=100"`
	UD29Text *string `json:"ud29_text,omitempty" validate:"max=100"`
	UD30Text *string `json:"ud30_text,omitempty" validate:"max=100"`
}

type StudentUDAreaDateRequest struct {
	UD31Date *tasscommon.Date `json:"ud31_date,omitempty"`
	UD32Date *tasscommon.Date `json:"ud32_date,omitempty"`
	UD40Date *tasscommon.Date `json:"ud33_date,omitempty"`
	UD33Date *tasscommon.Date `json:"ud34_date,omitempty"`
	UD34Date *tasscommon.Date `json:"ud35_date,omitempty"`
	UD35Date *tasscommon.Date `json:"ud36_date,omitempty"`
	UD36Date *tasscommon.Date `json:"ud37_date,omitempty"`
	UD37Date *tasscommon.Date `json:"ud38_date,omitempty"`
	UD38Date *tasscommon.Date `json:"ud39_date,omitempty"`
	UD39Date *tasscommon.Date `json:"ud40_date,omitempty"`
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
	UD1Flag  *string `json:"ud1_flg,omitempty" validate:"max=1"`
	UD2Flag  *string `json:"ud2_flg,omitempty" validate:"max=1"`
	UD3Flag  *string `json:"ud3_flg,omitempty" validate:"max=1"`
	UD4Flag  *string `json:"ud4_flg,omitempty" validate:"max=1"`
	UD5Flag  *string `json:"ud5_flg,omitempty" validate:"max=1"`
	UD6Flag  *string `json:"ud6_flg,omitempty" validate:"max=1"`
	UD7Flag  *string `json:"ud7_flg,omitempty" validate:"max=1"`
	UD8Flag  *string `json:"ud8_flg,omitempty" validate:"max=1"`
	UD9Flag  *string `json:"ud9_flg,omitempty" validate:"max=1"`
	UD10Flag *string `json:"ud10_flg,omitempty" validate:"max=1"`
}

type UDCodeResponse struct {
	UD11Code *string `json:"ud11_code,omitempty" validate:"max=3"`
	UD12Code *string `json:"ud12_code,omitempty" validate:"max=3"`
	UD13Code *string `json:"ud13_code,omitempty" validate:"max=3"`
	UD14Code *string `json:"ud14_code,omitempty" validate:"max=3"`
	UD15Code *string `json:"ud15_code,omitempty" validate:"max=3"`
	UD16Code *string `json:"ud16_code,omitempty" validate:"max=3"`
	UD17Code *string `json:"ud17_code,omitempty" validate:"max=3"`
	UD18Code *string `json:"ud18_code,omitempty" validate:"max=3"`
	UD19Code *string `json:"ud19_code,omitempty" validate:"max=3"`
	UD20Code *string `json:"ud20_code,omitempty" validate:"max=3"`
}

type UDTextResponse struct {
	UD21Text *string `json:"ud21_text,omitempty" validate:"max=100"`
	UD22Text *string `json:"ud22_text,omitempty" validate:"max=100"`
	UD23Text *string `json:"ud23_text,omitempty" validate:"max=100"`
	UD24Text *string `json:"ud24_text,omitempty" validate:"max=100"`
	UD25Text *string `json:"ud25_text,omitempty" validate:"max=100"`
	UD26Text *string `json:"ud26_text,omitempty" validate:"max=100"`
	UD27Text *string `json:"ud27_text,omitempty" validate:"max=100"`
	UD28Text *string `json:"ud28_text,omitempty" validate:"max=100"`
	UD29Text *string `json:"ud29_text,omitempty" validate:"max=100"`
	UD30Text *string `json:"ud30_text,omitempty" validate:"max=100"`
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
	UD1Flag  *string `json:"ud1_flg,omitempty" validate:"max=1"`
	UD2Flag  *string `json:"ud2_flg,omitempty" validate:"max=1"`
	UD3Flag  *string `json:"ud3_flg,omitempty" validate:"max=1"`
	UD4Flag  *string `json:"ud4_flg,omitempty" validate:"max=1"`
	UD5Flag  *string `json:"ud5_flg,omitempty" validate:"max=1"`
	UD6Flag  *string `json:"ud6_flg,omitempty" validate:"max=1"`
	UD7Flag  *string `json:"ud7_flg,omitempty" validate:"max=1"`
	UD8Flag  *string `json:"ud8_flg,omitempty" validate:"max=1"`
	UD9Flag  *string `json:"ud9_flg,omitempty" validate:"max=1"`
	UD10Flag *string `json:"ud10_flg,omitempty" validate:"max=1"`
}

type UDCodeRequest struct {
	UD11Code *string `json:"ud11_code,omitempty" validate:"max=3"`
	UD12Code *string `json:"ud12_code,omitempty" validate:"max=3"`
	UD13Code *string `json:"ud13_code,omitempty" validate:"max=3"`
	UD14Code *string `json:"ud14_code,omitempty" validate:"max=3"`
	UD15Code *string `json:"ud15_code,omitempty" validate:"max=3"`
	UD16Code *string `json:"ud16_code,omitempty" validate:"max=3"`
	UD17Code *string `json:"ud17_code,omitempty" validate:"max=3"`
	UD18Code *string `json:"ud18_code,omitempty" validate:"max=3"`
	UD19Code *string `json:"ud19_code,omitempty" validate:"max=3"`
	UD20Code *string `json:"ud20_code,omitempty" validate:"max=3"`
}

type UDTextRequest struct {
	UD21Text *string `json:"ud21_text,omitempty" validate:"max=100"`
	UD22Text *string `json:"ud22_text,omitempty" validate:"max=100"`
	UD23Text *string `json:"ud23_text,omitempty" validate:"max=100"`
	UD24Text *string `json:"ud24_text,omitempty" validate:"max=100"`
	UD25Text *string `json:"ud25_text,omitempty" validate:"max=100"`
	UD26Text *string `json:"ud26_text,omitempty" validate:"max=100"`
	UD27Text *string `json:"ud27_text,omitempty" validate:"max=100"`
	UD28Text *string `json:"ud28_text,omitempty" validate:"max=100"`
	UD29Text *string `json:"ud29_text,omitempty" validate:"max=100"`
	UD30Text *string `json:"ud30_text,omitempty" validate:"max=100"`
}

type StudentUDFieldOptionResponse struct {
	CompanyCode string `json:"cmpy_code"`
}

type StudentMCEECDYAResponse struct {
	CompanyCode            string               `json:"cmpy_code"`
	StudentCode            string               `json:"stud_code"`
	ArrivalYear            *int                 `json:"arrive_yr,omitempty"`
	Parent1LOTE            *string              `json:"mlote_code,omitempty"`
	Parent1NonSchoolLevel  *string              `json:"mnse_code,omitempty"`
	Parent1OccupationGroup *string              `json:"mocc_code,omitempty"`
	Parent1SchoolLevel     *string              `json:"mse_code,omitempty"`
	Parent2LOTE            *string              `json:"flote_code,omitempty"`
	Parent2NonSchoolLevel  *string              `json:"fnse_code,omitempty"`
	Parent2OccupationGroup *string              `json:"focc_code,omitempty"`
	Parent2SchoolLevel     *string              `json:"fse_code,omitempty"`
	IndiginousStatus       *string              `json:"s_indig_sts,omitempty"`
	CountryOfBirth         *string              `json:"scob_code,omitempty"`
	LOTE                   *string              `json:"slote_code,omitempty"`
	UpdatedOn              *tasscommon.DateTime `json:"update_on,omitempty"`
}

type UpdateStudentMCEECDYARequest struct {
	ArrivalYear            *int    `json:"arrive_yr,omitempty"`
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
	CompanyCode          string           `json:"cmpy_code"`
	LastOccurance        *tasscommon.Date `json:"last_occ_date,omitempty"`
	MedicalConditionCode string           `json:"mcond_code"`
	Severe               bool             `json:"severe_ind"`
	StudentCode          string           `json:"stud_code"`
	TreatmentDetails     *string          `json:"treat_text,omitempty"`
	Active               bool             `json:"active_flg"`
	UDFields             UDFields         `json:"ud_fields"`
	HasAttachment        bool             `json:"has_attachment"`
	HasNote              bool             `json:"has_note"`
}

type UDFields struct {
	UDText UDText `json:"ud_text"`
}

type UDText struct {
	UD1Text  *string `json:"ud1_text,omitempty" validate:"max=100"`
	UD2Text  *string `json:"ud2_text,omitempty" validate:"max=100"`
	UD3Text  *string `json:"ud3_text,omitempty" validate:"max=100"`
	UD4Text  *string `json:"ud4_text,omitempty" validate:"max=100"`
	UD5Text  *string `json:"ud5_text,omitempty" validate:"max=100"`
	UD6Text  *string `json:"ud6_text,omitempty" validate:"max=100"`
	UD7Text  *string `json:"ud7_text,omitempty" validate:"max=100"`
	UD8Text  *string `json:"ud8_text,omitempty" validate:"max=100"`
	UD9Text  *string `json:"ud9_text,omitempty" validate:"max=100"`
	UD10Text *string `json:"ud10_text,omitempty" validate:"max=100"`
}

type AddStudentMedicalConditionRequest struct {
	LastOccurance        *tasscommon.Date `json:"last_occ_date,omitempty"`
	MedicalConditionCode string           `json:"mcond_code" validate:"max=3"`
	Severe               bool             `json:"severe_ind"`
	TreatmentDetails     *string          `json:"treat_text,omitempty" validate:"max=4000"`
	Active               bool             `json:"active_flg"`
	UDFields             UDFields         `json:"ud_fields"`
}

type UpdateStudentMedicalConditionRequest struct {
	LastOccurance    *tasscommon.Date `json:"last_occ_date,omitempty"`
	Severe           bool             `json:"severe_ind"`
	TreatmentDetails *string          `json:"treat_text,omitempty" validate:"max=4000"`
	Active           bool             `json:"active_flg"`
	UDFields         UDFields         `json:"ud_fields"`
}

type StudentMedicalConditionNoteResponse struct {
	CompanyCode          string          `json:"cmpy_code"`
	StudentCode          string          `json:"stud_code"`
	MedicalConditionCode string          `json:"mcond_code"`
	Date                 tasscommon.Date `json:"note_date"`
	Text                 *string         `json:"note_text,omitempty"`
	ID                   string          `json:"note_uid"` // Must be a UUID
}

type AddStudentMedicalConditionNoteRequest struct {
	Date string  `json:"note_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Text *string `json:"note_text,omitempty" validate:"max=4000"`
}

type UpdateStudentMedicalConditionNoteRequest struct {
	Date string  `json:"note_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Text *string `json:"note_text,omitempty" validate:"max=4000"`
}

type StudentIllnessResponse struct {
	CompanyCode          string           `json:"cmpy_code"`
	StudentCode          string           `json:"stud_code"`
	IllnessUID           string           `json:"illness_uid"`
	IllnessDate          tasscommon.Date  `json:"ill_date"`
	IllnessTime          *tasscommon.Date `json:"ill_time,omitempty"`
	MedicalConditionCode *string          `json:"mcond_code,omitempty"`
	TreatmentCode        *string          `json:"treat_code,omitempty"`
	DischargeDate        *tasscommon.Date `json:"disch_date,omitempty"`
	DischargeTime        *tasscommon.Date `json:"disch_time,omitempty"`
	Hospitalised         bool             `json:"host_flg"`
	IllnessDescription   *string          `json:"ill_desc,omitempty"`
	IllnessNotes         *string          `json:"ill_note,omitempty"`
	HasMedications       bool             `json:"has_medications"`
}

type AddStudentIllnessRequest struct {
	IllnessDate          string `json:"ill_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	IllnessTime          string `json:"ill_time" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	MedicalConditionCode string `json:"mcond_code" validate:"max=3"`

	TreatmentCode      *string `json:"treat_code,omitempty" validate:"max=3"`
	DischargeDate      *string `json:"disch_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	DischargeTime      *string `json:"disch_time,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Hospitalised       bool    `json:"host_flg"`
	IllnessDescription *string `json:"ill_desc,omitempty" validate:"max=60"`
	IllnessNotes       *string `json:"ill_note,omitempty" validate:"max=4000"`
}

type UpdateStudentIllnessRequest struct {
	IllnessDate          string  `json:"ill_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	IllnessTime          string  `json:"ill_time" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	MedicalConditionCode string  `json:"mcond_code" validate:"max=3"`
	TreatmentCode        string  `json:"treat_code,omitempty" validate:"max=3"`
	DischargeDate        string  `json:"disch_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	DischargeTime        string  `json:"disch_time,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Hospitalised         bool    `json:"host_flg"`
	IllnessDescription   *string `json:"ill_desc,omitempty" validate:"max=60"`
	IllnessNotes         *string `json:"ill_note,omitempty" validate:"max=4000"`
}

type StudentImmunisationResponse struct {
	CompanyCode      string `json:"cmpy_code"`
	ImmunisationCode string `json:"imm_code"`
	ImmunisationYear *int   `json:"imm_year,omitempty"`
	StudentCode      string `json:"stud_code"`
}

type AddStudentImmunisationRequest struct {
	ImmunisationCode string `json:"imm_code" validate:"max=2"`
	ImmunisationYear *int   `json:"imm_year,omitempty"`
}

type UpdateStudentImmunisationRequest struct {
	ImmunisationYear *int `json:"imm_year,omitempty"`
}

type StudentImmunisationRegisterResponse struct {
	CompanyCode   string           `json:"cmpy_code"`
	AIRStateDate  *tasscommon.Date `json:"air_state_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	NextDueDate   *tasscommon.Date `json:"next_due_date,omitempty" validate:"datetime=2006-01-02"`  // TODO: Confirm date format
	StudentCode   string           `json:"stud_code"`
	StatusCode    *string          `json:"status_code,omitempty"`
	HasAttachment bool             `json:"has_attachment"`
}

type UpdateStudentImmunisationRegisterRequest struct {
	NextDueDate *string `json:"next_due_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
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
	EndDate               *tasscommon.Date      `json:"end_date,omitempty"`
	ExpiryDate            *tasscommon.Date      `json:"expiry_date,omitempty"`
	MedicalConditionCode  string                `json:"mcond_code"`
	FurtherDetails        *string               `json:"med_detl,omitempty"`
	MethodOfUse           *string               `json:"med_meth,omitempty"`
	Name                  *string               `json:"med_text,omitempty"`
	MedicationUID         string                `json:"medication_uid"` // Must be a UUID
	MinTimeBetweenDoses   *int                  `json:"min_time_between_doses,omitempty"`
	PrescribingDoctor     *string               `json:"script_doc,omitempty"`
	StartDate             *tasscommon.Date      `json:"start_date,omitempty"`
	StudentCode           string                `json:"stud_code"`
	StaffTrainingRequired bool                  `json:"training"`
	HasAttachment         bool                  `json:"has_attachment"`
	HasNote               bool                  `json:"has_note"`
	HasSchedule           bool                  `json:"has_schedule"`
}

type AddStudentMedicationRequest struct {
	Active                bool                 `json:"active_flg"`
	Administer            MedicationAdminister `json:"administer"`
	DoctorPhone           *string              `json:"doc_phone,omitempty" validate:"max=25"`
	EndDate               *string              `json:"end_date,omitempty" validate:"datetime=2006-01-02"`    // TODO: Confirm date format
	ExpiryDate            *string              `json:"expiry_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	FurtherDetails        *string              `json:"med_detl,omitempty" validate:"max=200"`
	MethodOfUse           *string              `json:"med_meth,omitempty" validate:"max=200"`
	Name                  *string              `json:"med_text,omitempty" validate:"max=200"`
	MinTimeBetweenDoses   *int                 `json:"min_time_between_doses,omitempty"`
	PrescribingDoctor     *string              `json:"script_doc,omitempty" validate:"max=30"`
	StartDate             *string              `json:"start_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	StaffTrainingRequired bool                 `json:"training"`
}

type UpdateStudentMedicationRequest struct {
	Active            bool                 `json:"active_flg"`
	Administer        MedicationAdminister `json:"administer"`
	DoctorPhone       *string              `json:"doc_phone,omitempty" validate:"max=25"`
	EndDate           *string              `json:"end_date,omitempty" validate:"datetime=2006-01-02"`    // TODO: Confirm date format
	ExpiryDate        *string              `json:"expiry_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	FurtherDetails    *string              `json:"med_detl,omitempty" validate:"max=200"`
	MethodOfUse       *string              `json:"med_meth,omitempty" validate:"max=200"`
	Name              *string              `json:"med_text,omitempty" validate:"max=200"`
	MedicationUID     string               `json:"medication_uid"` // Must be a UUID
	PrescribingDoctor *string              `json:"script_doc,omitempty" validate:"max=30"`
	StartDate         *string              `json:"start_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	StudentCode       string               `json:"stud_code"`
	Training          bool                 `json:"training"`
}

type StudentMedicationNoteResponse struct {
	CompanyCode          string          `json:"cmpy_code"`
	StudentCode          string          `json:"stud_code"`
	MedicalConditionCode string          `json:"mcond_code"`
	MedicationUID        string          `json:"medication_uid"` // Must be a UUID
	Date                 tasscommon.Date `json:"note_date"`
	Text                 *string         `json:"note_text,omitempty"`
	ID                   string          `json:"note_uid"` // Must be a UUID
}

type AddStudentMedicationNoteRequest struct {
	Date string  `json:"note_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Text *string `json:"note_text,omitempty" validate:"max=4000"`
}

type UpdateStudentMedicationNoteRequest struct {
	Date string  `json:"note_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Text *string `json:"note_text,omitempty" validate:"max=4000"`
}

type StudentMedicationScheduleResponse struct {
	CompanyCode          string       `json:"cmpy_code"`
	StudentCode          string       `json:"stud_code"`
	MedicalConditionCode string       `json:"mcond_code"`
	MedicationUID        string       `json:"medication_uid"` // Must be a UUID
	ID                   string       `json:"sched_uid"`      // Must be a UUID
	Dose                 *string      `json:"med_dose,omitempty"`
	DoseTime             string       `json:"med_time"`                                                 // TODO: Must be a time
	StartDate            *string      `json:"shed_start_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	EndDate              *string      `json:"shed_end_date,omitempty" validate:"datetime=2006-01-02"`   // TODO: Confirm date format
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
	Time      string `json:"med_time"`                                       // Must be a time
	StartDate string `json:"shed_start_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format

	EndDate *string      `json:"shed_end_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
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
	Time      string `json:"med_time"`                                       // Must be a time
	StartDate string `json:"shed_start_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format

	EndDate *string      `json:"shed_end_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Days    *DaysRequest `json:"days,omitempty"`
}

type StudentMedicalStandardNoteResponse struct {
	CompanyCode   string          `json:"cmpy_code"`
	StudentCode   string          `json:"stud_code"`
	Category      *string         `json:"note_cat,omitempty"`
	Date          tasscommon.Date `json:"note_date"`
	Text          *string         `json:"note_text,omitempty"`
	ID            string          `json:"note_uid"` // Must be a UUID
	HasAttachment bool            `json:"has_attachment"`
}

type AddStudentMedicalStandardNoteRequest struct {
	Category string `json:"note_cat" validate:"max=3"`
	Date     string `json:"note_date" validate:"datetime=2006-01-02T03:04:05.000"` // TODO: Confirm this works
	Text     string `json:"note_text" validate:"max=4000"`
}

type UpdateStudentMedicalStandardNoteRequest struct {
	Category string `json:"note_cat" validate:"max=3"`
	Date     string `json:"note_date" validate:"datetime=2006-01-02T03:04:05.000"` // TODO: Confirm this works
	Text     string `json:"note_text" validate:"max=4000"`
}

type StudentMedicalConfidentialNoteResponse struct {
	CompanyCode   string          `json:"cmpy_code"`
	StudentCode   string          `json:"stud_code"`
	Category      *string         `json:"note_cat,omitempty"`
	Date          tasscommon.Date `json:"note_date"`
	Text          *string         `json:"note_text,omitempty"`
	ID            string          `json:"note_uid"` // Must be a UUID
	HasAttachment bool            `json:"has_attachment"`
}

type AddStudentMedicalConfidentialNoteRequest struct {
	Category string `json:"note_cat" validate:"max=3"`
	Date     string `json:"note_date" validate:"datetime=2006-01-02T03:04:05.000"` // TODO: Confirm this works
	Text     string `json:"note_text" validate:"max=4000"`
}

type UpdateStudentMedicalConfidentialNoteRequest struct {
	Category string `json:"note_cat" validate:"max=3"`
	Date     string `json:"note_date" validate:"datetime=2006-01-02T03:04:05.000"` // TODO: Confirm this works
	Text     string `json:"note_text" validate:"max=4000"`
}

type StudentPractitionerResponse struct {
	CompanyCode          string  `json:"cmpy_code"`
	Name                 *string `json:"doct_name,omitempty"`
	Phone                *string `json:"doct_phone,omitempty"`
	PractitionerNumber   *int    `json:"prac_num,omitempty"`
	PractitionerTypeCode *string `json:"ptype_code,omitempty"`
	StudentCode          string  `json:"stud_code"`
}

type AddStudentPractitionerRequest struct {
	Name                 string `json:"doct_name" validate:"max=30"`
	PractitionerTypeCode string `json:"ptype_code" validate:"max=3"`

	Phone *string `json:"doct_phone,omitempty" validate:"max=25"`
}

type UpdateStudentPractitionerRequest struct {
	Name                 *string `json:"doct_name,omitempty" validate:"max=50"`
	Phone                *string `json:"doct_phone,omitempty" validate:"max=20"`
	PractitionerTypeCode *string `json:"ptype_code,omitempty" validate:"max=10"`
}

type StudentMedicalSupplementaryResponse struct {
	CompanyCode       string  `json:"cmpy_code"`
	StudentCode       string  `json:"stud_code"`
	Code              string  `json:"msupp_code"`
	AdditionalDetails *string `json:"comm_text,omitempty"`
}

type AddStudentMedicalSupplementaryRequest struct {
	Code              string  `json:"msupp_code" validate:"max=3"`
	AdditionalDetails *string `json:"comm_text,omitempty" validate:"max=200"`
}

type UpdateStudentMedicalSupplementaryRequest struct {
	AdditionalDetails *string `json:"comm_text,omitempty" validate:"max=200"`
}

// Assessment

type ActivityResponse struct {
	CompanyCode        string                       `json:"cmpy_code"`
	ActivityID         int                          `json:"activity_id"`
	Year               *string                      `json:"year,omitempty"`
	Period             *string                      `json:"period,omitempty"`
	SubjectCode        *string                      `json:"sub_code,omitempty"`
	YearGroup          *int                         `json:"year_grp,omitempty"`
	ActivityName       string                       `json:"activity_name"`
	TopicID            int                          `json:"topic_id"`
	TopicName          string                       `json:"topic_name"`
	AssessmentCriteria []AssessmentCriteriaResponse `json:"assessment_criteria"`
}

type AssessmentCriteriaResponse struct {
	ObjectCode        string                    `json:"obj_code"`
	ObjectDescription string                    `json:"obj_desc"`
	MaxValue          *int                      `json:"max_val,omitempty"`
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
	ValidResult string `json:"valid_result"`
	MinValue    *int   `json:"min_val,omitempty"`
	MaxValue    *int   `json:"max_val,omitempty"`
}

type ActivityStudentResponse struct {
	CompanyCode string  `json:"cmpy_code"`
	ActivityID  int     `json:"activity_id"`
	StudentCode string  `json:"stud_code"`
	ClassCode   *string `json:"class_code,omitempty"`
	YearGroup   *int    `json:"year_grp,omitempty"`
}

type ActivityStudentResultsResponse struct {
	CompanyCode      string            `json:"cmpy_code"`
	ActivityID       int               `json:"activity_id"`
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
	CompanyCode              string           `json:"cmpy_code"`
	StudentCode              string           `json:"stud_code"`
	AbsentDate               tasscommon.Date  `json:"absent_date"`
	AbsentTime               *tasscommon.Date `json:"absent_time,omitempty"`
	AbsentType               string           `json:"absent_type"`
	ReasonCode               *string          `json:"reas_code,omitempty"`
	DoctorsCertificateFlag   bool             `json:"dcert_flg"`
	ParentAcknowledgedFlag   bool             `json:"par_flg"`
	ParentAcknowledgmentDate *tasscommon.Date `json:"par_date,omitempty"`
	CorrespondanceSent       bool             `json:"corr_flg"`
	CorrespondanceSentDate   *tasscommon.Date `json:"corr_date,omitempty"`
	PeriodCode               *string          `json:"prd_code,omitempty"`
	SourceReference          *string          `json:"ref_num,omitempty"`
	AbsentFromTime           *tasscommon.Date `json:"abs_from_time,omitempty"`
	AbsentToTime             *tasscommon.Date `json:"abs_to_time,omitempty"`
	ID                       int              `json:"key_num"`
	TimetableID              *int             `json:"tt_id,omitempty"`
	Comment                  *string          `json:"note_text,omitempty"`
	YearGroup                *int             `json:"year_grp,omitempty"`
	Boarder                  bool             `json:"boarder"`
	House                    *string          `json:"house,omitempty"`
	PCTutorGroup             *string          `json:"pctut_grp,omitempty"`
	Gender                   *string          `json:"gender,omitempty"`
	CampusCode               *string          `json:"campus_code,omitempty"`
	HasAttachment            bool             `json:"has_attachment"`
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
	PersonPosition int     `json:"person_posn"`
	PreferredName  string  `json:"preferred_name"`
	Suffix         *string `json:"suffix,omitempty"`
	Surname        string  `json:"surname"`
	Title          *string `json:"title,omitempty"`
}

type StudentCommunicationRulesParentAddressResponse struct {
	AddressNumber  *int                                                  `json:"add_num"`
	AddressLine1   *string                                               `json:"addr1,omitempty"`
	AddressLine2   *string                                               `json:"addr2,omitempty"`
	AddressLine3   *string                                               `json:"addr3,omitempty"`
	Description    *string                                               `json:"addr_desc,omitempty"`
	Addresse       *string                                               `json:"addresse,omitempty"`
	BusinessPhone  *string                                               `json:"bus_phone,omitempty"`
	CallOrder      *int                                                  `json:"call_order,omitempty"`
	Country        *string                                               `json:"country,omitempty"`
	Email1         *string                                               `json:"e_mail1,omitempty"`
	Email2         *string                                               `json:"e_mail2,omitempty"`
	Fax            *string                                               `json:"fax,omitempty"`
	HomePhone      *string                                               `json:"home_phone,omitempty"`
	MobilePhone1   *string                                               `json:"mobile1,omitempty"`
	MobilePhone2   *string                                               `json:"mobile2,omitempty"`
	PersonPosition *int                                                  `json:"person_posn,omitempty"`
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
