package tassfinance

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetAccountCodeFormatRules, path: /{cmpy_code}/options/finance/generalledger/accounts/formatrules
func (c *Client) GetAccountCodeFormatRules() (CodeFormatRulesResponse, error) {
	// TODO: Implementation
	return CodeFormatRulesResponse{}, nil
}

// op: GetAllReportingCodeOptions, path: /{cmpy_code}/options/finance/generalledger/accounts/reportingcodes/reportingcodes
func (c *Client) GetAllReportingCodeOptions() ([]ReportingCodeResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllGeneralLedgerAccounts, path: /{cmpy_code}/options/finance/generalledger/accounts
func (c *Client) GetAllGeneralLedgerAccounts() ([]GeneralLedgerAccountResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddGeneralLedgerAccount, path: /{cmpy_code}/options/finance/generalledger/accounts
func (c *Client) AddGeneralLedgerAccount(payload AddGeneralLedgerAccountRequest) (GeneralLedgerAccountResponse, error) {
	// TODO: Implementation
	return GeneralLedgerAccountResponse{}, nil
}

// op: GetGeneralLedgerAccountByCode, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}
func (c *Client) GetGeneralLedgerAccount(accountCode string) (GeneralLedgerAccountResponse, error) {
	// TODO: Implementation
	return GeneralLedgerAccountResponse{}, nil
}

// op: UpdateGeneralLedgerAccount, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}
func (c *Client) UpdateGeneralLedgerAccount(accountCode string, payload UpdateGeneralLedgerAccountRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchGeneralLedgerAccount, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}
func (c *Client) PatchGeneralLedgerAccount(accountCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAccountBudgetsByCode, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/budgets
func (c *Client) GetAccountBudgets(accountCode string) ([]AccountBudgetResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddGeneralLedgerAccountBudget, , path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/budgets
func (c *Client) AddGeneralLedgerAccountBudget(accountCode string, payload AddAccountBudgetRequest) (AccountBudgetResponse, error) {
	// TODO: Implementation
	return AccountBudgetResponse{}, nil
}

// op: UpdateGeneralLedgerAccountBudget, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/budgets
func (c *Client) UpdateGeneralLedgerAccountBudget(accountCode string, payload UpdateAccountBudgetRequest) error {
	// TODO: Implementation
	return nil
}

// op: GetAccountBalancesByCode, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/balances
func (c *Client) GetAccountBalances(accountCode string) ([]AccountBalanceResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAccountTransactionsByCode, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/transactions/{year_num}/{period_num}
func (c *Client) GetAccountTransactions(accountCode string, yearNumber string, periodNumber string) ([]AccountTransactionResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAccountTransactionsByPeriod, path: /{cmpy_code}/options/finance/generalledger/accounts/transactions/{year_num}/{period_num}
func (c *Client) GetAllAccountTransactions(yearNumber string, periodNumber string) ([]AccountTransactionResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAccountReportingCodesByCode, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/reportingcodes
func (c *Client) GetAccountReportingCodes(accountCode string) (AccountReportingCodesResponse, error) {
	// TODO: Implementation
	return AccountReportingCodesResponse{}, nil
}

// op: UpdateAccountReportingCodes, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/reportingcodes
func (c *Client) UpdateAccountReportingCodes(accountCode string, payload UpdateAccountReportingCodesRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchAccountReportingCodes, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/reportingcodes
func (c *Client) PatchAccountReportingCodes(accountCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetAccountResponsibilitiesByCode, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/responsibilities
func (c *Client) GetAccountResponsibilities(accountCode string) ([]AccountResponsibilityResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddGeneralLedgerAccountResponsibility, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/responsibilities
func (c *Client) AddAccountResponsibility(accountCode string, payload AddAccountResponsibilityRequest) (AccountResponsibilityResponse, error) {
	// TODO: Implementation
	return AccountResponsibilityResponse{}, nil
}

// op: UpdateAccountResponsibility, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/responsibilities/{source_flg}/{user_code}
func (c *Client) UpdateAccountResponsibility(accountCode string, userType string, userCode string, payload UpdateAccountResponsibilityRequest) error {
	// TODO: Implementation
	return nil
}

// op: DeleteAccountResponsibility, path: /{cmpy_code}/options/finance/generalledger/accounts/{acct_code}/responsibilities/{source_flg}/{user_code}
func (c *Client) DeleteAccountResponsibility(accountCode string, userType string, userCode string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllJournals, path: /{cmpy_code}/options/finance/generalledger/journals
func (c *Client) GetAllJournals() ([]JournalResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddGeneralLedgerTaxJournal, path: /{cmpy_code}/options/finance/generalledger/journals/tax
func (c *Client) AddGeneralLedgerTaxJournal(payload AddTaxJournalRequest) (AddTaxJournalResponse, error) {
	// TODO: Implementation
	return AddTaxJournalResponse{}, nil
}

// op: UpdateGeneralLedgerTaxJournal, path: /{cmpy_code}/options/finance/generalledger/journals/tax/{jour_num}
func (c *Client) UpdateGeneralLedgerTaxJournal(journalNumber string, payload UpdateTaxJournalRequest) error {
	// TODO: Implementation
	return nil
}

// op: AddGeneralLedgerGeneralJournalpath: /{cmpy_code}/options/finance/generalledger/journals/general
func (c *Client) AddGeneralLedgerGeneralJournal(payload AddGeneralJournalRequest) (AddGeneralJournalResponse, error) {
	// TODO: Implementation
	return AddGeneralJournalResponse{}, nil
}

// op: UpdateGeneralLedgerGeneralJournalpath: /{cmpy_code}/options/finance/generalledger/journals/general/{jour_num}
func (c *Client) UpdateGeneralLedgerGeneralJournal(journalNumber string, payload UpdateGeneralJournalRequest) error {
	// TODO: Implementation
	return nil
}
