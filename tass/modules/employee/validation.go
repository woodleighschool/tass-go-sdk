package tassemployee

import (
	"fmt"
	"time"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

func (p AddEmployeePDActivityRequest) validate() error {
	// TODO: What format is the date in
	if _, err := time.Parse("", p.FinishDate); err != nil {
		return fmt.Errorf(tasscommon.DateError, "AddEmployeePDActivityRequest", "FinishDate", err)
	}
	if len(p.Status) > 3 {
		return fmt.Errorf(tasscommon.MaxLengthError, "AddEmployeePDActivityRequest", "Status", "3")
	}
	if len(p.Details) > 200 {
		return fmt.Errorf(tasscommon.MaxLengthError, "AddEmployeePDActivityRequest", "Details", "200")
	}
	if _, err := time.Parse("", p.StartDate); err != nil {
		return fmt.Errorf(tasscommon.DateError, "AddEmployeePDActivityRequest", "StartDate", err)
	}
	return nil
}
