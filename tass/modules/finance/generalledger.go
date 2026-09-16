package tassfinance

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetAllAccountYearPeriodOptions
func (c *Client) GetAllAccountYearPeriodOptions() ([]YearPeriodOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAccountBudgetOptions
func (c *Client) GetAllAccountBudgetOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllGroupCodeOptions
func (c *Client) GetAllGroupCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllTaxCodeOptions
func (c *Client) GetAllTaxCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllTypeCodesOptions
func (c *Client) GetAllTypeCodesOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAccountCodeFormatRules
func (c *Client) GetAccountCodeFormatRules() (CodeFormatRulesResponse, error) {
	// TODO: Implementation
	return CodeFormatRulesResponse{}, nil
}

// op: GetAllJournalTypesOptions
func (c *Client) GetAllJournalTypesOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllReportingCodeOptions
func (c *Client) GetAllReportingCodeOptions() ([]ReportingCodeResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllGeneralLedgerAccounts
func (c *Client) GetAllGeneralLedgerAccounts() ([]GeneralLedgerAccountsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddGeneralLedgerAccount
func (c *Client) AddGeneralLedgerAccount(payload AddGeneralLedgerAccountRequest) (GeneralLedgerAccountsResponse, error) {
	// TODO: Implementation
	return GeneralLedgerAccountsResponse{}, nil
}

// op: GetGeneralLedgerAccountByCode
func (c *Client) GetGeneralLedgerAccount(accountCode string) (GeneralLedgerAccountResponse, error) {
	// TODO: Implementation
	return GeneralLedgerAccountResponse{}, nil
}

// op: UpdateGeneralLedgerAccount
func (c *Client) UpdateGeneralLedgerAccount(accountCode string, payload UpdateGeneralLedgerAccountRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchGeneralLedgerAccount
func (c *Client) PatchGeneralLedgerAccount(accountCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAccountBudgetsByCode
func (c *Client) GetAccountBudgets(accountCode string) ([]AccountBudgetResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddGeneralLedgerAccountBudget
func (c *Client) AddGeneralLedgerAccountBudget(accountCode string, payload AddAccountBudgetRequest) (AccountBudgetResponse, error) {
	// TODO: Implementation
	return AccountBudgetResponse{}, nil
}

// op: UpdateGeneralLedgerAccountBudget
func (c *Client) UpdateGeneralLedgerAccountBudget(accountCode string, payload UpdateAccountBudgetRequest) error {
	// TODO: Implementation
	return nil
}

// op: GetAccountBalancesByCode
func (c *Client) GetAccountBalances(accountCode string) ([]AccountBalanceResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAccountTransactionsByCode
func (c *Client) GetAccountTransactions(accountCode string, yearNumber string, periodNumber string) ([]AccountTransactionResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAccountTransactionsByPeriod
func (c *Client) GetAllAccountTransactions(yearNumber string, periodNumber string) ([]AccountTransactionResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAccountReportingCodesByCode
func (c *Client) GetAccountReportingCodes(accountCode string) (AccountReportingCodesResponse, error) {
	// TODO: Implementation
	return AccountReportingCodesResponse{}, nil
}

// op: UpdateAccountReportingCodes
func (c *Client) UpdateAccountReportingCodes(accountCode string, payload UpdateAccountReportingCodesRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchAccountReportingCodes
func (c *Client) PatchAccountReportingCodes(accountCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAccountResponsibilitiesByCode
func (c *Client) GetAccountResponsibilities(accountCode string) ([]AccountResponsibilityResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddGeneralLedgerAccountResponsibility
func (c *Client) AddAccountResponsibility(accountCode string, payload AddAccountResponsibilityRequest) (AccountResponsibilityResponse, error) {
	// TODO: Implementation
	return AccountResponsibilityResponse{}, nil
}

// op: GetAllResponsibilityTeacherOptions
func (c *Client) GetAllResponsibilityTeacherOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllResponsibilityEmployeeOptions
func (c *Client) GetAllResponsibilityEmployeeOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllResponsibilityApprovalLevelOptions
func (c *Client) GetAllResponsibilityApprovalLevelOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: UpdateAccountResponsibility
func (c *Client) UpdateAccountResponsibility(accountCode string, userType string, userCode string, payload UpdateAccountResponsibilityRequest) error {
	// TODO: Implementation
	return nil
}

// op: DeleteAccountResponsibility
func (c *Client) DeleteAccountResponsibility(accountCode string, userType string, userCode string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllJournalTypeOptions
func (c *Client) GetAllJournalTypeOptions() ([]JournalTypeOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllGLAccountOptions
func (c *Client) GetAllGLAccountOptions() ([]GLAccountOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllTaxCodeOptions
func (c *Client) GetAllJournalTaxCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllYearPeriodOptions
func (c *Client) GetAllYearPeriodOptions() ([]YearPeriodOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllJournals
func (c *Client) GetAllJournals() ([]JournalResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddGeneralLedgerTaxJournal
func (c *Client) AddGeneralLedgerTaxJournal(payload AddTaxJournalRequest) (AddTaxJournalResponse, error) {
	// TODO: Implementation
	return AddTaxJournalResponse{}, nil
}

// op: UpdateGeneralLedgerTaxJournal
func (c *Client) UpdateGeneralLedgerTaxJournal(journalNumber string, payload UpdateTaxJournalRequest) error {
	// TODO: Implementation
	return nil
}

// op: AddGeneralLedgerGeneralJournal
func (c *Client) AddGeneralLedgerGeneralJournal(payload AddGeneralJournalRequest) (AddGeneralJournalResponse, error) {
	// TODO: Implementation
	return AddGeneralJournalResponse{}, nil
}

// op: UpdateGeneralLedgerGeneralJournal
func (c *Client) UpdateGeneralLedgerGeneralJournal(journalNumber string, payload UpdateGeneralJournalRequest) error {
	// TODO: Implementation
	return nil
}
