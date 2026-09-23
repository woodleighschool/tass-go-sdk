package tassfinance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// op: GetAllSupplierCredits, path: /{cmpy_code}/finance/accountspayable/credits
func (c *Client) GetAllSupplierCredits(ctx context.Context) ([]SupplierCreditResponse, error) {
	var result []SupplierCreditResponse
	body, err := c.t.request(ctx, http.MethodGet, "/finance/accountspayable/credits", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, err
}

// op: AddSupplierCredit, path: /{cmpy_code}/finance/accountspayable/credits
func (c *Client) AddSupplierCredit(ctx context.Context, payload AddSupplierCreditRequest) (SupplierCreditResponse, error) {
	var result SupplierCreditResponse
	if err := tasscommon.Validate(payload); err != nil {
		return SupplierCreditResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, "/finance/accountspayable/credits", nil, payload, http.StatusCreated)
	if err != nil {
		return SupplierCreditResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierCreditResponse{}, err
	}
	return result, nil
}

// op: GetSupplierCreditByID, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}
func (c *Client) GetSupplierCredit(ctx context.Context, creditNumber int) (SupplierCreditResponse, error) {
	var result SupplierCreditResponse
	url := fmt.Sprintf("/finance/accountspayable/credits/%d", creditNumber)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusCreated)
	if err != nil {
		return SupplierCreditResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierCreditResponse{}, err
	}
	return result, nil
}

// op: UpdateSupplierCredit, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}
func (c *Client) UpdateSupplierCredit(ctx context.Context, creditNumber int, payload UpdateSupplierCreditRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/credits/%d", creditNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: ApplySupplierCredit, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/apply
func (c *Client) ApplySupplierCredit(ctx context.Context, creditNumber int, payload ApplySupplierCreditRequest) (SupplierCreditApplyInvoicesResponse, error) {
	var result SupplierCreditApplyInvoicesResponse
	url := fmt.Sprintf("/finance/accountspayable/credits/%d/apply", creditNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return SupplierCreditApplyInvoicesResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusOK)
	if err != nil {
		return SupplierCreditApplyInvoicesResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierCreditApplyInvoicesResponse{}, err
	}
	return result, nil
}

// op: UnapplySupplierCredit, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/unapply
func (c *Client) UnapplySupplierCredit(ctx context.Context, creditNumber int) (SupplierCreditUnapplyInvoicesResponse, error) {
	var result SupplierCreditUnapplyInvoicesResponse
	url := fmt.Sprintf("/finance/accountspayable/credits/%d/unapply", creditNumber)
	body, err := c.t.request(ctx, http.MethodPost, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierCreditUnapplyInvoicesResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierCreditUnapplyInvoicesResponse{}, err
	}
	return result, nil
}

// op: GetSupplierCreditAttachments, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/attachments
func (c *Client) GetSupplierCreditAttachments(ctx context.Context, creditNumber int) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/finance/accountspayable/credits/%d/attachments", creditNumber)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddSupplierCreditAttachment, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/attachments
func (c *Client) AddSupplierCreditAttachment(ctx context.Context, creditNumber int, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/finance/accountspayable/credits/%d/attachments", creditNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}

// op: DownloadSupplierCreditAttachment, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/attachments/{attach_id}
func (c *Client) GetSupplierCreditAttachment(ctx context.Context, creditNumber int, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/finance/accountspayable/credits/%d/attachments/%s", creditNumber, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteSupplierCreditAttachment, path: /{cmpy_code}/finance/accountspayable/credits/{debit_num}/attachments/{attach_id}
func (c *Client) DeleteSupplierCreditAttachment(ctx context.Context, creditNumber int, attachmentID string) error {
	url := fmt.Sprintf("/finance/accountspayable/credits/%d/attachments/%s", creditNumber, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllSupplierInvoices, path: /{cmpy_code}/finance/accountspayable/invoices
func (c *Client) GetAllSupplierInvoices(ctx context.Context) ([]SupplierInvoiceResponse, error) {
	var result []SupplierInvoiceResponse
	body, err := c.t.request(ctx, http.MethodGet, "/finance/accountspayable/invoices", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: GetInvoiceHoldPaymentByID, path: /{cmpy_code}/finance/accountspayable/invoices/{vouch_code}/holdpayment
func (c *Client) GetInvoiceHoldPayment(ctx context.Context, invoiceNumber int) (SupplierInvoiceHoldPaymentResponse, error) {
	var result SupplierInvoiceHoldPaymentResponse
	url := fmt.Sprintf("/finance/accountspayable/invoices/%d/holdpayment", invoiceNumber)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierInvoiceHoldPaymentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierInvoiceHoldPaymentResponse{}, err
	}
	return result, nil
}

// op: UpdateInvoiceHoldPayment, path: /{cmpy_code}/finance/accountspayable/invoices/{vouch_code}/holdpayment
func (c *Client) UpdateInvoiceHoldPayment(ctx context.Context, invoiceNumber int, payload UpdateSupplierInvoiceHoldPaymentRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/invoices/%d/holdpayment", invoiceNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: AddSupplierInvoicePurchaseOrder, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder
func (c *Client) AddSupplierInvoicePurchaseOrder(ctx context.Context, payload AddSupplierInvoicePurchaseOrderRequest) (AddSupplierInvoicePurchaseOrderResponse, error) {
	var result AddSupplierInvoicePurchaseOrderResponse
	if err := tasscommon.Validate(payload); err != nil {
		return AddSupplierInvoicePurchaseOrderResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, "/finance/accountspayable/invoices/purchaseorder", nil, payload, http.StatusCreated)
	if err != nil {
		return AddSupplierInvoicePurchaseOrderResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return AddSupplierInvoicePurchaseOrderResponse{}, err
	}
	return result, nil
}

// op: AddSupplierInvoiceGeneralLedger, path: /{cmpy_code}/finance/accountspayable/invoices/generalledger
func (c *Client) AddSupplierInvoiceGeneralLedger(ctx context.Context, payload AddSupplierInvoiceGeneralLedgerRequest) (AddSupplierInvoiceGeneralLedgerResponse, error) {
	var result AddSupplierInvoiceGeneralLedgerResponse
	if err := tasscommon.Validate(payload); err != nil {
		return AddSupplierInvoiceGeneralLedgerResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, "/finance/accountspayable/invoices/generalledger", nil, nil, http.StatusCreated)
	if err != nil {
		return AddSupplierInvoiceGeneralLedgerResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return AddSupplierInvoiceGeneralLedgerResponse{}, err
	}
	return result, nil
}

// op: UpdateSupplierInvoicePurchaseOrder, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}
func (c *Client) UpdateSupplierInvoicePurchaseOrder(ctx context.Context, invoiceNumber int, payload UpdateSupplierInvoicePurchaseOrderRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/invoices/purchaseorder/%d", invoiceNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: UpdateSupplierInvoiceGeneralLedger, path: /{cmpy_code}/finance/accountspayable/invoices/generalledger/{vouch_code}
func (c *Client) UpdateSupplierInvoiceGeneralLedger(ctx context.Context, invoiceNumber int, payload UpdateSupplierInvoiceGeneralLedgerRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/invoices/generalledger/%d", invoiceNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: CancelSupplierInvoicePurchaseOrder, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/cancel
func (c *Client) CancelSupplierInvoicePurchaseOrder(ctx context.Context, invoiceNumber int) error {
	url := fmt.Sprintf("/finance/accountspayable/invoices/purchaseorder/%d/cancel", invoiceNumber)
	_, err := c.t.request(ctx, http.MethodPost, url, nil, nil, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

// op: CancelSupplierInvoiceGeneralLedger, path: /{cmpy_code}/finance/accountspayable/invoices/generalledger/{vouch_code}/cancel
func (c *Client) CancelSupplierInvoiceGeneralLedger(ctx context.Context, invoiceNumber int) error {
	url := fmt.Sprintf("/finance/accountspayable/invoices/generalledger/%d/cancel", invoiceNumber)
	_, err := c.t.request(ctx, http.MethodPost, url, nil, nil, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

// op: GetSupplierInvoicePurchaseOrderAttachments, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/attachments
func (c *Client) GetSupplierInvoicePurchaseOrderAttachments(ctx context.Context, invoiceNumber int) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/finance/accountspayable/invoices/purchaseorder/%d/attachments", invoiceNumber)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddSupplierInvoicePurchaseOrderAttachment, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/attachments
func (c *Client) AddSupplierInvoicePurchaseOrderAttachments(ctx context.Context, invoiceNumber int, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/finance/accountspayable/invoices/purchaseorder/%d/attachments", invoiceNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}

// op: DownloadSupplierInvoicePurchaseOrderAttachment, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/attachments/{attach_id}
func (c *Client) GetSupplierInvoicePurchaseOrderAttachment(ctx context.Context, invoiceNumber int, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/finance/accountspayable/invoices/purchaseorder/%d/attachments/%s", invoiceNumber, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteSupplierInvoicePurchaseOrderAttachment, path: /{cmpy_code}/finance/accountspayable/invoices/purchaseorder/{vouch_code}/attachments/{attach_id}
func (c *Client) DeleteSupplierInvoicePurchaseOrderAttachment(ctx context.Context, invoiceNumber int, attachmentID string) error {
	url := fmt.Sprintf("/finance/accountspayable/invoices/purchaseorder/%d/attachments/%s", invoiceNumber, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetSupplierInvoiceByID, path: /{cmpy_code}/finance/accountspayable/invoices/{vouch_code}
func (c *Client) GetSupplierInvoice(ctx context.Context, invoiceNumber int) (SupplierInvoiceResponse, error) {
	var result SupplierInvoiceResponse
	url := fmt.Sprintf("/finance/accountspayable/invoices/%d", invoiceNumber)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierInvoiceResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierInvoiceResponse{}, err
	}
	return result, nil
}

// op: GetAllSuppliers, path: /{cmpy_code}/finance/accountspayable/suppliers
func (c *Client) GetAllSuppliers(ctx context.Context) ([]SupplierResponse, error) {
	var result []SupplierResponse
	body, err := c.t.request(ctx, http.MethodGet, "/finance/accountspayable/suppliers", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddSuppliers, path: /{cmpy_code}/finance/accountspayable/suppliers
func (c *Client) AddSuppliers(ctx context.Context, payload AddSupplierRequest) (SupplierResponse, error) {
	var result SupplierResponse
	if err := tasscommon.Validate(payload); err != nil {
		return SupplierResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, "/finance/accountspayable/suppliers", nil, payload, http.StatusCreated)
	if err != nil {
		return SupplierResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierResponse{}, err
	}
	return result, nil
}

// op: GetSupplierByCode, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}
func (c *Client) GetSupplier(ctx context.Context, supplierCode string) (SupplierResponse, error) {
	var result SupplierResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s", supplierCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierResponse{}, err
	}
	return result, nil
}

// op: UpdateSupplier, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}
func (c *Client) UpdateSupplier(ctx context.Context, supplierCode string, payload UpdateSupplierRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s", supplierCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchSupplier, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}
func (c *Client) PatchSupplier(ctx context.Context, supplierCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s", supplierCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteSupplier, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}
func (c *Client) DeleteSupplier(ctx context.Context, supplierCode string) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s", supplierCode)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetSupplierContacts, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/contacts
func (c *Client) GetSupplierContacts(ctx context.Context, supplierCode string) (SupplierContactsResponse, error) {
	var result SupplierContactsResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/contacts", supplierCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierContactsResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierContactsResponse{}, err
	}
	return result, nil
}

// op: UpdateSupplierContacts, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/contacts
func (c *Client) UpdateSupplierContacts(ctx context.Context, supplierCode string, payload UpdateSupplierContactsRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/contacts", supplierCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchSupplierContacts, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/contacts
func (c *Client) PatchSupplierContacts(ctx context.Context, supplierCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/contacts", supplierCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetSupplierAccountInformation, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/accountinformation
func (c *Client) GetSupplierAccountInformation(ctx context.Context, supplierCode string) (SupplierAccountInformationResponse, error) {
	var result SupplierAccountInformationResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/accountsinformation", supplierCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierAccountInformationResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierAccountInformationResponse{}, err
	}
	return result, nil
}

// op: UpdateSupplierAccountInformation, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/accountinformation
func (c *Client) UpdateSupplierAccountInformation(ctx context.Context, supplierCode string, payload UpdateSupplierAccountInformationRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/accountinformation", supplierCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchSupplierAccountInformation, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/accountinformation
func (c *Client) PatchSupplierAccountInformation(ctx context.Context, supplierCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/accountinformation", supplierCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetSupplierPaymentInfo, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/paymentinformation
func (c *Client) GetSupplierPaymentInfo(ctx context.Context, supplierCode string) (SupplierPaymentInfoResponse, error) {
	var result SupplierPaymentInfoResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/paymentinformation", supplierCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierPaymentInfoResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierPaymentInfoResponse{}, err
	}
	return result, nil
}

// op: UpdateSupplierPaymentInfo, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/paymentinformation
func (c *Client) UpdateSupplierPaymentInfo(ctx context.Context, supplierCode string, payload UpdateSupplierPaymentInfoRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/paymentinformation", supplierCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchSupplierPaymentInfo, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/paymentinformation
func (c *Client) PatchSupplierPaymentInfo(ctx context.Context, supplierCode string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/paymentinformation", supplierCode)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetSupplierCreditStatus, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/creditstatus
func (c *Client) GetSupplierCreditStatus(ctx context.Context, supplierCode string) (SupplierCreditStatusResponse, error) {
	var result SupplierCreditStatusResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/creditstatus", supplierCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierCreditStatusResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierCreditStatusResponse{}, err
	}
	return result, nil
}

// op: UpdateSupplierCreditStatus, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/creditstatus
func (c *Client) UpdateSupplierCreditStatus(ctx context.Context, supplierCode string, payload UpdateSupplierCreditStatusRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/creditstatus", supplierCode)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllSupplierNotes, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes
func (c *Client) GetAllSupplierNotes(ctx context.Context, supplierCode string) ([]SupplierNoteResponse, error) {
	var result []SupplierNoteResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes", supplierCode)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes
func (c *Client) AddSupplierNote(ctx context.Context, supplierCode string, payload AddSupplierNoteRequest) (SupplierNoteResponse, error) {
	var result SupplierNoteResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes", supplierCode)
	if err := tasscommon.Validate(payload); err != nil {
		return SupplierNoteResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return SupplierNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierNoteResponse{}, err
	}
	return result, nil
}

// op: GetSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}
func (c *Client) GetSupplierNote(ctx context.Context, supplierCode string, noteID string) (SupplierNoteResponse, error) {
	var result SupplierNoteResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes/%s", supplierCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return SupplierNoteResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return SupplierNoteResponse{}, err
	}
	return result, nil
}

// op: UpdateSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}
func (c *Client) UpdateSupplierNote(ctx context.Context, supplierCode string, noteID string, payload UpdateSupplierNoteRequest) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes/%s", supplierCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: PatchSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}
func (c *Client) PatchSupplierNote(ctx context.Context, supplierCode string, noteID string, payload []tasscommon.Operation) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes/%s", supplierCode, noteID)
	_, err := c.t.request(ctx, http.MethodPatch, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: DeleteSupplierNote, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}
func (c *Client) DeleteSupplierNote(ctx context.Context, supplierCode string, noteID string) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes/%s", supplierCode, noteID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// op: GetAllSupplierNoteAttachments, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}/attachments
func (c *Client) GetAllSupplierNoteAttachments(ctx context.Context, supplierCode string, noteID string) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes/%s/attachments", supplierCode, noteID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: AddSupplierNoteAttachment, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}/attachments
func (c *Client) AddSupplierNoteAttachment(ctx context.Context, supplierCode string, noteID string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes/%s/attachments", supplierCode, noteID)
	if err := tasscommon.Validate(payload); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}

// op: DownloadSupplierNoteAttachment, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}/attachments/{attach_id}
func (c *Client) GetSupplierNoteAttachment(ctx context.Context, supplierCode string, noteID string, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes/%s/attachments/%s", supplierCode, noteID, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// op: DeleteSupplierNoteAttachment, path: /{cmpy_code}/finance/accountspayable/suppliers/{vend_code}/notes/{note_uid}/attachments/{attach_id}
func (c *Client) DeleteSupplierNoteAttachment(ctx context.Context, supplierCode string, noteID string, attachmentID string) error {
	url := fmt.Sprintf("/finance/accountspayable/suppliers/%s/notes/%s/attachments/%s", supplierCode, noteID, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}
