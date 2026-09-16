package tassfinance

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetAllPurchaseOrders
func (c *Client) GetAllPurchaseOrders() ([]PurchaseOrderResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddPurchaseOrder
func (c *Client) AddPurchaseOrder(payload AddPurchaseOrderRequest) (PurchaseOrderResponse, error) {
	// TODO: Implementation
	return PurchaseOrderResponse{}, nil
}

// op: GetPurchaseOrderById
func (c *Client) GetPurchaseOrderByID(orderNumber string) (PurchaseOrderResponse, error) {
	// TODO: Implementation
	return PurchaseOrderResponse{}, nil
}

// op: UpdatePurchaseOrder
func (c *Client) UpdatePurchaseOrder(orderNumber string, payload UpdatePurchaseOrderRequest) error {
	// TODO: Implementation
	return nil
}

// op: DeletePurchaseOrder
func (c *Client) DeletePurchaseOrder(orderNumber string) error {
	// TODO: Implementation
	return nil
}

// op: CancelPurchaseOrder
func (c *Client) CancelPurchaseOrder(orderNumber string) error {
	// TODO: Implementation
	return nil
}

// op: GetPurchaseOrderAttachments
func (c *Client) GetPurchaseOrderAttachments(orderNumber string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddPurchaseOrderAttachment
func (c *Client) AddPurchaseOrderAttachment(orderNumber string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadPurchaseOrderAttachment
func (c *Client) GetPurchaseOrderAttachment(orderNumber string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeletePurchaseOrderAttachment
func (c *Client) DeletePurchaseOrderAttachment(orderNumber string, attachmentID string) error {
	// TODO: Implementation
	return nil
}

// op: GetAllDeliveryPointOptions
func (c *Client) GetAllDeliveryPointOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetPurchaseOrderGLAccountOptions
func (c *Client) GetPurchaseOrderGLAccountOptions() ([]GLAccountOptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllPurchaseOrderStatusOptions
func (c *Client) GetAllPurchaseOrderStatusOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllPurchaseOrderTaxCodeOptions
func (c *Client) GetAllPurchaseOrderTaxCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllTermCodeOptions
func (c *Client) GetAllTermCodeOptions() ([]tasscommon.OptionsResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: GetAllVendorPurchaseOrderOptions
func (c *Client) GetAllVendorPurchaseOrderOptions() ([]tasscommon.OptionsResponseActive, error) {
	// TODO: Implementation
	return nil, nil
}
