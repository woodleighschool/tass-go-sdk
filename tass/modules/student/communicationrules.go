package tassstudent

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

func (c *Client) GetStudentCommunicationRules(studentCode string) (StudentCommunicationRulesResponse, error) {
	// TODO: Implementation
	return StudentCommunicationRulesResponse{}, nil
}

func (c *Client) GetStudentCommunicationRulesByCommType(studentCode string, communicationType string) (StudentCommunicationRulesResponse, error) {
	// TODO: Implementation
	return StudentCommunicationRulesResponse{}, nil
}

func (c *Client) GetAllStudentCommunicationRules() ([]StudentCommunicationRulesResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllStudentCommunicationRulesByCommType(communicationType string) ([]StudentCommunicationRulesResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllCommunicationRuleGenderOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

func (c *Client) GetAllCommunicationRuleTypeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}
