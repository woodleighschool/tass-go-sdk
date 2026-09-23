package tassfinance

import "time"

// Purchasing

type PurchaseOrderStatus string

const (
	OutstandingPO PurchaseOrderStatus = `O`
	PartialPO     PurchaseOrderStatus = `P`
	CancelledPO   PurchaseOrderStatus = `C`
	CompletePO    PurchaseOrderStatus = `F`
)

type PurchaseOrderResponse struct {
	SupplierCode       *string                              `json:"vend_code,omitempty"`
	OrderNumber        *int                                 `json:"order_num,omitempty"`
	Year               *int                                 `json:"year_num,omitempty"`
	Period             *int                                 `json:"period_num,omitempty"`
	OrderDate          *time.Time                           `json:"order_date,omitempty"`
	DueDate            *time.Time                           `json:"due_date,omitempty"`
	Status             PurchaseOrderStatus                  `json:"status_ind,omitempty"`
	Printed            bool                                 `json:"printed_flag"`
	AuthorisationCode  *string                              `json:"authorise_code,omitempty"`
	WarehouseCode      *string                              `json:"ware_code,omitempty"`
	Comment            *string                              `json:"comment_text,omitempty"`
	CancelDate         *time.Time                           `json:"cancel_date,omitempty"`
	RequisitionNumber  *int                                 `json:"req_num,omitempty"`
	HasAttachments     bool                                 `json:"has_attachments"`
	SupplierDetails    PurchaseOrderSupplierDetailsResponse `json:"supplier_details"`
	DeliveryDetails    PurchaseOrderDeliveryDetailsResponse `json:"delivery_details"`
	UDFields           PurchaseOrderUDFieldsResponse        `json:"ud_fields"`
	GeneralLedgerLines []PurchaseOrderGLLineResponse        `json:"gl_lines"`
}

type PurchaseOrderSupplierDetailsResponse struct {
	ContactName      *string `json:"supplier_contact_name,omitempty"`
	SupplierName1    *string `json:"supplier_name_text,omitempty"`
	SupplierName2    *string `json:"supplier_name_text2,omitempty"`
	AddressLine1     *string `json:"supplier_addr_line1,omitempty"`
	AddressLine2     *string `json:"supplier_addr_line2,omitempty"`
	AddressLine3     *string `json:"supplier_addr_line3,omitempty"`
	Suburb           *string `json:"supplier_town_suburb,omitempty"`
	State            *string `json:"supplier_state_code,omitempty"`
	PostCode         *string `json:"supplier_postcode,omitempty"`
	Country          *string `json:"supplier_country_text,omitempty"`
	PaymentTermsCode *string `json:"term_code,omitempty"`
	TaxCode          *string `json:"tax_code,omitempty"`
	ABN              *string `json:"abn_text,omitempty"`
}

type PurchaseOrderDeliveryDetailsResponse struct {
	DeliveryName *string `json:"delivery_name_text,omitempty"`
	AddressLine1 *string `json:"delivery_addr_line1,omitempty"`
	AddressLine2 *string `json:"delivery_addr_line2,omitempty"`
	AddressLine3 *string `json:"delivery_addr_line3,omitempty"`
	AddressLine4 *string `json:"delivery_addr_line4,omitempty"`
	Country      *string `json:"delivery_country_text,omitempty"`
}

type PurchaseOrderUDFieldsResponse struct {
	Flags PurchaseOrderUDFlagsResponse `json:"ud_flags"`
	Text  PurchaseOrderUDTextResponse  `json:"ud_text"`
}

type PurchaseOrderUDFlagsResponse struct {
	UD1Flag bool `json:"ud1_flg"`
	UD2Flag bool `json:"ud2_flg"`
	UD3Flag bool `json:"ud3_flg"`
}

type PurchaseOrderUDTextResponse struct {
	UD4Text *string `json:"ud4_text,omitempty"`
	UD5Text *string `json:"ud5_text,omitempty"`
	UD6Text *string `json:"ud6_text,omitempty"`
}

type POTaxType string

const (
	POTax         POTaxType = `T`
	POWithholding POTaxType = `W`
)

type PurchaseOrderGLLineResponse struct {
	Line                 *int      `json:"line_num"`
	Reference            *string   `json:"ref_text,omitempty"`
	OEM                  *string   `json:"oem_text,omitempty"`
	Description          *string   `json:"desc_text,omitempty"`
	AccountCode          string    `json:"acct_code"`
	OrderQuantity        *float64  `json:"order_qty,omitempty"`
	UnitCost             *float64  `json:"unit_cost_amt,omitempty"`
	UnitTax              *float64  `json:"unit_tax_amt,omitempty"`
	TaxCode              *string   `json:"tax_code,omitempty"`
	InvoicedQuantity     *float64  `json:"invoiced_qty,omitempty"`
	ExtendedCost         *float64  `json:"ext_cost_amt,omitempty"`
	ExtendedTax          *float64  `json:"ext_tax_amt,omitempty"`
	LineTotal            *float64  `json:"line_total_amt,omitempty"`
	TaxType              POTaxType `json:"tax_type,omitempty"`
	TaxPercentage        *float64  `json:"tax_per,omitempty"`
	GLAccountDescription *string   `json:"gl_acct_desc"`
}

type AddPurchaseOrderRequest struct {
	SupplierCode       string                        `json:"vend_code" validate:"max=8"`
	OrderDate          string                        `json:"order_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	WarehouseCode      string                        `json:"ware_code" validate:"max=3"`
	GeneralLedgerLines []PurchaseOrderGLLineResponse `json:"gl_lines"`

	DueDate           *string                              `json:"due_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Printed           bool                                 `json:"printed_flag"`
	AuthorisationCode *string                              `json:"authorise_code,omitempty" validate:"max=8"`
	Comment           *string                              `json:"comment_text,omitempty" validate:"max=250"`
	DeliveryDetails   PurchaseOrderDeliveryDetailsResponse `json:"delivery_details"`
	UDFields          PurchaseOrderUDFieldsResponse        `json:"ud_fields"`
}

type UpdatePurchaseOrderRequest struct {
	OrderDate string `json:"order_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format

	SupplierCode       *string                                    `json:"vend_code,omitempty" validate:"max=8"`
	DueDate            *string                                    `json:"due_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	AuthorisationCode  *string                                    `json:"authorise_code,omitempty" validate:"max=8"`
	Printed            bool                                       `json:"printed_flag"`
	WarehouseCode      *string                                    `json:"ware_code,omitempty" validate:"max=3"`
	Comment            *string                                    `json:"comment_text,omitempty" validate:"max=250"`
	DeliveryDetails    *UpdatePurchaseOrderDeliveryDetailsRequest `json:"delivery_details,omitempty"`
	UDFields           *UpdatePurchaseOrderUDFieldsRequest        `json:"ud_fields,omitempty"`
	GeneralLedgerLines []UpdatePurchaseOrderGLLineRequest         `json:"gl_lines,omitempty"`
}

type UpdatePurchaseOrderDeliveryDetailsRequest struct {
	Name         *string `json:"delivery_name_text,omitempty" validate:"max=40"`
	AddressLine1 *string `json:"delivery_addr_line1,omitempty" validate:"max=40"`
	AddressLine2 *string `json:"delivery_addr_line2,omitempty" validate:"max=40"`
	AddressLine3 *string `json:"delivery_addr_line3,omitempty" validate:"max=40"`
	AddressLine4 *string `json:"delivery_addr_line4,omitempty" validate:"max=40"`
	Country      *string `json:"delivery_country_text,omitempty" validate:"max=40"`
}

type UpdatePurchaseOrderUDFieldsRequest struct {
	Flags *UpdatePurchaseOrderUDFlagsRequest `json:"ud_flags,omitempty"`
	Text  *UpdatePurchaseOrderUDTextRequest  `json:"ud_text,omitempty"`
}

type UpdatePurchaseOrderUDFlagsRequest struct {
	UD1Flag bool `json:"ud1_flg"`
	UD2Flag bool `json:"ud2_flg"`
	UD3Flag bool `json:"ud3_flg"`
}

type UpdatePurchaseOrderUDTextRequest struct {
	UD4Text *string `json:"ud4_text,omitempty" validate:"max=20"`
	UD5Text *string `json:"ud5_text,omitempty" validate:"max=20"`
	UD6Text *string `json:"ud6_text,omitempty" validate:"max=20"`
}

type UpdatePurchaseOrderGLLineRequest struct {
	LineNumber        int      `json:"line_num"`
	Reference         *string  `json:"ref_text,omitempty" validate:"max=25"`
	SupplierReference *string  `json:"oem_text,omitempty" validate:"max=15"`
	Description       *string  `json:"desc_text,omitempty" validate:"max=4000"`
	AccountCode       *string  `json:"acct_code,omitempty" validate:"max=18"`
	OrderQuantity     *float64 `json:"order_qty,omitempty"`
	UnitCost          *float64 `json:"unit_cost_amt,omitempty"`
	UnitTax           *float64 `json:"unit_tax_amt,omitempty"`
	TaxCode           *string  `json:"tax_code,omitempty" validate:"max=3"`
}

// General Ledger

type CodeType string

const (
	SegmentType CodeType = `S`
	FixedType   CodeType = `F`
	AnyType     CodeType = `C`
)

type GLAccountOptionsResponse struct {
	Code        string `json:"code"`
	Description string `json:"desc"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
	StartPeriod int    `json:"start_perioud"`
	EndPeriod   int    `json:"end_period"`
}

type YearPeriodOptionsResponse struct {
	Year      *int       `json:"year_num,omitempty"`
	Period    *int       `json:"period_num,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

type CodeFormatRulesResponse struct {
	CompanyCode   string                  `json:"cmpy_code"`
	CodeStructure []CodeStructureResponse `json:"code_structure"`
}

type CodeStructureResponse struct {
	StartPosition           int               `json:"start_num"`
	Length                  int               `json:"length_num"`
	Description             string            `json:"desc_text"`
	Default                 string            `json:"default_text"`
	Type                    CodeType          `json:"type_ind"`
	DepartmentConsolidation bool              `json:"dept_consol_flg"`
	Segments                []SegmentResponse `json:"segments"`
}

type SegmentResponse struct {
	FlexCode    string `json:"flex_code"`
	Description string `json:"desc_text"`
	GroupCode   string `json:"group_code"`
}

type ReportingCodeResponse struct {
	CompanyCode  string                             `json:"cmpy_code"`
	FieldDetails []ReportingCodeOptionsFieldDetails `json:"field_details"`
}

type ReportingCodeOptionsFieldDetails struct {
	Name            string                                    `json:"field_name"`
	Description     string                                    `json:"field_desc"`
	ReferenceValues []ReportingCodeOptionsFieldReferenceValue `json:"reference_values"`
}

type ReportingCodeOptionsFieldReferenceValue struct {
	Code        string `json:"ref_code"`
	Description string `json:"ref_desc"`
}

type GeneralLedgerAccountResponse struct {
	CompanyCode    string  `json:"cmpy_code"`
	AccountCode    string  `json:"acct_code"`
	Description    string  `json:"desc_text"`
	StartYear      int     `json:"start_year_num"`
	StartPeriod    int     `json:"start_period_num"`
	EndYear        int     `json:"end_year_num"`
	EndPeriod      int     `json:"end_period_num"`
	GroupCode      *string `json:"group_code,omitempty"`
	ExternalCode   *string `json:"external_code,omitempty"`
	Type           string  `json:"type_ind"`
	DefaultTaxCode string  `json:"def_tax_code"`
}

type AddGeneralLedgerAccountRequest struct {
	AccountCode string `json:"acct_code" validate:"max=18"`
	Description string `json:"desc_text" validate:"max=40"`
	StartYear   int    `json:"start_year_num"`
	StartPeriod int    `json:"start_period_num"`
	EndYear     int    `json:"end_year_num"`
	EndPeriod   int    `json:"end_period_num"`
	Type        string `json:"type_ind" validate:"max=1"`

	GroupCode      *string `json:"group_code,omitempty" validate:"max=7"`
	ExternalCode   *string `json:"external_code,omitempty" validate:"max=20"`
	DefaultTaxCode *string `json:"def_tax_code,omitempty" validate:"max=3"`
}

type UpdateGeneralLedgerAccountRequest struct {
	AccountCode string `json:"acct_code" validate:"max=18"`
	Description string `json:"desc_text" validate:"max=40"`
	StartYear   int    `json:"start_year_num"`
	StartPeriod int    `json:"start_period_num"`
	EndYear     int    `json:"end_year_num"`
	EndPeriod   int    `json:"end_period_num"`
	Type        string `json:"type_ind" validate:"max=1"`

	GroupCode      *string `json:"group_code,omitempty" validate:"max=7"`
	ExternalCode   *string `json:"external_code,omitempty" validate:"max=20"`
	DefaultTaxCode *string `json:"def_tax_code,omitempty" validate:"max=3"`
}

type AccountBudgetResponse struct {
	CompanyCode  string                 `json:"cmpy_code"`
	Code         string                 `json:"acct_code"`
	Year         int                    `json:"year_num"`
	BudgetNumber int                    `json:"budget_num"`
	Periods      []PeriodBudgetResponse `json:"periods"`
}

type PeriodBudgetResponse struct {
	Period                   int     `json:"period_num"`
	Amount                   float64 `json:"period_amt"`
	YearToDatePreCloseAmount float64 `json:"ytd_pre_close_amt"`
	YearToDateBudget         float64 `json:"ytd_budget"`
	Variance                 float64 `json:"variance"`
}

type AddAccountBudgetRequest struct {
	Year         int    `json:"year_num"`
	BudgetNumber string `json:"budget_num"`

	Periods []PeriodBudgetRequest `json:"periods"`
}

type PeriodBudgetRequest struct {
	Period int     `json:"period_num"`
	Amount float64 `json:"period_amt"`
}

type UpdateAccountBudgetRequest struct {
	Year         int            `json:"year_num"`
	BudgetNumber string         `json:"budget_num"`
	Periods      []BudgetPeriod `json:"periods"`
}

type BudgetPeriod struct {
	Period int     `json:"period_num"`
	Amount float64 `json:"period_amt"`
}
type AccountBalanceResponse struct {
	CompanyCode    string                  `json:"cmpy_code"`
	AccountCode    string                  `json:"acct_code"`
	Year           int                     `json:"year_num"`
	PeriodBalances []PeriodBalanceResponse `json:"period_bals"`
}

type PeriodBalanceResponse struct {
	Period         int     `json:"period_num"`
	OpenAmount     float64 `json:"open_amt"`
	DebitAmount    float64 `json:"debit_amt"`
	CreditAmount   float64 `json:"credit_amt"`
	CloseAmount    float64 `json:"close_amt"`
	PreCloseAmount float64 `json:"pre_close_amt"`
}

type AccountTransactionResponse struct {
	CompanyCode           string    `json:"cmpy_code"`
	AccountCode           string    `json:"acct_code"`
	Year                  int       `json:"year_num"`
	Period                int       `json:"period_num"`
	JournalCode           *string   `json:"jour_code,omitempty"`
	JournalNumber         int       `json:"jour_num"`
	JournalSequenceNumber int       `json:"jour_seq_num"`
	Analysis              *string   `json:"analysis_text,omitempty"`
	Date                  time.Time `json:"tran_date"`
	Source                *string   `json:"ref_text"`
	ReferenceNumber       int       `json:"ref_num"`
	Description           *string   `json:"desc_text,omitempty"`
	DebitAmount           *float64  `json:"debit_amt,omitempty"`
	CreditAmount          *float64  `json:"credit_amt,omitempty"`
}

type AccountReportingCodesResponse struct {
	CompanyCode           string         `json:"cmpy_code"`
	AccountCode           string         `json:"acct_code"`
	AssociatedAccountCode *string        `json:"assoc_acct_code,omitempty"`
	ResponsibilityName    *string        `json:"resp_name,omitempty"`
	ResponsibilityEmail   *string        `json:"resp_e_mail,omitempty"`
	ReportingCodes        ReportingCodes `json:"rpt_codes"`
}

type ReportingCodes struct {
	RptCode1  *string `json:"rpt1_code,omitempty"`
	RptCode2  *string `json:"rpt2_code,omitempty"`
	RptCode3  *string `json:"rpt3_code,omitempty"`
	RptCode4  *string `json:"rpt4_code,omitempty"`
	RptCode5  *string `json:"rpt5_code,omitempty"`
	RptCode6  *string `json:"rpt6_code,omitempty"`
	RptCode7  *string `json:"rpt7_code,omitempty"`
	RptCode8  *string `json:"rpt8_code,omitempty"`
	RptCode9  *string `json:"rpt9_code,omitempty"`
	RptCode10 *string `json:"rpt10_code,omitempty"`
	RptCode11 *string `json:"rpt11_code,omitempty"`
	RptCode12 *string `json:"rpt12_code,omitempty"`
	RptCode13 *string `json:"rpt13_code,omitempty"`
	RptCode14 *string `json:"rpt14_code,omitempty"`
	RptCode15 *string `json:"rpt15_code,omitempty"`
	RptCode16 *string `json:"rpt16_code,omitempty"`
	RptCode17 *string `json:"rpt17_code,omitempty"`
	RptCode18 *string `json:"rpt18_code,omitempty"`
	RptCode19 *string `json:"rpt19_code,omitempty"`
	RptCode20 *string `json:"rpt20_code,omitempty"`
}

type UpdateAccountReportingCodesRequest struct {
	AssociatedAccountCode *string        `json:"assoc_acct_code,omitempty" validate:"max=18"`
	ResponsibilityName    *string        `json:"resp_name,omitempty" validate:"max=40"`
	ResponsibilityEmail   *string        `json:"resp_e_mail,omitempty" validate:"max=40"`
	ReportingCodes        ReportingCodes `json:"rpt_codes"`
}

type SourceFlag string

const (
	EmployeeSource SourceFlag = `E`
	TeacherSource  SourceFlag = `T`
)

type AccountResponsibilityResponse struct {
	CompanyCode   string     `json:"cmpy_code"`
	AccountCode   string     `json:"acct_code"`
	UserCode      string     `json:"user_code"`
	Source        SourceFlag `json:"source_flg"`
	ApprovalLevel int        `json:"resp_flg"`
}

type AddAccountResponsibilityRequest struct {
	UserCode      string     `json:"user_code" validate:"max=7"`
	Source        SourceFlag `json:"source_flg"`
	ApprovalLevel int        `json:"resp_flg"`
}

type UpdateAccountResponsibilityRequest struct {
	UserCode      string     `json:"user_code" validate:"max=7"`
	Source        SourceFlag `json:"source_flg"`
	ApprovalLevel int        `json:"resp_flg"`
}

type JournalTypeOptionsResponse struct {
	Code          string  `json:"code"`
	Description   *string `json:"desc,omitempty"`
	GeneralLedger bool    `json:"gl_flag"`
}

type JournalResponse struct {
	CompanyCode        string           `json:"cmpy_code"`
	Code               string           `json:"jour_code"`
	Number             int              `json:"jour_num"`
	Date               *time.Time       `json:"jour_date,omitempty"`
	Posted             bool             `json:"post_flag"`
	Year               int              `json:"year_num"`
	Period             int              `json:"period_num"`
	Comment1           *string          `json:"com1_text,omitempty"`
	Comment2           *string          `json:"com2_text,omitempty"`
	ControlAmount      float64          `json:"control_amt"`
	BankCode           *string          `json:"bank_code,omitempty"`
	BankDescription    *string          `json:"bank_desc,omitempty"`
	BankGLAccount      *string          `json:"bank_gl_acct,omitempty"`
	Warnings           []string         `json:"warnings"`
	GeneralLedgerLines []GLLineResponse `json:"gl_lines"`
}

type GLLineResponse struct {
	AccountCode          string   `json:"acct_code"`
	SequenceNumber       int      `json:"seq_num"`
	ReferenceText        *string  `json:"ref_text,omitempty"`
	ReferenceNumber      *int     `json:"ref_num,omitempty"`
	AnalysisText         *string  `json:"analysis_text,omitempty"`
	Description          *string  `json:"desc_text,omitempty"`
	DebitAmount          float64  `json:"debit_amt"`
	CreditAmount         float64  `json:"credit_amt"`
	TaxCode              *string  `json:"tax_code,omitempty"`
	TaxType              *string  `json:"tax_type,omitempty"`
	TaxAmount            *float64 `json:"tax_amt,omitempty"`
	TaxPercentage        *float64 `json:"tax_per,omitempty"`
	TaxNet               *float64 `json:"tax_net,omitempty"`
	GLAccountDescription *string  `json:"gl_acct_desc,omitempty"`
	Warnings             []string `json:"warnings"`
}

type AddTaxJournalRequest struct {
	JournalDate string `json:"jour_date"`
	Year        int    `json:"year_num"`
	Period      int    `json:"period_num"`

	Comment1           *string          `json:"com1_text,omitempty"`
	Comment2           *string          `json:"com2_text,omitempty"`
	ControlAmount      *float64         `json:"control_amt,omitempty"`
	GeneralLedgerLines []TaxJournalLine `json:"gl_lines"`
}

type TaxJournalLine struct {
	AccountCode  string  `json:"acct_code" validate:"max=18"`
	Description  string  `json:"desc_text" validate:"max=4000"`
	DebitAmount  float64 `json:"debit_amt"`
	CreditAmount float64 `json:"credit_amt"`
	TaxType      string  `json:"tax_type" validate:"max=1"`
	TaxCode      string  `json:"tax_code" validate:"max=3"`

	Reference *string `json:"ref_text,omitempty" validate:"max=10"`
	Analysis  *string `json:"analysis_text,omitempty" validate:"max=16"`
	TaxAmount float64 `json:"tax_amt,omitempty"`
}

type AddTaxJournalResponse struct {
	CompanyCode   string `json:"cmpy_code"`
	JournalNumber int    `json:"jour_num"`
}

type UpdateTaxJournalRequest struct {
	JournalDate string `json:"jour_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Year        int    `json:"year_num"`
	Period      int    `json:"period_num"`

	Comment1           *string        `json:"com1_text,omitempty"`
	Comment2           *string        `json:"com2_text,omitempty"`
	ControlAmount      *float64       `json:"control_amt,omitempty"`
	GeneralLedgerLines TaxJournalLine `json:"gl_lines"`
}

type AddGeneralJournalRequest struct {
	Code               string                `json:"jour_code"`
	Date               string                `json:"jour_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Year               int                   `json:"year_num"`
	Period             int                   `json:"period_num"`
	Comment1           *string               `json:"com1_text,omitempty"`
	Comment2           *string               `json:"com2_text,omitempty"`
	ControlAmount      *float64              `json:"control_amt,omitempty"`
	GeneralLedgerLines []StandardJournalLine `json:"gl_lines"`
}

type StandardJournalLine struct {
	AccountCode  string  `json:"acct_code" validate:"max=18"`
	Description  string  `json:"desc_text" validate:"max=4000"`
	DebitAmount  float64 `json:"debit_amt"`
	CreditAmount float64 `json:"credit_amt"`

	Reference *string `json:"ref_text,omitempty" validate:"max=10"`
	Analysis  *string `json:"analysis_text,omitempty" validate:"max=16"`
}

type AddGeneralJournalResponse struct {
	CompanyCode   string `json:"cmpy_code"`
	JournalNumber int    `json:"jour_num"`
}

type UpdateGeneralJournalRequest struct {
	Code   string `json:"jour_code"`
	Date   string `json:"jour_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Year   int    `json:"year_num"`
	Period int    `json:"period_num"`

	Comment1           *string               `json:"com1_text,omitempty"`
	Comment2           *string               `json:"com2_text,omitempty"`
	ControlAmount      *float64              `json:"control_amt,omitempty"`
	GeneralLedgerLines []StandardJournalLine `json:"gl_lines"`
}

// Accounts Payable

type SupplierCreditResponse struct {
	CompanyCode              string                                 `json:"cmpy_code"`
	SupplierCode             string                                 `json:"vend_code"`
	SupplierName             *string                                `json:"vend_name,omitempty"`
	DebitID                  *int                                   `json:"debit_num,omitempty"`
	DebitText                *string                                `json:"debit_text,omitempty"`
	DebitDate                *time.Time                             `json:"debit_date,omitempty"`
	TotalCreditAmount        *float64                               `json:"total_amt,omitempty"`
	TotalDistributedAmount   *float64                               `json:"dist_amt,omitempty"`
	TotalUndistributedAmount *float64                               `json:"undist_amt,omitempty"`
	TotalAppliedAmount       *float64                               `json:"apply_amt,omitempty"`
	TotalDiscountAmount      *float64                               `json:"disc_amt,omitempty"`
	Year                     *int                                   `json:"year_num,omitempty"`
	Period                   *int                                   `json:"period_num,omitempty"`
	Comment1                 *string                                `json:"com1_text,omitempty"`
	Comment2                 *string                                `json:"com2_text,omitempty"`
	Date                     *time.Time                             `json:"post_date,omitempty"`
	HasAttachments           *bool                                  `json:"has_attachments,omitempty"`
	GeneralLedgerLines       []SupplierCreditGLLineResponse         `json:"gl_lines"`
	AppliedInvoices          []SupplierCreditAppliedInvoiceResponse `json:"applied_invoices"`
}

type SupplierCreditGLLineResponse struct {
	AccountCode        *string  `json:"acct_code,omitempty"`
	AccountDescription *string  `json:"gl_acct_desc,omitempty"`
	LineDescription    *string  `json:"desc_text,omitempty"`
	TaxCode            *string  `json:"tax_code,omitempty"`
	TaxPercentage      *float64 `json:"tax_per,omitempty"`
	GrossAmount        *float64 `json:"dist_amt,omitempty"`
	NetAmount          *float64 `json:"dist_net,omitempty"`
	TaxAmount          *float64 `json:"dist_tax,omitempty"`
}

type SupplierCreditAppliedInvoiceResponse struct {
	InvoiceCode   int        `json:"vouch_code"`
	InvoiceDate   *time.Time `json:"vouch_date,omitempty"`
	InvoiceNumber *string    `json:"inv_text,omitempty"`
	TotalAmount   *float64   `json:"total_amt,omitempty"`
	PaidAmount    *float64   `json:"paid_amt,omitempty"`
	AppliedCredit *float64   `json:"apply_amt,omitempty"`
}

type AddSupplierCreditRequest struct {
	SupplierCode string  `json:"vend_code"`
	DebitText    string  `json:"debit_text"`
	DebitDate    string  `json:"debit_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	TotalAmount  float64 `json:"total_amt"`
	Year         int     `json:"year_num"`
	Period       int     `json:"period_num"`
	Comment1     string  `json:"com1_text" validate:"max=30"`

	Comment2           *string                          `json:"com2_text,omitempty" validate:"max=30"`
	GeneralLedgerLines []AddSupplierCreditGLLineRequest `json:"gl_lines"`
}

type AddSupplierCreditGLLineRequest struct {
	AccountCode string  `json:"acct_code" validate:"max=18"`
	Description string  `json:"desc_text" validate:"max=4000"`
	TaxCode     string  `json:"tax_code" validate:"max=3"`
	GrossAmount float64 `json:"dist_amt"`

	TaxAmount *float64 `json:"dist_tax,omitempty"`
}

type UpdateSupplierCreditRequest struct {
	SupplierCode string  `json:"vend_code" validate:"max=8"`
	DebitText    string  `json:"debit_text" validate:"max=25"`
	DebitDate    string  `json:"debit_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	TotalAmount  float64 `json:"total_amt"`
	Year         int     `json:"year_num"`
	Period       int     `json:"period_num"`
	Comment1     string  `json:"com1_text" validate:"max=30"`

	Comment2           *string                           `json:"com2_text,omitempty" validate:"max=30"`
	GeneralLedgerLines UpdateSupplierCreditGLLineRequest `json:"gl_lines"`
}

type UpdateSupplierCreditGLLineRequest struct {
	AccountCode string  `json:"acct_code" validate:"max=18"`
	Description string  `json:"desc_text" validate:"max=4000"`
	TaxCode     string  `json:"tax_code" validate:"max=3"`
	GrossAmount float64 `json:"dist_amt"`

	TaxAmount *float64 `json:"dist_tax,omitempty"`
}

type ApplySupplierCreditRequest struct {
	AppliedInvoices []CreditApplicationLineRequest `json:"applied_invoices,omitempty"`
}

type CreditApplicationLineRequest struct {
	InvoiceCode int     `json:"vouch_code"`
	ApplyAmount float64 `json:"apply_amt"`
}

type SupplierCreditApplyInvoicesResponse struct {
	AppliedInvoices []SupplierCreditInvoiceDetailsResponse `json:"applied_invoices"`
}

type SupplierCreditInvoiceDetailsResponse struct {
	Code        int        `json:"vouch_code"`
	Date        *time.Time `json:"vouch_date,omitempty"`
	Reference   *string    `json:"inv_text,omitempty"`
	ApplyAmount float64    `json:"apply_amt"`
}

type SupplierCreditUnapplyInvoicesResponse struct {
	UnappliedInvoices []SupplierCreditUnapplyInvoiceDetailsResponse `json:"unapplied_invoices"`
}

type SupplierCreditUnapplyInvoiceDetailsResponse struct {
	Code          int        `json:"vouch_code"`
	Date          *time.Time `json:"vouch_date,omitempty"`
	Reference     *string    `json:"inv_text,omitempty"`
	UnapplyAmount float64    `json:"unapply_amt"`
}

type SupplierInvoiceResponse struct {
	InvoiceCode         int                             `json:"vouch_code"`
	SupplierCode        string                          `json:"vend_code"`
	SupplierName        *string                         `json:"vend_name"`
	SupplierInvoiceCode string                          `json:"inv_text"`
	InvoiceDate         time.Time                       `json:"vouch_date"`
	TotalAmount         float64                         `json:"total_amt"`
	PaidAmount          float64                         `json:"paid_amt"`
	DistributedAmount   float64                         `json:"dist_amt"`
	DueDate             time.Time                       `json:"due_date"`
	PostedToGL          bool                            `json:"post_flag"`
	Year                int                             `json:"year_num"`
	Period              int                             `json:"period_num"`
	PaymentTermsCode    string                          `json:"term_code"`
	Comment1            string                          `json:"com1_text"`
	Comment2            string                          `json:"com2_text"`
	IsOutstanding       bool                            `json:"is_outstanding"`
	IsPOInvoice         bool                            `json:"is_purchase_order_invoice"`
	IsGLInvoice         bool                            `json:"is_gl_invoice"`
	InvoiceLines        []SupplierInvoiceLineResponse   `json:"invoice_lines"`
	AppliedCredits      []SupplierInvoiceCreditResponse `json:"applied_credits"`
}

type SupplierInvoiceLineResponse struct {
	LineNumber          int      `json:"line_num"`
	PONumber            *int     `json:"order_num,omitempty"`
	POLineNumber        *int     `json:"po_line_num,omitempty"`
	AccountCode         *string  `json:"acct_code,omitempty"`
	AccountDescription  *string  `json:"gl_acct_desc,omitempty"`
	Reference           *string  `json:"ref_text,omitempty"`
	SupplierReference   *string  `json:"oem_text,omitempty"`
	Description         *string  `json:"desc_text,omitempty"`
	TaxCode             *string  `json:"tax_code,omitempty"`
	TaxPercentage       *float64 `json:"tax_per,omitempty"`
	OutstandingQuantity *float64 `json:"outstanding_qty,omitempty"`
	OrderQuantity       *float64 `json:"order_qty,omitempty"`
	InvoicedQuantity    *float64 `json:"invoiced_qty,omitempty"`
	UnitCostAmount      *float64 `json:"unit_cost_amt,omitempty"`
	UnitTaxAmount       *float64 `json:"unit_tax_amt,omitempty"`
	ExtendedCodeAmount  *float64 `json:"ext_cost_amt,omitempty"`
	ExtendedTaxAmount   *float64 `json:"ext_tax_amt,omitempty"`
	LineTotalAmount     *float64 `json:"line_total_amt,omitempty"`
}

type SupplierInvoiceCreditResponse struct {
	ID          int       `json:"debit_num"`
	Reference   *string   `json:"debit_text,omitempty"`
	Date        time.Time `json:"credit_date"`
	ApplyAmount float64   `json:"apply_amt"`
}

type SupplierInvoiceHoldPaymentResponse struct {
	CompanyCode  string                                       `json:"cmpy_code"`
	InvoiceCode  int                                          `json:"vouch_code"`
	SupplierCode string                                       `json:"vend_code"`
	HoldCode     *string                                      `json:"hold_code,omitempty"`
	Assignees    []SupplierInvoiceHoldPaymentAssigneeResponse `json:"assignees"`
}

type SupplierInvoiceHoldPaymentAssigneeResponse struct {
	EmployeeCode string `json:"emp_code"`
}

type UpdateSupplierInvoiceHoldPaymentRequest struct {
	HoldCode string `json:"hold_code" validate:"max=2"`

	Assignees []SupplierInvoiceHoldPaymentAssigneeRequest `json:"assignees"`
}

type SupplierInvoiceHoldPaymentAssigneeRequest struct {
	EmployeeCode string `json:"emp_code" validate:"max=7"`
}

type AddSupplierInvoicePurchaseOrderRequest struct {
	SupplierCode        string                                       `json:"vend_code" validate:"max=8"`
	SupplierInvoiceCode string                                       `json:"inv_text" validate:"max=20"`
	Date                string                                       `json:"vouch_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	TotalAmount         float64                                      `json:"total_amt"`
	DueDate             string                                       `json:"due_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Year                int                                          `json:"year_num"`
	Period              int                                          `json:"period_num"`
	PurchaseOrderLines  []AddSupplierInvoicePurchaseOrderLineRequest `json:"po_lines"`

	PaymentTermsCode *string                               `json:"term_code,omitempty" validate:"max=3"`
	Comment1         *string                               `json:"com1_text,omitempty"`
	Comment2         *string                               `json:"com2_text,omitempty"`
	Assignees        []InvoicePurchaseOrderAssigneeRequest `json:"assignees"`
}

type AddSupplierInvoicePurchaseOrderLineRequest struct {
	PONumber       int     `json:"order_num"`
	POLineNumber   int     `json:"po_line_num"`
	AccountCode    string  `json:"acct_code" validate:"max=18"`
	OrderQuantity  float64 `json:"order_qty"`
	TaxCode        string  `json:"tax_code" validate:"max=3"`
	UnitCostAmount float64 `json:"unit_cost_amt"`
	UnitTaxAmount  float64 `json:"unit_tax_amt"`

	Reference         *string `json:"ref_text,omitempty" validate:"max=25"`
	SupplierReference *string `json:"oem_text,omitempty" validate:"max=15"`
	Description       *string `json:"desc_text,omitempty" validate:"max=4000"`
}

type InvoicePurchaseOrderAssigneeRequest struct {
	EmployeeCode string `json:"emp_code" validate:"max=7"`
}

type AddSupplierInvoicePurchaseOrderResponse struct {
	CompanyCode  string `json:"cmpy_code"`
	SupplierCode string `json:"vend_code"`
	InvoiceCode  int    `json:"vouch_code"`
}

type AddSupplierInvoiceGeneralLedgerRequest struct {
	SupplierCode          string                                       `json:"vend_code" validate:"max=8"`
	SupplierInvoiceNumber string                                       `json:"inv_text" validate:"max=20"`
	Date                  string                                       `json:"vouch_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	TotalAmount           float64                                      `json:"total_amt"`
	DueDate               string                                       `json:"due_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Year                  int                                          `json:"year_num"`
	Period                int                                          `json:"period_num"`
	GeneralLedgerLines    []AddSupplierInvoiceGeneralLedgerLineRequest `json:"gl_lines"`

	PaymentTermsCode *string `json:"term_code,omitempty" validate:"max=3"`
	Comment1         *string `json:"com1_text,omitempty"`
	Comment2         *string `json:"com2_text,omitempty"`
}

type AddSupplierInvoiceGeneralLedgerLineRequest struct {
	AccountCode     string  `json:"acct_code" validate:"max=18"`
	TaxCode         string  `json:"tax_code" validate:"max=3"`
	LineTotalAmount float64 `json:"line_total_amt"`

	TaxAmount   *float64 `json:"tax_amt,omitempty"`
	Description *string  `json:"desc_text,omitempty" validate:"max=4000"`
}

type AddSupplierInvoiceGeneralLedgerResponse struct {
	CompanyCode  string `json:"cmpy_code"`
	SupplierCode string `json:"vend_code"`
	InvoiceCode  int    `json:"vouch_code"`
}

type UpdateSupplierInvoiceGeneralLedgerRequest struct {
	SupplierInvoiceNumber *string                                         `json:"inv_text,omitempty" validate:"max=20"`
	Date                  *string                                         `json:"vouch_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	TotalAmount           *float64                                        `json:"total_amt,omitempty"`
	PaymentTermsCode      *string                                         `json:"term_code,omitempty" validate:"max=3"`
	DueDate               *string                                         `json:"due_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Year                  *int                                            `json:"year_num,omitempty"`
	Period                *int                                            `json:"period_num,omitempty"`
	Comment1              *string                                         `json:"com1_text,omitempty"`
	Comment2              *string                                         `json:"com2_text,omitempty"`
	GeneralLedgerLines    []UpdateSupplierInvoiceGeneralLedgerLineRequest `json:"gl_lines"`
}

type UpdateSupplierInvoiceGeneralLedgerLineRequest struct {
	AccountCode     string  `json:"acct_code" validate:"max=18"`
	TaxCode         string  `json:"tax_code" validate:"max=3"`
	LineTotalAmount float64 `json:"line_total_amt"`

	TaxAmount   *float64 `json:"tax_amt,omitempty"`
	Description *string  `json:"desc_text,omitempty" validate:"max=4000"`
}

type UpdateSupplierInvoicePurchaseOrderRequest struct {
	SupplierInvoiceNumber *string                                         `json:"inv_text,omitempty" validate:"max=20"`
	Date                  *string                                         `json:"vouch_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	TotalAmount           *float64                                        `json:"total_amt,omitempty"`
	PaymentTermsCode      *string                                         `json:"term_code,omitempty" validate:"max=3"`
	DueDate               *string                                         `json:"due_date,omitempty" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Year                  *int                                            `json:"year_num,omitempty"`
	Period                *int                                            `json:"period_num,omitempty"`
	Comment1              *string                                         `json:"com1_text,omitempty" validate:"max=30"`
	Comment2              *string                                         `json:"com2_text,omitempty" validate:"max=30"`
	PurchaseOrderLines    []UpdateSupplierInvoicePurchaseOrderLineRequest `json:"po_lines"`
}

type UpdateSupplierInvoicePurchaseOrderLineRequest struct {
	PONumber       int     `json:"order_num"`
	POLineNumber   int     `json:"po_line_num"`
	AccountCode    string  `json:"acct_code" validate:"max=18"`
	OrderQuantity  float64 `json:"order_qty"`
	TaxCode        string  `json:"tax_code" validate:"max=3"`
	UnitCostAmount float64 `json:"unit_cost_amt"`
	UnitTaxAmount  float64 `json:"unit_tax_amt"`

	Reference         *string `json:"ref_text,omitempty" validate:"max=25"`
	SupplierReference *string `json:"oem_text,omitempty" validate:"max=15"`
	Description       *string `json:"desc_text,omitempty" validate:"max=4000"`
}

type SupplierResponse struct {
	Code             *string `json:"vend_code,omitempty"`
	Name             *string `json:"name_text,omitempty"`
	Name2            *string `json:"name_text2,omitempty"`
	AddressLine1     *string `json:"addr1_text,omitempty"`
	AddressLine2     *string `json:"addr2_text,omitempty"`
	AddressLine3     *string `json:"addr3_text,omitempty"`
	City             *string `json:"city_text,omitempty"`
	State            *string `json:"state_code,omitempty"`
	PostCode         *string `json:"post_code,omitempty"`
	Country          *string `json:"country_text,omitempty"`
	WebAddress       *string `json:"web_address,omitempty"`
	Type             *string `json:"type_code,omitempty"`
	PaymentTermsCode *string `json:"term_code,omitempty"`
	TaxCode          *string `json:"tax_code,omitempty"`
	HoldCode         *string `json:"hold_code,omitempty"`
	ABN              *string `json:"abn_text,omitempty"`
	VATNumber        *string `json:"vat_number,omitempty"`
	WitholdTax       bool    `json:"withold_tax_ind"`
	Active           bool    `json:"active_flg"`
	Misc             bool    `json:"misc_flg"`
	TeacherKiosk     bool    `json:"tkiosk_flg"`
}

type AddSupplierRequest struct {
	Code             string `json:"vend_code" validate:"max=8"`
	Name             string `json:"name_text" validate:"max=30"`
	Type             string `json:"type_code" validate:"max=3"`
	PaymentTermsCode string `json:"term_code" validate:"max=3"`
	TaxCode          string `json:"tax_code" validate:"max=3"`
	HoldCode         string `json:"hold_code" validate:"max=2"`

	Name2        *string `json:"name_text2,omitempty" validate:"max=30"`
	AddressLine1 *string `json:"addr1_text,omitempty" validate:"max=40"`
	AddressLine2 *string `json:"addr2_text,omitempty" validate:"max=40"`
	AddressLine3 *string `json:"addr3_text,omitempty" validate:"max=40"`
	City         *string `json:"city_text,omitempty" validate:"max=40"`
	State        *string `json:"state_code,omitempty" validate:"max=6"`
	PostCode     *string `json:"post_code,omitempty" validate:"max=10"`
	Country      *string `json:"country_text,omitempty" validate:"max=20"`
	WebAddress   *string `json:"web_address,omitempty" validate:"max=60"`
	ABN          *string `json:"abn_text,omitempty" validate:"max=30"`
	VATNumber    *string `json:"vat_number,omitempty" validate:"max=13"`
	WitholdTax   bool    `json:"withold_tax_ind"`
	Active       bool    `json:"active_flg"`
	Misc         bool    `json:"misc_flg"`
	TeacherKiosk bool    `json:"tkiosk_flg"`
}

type UpdateSupplierRequest struct {
	Name             string `json:"name_text" validate:"max=30"`
	Type             string `json:"type_code" validate:"max=3"`
	PaymentTermsCode string `json:"term_code" validate:"max=3"`
	TaxCode          string `json:"tax_code" validate:"max=3"`
	HoldCode         string `json:"hold_code" validate:"max=2"`

	Code string `json:"vend_code"`

	Name2        *string `json:"name_text2,omitempty" validate:"max=30"`
	AddressLine1 *string `json:"addr1_text,omitempty" validate:"max=40"`
	AddressLine2 *string `json:"addr2_text,omitempty" validate:"max=40"`
	AddressLine3 *string `json:"addr3_text,omitempty" validate:"max=40"`
	City         *string `json:"city_text,omitempty" validate:"max=40"`
	State        *string `json:"state_code,omitempty" validate:"max=6"`
	PostCode     *string `json:"post_code,omitempty" validate:"max=10"`
	Country      *string `json:"country_text,omitempty" validate:"max=20"`
	WebAddress   *string `json:"web_address,omitempty" validate:"max=60"`
	ABN          *string `json:"abn_text,omitempty" validate:"max=30"`
	VATNumber    *string `json:"vat_number,omitempty" validate:"max=13"`
	WitholdTax   bool    `json:"withold_tax_ind"`
	Active       bool    `json:"active_flg"`
	Misc         bool    `json:"misc_flg"`
	TeacherKiosk bool    `json:"tkiosk_flg"`
}

type SupplierContactsResponse struct {
	AccountsPayable ContactDetails `json:"accounts_payable"`
	Purchasing      ContactDetails `json:"purchasing"`
}

type ContactDetails struct {
	Name           *string `json:"contact_text,omitempty"`
	Email          *string `json:"e_mail,omitempty"`
	PhoneExtension *string `json:"extension_text,omitempty"`
	Fax            *string `json:"fax_text,omitempty"`
	Mobile         *string `json:"mobile_text,omitempty"`
	Telephone      *string `json:"tele_text,omitempty"`
}

type UpdateSupplierContactsRequest struct {
	AccountsPayable ContactDetails `json:"accounts_payable"`
	Purchasing      ContactDetails `json:"purchasing"`
}

type UpdateContactDetails struct {
	Name           *string `json:"contact_text,omitempty" validate:"max=30"`
	Email          *string `json:"e_mail,omitempty" validate:"max=60"`
	PhoneExtension *string `json:"extension_text,omitempty" validate:"max=7"`
	Fax            *string `json:"fax_text,omitempty" validate:"max=30"`
	Mobile         *string `json:"mobile_text,omitempty" validate:"max=30"`
	Telephone      *string `json:"tele_text,omitempty" validate:"max=30"`
}

type SupplierAccountInformationResponse struct {
	AccountCode1 *string  `json:"usual_acct_code1,omitempty"`
	Percentage1  *float64 `json:"acct1_percent,omitempty"`
	AccountCode2 *string  `json:"usual_acct_code2,omitempty"`
	Percentage2  *float64 `json:"acct2_percent,omitempty"`
	AccountCode3 *string  `json:"usual_acct_code3,omitempty"`
	Percentage3  *float64 `json:"acct3_percent,omitempty"`
	AccountCode4 *string  `json:"usual_acct_code4,omitempty"`
	Percentage4  *float64 `json:"acct4_percent,omitempty"`
	AccountCode5 *string  `json:"usual_acct_code5,omitempty"`
	Percentage5  *float64 `json:"acct5_percent,omitempty"`
	AccountCode6 *string  `json:"usual_acct_code6,omitempty"`
	Percentage6  *float64 `json:"acct6_percent,omitempty"`
	AccountCode7 *string  `json:"usual_acct_code7,omitempty"`
	Percentage7  *float64 `json:"acct7_percent,omitempty"`
	AccountCode8 *string  `json:"usual_acct_code8,omitempty"`
	Percentage8  *float64 `json:"acct8_percent,omitempty"`
	AccountCode9 *string  `json:"usual_acct_code9,omitempty"`
	Percentage9  *float64 `json:"acct9_percent,omitempty"`
}

type UpdateSupplierAccountInformationRequest struct {
	AccountCode1 *string  `json:"usual_acct_code1,omitempty" validate:"max=18"`
	Percentage1  *float64 `json:"acct1_percent,omitempty"`
	AccountCode2 *string  `json:"usual_acct_code2,omitempty" validate:"max=18"`
	Percentage2  *float64 `json:"acct2_percent,omitempty"`
	AccountCode3 *string  `json:"usual_acct_code3,omitempty" validate:"max=18"`
	Percentage3  *float64 `json:"acct3_percent,omitempty"`
	AccountCode4 *string  `json:"usual_acct_code4,omitempty" validate:"max=18"`
	Percentage4  *float64 `json:"acct4_percent,omitempty"`
	AccountCode5 *string  `json:"usual_acct_code5,omitempty" validate:"max=18"`
	Percentage5  *float64 `json:"acct5_percent,omitempty"`
	AccountCode6 *string  `json:"usual_acct_code6,omitempty" validate:"max=18"`
	Percentage6  *float64 `json:"acct6_percent,omitempty"`
	AccountCode7 *string  `json:"usual_acct_code7,omitempty" validate:"max=18"`
	Percentage7  *float64 `json:"acct7_percent,omitempty"`
	AccountCode8 *string  `json:"usual_acct_code8,omitempty" validate:"max=18"`
	Percentage8  *float64 `json:"acct8_percent,omitempty"`
	AccountCode9 *string  `json:"usual_acct_code9,omitempty" validate:"max=18"`
	Percentage9  *float64 `json:"acct9_percent,omitempty"`
}

type SupplierPaymentInfoResponse struct {
	BankCode        *string `json:"bank_code,omitempty"`
	PaymentType     *string `json:"pay_type,omitempty"`
	BankBSBCode     *string `json:"bank_bsb_code,omitempty"`
	BankSortCode    *string `json:"bank_sort_code,omitempty"`
	BankAccountCode *string `json:"bank_acct_code,omitempty"`
	BankAccountName *string `json:"bank_acct_name,omitempty"`
	BACSReference   *string `json:"bacs_ref,omitempty"`
	OurAccountCode  *string `json:"our_acct_code,omitempty"`
}

type UpdateSupplierPaymentInfoRequest struct {
	BankCode        *string `json:"bank_code,omitempty" validate:"max=9"`
	PaymentType     *string `json:"pay_type,omitempty" validate:"max=2"`
	BankBSBCode     *string `json:"bank_bsb_code,omitempty" validate:"max=7"`
	BankSortCode    *string `json:"bank_sort_code,omitempty" validate:"max=8"`
	BankAccountCode *string `json:"bank_acct_code,omitempty" validate:"max=20"`
	BankAccountName *string `json:"bank_acct_name,omitempty" validate:"max=64"`
	BACSReference   *string `json:"bacs_ref,omitempty" validate:"max=60"`
	OurAccountCode  *string `json:"our_acct_code,omitempty" validate:"max=21"`
}

type SupplierCreditStatusResponse struct {
	CreditLimit                     *float64   `json:"limit_amt,omitempty"`
	AccountBalance                  *float64   `json:"bal_amt,omitempty"`
	HighestBalance                  *float64   `json:"highest_bal_amt,omitempty"`
	CurrentAmount                   *float64   `json:"curr_amt,omitempty"`
	OverdueAmountUnder30Days        *float64   `json:"over1_amt,omitempty"`
	OverdueAmountBetween31And60Days *float64   `json:"over30_amt,omitempty"`
	OverdueAmountBetween61And90Days *float64   `json:"over60_amt,omitempty"`
	OverdueAmountOver90Days         *float64   `json:"over90_amt,omitempty"`
	SetupDate                       *time.Time `json:"setup_date,omitempty"`
	LastDebitDate                   *time.Time `json:"last_debit_date,omitempty"`
	LastPurchaseOrderDate           *time.Time `json:"last_po_date,omitempty"`
	LastInvoiceDate                 *time.Time `json:"last_vouc_date,omitempty"`
	LastPaymentDate                 *time.Time `json:"last_payment_date,omitempty"`
}

type UpdateSupplierCreditStatusRequest struct {
	CreditLimit *float64 `json:"limit_amt,omitempty"`
}

type SupplierNoteResponse struct {
	CompanyCode    string     `json:"cmpy_code"`
	SupplierCode   string     `json:"vend_code"`
	ID             string     `json:"note_uid"` // Must be a UUID
	Date           *time.Time `json:"note_date,omitempty"`
	Category       *string    `json:"note_cat,omitempty"`
	Text           *string    `json:"note_text,omitempty"`
	HasAttachments bool       `json:"has_attachments"`
}

type AddSupplierNoteRequest struct {
	Date     string `json:"note_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Category string `json:"note_cat" validate:"max=3"`
	Text     string `json:"note_text" validate:"max=4000"`
}

type UpdateSupplierNoteRequest struct {
	Date     string `json:"note_date" validate:"datetime=2006-01-02"` // TODO: Confirm date format
	Category string `json:"note_cat" validate:"max=3"`
	Text     string `json:"note_text" validate:"max=4000"`
}
