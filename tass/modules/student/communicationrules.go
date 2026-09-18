package tassstudent

// op: GetStudentCommunicationRules, path: /{cmpy_code}/students/{stud_code}/communicationrules
func (c *Client) GetStudentCommunicationRules(studentCode string) (StudentCommunicationRulesResponse, error) {
	// TODO: Implementation
	return StudentCommunicationRulesResponse{}, nil
}

// op: GetStudentCommunicationRulesByCommTypeCode, path: /{cmpy_code}/students/{stud_code}/communicationrules/{commtype_code}
func (c *Client) GetStudentCommunicationRulesByCommType(studentCode string, communicationType string) (StudentCommunicationRulesResponse, error) {
	// TODO: Implementation
	return StudentCommunicationRulesResponse{}, nil
}

// op: GetAllStudentCommunicationRules, path: /{cmpy_code}/students/communicationrules
func (c *Client) GetAllStudentCommunicationRules() ([]StudentCommunicationRulesResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllStudentCommunicationRulesByCommType, path: /{cmpy_code}/students/communicationrules/{commtype_code}
func (c *Client) GetAllStudentCommunicationRulesByCommType(communicationType string) ([]StudentCommunicationRulesResponse, error) {
	// TODO: Implementation
	return nil, nil
}
