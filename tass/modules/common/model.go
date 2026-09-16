package tasscommon

import "time"

const FileSizeValidation string = `^-?(?:0|[1-9]\\d*)$`

type FileDetails struct {
	Name         *string    `json:"file_name,omitempty"`
	Size         *string    `json:"file_size,omitempty"`
	DateUploaded *time.Time `json:"date_uploaded,omitempty"`
	AttachmentID *string    `json:"attach_id,omitempty"`
}

type FileResponse struct {
	FileName     *string `json:"file_name,omitempty"`
	FileSize     *string `json:"file_size,omitempty"`
	DateUploaded *string `json:"date_uploaded,omitempty"`
	AttachmentID *string `json:"attach_id,omitempty"`
}

type FileRequest struct {
	// TODO: Implementation
}

type OptionsResponse struct {
	Code        string `json:"code"`
	Description string `json:"desc"`
}

type OptionsResponseActive struct {
	Code        string `json:"code"`
	Description string `json:"desc"`
	Active      bool   `json:"is_active"`
}

const FieldNumberValidation string = `^-?(?:0|[1-9]\\d*)$`

type AttachmentDetails struct {
	FieldNumber string        `json:"field_number"`
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
