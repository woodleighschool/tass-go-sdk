package tassfinance

import (
	"context"
	"encoding/json"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// Accounts Payable

// op: GetAllAPCreditTaxCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/credits/taxcodes
func (c *Client) GetAllAPCreditTaxCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/credits/taxcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPCreditYearPeriodOptions, path: /{cmpy_code}/options/finance/accountspayable/credits/yearsperiods
func (c *Client) GetAllAPCreditYearPeriodOptions(ctx context.Context) ([]YearPeriodOptionsResponse, error) {
	var result []YearPeriodOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/credits/yearsperiods", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPCreditGLAccountOptions, path: /{cmpy_code}/options/finance/accountspayable/credits/glaccounts
func (c *Client) GetAllAPCreditGLAccountOptions(ctx context.Context) ([]GLAccountOptionsResponse, error) {
	var result []GLAccountOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/credits/glaccounts", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPCreditSupplierOptions, path: /{cmpy_code}/options/finance/accountspayable/credits/suppliers
func (c *Client) GetAllAPCreditSupplierOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/credits/suppliers", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPInvoiceYearPeriodOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/yearsperiods
func (c *Client) GetAllAPInvoiceYearPeriodOptions(ctx context.Context) ([]YearPeriodOptionsResponse, error) {
	var result []YearPeriodOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/invoices/yearsperiods", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPInvoiceGLAccountOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/glaccounts
func (c *Client) GetAllAPInvoiceGLAccountOptions(ctx context.Context) ([]GLAccountOptionsResponse, error) {
	var result []GLAccountOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/invoices/glaccounts", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPInvoiceTermCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/termcodes
func (c *Client) GetAllAPInvoiceTermCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/invoices/termcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPInvoiceTaxCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/taxcodes
func (c *Client) GetAllAPInvoiceTaxCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/invoices/taxcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPInvoiceHoldCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/holdcodes
func (c *Client) GetAllAPInvoiceHoldCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/invoices/holdcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAPInvoiceApproverOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/approvers
func (c *Client) GetAllAPInvoiceApproverOptions(ctx context.Context) ([]tasscommon.OptionsResponseActive, error) {
	var result []tasscommon.OptionsResponseActive
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/invoices/approvers", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllSupplierTypeOptions, path: /{cmpy_code}/options/finance/accountspayable/suppliers/types
func (c *Client) GetAllSupplierTypeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/suppliers/types", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllSupplierTermCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/suppliers/termcodes
func (c *Client) GetAllSupplierTermCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/suppliers/termcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllSupplierTaxCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/suppliers/taxcodes
func (c *Client) GetAllSupplierTaxCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/suppliers/taxcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllSupplierHoldCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/suppliers/holdcodes
func (c *Client) GetAllSupplierHoldCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/suppliers/holdcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllSupplierGLAccountOptions, path: /{cmpy_code}/options/finance/accountspayable/suppliers/glaccounts
func (c *Client) GetAllSupplierGLAccountOptions(ctx context.Context) ([]GLAccountOptionsResponse, error) {
	var result []GLAccountOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/suppliers/glaccounts", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllSupplierBankAccountOptions, path: /{cmpy_code}/options/finance/accountspayable/suppliers/bankaccounts
func (c *Client) GetAllSupplierBankAccountOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/suppliers/bankaccounts", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllSupplierNoteCategoryOptions, path: /{cmpy_code}/options/finance/accountspayable/suppliers/notecategories
func (c *Client) GetAllSupplierNoteCategoryOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/accountspayable/suppliers/notecategories", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// General Ledger

// op: GetAccountCodeFormatRules, path: /{cmpy_code}/options/finance/generalledger/accounts/formatrules
func (c *Client) GetAccountCodeFormatRules(ctx context.Context) (CodeFormatRulesResponse, error) {
	var result CodeFormatRulesResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/formatrules", nil, nil, http.StatusOK)
	if err != nil {
		return CodeFormatRulesResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return CodeFormatRulesResponse{}, err
	}
	return result, nil
}

// op: GetAllReportingCodeOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/reportingcodes/reportingcodes
func (c *Client) GetAllReportingCodeOptions(ctx context.Context) (ReportingCodeResponse, error) {
	var result ReportingCodeResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/reportingcodes/reportingcodes", nil, nil, http.StatusOK)
	if err != nil {
		return ReportingCodeResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return ReportingCodeResponse{}, err
	}
	return result, nil
}

// op: GetAllAccountYearPeriodOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/yearsperiods
func (c *Client) GetAllAccountYearPeriodOptions(ctx context.Context) ([]YearPeriodOptionsResponse, error) {
	var result []YearPeriodOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/yearsperiods", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllAccountBudgetOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/budgets
func (c *Client) GetAllAccountBudgetOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/budgets", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllGroupCodesOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/groupcodes
func (c *Client) GetAllGroupCodesOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/groupcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllTaxCodeOptions, path: /{cmpy_code}/options/finance/generalledger/journals/taxcodes
func (c *Client) GetAllTaxCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/taxcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllTypeCodesOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/typecodes
func (c *Client) GetAllTypeCodesOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/typecodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllJournalTypesOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/transactions/journaltypes
func (c *Client) GetAllJournalTypesOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/transactions/journaltypes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllResponsibilityTeacherOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/responsibilities/teachers
func (c *Client) GetAllResponsibilityTeacherOptions(ctx context.Context) ([]tasscommon.OptionsResponseActive, error) {
	var result []tasscommon.OptionsResponseActive
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/responsibilities/teachers", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllResponsibilityEmployeeOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/responsibilities/employees
func (c *Client) GetAllResponsibilityEmployeeOptions(ctx context.Context) ([]tasscommon.OptionsResponseActive, error) {
	var result []tasscommon.OptionsResponseActive
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/responsibilities/employees", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllResponsibilityApprovalLevelOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/responsibilities/approvallevels
func (c *Client) GetAllResponsibilityApprovalLevelOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/accounts/responsibilities/approvallevels", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllJournalTypeOptions, path: /{cmpy_code}/options/finance/generalledger/journals/types
func (c *Client) GetAllJournalTypeOptions(ctx context.Context) ([]JournalTypeOptionsResponse, error) {
	var result []JournalTypeOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/journals/types", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllGLAccountOptions, path: /{cmpy_code}/options/finance/generalledger/journals/glaccounts
func (c *Client) GetAllGLAccountOptions(ctx context.Context) ([]GLAccountOptionsResponse, error) {
	var result []GLAccountOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/journals/glaccounts", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllTaxCodeOptions, path: /{cmpy_code}/options/finance/generalledger/journals/taxcodes
func (c *Client) GetAllJournalTaxCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/journals/taxcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllYearPeriodOptions, path: /{cmpy_code}/options/finance/generalledger/journals/yearsperiods
func (c *Client) GetAllYearPeriodOptions(ctx context.Context) ([]YearPeriodOptionsResponse, error) {
	var result []YearPeriodOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/generalledger/journals/yearsperiods", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Purchasing

// op: GetAllDeliveryPointOptions, path: /{cmpy_code}/options/finance/purchasing/deliverypoints
func (c *Client) GetAllDeliveryPointOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/purchasing/deliverypoints", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetPurchaseOrderGLAccountOptions, path: /{cmpy_code}/options/finance/purchasing/glaccounts
func (c *Client) GetPurchaseOrderGLAccountOptions(ctx context.Context) ([]GLAccountOptionsResponse, error) {
	var result []GLAccountOptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/purchasing/glaccounts", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllPurchaseOrderStatusOptions, path: /{cmpy_code}/options/finance/purchasing/statuses
func (c *Client) GetAllPurchaseOrderStatusOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/purchasing/statuses", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllPurchaseOrderTaxCodeOptions, path: /{cmpy_code}/options/finance/purchasing/taxcodes
func (c *Client) GetAllPurchaseOrderTaxCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/purchasing/taxcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllTermCodeOptions, path: /{cmpy_code}/options/finance/purchasing/termcodes
func (c *Client) GetAllTermCodeOptions(ctx context.Context) ([]tasscommon.OptionsResponse, error) {
	var result []tasscommon.OptionsResponse
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/purchasing/termcodes", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetAllVendorPurchaseOrderOptions, path: /{cmpy_code}/options/finance/purchasing/vendors
func (c *Client) GetAllVendorPurchaseOrderOptions(ctx context.Context) ([]tasscommon.OptionsResponseActive, error) {
	var result []tasscommon.OptionsResponseActive
	body, err := c.t.request(ctx, http.MethodGet, "/options/finance/purchasing/vendors", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
