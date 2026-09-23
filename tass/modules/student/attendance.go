package tassstudent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetAllStudentAttendances, path=/{cmpy_code}/students/attendance
func (c *Client) GetAllStudentAttendances(ctx context.Context) ([]StudentAttendanceResponse, error) {
	var result []StudentAttendanceResponse
	body, err := c.t.request(ctx, http.MethodGet, "/students/absences", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetStudentAttendanceAttachments, path=/{cmpy_code}/students/{stud_code}/attendance/{key_num}/attachments
func (c *Client) GetStudentAttendanceAttachments(ctx context.Context, studentCode string, absenceID int) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/students/%s/attendance/%d/attachments", studentCode, absenceID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DownloadStudentAttendanceAttachment, path=/{cmpy_code}/students/{stud_code}/attendance/{key_num}/attachments/{attach_id}
func (c *Client) GetStudentAttendanceAttachment(ctx context.Context, studentCode string, absenceID int, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/students/%s/attendance/%d/attachments/%s", studentCode, absenceID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
