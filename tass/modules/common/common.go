package tasscommon

import "time"

const FileSizeValidation string = `^-?(?:0|[1-9]\\d*)$`

type FileDetails struct {
	Name         *string    `json:"file_name,omitempty"`
	Size         *int       `json:"file_size,omitempty"`
	DateUploaded *time.Time `json:"date_uploaded,omitempty"`
	AttachmentID *string    `json:"attach_id,omitempty"`
}

type FileResponse struct {
	FileName     *string    `json:"file_name,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
	DateUploaded *time.Time `json:"date_uploaded,omitempty"`
	AttachmentID *string    `json:"attach_id,omitempty"`
}

type IFormFile string

// Custom struct
type FileRequest struct {
	FileName    *string   `json:"file_name,omitempty"`
	FileContent IFormFile `json:"file_content,omitempty"`
}

type AttachmentDetails struct {
	FieldNumber int           `json:"field_number"`
	Files       []FileDetails `json:"files"`
}

type NewAttachmentResponse struct {
	AttachmentID string `json:"attach_id"`
}

type OperationType int

type Operation struct {
	Value         any           `json:"value"`
	OperationType OperationType `json:"operationType"`
	Path          *string       `json:"path,omitempty"`
	Operation     *string       `json:"op,omitempty"`
	From          *string       `json:"from,omitempty"`
}

type UDAreaFieldsResponse struct {
	Flags       UDAreaFlagResponse       `json:"ud_flags"`
	Codes       UDAreaCodeResponse       `json:"ud_codes"`
	Text        UDAreaTextResponse       `json:"ud_text"`
	Dates       UDAreaDateResponse       `json:"ud_dates"`
	Attachments UDAreaAttachmentResponse `json:"ud_attachments"`
}

type UDAreaFlagResponse struct {
	UD1Flag  *string `json:"ud1_flag,omitempty"`
	UD2Flag  *string `json:"ud2_flag,omitempty"`
	UD3Flag  *string `json:"ud3_flag,omitempty"`
	UD4Flag  *string `json:"ud4_flag,omitempty"`
	UD5Flag  *string `json:"ud5_flag,omitempty"`
	UD6Flag  *string `json:"ud6_flag,omitempty"`
	UD7Flag  *string `json:"ud7_flag,omitempty"`
	UD8Flag  *string `json:"ud8_flag,omitempty"`
	UD9Flag  *string `json:"ud9_flag,omitempty"`
	UD10Flag *string `json:"ud10_flag,omitempty"`
}

type UDAreaCodeResponse struct {
	UD11Code *string `json:"ud11_code,omitempty"`
	UD12Code *string `json:"ud12_code,omitempty"`
	UD13Code *string `json:"ud13_code,omitempty"`
	UD14Code *string `json:"ud14_code,omitempty"`
	UD15Code *string `json:"ud15_code,omitempty"`
	UD16Code *string `json:"ud16_code,omitempty"`
	UD17Code *string `json:"ud17_code,omitempty"`
	UD18Code *string `json:"ud18_code,omitempty"`
	UD19Code *string `json:"ud19_code,omitempty"`
	UD20Code *string `json:"ud20_code,omitempty"`
}

type UDAreaTextResponse struct {
	UD21Text *string `json:"ud21_text,omitempty"`
	UD22Text *string `json:"ud22_text,omitempty"`
	UD23Text *string `json:"ud23_text,omitempty"`
	UD24Text *string `json:"ud24_text,omitempty"`
	UD25Text *string `json:"ud25_text,omitempty"`
	UD26Text *string `json:"ud26_text,omitempty"`
	UD27Text *string `json:"ud27_text,omitempty"`
	UD28Text *string `json:"ud28_text,omitempty"`
	UD29Text *string `json:"ud29_text,omitempty"`
	UD30Text *string `json:"ud30_text,omitempty"`
}

type UDAreaDateResponse struct {
	UD31Date *time.Time `json:"ud31_date,omitempty"`
	UD32Date *time.Time `json:"ud32_date,omitempty"`
	UD33Date *time.Time `json:"ud33_date,omitempty"`
	UD34Date *time.Time `json:"ud34_date,omitempty"`
	UD35Date *time.Time `json:"ud35_date,omitempty"`
	UD36Date *time.Time `json:"ud36_date,omitempty"`
	UD37Date *time.Time `json:"ud37_date,omitempty"`
	UD38Date *time.Time `json:"ud38_date,omitempty"`
	UD39Date *time.Time `json:"ud39_date,omitempty"`
	UD40Date *time.Time `json:"ud40_date,omitempty"`
}

type UDAreaAttachmentResponse struct {
	UD41AttachmentDetails AttachmentDetails `json:"ud41_attachment_details"`
	UD42AttachmentDetails AttachmentDetails `json:"ud42_attachment_details"`
	UD43AttachmentDetails AttachmentDetails `json:"ud43_attachment_details"`
	UD44AttachmentDetails AttachmentDetails `json:"ud44_attachment_details"`
	UD45AttachmentDetails AttachmentDetails `json:"ud45_attachment_details"`
	UD46AttachmentDetails AttachmentDetails `json:"ud46_attachment_details"`
	UD47AttachmentDetails AttachmentDetails `json:"ud47_attachment_details"`
	UD48AttachmentDetails AttachmentDetails `json:"ud48_attachment_details"`
	UD49AttachmentDetails AttachmentDetails `json:"ud49_attachment_details"`
	UD50AttachmentDetails AttachmentDetails `json:"ud50_attachment_details"`
}

type UDAreaFieldsRequest struct {
	Flags UDAreaFlagRequest `json:"ud_flags"`
	Codes UDAreaCodeRequest `json:"ud_codes"`
	Text  UDAreaTextRequest `json:"ud_text"`
	Dates UDAreaDateRequest `json:"ud_dates"`
}

type UDAreaFlagRequest struct {
	UD1Flag  *string `json:"ud1_flag,omitempty"`  // Max length 1
	UD2Flag  *string `json:"ud2_flag,omitempty"`  // Max length 1
	UD3Flag  *string `json:"ud3_flag,omitempty"`  // Max length 1
	UD4Flag  *string `json:"ud4_flag,omitempty"`  // Max length 1
	UD5Flag  *string `json:"ud5_flag,omitempty"`  // Max length 1
	UD6Flag  *string `json:"ud6_flag,omitempty"`  // Max length 1
	UD7Flag  *string `json:"ud7_flag,omitempty"`  // Max length 1
	UD8Flag  *string `json:"ud8_flag,omitempty"`  // Max length 1
	UD9Flag  *string `json:"ud9_flag,omitempty"`  // Max length 1
	UD10Flag *string `json:"ud10_flag,omitempty"` // Max length 1
}

type UDAreaCodeRequest struct {
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

type UDAreaTextRequest struct {
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

type UDAreaDateRequest struct {
	UD31Date *time.Time `json:"ud31_date,omitempty"` // Max length 100
	UD32Date *time.Time `json:"ud32_date,omitempty"` // Max length 100
	UD33Date *time.Time `json:"ud33_date,omitempty"` // Max length 100
	UD34Date *time.Time `json:"ud34_date,omitempty"` // Max length 100
	UD35Date *time.Time `json:"ud35_date,omitempty"` // Max length 100
	UD36Date *time.Time `json:"ud36_date,omitempty"` // Max length 100
	UD37Date *time.Time `json:"ud37_date,omitempty"` // Max length 100
	UD38Date *time.Time `json:"ud38_date,omitempty"` // Max length 100
	UD39Date *time.Time `json:"ud39_date,omitempty"` // Max length 100
	UD40Date *time.Time `json:"ud40_date,omitempty"` // Max length 100
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
	SortOrder   int     `json:"sort_order"`
}

type UDCodeFieldDetails struct {
	ReferenceValues []UDFieldReferenceValue `json:"reference_values"`

	UDFieldDetails
}

type UDFieldReferenceValue struct {
	Code        *string `json:"ud_code,omitempty"`
	Description string  `json:"ud_desc"`
	SortOrder   int     `json:"sort_order"`
}

type UDAttachmentFieldDetails struct {
	FieldNumber int `json:"field_number"`

	UDFieldDetails
}
