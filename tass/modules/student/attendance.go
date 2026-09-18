package tassstudent

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetAllStudentAttendances, path: /{cmpy_code}/students/attendance
func (c *Client) GetAllStudentAttendances() ([]StudentAttendanceResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetStudentAttendanceAttachments, path: /{cmpy_code}/students/{stud_code}/attendance/{key_num}/attachments
func (c *Client) GetStudentAttendanceAttachments(studentCode string, absenceID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DownloadStudentAttendanceAttachment, path: /{cmpy_code}/students/{stud_code}/attendance/{key_num}/attachments/{attach_id}
func (c *Client) GetStudentAttendanceAttachment(studentCode string, absenceID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}
