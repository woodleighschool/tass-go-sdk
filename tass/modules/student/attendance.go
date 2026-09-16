package tassstudent

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

func (c *Client) GetAllStudentAttendances() ([]StudentAttendanceResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentAttendanceAttachments(studentCode string, absenceID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentAttendanceAttachment(studentCode string, absenceID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentAbsenceReasons() ([]AbsenceReasonOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetStudentAbsenceTypes() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllCampusOptionsForAttendance() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllHouseOptionsForAttendance() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllYearGroupOptionsForAttendance() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}
