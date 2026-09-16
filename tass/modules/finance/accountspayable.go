package tassfinance

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetAllSupplierCredits, path: /{cmpy_code}/finance/accountspayable/credits
func (c *Client) GetAllSupplierCredits() ([]SupplierCreditResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddSupplierCredit, path: /{cmpy_code}/finance/accountspayable/credits
func (c *Client) AddSupplierCredit(payload AddSupplierCreditRequest) (SupplierCreditResponse, error) {
	// TODO: Implementation
	return SupplierCreditResponse{}, nil
}

// op: GetSupplierCreditByID, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}
func (c *Client) GetSupplierCreditByID(creditNumber string) (SupplierCreditResponse, error) {
	// TODO: Implementation
	return SupplierCreditResponse{}, nil
}

// op: UpdateSupplierCredit, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}
func (c *Client) UpdateSupplierCredit(creditNumber string, payload UpdateSupplierCreditRequest) error {
	// TODO: Implementation
	return nil
}

// op: ApplySupplierCredit, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/apply
func (c *Client) ApplySupplierCredit(creditNumber string) (ApplySupplerCreditResponse, error) {
	// TODO: Implementation
	return ApplySupplerCreditResponse{}, nil
}

// op: UnapplySupplierCredit, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/unapply
func (c *Client) UnapplySupplierCredit(creditNumber string) (SupplierCreditUnapplyInvoicesResponse, error) {
	// TODO: Implementation
	return SupplierCreditUnapplyInvoicesResponse{}, nil
}

// op: GetSupplierCreditAttachments, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/attachments
func (c *Client) GetSupplierCreditAttachments(creditNumber string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddSupplierCreditAttachment, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/attachments
func (c *Client) AddSupplierCreditAttachment(creditNumber string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadSupplierCreditAttachment, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/attachments/{attach_id}
func (c *Client) GetSupplierCreditAttachment(creditNumber string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteSupplierCreditAttachment, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/attachments/{attach_id}
func (c *Client) DeleteSupplierCreditAttachment(creditNumber string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllAPCreditTaxCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/credits/taxcodes
func (c *Client) GetAllAPCreditTaxCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPCreditYearPeriodOptions, path: /{cmpy_code}/options/finance/accountspayable/credits/yearsperiods
func (c *Client) GetAllAPCreditYearPeriodOptions() ([]YearPeriodOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPCreditGLAccountOptions, path: /{cmpy_code}/options/finance/accountspayable/credits/glaccounts
func (c *Client) GetAllAPCreditGLAccountOptions() ([]GLAccountOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPCreditSupplierOptions, path: /{cmpy_code}/options/finance/accountspayable/credits/suppliers
func (c *Client) GetAllAPCreditSupplierOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllSupplierInvoices, path: /{cmpy_code}/finance/accountspayable/invoices
func (c *Client) GetAllSupplierInvoices() ([]SupplierInvoiceResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPInvoiceYearPeriodOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/yearsperiods
func (c *Client) GetAllAPInvoiceYearPeriodOptions() ([]YearPeriodOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPInvoiceGLAccountOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/glaccounts
func (c *Client) GetAllAPInvoiceGLAccountOptions() ([]GLAccountOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPInvoiceTermCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/termcodes
func (c *Client) GetAllAPInvoiceTermCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPInvoiceTaxCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/taxcodes
func (c *Client) GetAllAPInvoiceTaxCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPInvoiceHoldCodeOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/holdcodes
func (c *Client) GetAllAPInvoiceHoldCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllAPInvoiceApproverOptions, path: /{cmpy_code}/options/finance/accountspayable/invoices/approvers
func (c *Client) GetAllAPInvoiceApproverOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetInvoiceHoldPaymentByID, path: /{cmpy_code}/finance/accountspayable/invoices/{vouch_code}/holdpayment
func (c *Client) GetInvoiceHoldPayment(invoiceNumber string) ([]SupplierInvoiceHoldPaymentResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: UpdateInvoiceHoldPayment, path: /{cmpy_code}/finance/accountspayable/invoices/{vouch_code}/holdpayment
func (c *Client) UpdateInvoiceHoldPayment(invoiceNumber string, payload UpdateSupplierInvoiceHoldPaymentRequest) error {
	// TODO: Implementation
	return nil
}

// op: AddSupplierInvoicePurchaseOrder, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder
func (c *Client) AddSupplierInvoicePurchaseOrder(payload AddSupplierInvoicePurchaseOrderRequest) (AddSupplierInvoicePurchaseOrderResponse, error) {
	// TODO: Implementation
	return AddSupplierInvoicePurchaseOrderResponse{}, nil
}
