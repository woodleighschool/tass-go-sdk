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
func (c *Client) ApplySupplierCredit(creditNumber string, payload ApplySupplierCreditRequest) (SupplierCreditApplyInvoicesResponse, error) {
	// TODO: Implementation
	return SupplierCreditApplyInvoicesResponse{}, nil
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

// op: GetAllSupplierInvoices, path: /{cmpy_code}/finance/accountspayable/invoices
func (c *Client) GetAllSupplierInvoices() ([]SupplierInvoiceResponse, error) {
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

// op: AddSupplierInvoiceGeneralLedger, path: /{cmpy_code}/finance/accountspayable/invoices/generalledger
func (c *Client) AddSupplierInvoiceGeneralLedger(payload AddSupplierInvoiceGeneralLedgerRequest) (AddSupplierInvoiceGeneralLedgerResponse, error) {
	// TODO: Implementation
	return AddSupplierInvoiceGeneralLedgerResponse{}, nil
}

// op: UpdateSupplierInvoiceGeneralLedger, path: /{cmpy_code}/finance/accountspayable/invoices/generalledger/{vouch_code}
func (c *Client) UpdateSupplierInvoiceGeneralLedger(invoiceNumber string, payload UpdateSupplierInvoiceGeneralLedgerRequest) error {
	// TODO: Implementation
	return nil
}

// op: UpdateSupplierInvoicePurchaseOrder, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}
func (c *Client) UpdateSupplierInvoicePurchaseOrder(invoiceNumber string, payload UpdateSupplierInvoicePurchaseOrderRequest) error {
	// TODO: Implementation
	return nil
}

// op: CancelSupplierInvoicePurchaseOrder, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/cancel
func (c *Client) CancelSupplierInvoicePurchaseOrder(invoiceNumber string) error {
	// TODO: Implementation
	return nil
}

// op: CancelSupplierInvoiceGeneralLedger, path: /{cmpy_code}/finance/accountspayable/invoices/generalledger/{vouch_code}/cancel
func (c *Client) CancelSupplierInvoiceGeneralLedger(invoiceNumber string) error {
	// TODO: Implementation
	return nil
}

// op: GetSupplierInvoicePurchaseOrderAttachments, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/attachments
func (c *Client) GetSupplierInvoicePurchaseOrderAttachments(invoiceNumber string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddSupplierInvoicePurchaseOrderAttachment, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/attachments
func (c *Client) AddSupplierInvoicePurchaseOrderAttachments(invoiceNumber string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadSupplierInvoicePurchaseOrderAttachment, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/attachments/{attach_id}
func (c *Client) GetSupplierInvoicePurchaseOrderAttachment(invoiceNumber string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteSupplierInvoicePurchaseOrderAttachment, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/attachments/{attach_id}
func (c *Client) DeleteSupplierInvoicePurchaseOrderAttachment(invoiceNumber string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetSupplierInvoiceByID, path: /{cmpy_code}/finance/accountspayable/invoices/{vouch_code}
func (c *Client) GetSupplierInvoice(invoiceNumber string) (SupplierInvoiceResponse, error) {
	// TODO: Implementation
	return SupplierInvoiceResponse{}, nil
}

// op: GetAllSuppliers, path: /{cmpy_code}/finance/accountspayable/suppliers
func (c *Client) GetAllSuppliers() ([]SupplierResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddSuppliers, path: /{cmpy_code}/finance/accountspayable/suppliers
func (c *Client) AddSuppliers(payload AddSupplierRequest) (SupplierResponse, error) {
	// TODO: Implementation
	return SupplierResponse{}, nil
}

// op: GetSupplierByCode, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}
func (c *Client) GetSupplier(supplierCode string) (SupplierResponse, error) {
	// TODO: Implementation
	return SupplierResponse{}, nil
}

// op: UpdateSupplier, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}
func (c *Client) UpdateSupplier(supplierCode string, payload UpdateSupplierRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchSupplier, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}
func (c *Client) PatchSupplier(supplierCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteSupplier, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}
func (c *Client) DeleteSupplier(supplierCode string) error {
	// TODO: Implementation
	return nil
}

// op: GetSupplierContacts, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/contacts
func (c *Client) GetSupplierContacts(supplierCode string) (SupplierContactsResponse, error) {
	// TODO: Implementation
	return SupplierContactsResponse{}, nil
}

// op: UpdateSupplierContacts, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/contacts
func (c *Client) UpdateSupplierContacts(supplierCode string, payload UpdateSupplierContactsRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchSupplierContacts, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/contacts
func (c *Client) PatchSupplierContacts(supplierCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetSupplierAccountInformation, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/accountinformation
func (c *Client) GetSupplierAccountInformation(supplierCode string) (SupplierAccountInformationResponse, error) {
	// TODO: Implementation
	return SupplierAccountInformationResponse{}, nil
}

// op: UpdateSupplierAccountInformation, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/accountinformation
func (c *Client) UpdateSupplierAccountInformation(supplierCode string, payload UpdateSupplierAccountInformationRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchSupplierAccountInformation, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/accountinformation
func (c *Client) PatchSupplierAccountInformation(supplierCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetSupplierPaymentInfo, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/paymentinformation
func (c *Client) GetSupplierPaymentInfo(supplierCode string) (SupplierPaymentInfoResponse, error) {
	// TODO: Implementation
	return SupplierPaymentInfoResponse{}, nil
}

// op: UpdateSupplierPaymentInfo, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/paymentinformation
func (c *Client) UpdateSupplierPaymentInfo(supplierCode string, payload UpdateSupplierPaymentInfoRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchSupplierPaymentInfo, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/paymentinformation
func (c *Client) PatchSupplierPaymentInfo(supplierCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: GetSupplierCreditStatus, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/creditstatus
func (c *Client) GetSupplierCreditStatus(supplierCode string) (SupplierCreditStatusResponse, error) {
	// TODO: Implementation
	return SupplierCreditStatusResponse{}, nil
}

// op: UpdateSupplierCreditStatus, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/creditstatus
func (c *Client) UpdateSupplierCreditStatus(supplierCode string, payload UpdateSupplierCreditStatusRequest) error {
	// TODO: Implementation
	return nil
}

// op: GetAllSupplierNotes, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes
func (c *Client) GetAllSupplierNotes(supplierCode string) ([]SupplierNoteResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes
func (c *Client) AddSupplierNote(supplierCode string, payload AddSupplierNoteRequest) (SupplierNoteResponse, error) {
	// TODO: Implementation
	return SupplierNoteResponse{}, nil
}

// op: GetSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}
func (c *Client) GetSupplierNote(supplierCode string) (SupplierNoteResponse, error) {
	// TODO: Implementation
	return SupplierNoteResponse{}, nil
}

// op: UpdateSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}
func (c *Client) UpdateSupplierNote(supplierCode string, payload UpdateSupplierNoteRequest) error {
	// TODO: Implementation
	return nil
}

// op: PatchSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}
func (c *Client) PatchSupplierNote(supplierCode string, payload tasscommon.Operation) error {
	// TODO: Implementation
	return nil
}

// op: DeleteSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}
func (c *Client) DeleteSupplierNote(supplierCode string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllSupplierNoteAttachments, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}/attachments
func (c *Client) GetAllSupplierNoteAttachments(supplierCode string, noteID string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddSupplierNoteAttachment, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}/attachments
func (c *Client) AddSupplierNoteAttachment(supplierCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadSupplierNoteAttachment, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}/attachments/{attach_id}
func (c *Client) GetSupplierNoteAttachment(supplierCode string, noteID string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeleteSupplierNoteAttachment, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteSupplierNoteAttachment(supplierCode string, noteID string, attachmentID string) error {
	// TODO: Implementation
	return nil
}
