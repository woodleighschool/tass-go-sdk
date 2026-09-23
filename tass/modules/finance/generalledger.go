package tassfinance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetAllGeneralLedgerAccounts, path=/{cmpy_code}/finance/generalledger/accounts
func (c *Client) GetAllGeneralLedgerAccounts(ctx context.Context) ([]GeneralLedgerAccountResponse, error) {
	var result []GeneralLedgerAccountResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/finance/generalledger/accounts", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddGeneralLedgerAccount, path=/{cmpy_code}finance/generalledger/accounts
func (c *Client) AddGeneralLedgerAccount(ctx context.Context, payload AddGeneralLedgerAccountRequest) (GeneralLedgerAccountResponse, error) {
	var result GeneralLedgerAccountResponse
	if err := tasscommon.Validate(payload); err != nil {
		return GeneralLedgerAccountResponse{}, err
	}
	body, err := c.t.Request(ctx, http.MethodPost, "/finance/generalledger/accounts", nil, payload, http.StatusCreated)
	if err != nil {
		return GeneralLedgerAccountResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return GeneralLedgerAccountResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetGeneralLedgerAccountByCode, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}
func (c *Client) GetGeneralLedgerAccount(ctx context.Context, accountCode string) (GeneralLedgerAccountResponse, error) {
	var result GeneralLedgerAccountResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/%s", accountCode)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return GeneralLedgerAccountResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return GeneralLedgerAccountResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateGeneralLedgerAccount, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}
func (c *Client) UpdateGeneralLedgerAccount(ctx context.Context, accountCode string, payload UpdateGeneralLedgerAccountRequest) error {
	url := fmt.Sprintf("/finance/generalledger/accounts/%s", accountCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.Request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchGeneralLedgerAccount, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}
func (c *Client) PatchGeneralLedgerAccount(ctx context.Context, accountCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/finance/generalledger/accounts/%s", accountCode)
	_, err := c.t.Request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAccountBudgetsByCode, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/budgets
func (c *Client) GetAccountBudgets(ctx context.Context, accountCode string) ([]AccountBudgetResponse, error) {
	var result []AccountBudgetResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/budgets", accountCode)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddGeneralLedgerAccountBudget, , path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/budgets
func (c *Client) AddGeneralLedgerAccountBudget(ctx context.Context, accountCode string, payload AddAccountBudgetRequest) (AccountBudgetResponse, error) {
	var result AccountBudgetResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/budgets", accountCode)
	if err := tasscommon.Validate(payload); err != nil {
		return AccountBudgetResponse{}, err
	}
	body, err := c.t.Request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return AccountBudgetResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return AccountBudgetResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateGeneralLedgerAccountBudget, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/budgets
func (c *Client) UpdateGeneralLedgerAccountBudget(ctx context.Context, accountCode string, payload UpdateAccountBudgetRequest) error {
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/budgets", accountCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.Request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAccountBalancesByCode, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/balances
func (c *Client) GetAccountBalances(ctx context.Context, accountCode string) ([]AccountBalanceResponse, error) {
	var result []AccountBalanceResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/balances", accountCode)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAccountTransactionsByCode, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/transactions/{year_num}/{period_num}
func (c *Client) GetAccountTransactions(ctx context.Context, accountCode string, yearNumber int, periodNumber int) ([]AccountTransactionResponse, error) {
	var result []AccountTransactionResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/transactions/%d/%d", accountCode, yearNumber, periodNumber)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAccountTransactionsByPeriod, path=/{cmpy_code}/finance/generalledger/accounts/transactions/{year_num}/{period_num}
func (c *Client) GetAllAccountTransactions(ctx context.Context, yearNumber int, periodNumber int) ([]AccountTransactionResponse, error) {
	var result []AccountTransactionResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/transactions/%d/%d", yearNumber, periodNumber)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=GetAccountReportingCodesByCode, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/reportingcodes
func (c *Client) GetAccountReportingCodes(ctx context.Context, accountCode string) (AccountReportingCodesResponse, error) {
	var result AccountReportingCodesResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/reportingcodes", accountCode)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return AccountReportingCodesResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return AccountReportingCodesResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateAccountReportingCodes, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/reportingcodes
func (c *Client) UpdateAccountReportingCodes(ctx context.Context, accountCode string, payload UpdateAccountReportingCodesRequest) error {
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/reportingcodes", accountCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.Request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=PatchAccountReportingCodes, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/reportingcodes
func (c *Client) PatchAccountReportingCodes(ctx context.Context, accountCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/reportingcodes", accountCode)
	_, err := c.t.Request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAccountResponsibilitiesByCode, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/responsibilities
func (c *Client) GetAccountResponsibilities(ctx context.Context, accountCode string) ([]AccountResponsibilityResponse, error) {
	var result []AccountResponsibilityResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/responsibilities", accountCode)
	body, err := c.t.Request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddGeneralLedgerAccountResponsibility, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/responsibilities
func (c *Client) AddAccountResponsibility(ctx context.Context, accountCode string, payload AddAccountResponsibilityRequest) (AccountResponsibilityResponse, error) {
	var result AccountResponsibilityResponse
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/responsibilities", accountCode)
	if err := tasscommon.Validate(payload); err != nil {
		return AccountResponsibilityResponse{}, err
	}
	body, err := c.t.Request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return AccountResponsibilityResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return AccountResponsibilityResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateAccountResponsibility, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/responsibilities/{source_flg}/{user_code}
func (c *Client) UpdateAccountResponsibility(ctx context.Context, accountCode string, userType string, userCode string, payload UpdateAccountResponsibilityRequest) error {
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/responsibilities/%s/%s", accountCode, userType, userCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.Request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeleteAccountResponsibility, path=/{cmpy_code}/finance/generalledger/accounts/{acct_code}/responsibilities/{source_flg}/{user_code}
func (c *Client) DeleteAccountResponsibility(ctx context.Context, accountCode string, userType string, userCode string) error {
	url := fmt.Sprintf("/finance/generalledger/accounts/%s/responsibilities/%s/%s", accountCode, userType, userCode)
	_, err := c.t.Request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetAllJournals, path=/{cmpy_code}/finance/generalledger/journals
func (c *Client) GetAllJournals(ctx context.Context) ([]JournalResponse, error) {
	var result []JournalResponse
	body, err := c.t.Request(ctx, http.MethodGet, "/finance/generalledger/journals", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddGeneralLedgerTaxJournal, path=/{cmpy_code}/finance/generalledger/journals/tax
func (c *Client) AddGeneralLedgerTaxJournal(ctx context.Context, payload AddTaxJournalRequest) (AddTaxJournalResponse, error) {
	var result AddTaxJournalResponse
	if err := tasscommon.Validate(payload); err != nil {
		return AddTaxJournalResponse{}, err
	}
	body, err := c.t.Request(ctx, http.MethodPost, "/finance/generalledger/journals/tax", nil, payload, http.StatusCreated)
	if err != nil {
		return AddTaxJournalResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return AddTaxJournalResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdateGeneralLedgerTaxJournal, path=/{cmpy_code}/finance/generalledger/journals/tax/{jour_num}
func (c *Client) UpdateGeneralLedgerTaxJournal(ctx context.Context, journalNumber int, payload UpdateTaxJournalRequest) error {
	url := fmt.Sprintf("/finance/generalledger/journals/tax/%d", journalNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.Request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddGeneralLedgerGeneralJournal(ctx context.Context, payload AddGeneralJournalRequest) (AddGeneralJournalResponse, error) {
	var result AddGeneralJournalResponse
	if err := tasscommon.Validate(payload); err != nil {
		return AddGeneralJournalResponse{}, err
	}
	body, err := c.t.Request(ctx, http.MethodPost, "/finance/generalledger/journals/general", nil, payload, http.StatusCreated)
	if err != nil {
		return AddGeneralJournalResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return AddGeneralJournalResponse{}, err
	}
	return result, nil
}

func (c *Client) UpdateGeneralLedgerGeneralJournal(ctx context.Context, journalNumber int, payload UpdateGeneralJournalRequest) error {
	url := fmt.Sprintf("/finance/generalledger/journals/general/%d", journalNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.Request(ctx, http.MethodPut, url, nil, payload, http.StatusCreated)
	if err != nil {
		return err
	}
	return nil
}
