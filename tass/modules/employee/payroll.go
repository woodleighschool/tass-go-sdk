package tassemployee

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// op: GetAllEmployeeLeaveEntitlements, path: /{cmpy_code}/payroll/leaveentitlements/{emp_code}
func (c *Client) GetAllEmployeeLeaveEntitlements(ctx context.Context, employeeCode string) ([]EmployeeLeaveBalanceResponse, error) {
	var result []EmployeeLeaveBalanceResponse
	url := fmt.Sprintf("/payroll/leaveentitlements/%s", employeeCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
